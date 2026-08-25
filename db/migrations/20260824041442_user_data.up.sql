CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 1. SEED USERS
INSERT INTO users (
    id, name, email, password_hash, default_avatar_url, phone_numbers, role_profiles, active_role, subscription_plan
) VALUES 
(
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Rian Hidayat', 'rian.hidayat@example.com',
    '$2a$10$42w7Xp3M9Ie8S5rJ0vY8e.LqC6U9rT01v5Y1wW819aK1M9L8N7O6P',
    'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200&h=200&fit=crop',
    '[{"number": "+6281122334455", "is_primary": true, "roles": ["stall_owner", "supplier"]}]'::jsonb,
    '{"stall_owner": {"display_name": "Rian Property Group", "avatar_url": "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200&h=200&fit=crop"}}'::jsonb,
    'stall_owner', 'all_access'
),
(
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'Budi Santoso', 'budi.santoso@example.com',
    '$2a$10$42w7Xp3M9Ie8S5rJ0vY8e.LqC6U9rT01v5Y1wW819aK1M9L8N7O6P',
    'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=200&h=200&fit=crop',
    '[{"number": "+6281234567890", "is_primary": true, "roles": ["stall_owner"]}]'::jsonb,
    '{"stall_owner": {"display_name": "Budi Commercial Space", "avatar_url": "https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=200&h=200&fit=crop"}}'::jsonb,
    'stall_owner', 'single_role'
),
(
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 'Jakarta Event Management', 'info@jakartaevent.com',
    '$2a$10$42w7Xp3M9Ie8S5rJ0vY8e.LqC6U9rT01v5Y1wW819aK1M9L8N7O6P',
    'https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?w=200&h=200&fit=crop',
    '[{"number": "+6281399887766", "is_primary": true, "roles": ["stall_owner"]}]'::jsonb,
    '{"stall_owner": {"display_name": "Jakarta EO Official", "avatar_url": "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?w=200&h=200&fit=crop"}}'::jsonb,
    'stall_owner', 'all_access'
),
(
    'd0eebc99-9c0b-4ef8-bb6d-6bb9bd380a44', 'Siti Aminah', 'siti.aminah@example.com',
    '$2a$10$42w7Xp3M9Ie8S5rJ0vY8e.LqC6U9rT01v5Y1wW819aK1M9L8N7O6P',
    'https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=200&h=200&fit=crop',
    '[{"number": "+6285611223344", "is_primary": true, "roles": ["tenant"]}]'::jsonb,
    '{"tenant": {"display_name": "Kedai Kopi Siti", "avatar_url": "https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=200&h=200&fit=crop"}}'::jsonb,
    'tenant', 'free'
);

-- 2. SEED USER IDENTITY PROFILES (KYC)
INSERT INTO user_identity_profiles (id, user_id, full_name_ktp, nik, ktp_photo_url, domicile_city)
SELECT 
    gen_random_uuid(), id, name, '327301' || floor(random() * 8999999999 + 1000000000)::text,
    'https://images.unsplash.com/photo-1589829545856-d10d557cf95f?w=600&h=400&fit=crop', 'Bandung'
FROM users;

-- 3. SEED BANK ACCOUNTS (Untuk Payout Escrow)
INSERT INTO bank_accounts (id, user_id, bank_code, bank_name, account_number, account_holder_name, is_primary)
SELECT 
    gen_random_uuid(), id, 'BCA', 'Bank Central Asia',
    floor(random() * 8999999999 + 1000000000)::text, name, TRUE
FROM users;