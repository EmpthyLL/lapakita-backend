-- 1. USERS & ACCOUNT SYSTEM
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(32) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    is_kyc_verified BOOLEAN DEFAULT FALSE,
    active_role VARCHAR(32) DEFAULT 'tenant', -- tenant, owner, supplier
    subscription_plan VARCHAR(32) DEFAULT 'free', -- free, single_role, all_access
    subscription_expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. BUSINESS TYPES (Master Preset & Category Data)
CREATE TABLE business_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    group_name VARCHAR(128) NOT NULL, -- "F&B (Food & Beverages)", "Retail & Commerce", "Services"
    label VARCHAR(255) NOT NULL,
    
    -- Baseline Financial Parameters
    default_bep_months INT NOT NULL DEFAULT 6,
    default_capital NUMERIC(15, 2) NOT NULL DEFAULT 35000000.00,
    
    -- Physical Property Presets (Template UI Filters)
    recommended_property_types JSONB DEFAULT '[]'::jsonb, -- e.g. ["shophouse", "mall-island"]
    recommended_placement VARCHAR(32) NOT NULL DEFAULT 'indoor', -- "indoor", "semi-outdoor", "outdoor"
    min_size_sqm NUMERIC(6, 2) NOT NULL DEFAULT 12.00,
    max_size_sqm NUMERIC(6, 2) NOT NULL DEFAULT 40.00,
    min_floors INT NOT NULL DEFAULT 1,
    max_floors INT NOT NULL DEFAULT 1,
    
    -- Facilities & Landmark Tags (Array of Strings)
    recommended_facilities JSONB DEFAULT '[]'::jsonb, -- e.g. ["power", "water"]
    recommended_landmarks JSONB DEFAULT '[]'::jsonb,  -- e.g. ["campus", "office"]
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. BUSINESS PROFILES (Tenant Businesses - e.g. "Kedai Kopi 90")
CREATE TABLE businesses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    business_type_id UUID NOT NULL REFERENCES business_types(id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    logo_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 4. SUPPLIER PROFILES & TARGET MARKET DEFAULTS
CREATE TABLE supplier_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    company_name VARCHAR(255) NOT NULL,
    description TEXT,
    logo_url TEXT,
    default_target_business_type_id UUID REFERENCES business_types(id) ON DELETE SET NULL,
    target_business_type_ids JSONB DEFAULT '[]'::jsonb, -- Array UUID untuk multi-target
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 5. STALLS (Physical Properties Listed by Owners)
CREATE TABLE stalls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    
    -- Specs & Placement
    property_type VARCHAR(64) NOT NULL,
    placement VARCHAR(32) NOT NULL,
    size_sqm NUMERIC(8, 2) NOT NULL,
    length_meters NUMERIC(6, 2),
    width_meters NUMERIC(6, 2),
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

    -- Multi-Period Rates & Financials
    monthly_rate NUMERIC(15, 2),
    quarterly_rate NUMERIC(15, 2),
    semesterly_rate NUMERIC(15, 2),
    yearly_rate NUMERIC(15, 2),
    cheapest_price_amount NUMERIC(15, 2) NOT NULL,
    cheapest_price_period VARCHAR(32) NOT NULL,
    security_deposit NUMERIC(15, 2) NOT NULL,
    allowed_payment_cycles JSONB NOT NULL DEFAULT '["month"]'::jsonb,

    -- Rules & Media
    minimum_lease_months INT DEFAULT 1,
    start_date_options JSONB NOT NULL DEFAULT '["1", "15", "eom"]'::jsonb,
    utility_terms TEXT,
    facility_values JSONB DEFAULT '[]'::jsonb,
    house_rules JSONB DEFAULT '[]'::jsonb,
    media JSONB NOT NULL,

    -- Ratings & State
    rating_avg NUMERIC(3, 2) DEFAULT 0.00,
    review_count INT DEFAULT 0,
    is_published BOOLEAN DEFAULT FALSE,
    favorited_by_user_ids JSONB DEFAULT '[]'::jsonb,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_stalls_city ON stalls(city);
CREATE INDEX idx_stalls_province ON stalls(province);
CREATE INDEX idx_stalls_is_published ON stalls(is_published);
CREATE INDEX idx_stalls_lat_lon ON stalls(latitude, longitude);

-- 6. LEASE CONTRACTS & RENTAL HISTORY
CREATE TABLE lease_contracts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_id UUID NOT NULL REFERENCES stalls(id),
    tenant_id UUID NOT NULL REFERENCES users(id),
    owner_id UUID NOT NULL REFERENCES users(id),
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

-- 7. STALL REVIEWS & OWNER RESPONSES
CREATE TABLE stall_reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    stall_id UUID NOT NULL REFERENCES stalls(id) ON DELETE CASCADE,
    tenant_id UUID NOT NULL REFERENCES users(id),
    lease_contract_id UUID REFERENCES lease_contracts(id),
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT NOT NULL,
    photos JSONB DEFAULT '[]'::jsonb,
    owner_response JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 8. POS CASHIER SYSTEM
CREATE TABLE pos_staff_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    pin_code_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pos_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pos_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    category_id UUID REFERENCES pos_categories(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(64),
    price NUMERIC(15, 2) NOT NULL,
    cost_price NUMERIC(15, 2) DEFAULT 0.00,
    stock_quantity INT DEFAULT 0,
    track_stock BOOLEAN DEFAULT TRUE,
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pos_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    business_id UUID NOT NULL REFERENCES businesses(id) ON DELETE CASCADE,
    staff_id UUID REFERENCES pos_staff_accounts(id),
    receipt_number VARCHAR(64) UNIQUE NOT NULL,
    total_amount NUMERIC(15, 2) NOT NULL,
    payment_method VARCHAR(32) NOT NULL,
    items_json JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 9. SUPPLIER B2B MARKETPLACE
CREATE TABLE supplier_catalogs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    supplier_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    target_business_type_id UUID REFERENCES business_types(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    unit_type VARCHAR(32) NOT NULL,
    price_per_unit NUMERIC(15, 2) NOT NULL,
    moq INT DEFAULT 1,
    tiered_pricing JSONB,
    is_available BOOLEAN DEFAULT TRUE,
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE supplier_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES users(id),
    business_id UUID NOT NULL REFERENCES businesses(id),
    supplier_id UUID NOT NULL REFERENCES users(id),
    order_number VARCHAR(64) UNIQUE NOT NULL,
    total_amount NUMERIC(15, 2) NOT NULL,
    status VARCHAR(32) DEFAULT 'pending',
    items_json JSONB NOT NULL,
    digital_delivery_note_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 10. CONTACT INQUIRIES & PUBLIC CMS DATA
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
    category_id VARCHAR(32) NOT NULL,
    sub_topic_title VARCHAR(255) NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    role_type VARCHAR(32),
    sort_order INT DEFAULT 0
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