-- DOWN MIGRATION (Corrected Order with Foreign Key Dependencies)
-- 1. Drop CMS & Inquiry Tables
DROP TABLE IF EXISTS cms_legal_documents CASCADE;

DROP TABLE IF EXISTS cms_public_faqs CASCADE;

DROP TABLE IF EXISTS contact_inquiries CASCADE;

-- 2. Drop Supplier B2B Marketplace Tables
DROP TABLE IF EXISTS supplier_orders CASCADE;

DROP TABLE IF EXISTS supplier_catalogs CASCADE;

DROP TABLE IF EXISTS supplier_profiles CASCADE;

-- 3. Drop POS Cashier System Tables
DROP TABLE IF EXISTS pos_transactions CASCADE;

DROP TABLE IF EXISTS pos_items CASCADE;

DROP TABLE IF EXISTS pos_categories CASCADE;

DROP TABLE IF EXISTS pos_staff_accounts CASCADE;

-- 4. Drop Stall Reviews & Lease Contracts
DROP TABLE IF EXISTS stall_reviews CASCADE;

DROP TABLE IF EXISTS lease_contracts CASCADE;

-- 5. Drop Stalls, Businesses & Business Types
DROP TABLE IF EXISTS stalls CASCADE;

DROP TABLE IF EXISTS businesses CASCADE;

DROP TABLE IF EXISTS business_types CASCADE;

-- 6. Drop Users Table
DROP TABLE IF EXISTS users CASCADE;