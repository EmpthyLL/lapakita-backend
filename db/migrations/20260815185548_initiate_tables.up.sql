CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =============================================================================
-- 1. USERS, IDENTITIES & BANK ACCOUNTS
-- =============================================================================

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(32) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    
    -- Active Context & Platform Subscription
    active_role VARCHAR(32) DEFAULT 'tenant', -- tenant, stall_owner, supplier
    subscription_plan VARCHAR(32) DEFAULT 'free', -- free, single_role, all_access
    subscription_expires_at TIMESTAMP WITH TIME ZONE,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);

-- Data Identitas Resmi User (Tenant & Stall Owner)
CREATE TABLE user_identity_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    full_name_ktp VARCHAR(255) NOT NULL,
    nik VARCHAR(16) UNIQUE NOT NULL,
    ktp_photo_url TEXT NOT NULL, -- Foto KTP Ber-watermark
    domicile_city VARCHAR(128),  -- Opsional (Cukup Kota)
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_identity_profiles_user_id ON user_identity_profiles(user_id);

-- Rekening Bank User (Dapat dihapus permanen / Hard Delete)
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
    group_name VARCHAR(128) NOT NULL, -- e.g. "F&B", "Retail & Commerce", "Services"
    label VARCHAR(255) NOT NULL,
    
    -- Baseline Financial Parameters
    default_bep_months INT NOT NULL DEFAULT 6,
    default_capital NUMERIC(15, 2) NOT NULL DEFAULT 35000000.00,
    
    -- Physical Property Presets
    recommended_property_types JSONB DEFAULT '[]'::jsonb,
    recommended_placement VARCHAR(32) NOT NULL DEFAULT 'indoor',
    min_size_sqm NUMERIC(6, 2) NOT NULL DEFAULT 12.00,
    max_size_sqm NUMERIC(6, 2) NOT NULL DEFAULT 40.00,
    min_floors INT NOT NULL DEFAULT 1,
    max_floors INT NOT NULL DEFAULT 1,
    
    -- Facilities & Landmark Tags
    recommended_facilities JSONB DEFAULT '[]'::jsonb,
    recommended_landmarks JSONB DEFAULT '[]'::jsonb,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE TABLE businesses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Menggunakan user_id (Bukan owner_id)
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
-- 3. SUPPLIER PROFILES & REGULARS (FOLLOWERS)
-- =============================================================================

CREATE TABLE supplier_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    supplier_code VARCHAR(32) UNIQUE NOT NULL, -- Kode Unik Supplier (e.g. "SUP-88291")
    company_name VARCHAR(255) NOT NULL,
    description TEXT,
    logo_url TEXT,
    
    -- Dokumen Legal Usaha Supplier
    owner_ktp_photo_url TEXT NOT NULL, -- Foto KTP PJ Ber-watermark
    business_document_photo_url TEXT,  -- NIB / NPWP / Surat Izin Usaha Ber-watermark (Opsional)
    
    -- Verifikasi Admin & Target Market
    is_verified_by_admin BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    target_business_type_ids JSONB NOT NULL DEFAULT '[]'::jsonb,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_supplier_profiles_user_id ON supplier_profiles(user_id);
CREATE INDEX idx_supplier_profiles_supplier_code ON supplier_profiles(supplier_code);

-- Sistem Pelanggan Supplier (Tenant Follows Supplier)
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
-- 4. STALLS (PHYSICAL PROPERTIES LISTED BY STALL OWNERS)
-- =============================================================================

CREATE TABLE stalls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Istilah Spesifik Pemilik Lapak
    title VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Specs & Placement
    property_type VARCHAR(64) NOT NULL,
    placement VARCHAR(32) NOT NULL,
    size_sqm NUMERIC(8, 2) NOT NULL,
    length_meters NUMERIC(6, 2),
    width_meters NUMERIC(6, 2),
    electricity_capacity_va INT DEFAULT 1300,

    -- Location Data (Geoapify Clean Components)
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

    -- Multi-Period Rates & Financials
    allowed_payment_cycles JSONB NOT NULL DEFAULT '["month"]'::jsonb,
    monthly_rate NUMERIC(15, 2),
    quarterly_rate NUMERIC(15, 2),
    semesterly_rate NUMERIC(15, 2),
    yearly_rate NUMERIC(15, 2),
    security_deposit NUMERIC(15, 2) NOT NULL,

    -- Rules & Media Separation
    minimum_lease_months INT DEFAULT 1,
    start_date_options JSONB NOT NULL DEFAULT '["1", "15", "eom"]'::jsonb,
    utility_terms TEXT,
    facility_values JSONB DEFAULT '[]'::jsonb,
    house_rules JSONB DEFAULT '[]'::jsonb,
    
    display_media JSONB NOT NULL,    -- { mainImage, facilityImages: [], virtualTour360Url }
    legal_documents JSONB DEFAULT '[]'::jsonb, -- { certificateType, documentPhotos: [] }

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
CREATE INDEX idx_stalls_province ON stalls(province);
CREATE INDEX idx_stalls_is_published ON stalls(is_published);
CREATE INDEX idx_stalls_lat_lon ON stalls(latitude, longitude);
CREATE INDEX idx_stalls_deleted_at ON stalls(deleted_at);

-- =============================================================================
-- 5. LEASE CONTRACTS & REVIEWS
-- =============================================================================

CREATE TABLE lease_contracts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_id UUID NOT NULL REFERENCES stalls(id),
    tenant_user_id UUID NOT NULL REFERENCES users(id),
    stall_owner_id UUID NOT NULL REFERENCES users(id),
    business_id UUID REFERENCES businesses(id),
    
    -- Terms Lock
    selected_period VARCHAR(32) NOT NULL,
    agreed_rent_rate NUMERIC(15, 2) NOT NULL,
    agreed_security_deposit NUMERIC(15, 2) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    billing_cycle VARCHAR(32) NOT NULL,
    
    -- Status & Escrow
    status VARCHAR(32) DEFAULT 'pending_approval',
    escrow_deposit_status VARCHAR(32) DEFAULT 'held',
    deposit_claimed_amount NUMERIC(15, 2) DEFAULT 0.00,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_lease_contracts_tenant_user_id ON lease_contracts(tenant_user_id);
CREATE INDEX idx_lease_contracts_stall_owner_id ON lease_contracts(stall_owner_id);
CREATE INDEX idx_lease_contracts_stall_id ON lease_contracts(stall_id);
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
-- 6. POS CASHIER SYSTEM (WITH PROMO & FAVORITE FEATURES)
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
    
    -- Financials & Promo Support
    price NUMERIC(15, 2) NOT NULL,
    discount_price NUMERIC(15, 2) DEFAULT NULL, -- Harga Promo / Diskon (Opsional)
    cost_price NUMERIC(15, 2) DEFAULT 0.00,
    
    -- Inventory & Tags
    stock_quantity INT DEFAULT 0,
    track_stock BOOLEAN DEFAULT TRUE,
    is_favorite BOOLEAN DEFAULT FALSE,  -- Ditandai cepat di kasir
    is_bestseller BOOLEAN DEFAULT FALSE, -- Ditandai produk terlaris
    
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX idx_pos_items_business_id ON pos_items(business_id);
CREATE INDEX idx_pos_items_deleted_at ON pos_items(deleted_at);

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
    favorited_by_user_ids JSONB DEFAULT '[]'::jsonb, -- Favorit Produk Supplier
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX idx_supplier_catalogs_supplier_user_id ON supplier_catalogs(supplier_user_id);
CREATE INDEX idx_supplier_catalogs_deleted_at ON supplier_catalogs(deleted_at);

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
CREATE INDEX idx_supplier_orders_deleted_at ON supplier_orders(deleted_at);

-- =============================================================================
-- 8. NOTIFICATIONS & CMS
-- =============================================================================

-- Notifikasi Menggunakan Hard Delete
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
CREATE INDEX idx_notifications_is_read ON notifications(is_read);

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

-- Public FAQ Menggunakan Hard Delete
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