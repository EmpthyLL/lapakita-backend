CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =============================================================================
-- ENUM DEFINITIONS (Core Types)
-- =============================================================================

-- Level Kontrol Operasional & Dependensi Lapak
CREATE TYPE stall_permanence_type AS ENUM ('permanent', 'semi-permanent', 'temporary');
CREATE TYPE stall_placement_type AS ENUM ('indoor', 'semi-outdoor', 'outdoor');

CREATE TYPE event_operating_days_type AS ENUM ('everyday', 'weekends', 'weekdays', 'flexible');
CREATE TYPE attendance_requirement_type AS ENUM ('mandatory_full', 'flexible_days');
CREATE TYPE cancellation_policy_type AS ENUM ('pro_rata', 'deposit_refundable', 'non_refundable');

-- =============================================================================
-- 1. USERS, IDENTITIES & MULTI-ROLE PROFILES
-- =============================================================================

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL, -- Primary Identity & Login ID
    password_hash VARCHAR(255) NOT NULL,
    default_avatar_url TEXT,
    
    -- Telepon Kontak (Multi-Nomor: Primary, Secondary, WhatsApp)
    phone_numbers JSONB NOT NULL DEFAULT '[{"number": "", "is_primary": true, "roles": []}]'::jsonb,
    
    -- Profil Kustom Per Role (Avatar & Display Name Independen Per Role)
    role_profiles JSONB DEFAULT '{}'::jsonb,
    
    -- Active Context & Platform Subscription
    active_role VARCHAR(32) DEFAULT 'tenant', -- tenant, owner, supplier
    subscription_plan VARCHAR(32) DEFAULT 'free', -- free, single_role, all_access
    subscription_expires_at TIMESTAMP WITH TIME ZONE,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

CREATE TABLE user_identity_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    document_type VARCHAR(32) NOT NULL DEFAULT 'national_id', -- 'national_id', 'passport', 'residence_permit'
    full_name_identity VARCHAR(255) NOT NULL,                  -- Nama lengkap sesuai dokumen
    document_number VARCHAR(64) NOT NULL,                      -- NIK / No. Paspor / No. Izin Tinggal
    document_photo_url TEXT NOT NULL,                          -- Foto Dokumen ber-watermark
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexing untuk mempercepat pencarian query
CREATE INDEX idx_user_identity_doc_num ON user_identity_profiles(document_number);
CREATE INDEX idx_user_identity_user_id ON user_identity_profiles(user_id);

-- Rekening Bank User (Penarikan Dana Escrow / Payout)
CREATE TABLE bank_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    bank_code VARCHAR(32) NOT NULL,        -- e.g. "BCA", "MANDIRI", "BRI", "BNI"
    bank_name VARCHAR(128) NOT NULL,       -- e.g. "Bank Central Asia"
    account_number VARCHAR(64) NOT NULL,
    account_holder_name VARCHAR(255) NOT NULL,
    
    is_primary BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bank_accounts_user_id ON bank_accounts(user_id);

-- =============================================================================
-- 2. BUSINESS TYPES & TENANT BUSINESS PROFILES
-- =============================================================================

CREATE TABLE business_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Label & Group Name Multi-Bahasa (I18n)
    label_lang JSONB NOT NULL DEFAULT '{"en": "", "id": ""}'::jsonb,
    group_name_lang JSONB NOT NULL DEFAULT '{"en": "", "id": ""}'::jsonb,
    
    -- Financial Benchmarks
    default_bep_months INT NOT NULL DEFAULT 6,
    default_capital NUMERIC(15, 2) NOT NULL DEFAULT 35000000.00,
    avg_gross_margin_ratio NUMERIC(5, 4) NOT NULL DEFAULT 0.5000,
    industry_rent_to_revenue_ratio NUMERIC(5, 4) NOT NULL DEFAULT 0.1500,
    
    -- Presets Per Permanence Tab (Permanent, Semi-Permanent, Temporary)
    permanence_presets JSONB NOT NULL DEFAULT '{}'::jsonb,
    
    -- Target Landmark Tags
    recommended_landmarks JSONB DEFAULT '[]'::jsonb,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE TABLE businesses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    business_type_id UUID NOT NULL REFERENCES business_types(id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    logo_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX idx_businesses_user_id ON businesses(user_id);
CREATE INDEX idx_businesses_deleted_at ON businesses(deleted_at);

-- =============================================================================
-- 3. SUPPLIER PROFILES & REGULARS
-- =============================================================================

CREATE TABLE supplier_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    supplier_code VARCHAR(32) UNIQUE NOT NULL, -- e.g. "SUP-88291"
    company_name VARCHAR(255) NOT NULL,
    description TEXT,
    logo_url TEXT,
    
    supplier_ktp_photo_url TEXT NOT NULL,
    business_document_photo_url TEXT,
    
    is_verified_by_admin BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    target_business_type_ids JSONB NOT NULL DEFAULT '[]'::jsonb,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_supplier_profiles_user_id ON supplier_profiles(user_id);
CREATE INDEX idx_supplier_profiles_supplier_code ON supplier_profiles(supplier_code);

CREATE TABLE supplier_regulars (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    business_id UUID REFERENCES businesses(id) ON DELETE SET NULL,
    supplier_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(tenant_user_id, supplier_user_id)
);

CREATE INDEX idx_supplier_regulars_tenant_user_id ON supplier_regulars(tenant_user_id);
CREATE INDEX idx_supplier_regulars_supplier_user_id ON supplier_regulars(supplier_user_id);

-- =============================================================================
-- 4. PHYSICAL STALLS (SINGLE SOURCE OF TRUTH UNTUK SEMUA LAPAK & BAZAAR)
-- =============================================================================

CREATE TABLE stalls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Klasifikasi Properti
    property_type VARCHAR(64) NOT NULL,                  -- e.g., 'shophouse', 'mall-shop', 'bazaar-booth'
    permanence_type stall_permanence_type NOT NULL DEFAULT 'permanent',
    placement stall_placement_type NOT NULL DEFAULT 'indoor',
    
    -- Physical Specs (Opsional untuk Semi / Temporary)
    size_sqm NUMERIC(8, 2) DEFAULT NULL,
    length_meters NUMERIC(6, 2) DEFAULT NULL,
    width_meters NUMERIC(6, 2) DEFAULT NULL,
    floor_level INT DEFAULT 1,
    electricity_capacity_va INT DEFAULT 1300,

    -- ── 1. SEMI-PERMANENT CONTEXT (Parent Complex & Operating Hours) ──
    parent_complex_name VARCHAR(255) DEFAULT NULL,       -- e.g. "Plaza Margonda", "Pasar Johar"
    operating_hours JSONB DEFAULT NULL,                  -- e.g. {"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}

    -- ── 2. TEMPORARY CONTEXT (Event Schedule & Slots) ──
    event_schedule JSONB DEFAULT NULL,                   -- e.g. {"event_name": "Ramadan Fest", "start_date": "2026-03-20", "end_date": "2026-03-23", "registration_deadline": "2026-03-15"}
    slot_info JSONB DEFAULT NULL,                        -- e.g. {"total_slots": 20, "available_slots": 6}

    -- Location Data
    street_address TEXT NOT NULL,
    suburb VARCHAR(128),
    district VARCHAR(128),
    city VARCHAR(128) NOT NULL,
    province VARCHAR(128) NOT NULL,
    country VARCHAR(128) NOT NULL DEFAULT 'Indonesia',
    country_code VARCHAR(8) DEFAULT 'ID',
    postal_code VARCHAR(16),
    latitude NUMERIC(10, 8),
    longitude NUMERIC(11, 8),
    map_url TEXT,
    embedded_map_url TEXT,
    nearby_landmarks JSONB DEFAULT '[]'::jsonb,

    -- Multi-Cycle Rates
    allowed_payment_cycles JSONB NOT NULL DEFAULT '["month"]'::jsonb, -- e.g. ["day", "month"]
    daily_rate NUMERIC(15, 2) DEFAULT NULL,
    monthly_rate NUMERIC(15, 2) DEFAULT NULL,
    quarterly_rate NUMERIC(15, 2) DEFAULT NULL,
    semesterly_rate NUMERIC(15, 2) DEFAULT NULL,
    yearly_rate NUMERIC(15, 2) DEFAULT NULL,
    security_deposit NUMERIC(15, 2) NOT NULL DEFAULT 0.00,

    -- ── 3. LEASE TERMS & OPTIONS ──
    minimum_lease_months INT DEFAULT NULL,                -- Digunakan oleh Permanent & Semi-Permanent
    minimum_lease_days INT DEFAULT NULL,                  -- Digunakan oleh Temporary Event
    start_date_options JSONB DEFAULT '[]'::jsonb,        -- Permanent: ["1", "15", "eom"], Temp: ["event_day_1", "event_day_2", "event_week_1"]
    
    -- Aturan Sewa Event Khusus Temporary (Nullable untuk Permanent & Semi)
    event_operating_days event_operating_days_type DEFAULT NULL,
    event_attendance_requirement attendance_requirement_type DEFAULT NULL,
    event_cancellation_policy cancellation_policy_type DEFAULT NULL,

    utility_terms TEXT,

    -- Facilities & Matchmaking
    facility_values JSONB DEFAULT '[]'::jsonb,           -- e.g. ["power", "water", "trash-area"]
    allowed_business_type_ids JSONB DEFAULT '[]'::jsonb, -- Filter bisnis yang diizinkan
    house_rules JSONB DEFAULT '[]'::jsonb,
    
    display_media JSONB NOT NULL,                        -- { mainImage, facilityImages: [] }
    legal_documents JSONB DEFAULT '[]'::jsonb,

    -- Ratings & State
    rating_avg NUMERIC(3, 2) DEFAULT 0.00,
    review_count INT DEFAULT 0,
    is_published BOOLEAN DEFAULT FALSE,
    favorited_by_user_ids JSONB DEFAULT '[]'::jsonb,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX idx_stalls_city ON stalls(city);
CREATE INDEX idx_stalls_permanence ON stalls(permanence_type);
CREATE INDEX idx_stalls_property_type ON stalls(property_type);
CREATE INDEX idx_stalls_is_published ON stalls(is_published);
CREATE INDEX idx_stalls_lat_lon ON stalls(latitude, longitude);
CREATE INDEX idx_stalls_deleted_at ON stalls(deleted_at);

-- =============================================================================
-- 5. LEASE CONTRACTS & REVIEWS
-- =============================================================================

CREATE TABLE lease_contracts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Relasi Entitas Utama
    stall_id UUID NOT NULL REFERENCES stalls(id),
    tenant_user_id UUID NOT NULL REFERENCES users(id),
    stall_owner_id UUID NOT NULL REFERENCES users(id),
    business_id UUID REFERENCES businesses(id) ON DELETE SET NULL,
    
    -- Snapshots Identitas Dokumen Legal (Di-lock saat pengajuan)
    tenant_identity_snapshot JSONB NOT NULL, 
    owner_identity_snapshot JSONB NOT NULL, 

    -- Rincian Nilai & Durasi Sewa
    selected_period VARCHAR(32) NOT NULL,            -- e.g. "1_year", "3_months", "event_4_days"
    agreed_rent_rate NUMERIC(15, 2) NOT NULL,
    agreed_security_deposit NUMERIC(15, 2) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    billing_cycle VARCHAR(32) NOT NULL,               -- 'day', 'month', 'quarter', 'year'

    -- Digital Signature Audit Trail
    tenant_signature JSONB NOT NULL,
    owner_signature JSONB DEFAULT NULL,
    contract_hash_sha256 VARCHAR(64) DEFAULT NULL,

    -- Status Kontrak Utama (Tanpa Atribut Escrow Ribet)
    status VARCHAR(32) NOT NULL DEFAULT 'pending_approval', 
    -- 'pending_approval', 'payment_pending', 'active', 'rejected', 'cancelled', 'completed'
    
    -- Simple Deposit Refund Tracker
    is_deposit_refunded BOOLEAN DEFAULT FALSE,
    refunded_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_lease_contracts_tenant ON lease_contracts(tenant_user_id);
CREATE INDEX idx_lease_contracts_owner ON lease_contracts(stall_owner_id);
CREATE INDEX idx_lease_contracts_stall ON lease_contracts(stall_id);
CREATE INDEX idx_lease_contracts_status ON lease_contracts(status);

CREATE TABLE stall_reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_id UUID NOT NULL REFERENCES stalls(id) ON DELETE CASCADE,
    tenant_user_id UUID NOT NULL REFERENCES users(id),
    lease_contract_id UUID REFERENCES lease_contracts(id),
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT NOT NULL,
    photos JSONB DEFAULT '[]'::jsonb,
    owner_response JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_stall_reviews_stall_id ON stall_reviews(stall_id);

-- =============================================================================
-- INSPEKSI, KLAIM KERUSAKAN & PERMINTAAN KUNCI LAPAK
-- =============================================================================

-- 1. INSPEKSI KONDISI LAPAK (Move-in / Move-out Check)
CREATE TABLE stall_inspections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lease_contract_id UUID NOT NULL REFERENCES lease_contracts(id) ON DELETE CASCADE,
    
    inspector_role VARCHAR(16) NOT NULL,                 -- 'owner' atau 'tenant'
    inspection_type VARCHAR(32) NOT NULL,                -- 'move_in' atau 'move_out'
    notes TEXT,
    photos JSONB NOT NULL DEFAULT '[]'::jsonb,           -- Bukti foto kondisi fisik
    checked_items JSONB NOT NULL DEFAULT '{}'::jsonb,    -- e.g. {"lock_cylinder": "good", "electricity": "ok"}
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_stall_inspections_contract ON stall_inspections(lease_contract_id);


-- 2. LAPORAN & KLAIM KERUSAKAN (Damage & Utility Arrears Claims)
CREATE TABLE damage_claims (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lease_contract_id UUID NOT NULL REFERENCES lease_contracts(id) ON DELETE CASCADE,
    
    claim_reason TEXT NOT NULL,                          -- Rincian kerusakan / tunggakan utilitas
    claimed_amount NUMERIC(15, 2) NOT NULL,             -- Nominal pemotongan deposit yang diminta Owner
    proof_photos JSONB NOT NULL DEFAULT '[]'::jsonb,     -- Bukti foto kerusakan
    
    status VARCHAR(32) NOT NULL DEFAULT 'pending',       -- 'pending', 'accepted_by_tenant', 'appealed', 'resolved'
    tenant_response_notes TEXT DEFAULT NULL,
    
    resolved_amount NUMERIC(15, 2) DEFAULT NULL,         -- Nominal kesepakatan akhir
    resolved_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_damage_claims_contract ON damage_claims(lease_contract_id);
CREATE INDEX idx_damage_claims_status ON damage_claims(status);


-- 3. PERMINTAAN KUNCI & AKSES (Key Handover & Loss Management)
CREATE TABLE key_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lease_contract_id UUID NOT NULL REFERENCES lease_contracts(id) ON DELETE CASCADE,
    
    request_type VARCHAR(32) NOT NULL,                  -- 'initial_handover', 'duplicate_spare', 'lost_key'
    notes TEXT,
    fee_amount NUMERIC(15, 2) DEFAULT 0.00,             -- Biaya jika duplikat / ganti silinder akibat kelalaian
    
    status VARCHAR(32) NOT NULL DEFAULT 'pending',       -- 'pending', 'approved', 'completed', 'rejected'
    handled_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_key_requests_contract ON key_requests(lease_contract_id);

-- =============================================================================
-- 6. POS CASHIER SYSTEM
-- =============================================================================

CREATE TABLE pos_staff_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    pin_code_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pos_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pos_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    category_id UUID REFERENCES pos_categories(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(64),
    
    price NUMERIC(15, 2) NOT NULL,
    discount_price NUMERIC(15, 2) DEFAULT NULL,
    cost_price NUMERIC(15, 2) DEFAULT 0.00,
    
    stock_quantity INT DEFAULT 0,
    track_stock BOOLEAN DEFAULT TRUE,
    is_favorite BOOLEAN DEFAULT FALSE,
    is_bestseller BOOLEAN DEFAULT FALSE,
    
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_pos_items_business_id ON pos_items(business_id);

CREATE TABLE pos_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    staff_id UUID REFERENCES pos_staff_accounts(id),
    receipt_number VARCHAR(64) UNIQUE NOT NULL,
    total_amount NUMERIC(15, 2) NOT NULL,
    payment_method VARCHAR(32) NOT NULL,
    items_json JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_pos_transactions_business_id ON pos_transactions(business_id);

-- =============================================================================
-- 7. SUPPLIER B2B MARKETPLACE
-- =============================================================================

CREATE TABLE supplier_catalogs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    supplier_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    target_business_type_ids JSONB DEFAULT '[]'::jsonb, 
    name VARCHAR(255) NOT NULL,
    description TEXT,
    unit_type VARCHAR(32) NOT NULL,
    price_per_unit NUMERIC(15, 2) NOT NULL,
    moq INT DEFAULT 1,
    tiered_pricing JSONB,
    is_available BOOLEAN DEFAULT TRUE,
    image_url TEXT,
    favorited_by_user_ids JSONB DEFAULT '[]'::jsonb,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_supplier_catalogs_supplier_user_id ON supplier_catalogs(supplier_user_id);

CREATE TABLE supplier_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_user_id UUID NOT NULL REFERENCES users(id),
    business_id UUID NOT NULL REFERENCES businesses(id),
    supplier_user_id UUID NOT NULL REFERENCES users(id),
    order_number VARCHAR(64) UNIQUE NOT NULL,
    total_amount NUMERIC(15, 2) NOT NULL,
    status VARCHAR(32) DEFAULT 'pending',
    items_json JSONB NOT NULL,
    digital_delivery_note_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_supplier_orders_tenant_user_id ON supplier_orders(tenant_user_id);
CREATE INDEX idx_supplier_orders_supplier_user_id ON supplier_orders(supplier_user_id);

-- =============================================================================
-- 8. NOTIFICATIONS & CMS
-- =============================================================================

CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    type VARCHAR(64) NOT NULL,
    target_role VARCHAR(32) NOT NULL DEFAULT 'tenant',
    
    action_url TEXT,
    metadata JSONB DEFAULT '{}'::jsonb,
    
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_notifications_user_id ON notifications(user_id);

CREATE TABLE contact_inquiries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    whatsapp VARCHAR(32),
    persona VARCHAR(64) NOT NULL,
    inquiry_type VARCHAR(64) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(32) DEFAULT 'unread',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cms_public_faqs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lang VARCHAR(8) NOT NULL DEFAULT 'en',
    sub_topic_title VARCHAR(255) NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    role_type VARCHAR(32),
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cms_legal_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    doc_type VARCHAR(32) NOT NULL,
    lang VARCHAR(8) NOT NULL DEFAULT 'en',
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    sections_json JSONB NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(doc_type, lang)
);

-- =============================================================================
-- 9. GENERATED REPORTS (Multi-Role Analysis History)
-- =============================================================================

CREATE TABLE generated_reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    role_type VARCHAR(32) NOT NULL,             -- 'tenant', 'stall_owner', 'supplier'
    report_type VARCHAR(64) NOT NULL,           -- 'multi_timeline_forecast', 'vacancy_loss_analysis', 'opportunity_gap_analysis'
    
    business_id UUID REFERENCES businesses(id) ON DELETE SET NULL,
    stall_id UUID REFERENCES stalls(id) ON DELETE SET NULL,
    
    title VARCHAR(255) NOT NULL,                -- e.g. "Forecast Q3 2026 - Coffee Shop A"
    
    input_snapshot JSONB NOT NULL DEFAULT '{}'::jsonb,
    result_payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_generated_reports_user_role ON generated_reports(user_id, role_type);
CREATE INDEX idx_generated_reports_created_at ON generated_reports(created_at DESC);