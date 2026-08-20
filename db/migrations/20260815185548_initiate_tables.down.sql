-- DOWN MIGRATION (Strict Dependency Order)
-- 1. Drop CMS & Inquiry Tables
DROP TABLE IF EXISTS cms_legal_documents CASCADE;

DROP TABLE IF EXISTS cms_public_faqs CASCADE;

DROP TABLE IF EXISTS contact_inquiries CASCADE;

-- 2. Drop Notifications
DROP TABLE IF EXISTS notifications CASCADE;

-- 3. Drop Supplier B2B Marketplace Tables
DROP TABLE IF EXISTS supplier_orders CASCADE;

DROP TABLE IF EXISTS supplier_catalogs CASCADE;

DROP TABLE IF EXISTS supplier_regulars CASCADE;

DROP TABLE IF EXISTS supplier_profiles CASCADE;

-- 4. Drop POS Cashier System Tables
DROP TABLE IF EXISTS pos_transactions CASCADE;

DROP TABLE IF EXISTS pos_items CASCADE;

DROP TABLE IF EXISTS pos_categories CASCADE;

DROP TABLE IF EXISTS pos_staff_accounts CASCADE;

-- 5. Drop Stall Reviews & Lease Contracts
DROP TABLE IF EXISTS stall_reviews CASCADE;

DROP TABLE IF EXISTS lease_contracts CASCADE;

-- 6. Drop Stalls, Businesses & Business Types
DROP TABLE IF EXISTS stalls CASCADE;

DROP TABLE IF EXISTS businesses CASCADE;

DROP TABLE IF EXISTS business_types CASCADE;

-- 7. Drop Bank Accounts, Identities & Users
DROP TABLE IF EXISTS bank_accounts CASCADE;

DROP TABLE IF EXISTS user_identity_profiles CASCADE;

DROP TABLE IF EXISTS users CASCADE;

-- 8. Drop Generated Reports
DROP TABLE IF EXISTS generated_reports CASCADE;

-- 9. Drop Custom Enum Types
DROP TYPE IF EXISTS stall_permanence_type CASCADE;

DROP TYPE IF EXISTS stall_placement_type CASCADE;