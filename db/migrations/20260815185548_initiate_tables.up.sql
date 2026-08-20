CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =============================================================================
-- ENUM DEFINITIONS (Core Types)
-- =============================================================================

-- Level Kontrol Operasional & Dependensi Lapak
CREATE TYPE stall_permanence_type AS ENUM (
    'permanent',        -- Independent / Standalone (Mandiri, 24/7 Access, Tanpa Induk)
    'semi-permanent',   -- Managed Complex (Terikat Jam Buka/Tutup Pengelola Induk: Mall, Pasar, Foodcourt)
    'temporary'         -- Temporary & Event Spots (Bazaar Pop-Up, Kakilima, Food Truck)
);

CREATE TYPE stall_placement_type AS ENUM (
    'indoor',
    'semi-outdoor',
    'outdoor'
);

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
    phone_numbers JSONB NOT NULL DEFAULT '[{"type": "primary", "number": "", "is_whatsapp": true}]'::jsonb,
    
    -- Profil Kustom Per Role (Avatar & Display Name Independen Per Role)
    -- Format JSON: { "tenant": { "name": "...", "avatar_url": "..." }, "stall_owner": { ... }, "supplier": { ... } }
    role_profiles JSONB DEFAULT '{}'::jsonb,
    
    -- Active Context & Platform Subscription
    active_role VARCHAR(32) DEFAULT 'tenant', -- tenant, stall_owner, supplier
    subscription_plan VARCHAR(32) DEFAULT 'free', -- free, single_role, all_access
    subscription_expires_at TIMESTAMP WITH TIME ZONE,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

-- Data Identitas Resmi User (KYC untuk Tenant & Stall Owner)
CREATE TABLE user_identity_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    full_name_ktp VARCHAR(255) NOT NULL,
    nik VARCHAR(16) UNIQUE NOT NULL,
    ktp_photo_url TEXT NOT NULL, -- Foto KTP Ber-watermark
    domicile_city VARCHAR(128),
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_identity_profiles_user_id ON user_identity_profiles(user_id);

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
-- 2. BUSINESS TYPES (WITH I18N SUPPORT) & TENANT BUSINESS PROFILES
-- =============================================================================

CREATE TABLE business_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(128) UNIQUE NOT NULL,    -- e.g. "full-service-restaurant", "coffee-shop-cafe"
    
    -- Label & Group Name Multi-Bahasa (I18n)
    -- Format JSONB: {"en": "Full-Service Restaurant", "id": "Restoran Layanan Penuh"}
    label_lang JSONB NOT NULL DEFAULT '{"en": "", "id": ""}'::jsonb,
    group_name_lang JSONB NOT NULL DEFAULT '{"en": "", "id": ""}'::jsonb,
    
    -- Financial Benchmarks
    default_bep_months INT NOT NULL DEFAULT 6,
    default_capital NUMERIC(15, 2) NOT NULL DEFAULT 35000000.00,
    avg_gross_margin_ratio NUMERIC(5, 4) NOT NULL DEFAULT 0.5000,
    industry_rent_to_revenue_ratio NUMERIC(5, 4) NOT NULL DEFAULT 0.1500,
    
    -- Presets Per Permanence Tab (Permanent/Independent, Semi-Permanent/Managed, Temporary)
    -- Menyimpan JSONB struktur permanencePresets dari TypeScript
    permanence_presets JSONB NOT NULL DEFAULT '{}'::jsonb,
    
    -- Target Landmark Tags
    recommended_landmarks JSONB DEFAULT '[]'::jsonb,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX idx_business_types_slug ON business_types(slug);

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
    
    owner_ktp_photo_url TEXT NOT NULL,
    business_document_photo_url TEXT,
    
    is_verified_by_admin BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
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
-- 4. PHYSICAL STALLS (PROPERTY LISTINGS)
-- =============================================================================

CREATE TABLE stalls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Physical Property Classification
    property_type VARCHAR(64) NOT NULL,                  -- e.g., 'shophouse', 'mall-shop', 'open-market-stall'
    permanence_type stall_permanence_type NOT NULL DEFAULT 'permanent',
    placement stall_placement_type NOT NULL DEFAULT 'indoor',
    
    -- Physical Specs
    size_sqm NUMERIC(8, 2) NOT NULL,
    length_meters NUMERIC(6, 2),
    width_meters NUMERIC(6, 2),
    floor_level INT DEFAULT 1,
    electricity_capacity_va INT DEFAULT 1300,

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

    -- Rates & Payment Cycles
    allowed_payment_cycles JSONB NOT NULL DEFAULT '["month"]'::jsonb,
    daily_rate NUMERIC(15, 2),
    monthly_rate NUMERIC(15, 2),
    quarterly_rate NUMERIC(15, 2),
    semesterly_rate NUMERIC(15, 2),
    yearly_rate NUMERIC(15, 2),
    security_deposit NUMERIC(15, 2) NOT NULL,

    -- Terms & Contextual Facilities
    minimum_lease_months INT DEFAULT 1,
    start_date_options JSONB NOT NULL DEFAULT '["1", "15", "eom"]'::jsonb,
    utility_terms TEXT,
    facility_values JSONB DEFAULT '[]'::jsonb,           -- e.g. ["power", "water", "wifi"]
    allowed_business_type_ids JSONB DEFAULT '[]'::jsonb, -- Matchmaking filter
    house_rules JSONB DEFAULT '[]'::jsonb,
    
    display_media JSONB NOT NULL,
    legal_documents JSONB DEFAULT '[]'::jsonb,

    -- Rating & State
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
-- 5. BAZAARS & POP-UP EVENTS (SHORT-TERM EVENT DOMAIN)
-- =============================================================================

CREATE TABLE bazaars (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organizer_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Event Schedule & Registration Window
    registration_start_date DATE NOT NULL,
    registration_end_date DATE NOT NULL,
    event_start_date DATE NOT NULL,
    event_end_date DATE NOT NULL,
    
    -- Event Location
    venue_name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    city VARCHAR(128) NOT NULL,
    latitude NUMERIC(10, 8),
    longitude NUMERIC(11, 8),
    
    -- Allowed Business Types
    allowed_business_type_ids JSONB DEFAULT '[]'::jsonb,
    
    banner_url TEXT,
    layout_map_url TEXT,
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bazaars_event_dates ON bazaars(event_start_date, event_end_date);

-- Booth Slots inside Event Bazaar
CREATE TABLE bazaar_booths (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bazaar_id UUID NOT NULL REFERENCES bazaars(id) ON DELETE CASCADE,
    booth_code VARCHAR(32) NOT NULL,          -- e.g. "A-01", "VIP-03"
    size_sqm NUMERIC(6, 2) NOT NULL,
    placement stall_placement_type DEFAULT 'indoor',
    
    -- Financials (Full Event Duration Package Rate)
    total_price_amount NUMERIC(15, 2) NOT NULL,
    included_facilities JSONB DEFAULT '["power", "trash-area"]'::jsonb,
    
    -- Status
    status VARCHAR(32) DEFAULT 'available',    -- 'available', 'booked', 'occupied'
    booked_by_user_id UUID REFERENCES users(id),
    booked_business_id UUID REFERENCES businesses(id),
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bazaar_booths_bazaar_status ON bazaar_booths(bazaar_id, status);

-- =============================================================================
-- 6. LEASE CONTRACTS & REVIEWS
-- =============================================================================

CREATE TABLE lease_contracts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_id UUID NOT NULL REFERENCES stalls(id),
    tenant_user_id UUID NOT NULL REFERENCES users(id),
    stall_owner_id UUID NOT NULL REFERENCES users(id),
    business_id UUID REFERENCES businesses(id),
    
    selected_period VARCHAR(32) NOT NULL,
    agreed_rent_rate NUMERIC(15, 2) NOT NULL,
    agreed_security_deposit NUMERIC(15, 2) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    billing_cycle VARCHAR(32) NOT NULL,
    
    status VARCHAR(32) DEFAULT 'pending_approval',
    escrow_deposit_status VARCHAR(32) DEFAULT 'held',
    deposit_claimed_amount NUMERIC(15, 2) DEFAULT 0.00,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_lease_contracts_tenant_user_id ON lease_contracts(tenant_user_id);
CREATE INDEX idx_lease_contracts_stall_owner_id ON lease_contracts(stall_owner_id);
CREATE INDEX idx_lease_contracts_stall_id ON lease_contracts(stall_id);

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
-- 7. POS CASHIER SYSTEM
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
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
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
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
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
-- 8. SUPPLIER B2B MARKETPLACE
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
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
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
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX idx_supplier_orders_tenant_user_id ON supplier_orders(tenant_user_id);
CREATE INDEX idx_supplier_orders_supplier_user_id ON supplier_orders(supplier_user_id);

-- =============================================================================
-- 9. NOTIFICATIONS & CMS
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
    read_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    
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

-- FAQ Publik (Bahasa Indonesia & Inggris via lang)
CREATE TABLE cms_public_faqs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lang VARCHAR(8) NOT NULL DEFAULT 'en',
    category_id VARCHAR(32) NOT NULL,
    sub_topic_title VARCHAR(255) NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    role_type VARCHAR(32),
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Dokumen Legal Publik (Sektor Bahasa Menggunakan lang)
CREATE TABLE cms_legal_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    doc_type VARCHAR(32) NOT NULL,
    lang VARCHAR(8) NOT NULL DEFAULT 'en',
    title VARCHAR(255) NOT NULL,
    intro TEXT NOT NULL,
    sections_json JSONB NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(doc_type, lang)
);

-- =============================================================================
-- 10. GENERATED REPORTS (Multi-Role Analysis History)
-- =============================================================================

CREATE TABLE generated_reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    -- Role Context & Jenis Laporan
    role_type VARCHAR(32) NOT NULL,             -- 'tenant', 'stall_owner', 'supplier'
    report_type VARCHAR(64) NOT NULL,           -- 'multi_timeline_forecast', 'vacancy_loss_analysis', 'opportunity_gap_analysis'
    
    -- Entity Reference (Opsional: terikat ke Bisnis atau Lapak spesifik)
    business_id UUID REFERENCES businesses(id) ON DELETE SET NULL,
    stall_id UUID REFERENCES stalls(id) ON DELETE SET NULL,
    
    title VARCHAR(255) NOT NULL,                -- e.g. "Forecast Q3 2026 - Coffee Shop A"
    
    -- Flexible JSON Payload
    -- Menampung snapshot input (Preset/Excel summary) & hasil keluaran AI/Forecast
    input_snapshot JSONB NOT NULL DEFAULT '{}'::jsonb,
    result_payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_generated_reports_user_role ON generated_reports(user_id, role_type);
CREATE INDEX idx_generated_reports_created_at ON generated_reports(created_at DESC);