CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ==============================================================================
-- 1. PERMANENT STALLS (15 Additional Data)
-- ==============================================================================

-- 1.1 Permanent: Ruko Boulevard Kelapa Gading
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruko Commercial Strip Boulevard Raya Kelapa Gading',
    'Ruko 3 lantai di jalur komersial utama Kelapa Gading. Visibilitas luar biasa, cocok untuk bank, showroom, klinik kecantikan, atau restoran.',
    'shophouse', 'permanent', 'indoor',
    72.00, 12.00, 6.00, 1, 6600,
    'Jl. Boulevard Raya Blok LA-06', 'Kelapa Gading Barat', 'Kelapa Gading', 'Jakarta Utara', 'DKI Jakarta', 'Indonesia', 'ID', '14240',
    -6.15810000, 106.90520000, 'https://maps.google.com/?q=-6.1581,106.9052', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "mall", "name": "Mall Kelapa Gading", "distanceKm": 0.4}]'::jsonb,
    '["month", "quarter", "year"]'::jsonb, 12000000.00, 34000000.00, 130000000.00, 10000000.00,
    6, '["1", "15"]'::jsonb, 'Listrik PLN & air PDAM sesuai meteran.',
    '["power", "high-power", "water", "drainage", "parking", "toilet", "security", "cctv"]'::jsonb,
    '["Renovasi fasad wajib dikonfirmasi terlebih dahulu."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 32, TRUE
);

-- 1.2 Permanent: Kios Garasi Dipatiukur Bandung
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Kios Garasi Usaha Dipatiukur UNPAD',
    'Alih fungsi garasi depan rumah menjadi kios usaha kuliner/kopi takeaway. Sangat dekat dengan kampus UNPAD Dipatiukur.',
    'garage-driveway', 'permanent', 'semi-outdoor',
    18.00, 6.00, 3.00, 1, 2200,
    'Jl. Dipatiukur No. 82', 'Lebakgede', 'Coblong', 'Bandung', 'Jawa Barat', 'Indonesia', 'ID', '40132',
    -6.89210000, 107.61810000, 'https://maps.google.com/?q=-6.8921,107.6181', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Padjadjaran (UNPAD)", "distanceKm": 0.2}]'::jsonb,
    '["month"]'::jsonb, 3200000.00, 1000000.00,
    2, '["1", "15"]'::jsonb, 'Listrik token mandiri.',
    '["power", "water", "wifi", "parking", "trash-area"]'::jsonb,
    '["Menjaga kebersihan selasar depan kios."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1521017432531-fbd92d768814?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 14, TRUE
);

-- 1.3 Permanent: Ruko Dinoyo Malang
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruko Komersial Soekarno Hatta Dinoyo Malang',
    'Ruko strategis di kawasan mahasiswa Soehat Malang. Sangat ramai lalu lintas, ideal untuk kafe, distro, atau fotokopi/percetakan.',
    'shophouse', 'permanent', 'indoor',
    40.00, 10.00, 4.00, 1, 3500,
    'Jl. Soekarno Hatta No. 45', 'Jatimulyo', 'Lowokwaru', 'Malang', 'Jawa Timur', 'Indonesia', 'ID', '65141',
    -7.94810000, 112.61820000, 'https://maps.google.com/?q=-7.9481,112.6182', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Brawijaya (UB)", "distanceKm": 0.6}]'::jsonb,
    '["month", "quarter", "year"]'::jsonb, 4500000.00, 13000000.00, 50000000.00, 3000000.00,
    3, '["1", "15"]'::jsonb, 'Biaya listrik & air dibayar terpisah.',
    '["power", "water", "drainage", "parking", "toilet", "security"]'::jsonb,
    '["Diwajibkan membuang sampah pada tempatnya."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1497366216548-37526070297c?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 19, TRUE
);

-- 1.4 Permanent: Kios Pasar Santa Jakarta
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Kios Retail Modern Pasar Santa Lantai 1',
    'Kios dalam area komunitas kreatif Pasar Santa. Sangat populer di kalangan anak muda untuk thrift shop, piringan hitam, atau kopi kurasi.',
    'market-kiosk', 'permanent', 'indoor',
    12.00, 4.00, 3.00, 1, 1300,
    'Jl. Cipaku I No. 1, Lt. 1 A-14', 'Petogogan', 'Kebayoran Baru', 'Jakarta Selatan', 'DKI Jakarta', 'Indonesia', 'ID', '12170',
    -6.23810000, 106.81210000, 'https://maps.google.com/?q=-6.2381,106.8121', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Kawasan SCBD", "distanceKm": 1.2}]'::jsonb,
    '["month"]'::jsonb, 2100000.00, 1000000.00,
    1, '["1"]'::jsonb, 'Listrik flat 150rb per bulan.',
    '["power", "wifi", "toilet", "security"]'::jsonb,
    '["Jam operasional kios bebas hingga jam 22:00 WIB."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 11, TRUE
);

-- 1.5 Permanent: Standalone Building Ubud Bali
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruang Usaha Art & Boutique Monkey Forest Ubud',
    'Bangunan unik nuansa tropis di jalan utama Monkey Forest Ubud. Ideal untuk galeri lukisan, studio yoga, toko perhiasan, atau boutique.',
    'standalone-building', 'permanent', 'indoor',
    65.00, 13.00, 5.00, 1, 5500,
    'Jl. Monkey Forest No. 42', 'Ubud', 'Ubud', 'Gianyar', 'Bali', 'Indonesia', 'ID', '80571',
    -8.51210000, 115.26210000, 'https://maps.google.com/?q=-8.5121,115.2621', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Sacred Monkey Forest Sanctuary", "distanceKm": 0.3}]'::jsonb,
    '["month", "year"]'::jsonb, 11000000.00, 120000000.00, 10000000.00,
    6, '["1"]'::jsonb, 'Listrik & air sumur bor lancar.',
    '["power", "water", "air-conditioner", "wifi", "parking", "toilet"]'::jsonb,
    '["Wajib merawat keasrian taman kayu di halaman."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    5.0, 29, TRUE
);

-- 1.6 Permanent: Kios Pinggir Jalan Pettarani Makassar
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Street Kiosk Utama AP Pettarani',
    'Kios kontainer tepi jalan tol melayang AP Pettarani. Lokasi lalu lintas sibuk, cocok untuk usaha pulsa/service HP atau snack kencang.',
    'street-kiosk', 'permanent', 'outdoor',
    10.00, 5.00, 2.00, 1, 1300,
    'Jl. AP Pettarani No. 88', 'Buakana', 'Rappocini', 'Makassar', 'Sulawesi Selatan', 'Indonesia', 'ID', '90222',
    -5.14810000, 119.43210000, 'https://maps.google.com/?q=-5.1481,119.4321', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Kantor SULSEL", "distanceKm": 0.5}]'::jsonb,
    '["month"]'::jsonb, 1800000.00, 800000.00,
    1, '["1", "15"]'::jsonb, 'Listrik PLN token.',
    '["power", "trash-area", "security"]'::jsonb,
    '["Akses penerangan malam siap pakai."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1578916171728-46686eac8d58?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.5, 8, TRUE
);

-- 1.7 Permanent: Ruko Merr Surabaya
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruko Komersial Koridor MERR Kalijudan Surabaya',
    'Ruko baru modern di jalur lingkar dalam timur MERR Surabaya. Lahan parkir khusus pengunjung luas.',
    'shophouse', 'permanent', 'indoor',
    54.00, 12.00, 4.50, 1, 4400,
    'Jl. Dr. Ir. H. Soekarno No. 120', 'Kalijudan', 'Mulyorejo', 'Surabaya', 'Jawa Timur', 'Indonesia', 'ID', '60114',
    -7.25810000, 112.78210000, 'https://maps.google.com/?q=-7.2581,112.7821', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Airlangga (UNAIR) Kampus C", "distanceKm": 1.0}]'::jsonb,
    '["month", "quarter", "year"]'::jsonb, 6000000.00, 17000000.00, 65000000.00, 4000000.00,
    3, '["1", "15"]'::jsonb, 'Listrik PLN 4400 VA & PDAM.',
    '["power", "high-power", "water", "drainage", "parking", "toilet", "security"]'::jsonb,
    '["Area parkir digunakan bersama tenant lain secara tertib."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 23, TRUE
);

-- 1.8 Permanent: Garasi Komersial Jakal KM 8 Sleman
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Garasi Usaha Kuliner Jalan Kaliurang KM 8.5',
    'Space garasi komersial hook tepi Jalan Kaliurang. Dekat kawasan hunian mahasiswa UII dan UGM.',
    'garage-driveway', 'permanent', 'semi-outdoor',
    22.00, 5.50, 4.00, 1, 2200,
    'Jl. Kaliurang KM 8.5 No. 14', 'Sinduharjo', 'Ngaglik', 'Sleman', 'DI Yogyakarta', 'Indonesia', 'ID', '55581',
    -7.73210000, 110.39210000, 'https://maps.google.com/?q=-7.7321,110.3921', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Islam Indonesia (UII)", "distanceKm": 4.0}]'::jsonb,
    '["month"]'::jsonb, 2200000.00, 1000000.00,
    1, '["1"]'::jsonb, 'Listrik sub-meter.',
    '["power", "water", "wifi", "parking"]'::jsonb,
    '["Boleh menambah dekorasi kanopi portabel."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1521017432531-fbd92d768814?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 16, TRUE
);

-- 1.9 Permanent: Ruko Gatot Subroto Denpasar
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruko Komersial Jalan Gatot Subroto Tengah Denpasar',
    'Ruko 2 lantai pusat bisnis Gatsu Denpasar. Lahan komersial ramai untuk distributor, toko sparepart, atau kantor cabang.',
    'shophouse', 'permanent', 'indoor',
    48.00, 12.00, 4.00, 1, 3500,
    'Jl. Gatot Subroto Tengah No. 210', 'Tonja', 'Denpasar Utara', 'Denpasar', 'Bali', 'Indonesia', 'ID', '80235',
    -8.63810000, 115.22810000, 'https://maps.google.com/?q=-8.6381,115.2281', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Pusat Bisnis Gatsu", "distanceKm": 0.2}]'::jsonb,
    '["month", "quarter", "year"]'::jsonb, 5500000.00, 15500000.00, 60000000.00, 3500000.00,
    3, '["1", "15"]'::jsonb, 'Listrik PLN & air sumur bor.',
    '["power", "water", "parking", "toilet", "security"]'::jsonb,
    '["Parkir bersama tertib tidak menghalangi ruko sebelah."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 12, TRUE
);

-- 1.10 Permanent: Kios Pasar Tradisional Pasar Minggu
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Kios Sembako Pasar Minggu Jakarta Selatan',
    'Kios permanen area basah/kering Pasar Minggu. Dekat dengan akses komuter Stasiun Pasar Minggu.',
    'market-kiosk', 'permanent', 'indoor',
    15.00, 5.00, 3.00, 1, 1300,
    'Jl. Pasar Minggu Raya, Blok B-12', 'Pasar Minggu', 'Pasar Minggu', 'Jakarta Selatan', 'DKI Jakarta', 'Indonesia', 'ID', '12510',
    -6.28420000, 106.84510000, 'https://maps.google.com/?q=-6.2842,106.8451', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun Pasar Minggu", "distanceKm": 0.2}]'::jsonb,
    '["month"]'::jsonb, 1700000.00, 500000.00,
    1, '["1"]'::jsonb, 'Biaya air & kebersihan dikelola PD Pasar Jaya.',
    '["power", "water", "trash-area", "security"]'::jsonb,
    '["Mengikuti aturan operasional pasar."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.5, 10, TRUE
);

-- 1.11 Permanent: Ruko BSB City Semarang
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruko Komersial BSB City Mijen Semarang',
    'Ruko baru di kawasan mandiri BSB City Semarang. Lingkungan rapi dekat perumahan elite dan kampus UNIKA BSB.',
    'shophouse', 'permanent', 'indoor',
    45.00, 9.00, 5.00, 1, 3500,
    'Kawasan Commercial BSB City Blok A-02', 'Pesantren', 'Mijen', 'Semarang', 'Jawa Tengah', 'Indonesia', 'ID', '50212',
    -7.03810000, 110.33210000, 'https://maps.google.com/?q=-7.0381,110.3321', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Unika Soegijapranata BSB", "distanceKm": 0.8}]'::jsonb,
    '["month", "quarter", "year"]'::jsonb, 3800000.00, 10500000.00, 40000000.00, 2500000.00,
    3, '["1", "15"]'::jsonb, 'Listrik & air ditanggung tenant.',
    '["power", "water", "parking", "toilet", "security"]'::jsonb,
    '["Jagalah ketertiban area kawasan perumahan BSB."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1497366216548-37526070297c?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 15, TRUE
);

-- 1.12 Permanent: Kios Garasi Serpong BSD
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Kios Garasi Usaha Komersial BSD Sektor 1',
    'Garasi rumah sudut komersial di BSD Sektor 1. Akses mudah dari jalan raya, cocok untuk barbershop atau minimarket modal kecil.',
    'garage-driveway', 'permanent', 'indoor',
    25.00, 5.00, 5.00, 1, 2200,
    'Jl. Griya Loka Raya Sektor 1.2', 'Rawa Buntu', 'Serpong', 'Tangerang Selatan', 'Banten', 'Indonesia', 'ID', '15318',
    -6.30120000, 106.67810000, 'https://maps.google.com/?q=-6.3012,106.6781', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun Rawa Buntu", "distanceKm": 0.5}]'::jsonb,
    '["month"]'::jsonb, 2500000.00, 1000000.00,
    1, '["1"]'::jsonb, 'Listrik kWh token.',
    '["power", "water", "wifi", "parking"]'::jsonb,
    '["Bisa dipasang papan nama ukuran sedang."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1521017432531-fbd92d768814?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 13, TRUE
);

-- 1.13 Permanent: Ruko Ahmad Yani Banjarmasin
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruko Komersial Jalan Ahmad Yani KM 3 Banjarmasin',
    'Ruko komersial di arteri utama Kota Banjarmasin. Akses mudah dari pusat kota maupun luar daerah.',
    'shophouse', 'permanent', 'indoor',
    50.00, 10.00, 5.00, 1, 3500,
    'Jl. Ahmad Yani KM 3.5 No. 45', 'Kuripan', 'Banjarmasin Timur', 'Banjarmasin', 'Kalimantan Selatan', 'Indonesia', 'ID', '70234',
    -3.32810000, 114.60810000, 'https://maps.google.com/?q=-3.3281,114.6081', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Duta Mall Banjarmasin", "distanceKm": 1.0}]'::jsonb,
    '["month", "year"]'::jsonb, 4500000.00, 50000000.00, 3000000.00,
    3, '["1", "15"]'::jsonb, 'Listrik & air ditanggung penyewa.',
    '["power", "water", "parking", "toilet", "security"]'::jsonb,
    '["Aturan sewa sesuai kesepakatan tertulis."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 11, TRUE
);

-- 1.14 Permanent: Standalone Space Margonda Depok
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Bangunan Komersial Mandiri Margonda Raya',
    'Bangunan komersial independen 1 lantai halaman luas di Margonda Raya Depok. Cocok untuk kafe outdoor, resto cepat saji, atau bengkel modifikasi.',
    'standalone-building', 'permanent', 'indoor',
    80.00, 16.00, 5.00, 1, 5500,
    'Jl. Margonda Raya No. 350', 'Kemirimuka', 'Beji', 'Depok', 'Jawa Barat', 'Indonesia', 'ID', '16423',
    -6.37810000, 106.83210000, 'https://maps.google.com/?q=-6.3781,106.8321', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Indonesia (UI)", "distanceKm": 0.8}]'::jsonb,
    '["month", "quarter"]'::jsonb, 8000000.00, 22000000.00, 5000000.00,
    3, '["1", "15"]'::jsonb, 'Listrik PLN pascabayar.',
    '["power", "high-power", "water", "drainage", "parking", "toilet", "security"]'::jsonb,
    '["Izin penggunaan area luar dikordinasikan dahulu."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 20, TRUE
);

-- 1.15 Permanent: Kios Container Cihampelas Bandung
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Container Kiosk Cihampelas Walk Walkway',
    'Kios kontainer baja modern di pinggir Jalan Cihampelas Bandung. Lalu lintas pejalan kaki padat wisatawan belanja.',
    'street-kiosk', 'permanent', 'semi-outdoor',
    12.00, 4.00, 3.00, 1, 1300,
    'Jl. Cihampelas No. 160', 'Cipaganti', 'Coblong', 'Bandung', 'Jawa Barat', 'Indonesia', 'ID', '40131',
    -6.89420000, 107.60410000, 'https://maps.google.com/?q=-6.8942,107.6041', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "mall", "name": "Ciwalk Mall", "distanceKm": 0.1}]'::jsonb,
    '["month"]'::jsonb, 2500000.00, 1000000.00,
    1, '["1", "15"]'::jsonb, 'Listrik flat 200rb/bulan.',
    '["power", "trash-area", "security"]'::jsonb,
    '["Menjaga kerapihan etalase kounter."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1578916171728-46686eac8d58?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 17, TRUE
);


-- ==============================================================================
-- 2. SEMI-PERMANENT STALLS (15 Additional Data)
-- ==============================================================================

-- 2.1 Semi-Permanent: Island Mall Pondok Indah Mall 2
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Island Kiosk Premium Level 2 PIM 2',
    'Spot island premium di lantai 2 Pondok Indah Mall 2 Jakarta Selatan. Target pasar pengunjung segmen menengah ke atas.',
    'mall-island', 'semi-permanent', 'indoor',
    3500, 'Pondok Indah Mall 2', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. Metro Pondok Indah, Lt. 2', 'Pondok Pinang', 'Kebayoran Lama', 'Jakarta Selatan', 'DKI Jakarta', 'Indonesia', 'ID', '12310',
    -6.26510000, 106.78410000, 'https://maps.google.com/?q=-6.2651,106.7841', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Pondok Indah Office Tower", "distanceKm": 0.2}]'::jsonb,
    '["month", "quarter"]'::jsonb, 9000000.00, 25500000.00, 5000000.00,
    3, '["1", "15"]'::jsonb, 'Termasuk service charge mall.',
    '["power", "air-conditioner", "wifi", "toilet", "cleaning-service", "security"]'::jsonb,
    '["Tampilan island booth wajib elegan sesuai standar PIM."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    5.0, 38, TRUE
);

-- 2.2 Semi-Permanent: Counter Foodcourt Mall Ciputra Semarang
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Counter Kuliner Foodcourt Mall Ciputra Simpang Lima',
    'Counter siap pakai di pusat kuliner Foodcourt Mall Ciputra Semarang. Lalu lintas pengunjung padat di pusat Simpang Lima.',
    'food-court-counter', 'semi-permanent', 'indoor',
    3500, 'Mall Ciputra Semarang', '{"opening_time": "10:00", "closing_time": "21:30", "is_24_hours": false}'::jsonb,
    'Jl. Simpang Lima No. 1, Lt. 2', 'Pekunden', 'Semarang Tengah', 'Semarang', 'Jawa Tengah', 'Indonesia', 'ID', '50134',
    -6.99010000, 110.42310000, 'https://maps.google.com/?q=-6.9901,110.4231', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Lapangan Simpang Lima", "distanceKm": 0.0}]'::jsonb,
    '["month"]'::jsonb, 4000000.00, 1500000.00,
    1, '["1", "15"]'::jsonb, 'Bagi hasil utilitas dikelola pihak mall.',
    '["power", "water", "air-conditioner", "wifi", "toilet", "cleaning-service"]'::jsonb,
    '["Menjaga kebersihan area konter masing-masing."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 18, TRUE
);

-- 2.3 Semi-Permanent: Kios Stasiun KRL Tebet
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Kios Komersial Area Gate Stasiun KRL Tebet',
    'Spot usaha di dalam gate stasiun KRL Tebet. Arus ribuan penumpang komuter Jabodetabek jam berangkat dan pulang kerja.',
    'transit-kiosk', 'semi-permanent', 'indoor',
    2200, 'Stasiun KRL Tebet', '{"opening_time": "05:00", "closing_time": "23:00", "is_24_hours": false}'::jsonb,
    'Stasiun KRL Tebet, Hall Utama B-03', 'Tebet Timur', 'Tebet', 'Jakarta Selatan', 'DKI Jakarta', 'Indonesia', 'ID', '12820',
    -6.22610000, 106.85810000, 'https://maps.google.com/?q=-6.2261,106.8581', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun KRL Tebet", "distanceKm": 0.0}]'::jsonb,
    '["month"]'::jsonb, 4500000.00, 2000000.00,
    1, '["1"]'::jsonb, 'Biaya listrik & kebersihan stasiun included.',
    '["power", "air-conditioner", "security", "cctv"]'::jsonb,
    '["Hanya menjual makanan/minuman kemasan rapi."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 26, TRUE
);

-- 2.4 Semi-Permanent: Island Mall Supermal Karawaci
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Island Kiosk Ground Floor Supermal Karawaci',
    'Booth island di area Ground Floor Supermal Karawaci dekat Timezone. Lalu lintas ramai anak muda dan keluarga.',
    'mall-island', 'semi-permanent', 'indoor',
    2200, 'Supermal Karawaci', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. Boulevard Diponegoro No. 105, GF-12', 'Bencongan', 'Kelapa Dua', 'Tangerang', 'Banten', 'Indonesia', 'ID', '15810',
    -6.22810000, 106.60810000, 'https://maps.google.com/?q=-6.2281,106.6081', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Pelita Harapan (UPH)", "distanceKm": 0.5}]'::jsonb,
    '["month", "quarter"]'::jsonb, 4200000.00, 11800000.00, 2000000.00,
    1, '["1", "15"]'::jsonb, 'Listrik token isi mandiri.',
    '["power", "air-conditioner", "wifi", "toilet", "security"]'::jsonb,
    '["Mendorong penjualan dengan pelayanan ramah."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 17, TRUE
);

-- 2.5 Semi-Permanent: Stand Kantin Kampus ITB Ganesha
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Stand Kantin Barat Kampus ITB Ganesha Bandung',
    'Kios makanan di area Kantin Barat ITB Ganesha. Konsumen pasti ribuan mahasiswa, dosen, dan karyawan kampus.',
    'campus-canteen-counter', 'semi-permanent', 'indoor',
    2200, 'Kantin Barat ITB', '{"opening_time": "07:00", "closing_time": "18:00", "is_24_hours": false}'::jsonb,
    'Jl. Ganesha No. 10, Kantin Barat Stand 05', 'Lebak Siliwangi', 'Coblong', 'Bandung', 'Jawa Barat', 'Indonesia', 'ID', '40132',
    -6.89120000, 107.61010000, 'https://maps.google.com/?q=-6.8912,107.6101', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Institut Teknologi Bandung", "distanceKm": 0.0}]'::jsonb,
    '["month"]'::jsonb, 2800000.00, 1000000.00,
    1, '["1"]'::jsonb, 'Listrik & air gratis dari koperasi kampus.',
    '["power", "water", "wifi", "toilet", "cleaning-service"]'::jsonb,
    '["Menjaga harga terjangkau bagi mahasiswa."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 31, TRUE
);

-- 2.6 Semi-Permanent: Island Mall Mal Bali Galeria Denpasar
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Island Booth GF Mal Bali Galeria Kuta',
    'Kios island strategis di atrium Mal Bali Galeria Simpang Siur Kuta. Ramai wisatawan dan keluarga Bali.',
    'mall-island', 'semi-permanent', 'indoor',
    2200, 'Mal Bali Galeria', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. By Pass Ngurah Rai, GF Koridor Utama', 'Kuta', 'Kuta', 'Badung', 'Bali', 'Indonesia', 'ID', '80361',
    -8.71810000, 115.18210000, 'https://maps.google.com/?q=-8.7181,115.1821', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Patung Simpang Dewa Ruci", "distanceKm": 0.1}]'::jsonb,
    '["month", "quarter"]'::jsonb, 4800000.00, 13500000.00, 2000000.00,
    1, '["1", "15"]'::jsonb, 'Biaya utilitas listrik teratur bulanan.',
    '["power", "air-conditioner", "wifi", "toilet", "security"]'::jsonb,
    '["Penyewa menjaga kebersihan island booth."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 20, TRUE
);

-- 2.7 Semi-Permanent: Counter Foodcourt Plaza Surabaya
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Counter F&B Foodcourt Delta Plaza Surabaya',
    'Counter kuliner siap pakai di Foodcourt Delta Plaza Surabaya. Pengunjung keluarga dan pelajar stabil.',
    'food-court-counter', 'semi-permanent', 'indoor',
    3500, 'Plaza Surabaya (Delta)', '{"opening_time": "10:00", "closing_time": "21:30", "is_24_hours": false}'::jsonb,
    'Jl. Pemuda No. 33-37, Lt. 3', 'Embong Kaliasin', 'Genteng', 'Surabaya', 'Jawa Timur', 'Indonesia', 'ID', '60271',
    -7.26510000, 112.74820000, 'https://maps.google.com/?q=-7.2651,112.7482', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun Surabaya Gubeng", "distanceKm": 0.4}]'::jsonb,
    '["month"]'::jsonb, 3500000.00, 1500000.00,
    1, '["1"]'::jsonb, 'Pelayanan cuci piring ditangani pengelola.',
    '["power", "water", "air-conditioner", "wifi", "toilet", "cleaning-service"]'::jsonb,
    '["Menerapkan standar higienitas F&B."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 15, TRUE
);

-- 2.8 Semi-Permanent: Rest Area KM 260B Brebes
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Kios UMKM Rest Area Heritage KM 260B Banjaratma',
    'Kios komersial unik di Rest Area Pabrik Gula Banjaratma Tol Pejagan-Pemalang. Ikonik dan sangat diminati pemudik / pelancong.',
    'rest-area-kiosk', 'semi-permanent', 'indoor',
    2200, 'Rest Area KM 260B Banjaratma', '{"opening_time": "00:00", "closing_time": "23:59", "is_24_hours": true}'::jsonb,
    'Rest Area KM 260B Tol Pejagan-Pemalang, Stand A-15', 'Banjaratma', 'Bulakamba', 'Brebes', 'Jawa Tengah', 'Indonesia', 'ID', '52253',
    -6.88210000, 108.96210000, 'https://maps.google.com/?q=-6.8821,108.9621', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Bangunan Cagar Budaya Pabrik Gula", "distanceKm": 0.0}]'::jsonb,
    '["month"]'::jsonb, 3000000.00, 1000000.00,
    1, '["1"]'::jsonb, 'Listrik & air disiapkan pengelola rest area.',
    '["power", "water", "parking", "toilet", "security"]'::jsonb,
    '["Buka 24 jam mendukung kelancaran pengguna tol."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1556742049-0cfed4f6a45d?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 35, TRUE
);

-- 2.9 Semi-Permanent: Island Mall Palembang Icon
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Island Kiosk Level 1 Palembang Icon Mall',
    'Booth island modern di lantai 1 Palembang Icon Mall dekat danau. Mall paling hits di Palembang.',
    'mall-island', 'semi-permanent', 'indoor',
    2200, 'Palembang Icon Mall', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. POM IX, Lt. 1 Kiosk 08', 'Lorok Pakjo', 'Ilir Barat I', 'Palembang', 'Sumatera Selatan', 'Indonesia', 'ID', '30137',
    -2.98120000, 104.75210000, 'https://maps.google.com/?q=-2.9812,104.7521', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Stadion Bumi Sriwijaya", "distanceKm": 0.2}]'::jsonb,
    '["month", "quarter"]'::jsonb, 4000000.00, 11000000.00, 2000000.00,
    1, '["1", "15"]'::jsonb, 'Listrik token mandiri.',
    '["power", "air-conditioner", "wifi", "toilet", "security"]'::jsonb,
    '["Mengikuti standar operasional mall."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 14, TRUE
);

-- 2.10 Semi-Permanent: Counter Foodcourt MOG Malang
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Counter F&B Foodcourt Mall Olympic Garden Malang',
    'Counter kuliner di area Foodcourt MOG Malang. Dekat Stadion Gajayana, ramai pengunjung anak muda dan keluarga.',
    'food-court-counter', 'semi-permanent', 'indoor',
    3500, 'Mall Olympic Garden (MOG)', '{"opening_time": "10:00", "closing_time": "21:30", "is_24_hours": false}'::jsonb,
    'Jl. Kawi No. 24, Lt. 3 FC-06', 'Kauman', 'Klojen', 'Malang', 'Jawa Timur', 'Indonesia', 'ID', '65119',
    -7.97810000, 112.62510000, 'https://maps.google.com/?q=-7.9781,112.6251', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Stadion Gajayana Malang", "distanceKm": 0.1}]'::jsonb,
    '["month"]'::jsonb, 3200000.00, 1000000.00,
    1, '["1"]'::jsonb, 'Utilitas & kebersihan meja dikelola mall.',
    '["power", "water", "air-conditioner", "wifi", "toilet", "cleaning-service"]'::jsonb,
    '["Menjaga kerapihan peralatan masak."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 13, TRUE
);

-- 2.11 Semi-Permanent: Island Mall Centre Point Medan
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Island Booth GF Mall Centre Point Medan',
    'Lapak island di Ground Floor Mall Centre Point Medan dekat pintu masuk stasiun kereta bandara.',
    'mall-island', 'semi-permanent', 'indoor',
    2200, 'Centre Point Mall Medan', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. Jawa No. 8, GF Kiosk 12', 'Ganggawa', 'Medan Timur', 'Medan', 'Sumatera Utara', 'Indonesia', 'ID', '20136',
    -3.59120000, 98.68120000, 'https://maps.google.com/?q=-3.5912,98.6812', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun Kereta Bandara Medan", "distanceKm": 0.1}]'::jsonb,
    '["month", "quarter"]'::jsonb, 5500000.00, 15500000.00, 2500000.00,
    1, '["1", "15"]'::jsonb, 'Listrik PLN token mandiri.',
    '["power", "air-conditioner", "wifi", "toilet", "security"]'::jsonb,
    '["Kios harus selalu rapi dan siap melayani."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 22, TRUE
);

-- 2.12 Semi-Permanent: Kios Stasiun MRT Blok M
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Kios Komersial Concourse MRT Blok M BCA',
    'Spot usaha di area concourse stasiun MRT Blok M. Terhubung langsung ke Blok M Plaza.',
    'transit-kiosk', 'semi-permanent', 'indoor',
    3500, 'Stasiun MRT Blok M BCA', '{"opening_time": "06:00", "closing_time": "23:00", "is_24_hours": false}'::jsonb,
    'Stasiun MRT Blok M Area Concourse A-01', 'Kramat Pela', 'Kebayoran Baru', 'Jakarta Selatan', 'DKI Jakarta', 'Indonesia', 'ID', '12130',
    -6.24420000, 106.79810000, 'https://maps.google.com/?q=-6.2442,106.7981', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun MRT Blok M", "distanceKm": 0.0}]'::jsonb,
    '["month"]'::jsonb, 5500000.00, 2500000.00,
    2, '["1"]'::jsonb, 'Include utilitas & fasilitas stasiun.',
    '["power", "air-conditioner", "security", "cctv"]'::jsonb,
    '["Dilarang menggunakan bahan berbau menyengat."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 31, TRUE
);

-- 2.13 Semi-Permanent: Counter Foodcourt Botani Square Bogor
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Counter F&B Foodcourt Botani Square Mall Bogor',
    'Counter kuliner siap pakai di Foodcourt Botani Square Bogor dekat Tugu Kujang. Sangat ramai pelajar & mahasiswa IPB.',
    'food-court-counter', 'semi-permanent', 'indoor',
    3500, 'Botani Square Mall', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. Pajajaran No. 69, Lt. 2 FC-10', 'Tegallega', 'Bogor Tengah', 'Bogor', 'Jawa Barat', 'Indonesia', 'ID', '16127',
    -6.60120000, 106.80810000, 'https://maps.google.com/?q=-6.6012,106.8081', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "IPB Kampus Baranangsiang", "distanceKm": 0.2}]'::jsonb,
    '["month", "quarter"]'::jsonb, 4200000.00, 11800000.00, 2000000.00,
    1, '["1", "15"]'::jsonb, 'Pembersihan meja oleh pengelola mall.',
    '["power", "water", "air-conditioner", "wifi", "toilet", "cleaning-service"]'::jsonb,
    '["Wajib memakai celemek dan sarung tangan saat memasak."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 19, TRUE
);

-- 2.14 Semi-Permanent: Island Mall Cihampelas Walk Bandung
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Island Booth Open Area Ciwalk Bandung',
    'Booth island terbuka di area hijau Ciwalk Bandung. Suasana sejuk dan nyaman bagi wisatawan belanja.',
    'mall-island', 'semi-permanent', 'semi-outdoor',
    2200, 'Cihampelas Walk (Ciwalk)', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. Cihampelas No. 160, Area Broadway', 'Cipaganti', 'Coblong', 'Bandung', 'Jawa Barat', 'Indonesia', 'ID', '40131',
    -6.89510000, 107.60420000, 'https://maps.google.com/?q=-6.8951,107.6042', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Kawasan Wisata Belanja Cihampelas", "distanceKm": 0.0}]'::jsonb,
    '["month"]'::jsonb, 3800000.00, 1500000.00,
    1, '["1", "15"]'::jsonb, 'Listrik PLN token mandiri.',
    '["power", "wifi", "toilet", "security"]'::jsonb,
    '["Booth ditutup terpal rapi saat mall tutup."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 16, TRUE
);

-- 2.15 Semi-Permanent: Counter Foodcourt Solo Paragon
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, parent_complex_name, operating_hours,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Counter Foodcourt Foodtown Solo Paragon Mall',
    'Counter kuliner di Foodtown Solo Paragon Lifestyle Mall. Mall teramai di kota Surakarta.',
    'food-court-counter', 'semi-permanent', 'indoor',
    3500, 'Solo Paragon Lifestyle Mall', '{"opening_time": "10:00", "closing_time": "21:30", "is_24_hours": false}'::jsonb,
    'Jl. Yosodipuro No. 133, Lt. 1', 'Mangkubumen', 'Banjarsari', 'Surakarta', 'Jawa Tengah', 'Indonesia', 'ID', '57139',
    -7.56120000, 110.81210000, 'https://maps.google.com/?q=-7.5612,110.8121', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Stasiun Solo Balapan", "distanceKm": 1.2}]'::jsonb,
    '["month"]'::jsonb, 3000000.00, 1000000.00,
    1, '["1"]'::jsonb, 'Fasilitas utilitas siap pakai.',
    '["power", "water", "air-conditioner", "wifi", "toilet", "cleaning-service"]'::jsonb,
    '["Menjaga kebersihan area saji."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 11, TRUE
);


-- ==============================================================================
-- 3. TEMPORARY STALLS (15 Additional Data)
-- ==============================================================================

-- 3.1 Temporary: Pop-Up Booth JakCloth Jakarta
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Festival Clothing JakCloth Senayan',
    'Booth distro & fashion apparel pada event clothing terbesar JakCloth di Parkir Timur Senayan.',
    'bazaar-booth', 'temporary', 'outdoor',
    1300,
    '{"event_name": "JakCloth Year End Sale 2026", "start_date": "2026-12-23", "end_date": "2026-12-28", "registration_deadline": "2026-12-10"}'::jsonb,
    '{"total_slots": 100, "available_slots": 22}'::jsonb,
    'GBK Parkir Timur Senayan, Area Distro C-12', 'Gelora', 'Tanah Abang', 'Jakarta Pusat', 'DKI Jakarta', 'Indonesia', 'ID', '10270',
    -6.21820000, 106.80210000, 'https://maps.google.com/?q=-6.2182,106.8021', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Kawasan SCBD", "distanceKm": 0.8}]'::jsonb,
    '["day"]'::jsonb, 450000.00, 400000.00,
    6, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'non_refundable',
    'Listrik 1300 VA & tenda sarnafil 3x3m.',
    '["power", "seating", "parking", "trash-area", "security"]'::jsonb,
    '["Barang jualan dikirim saat loading day H-1."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565123409695-7b5ef63a2efb?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 25, TRUE
);

-- 3.2 Temporary: Food Truck Park Bintaro Xchange
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, monthly_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Outdoor Food Truck Park BXC Mall Parkir Barat',
    'Spot parkir food truck di halaman outdoor Bintaro Jaya Xchange Mall. Ramai keluarga & anak muda saat akhir pekan.',
    'food-truck-spot', 'temporary', 'outdoor',
    2200,
    '{"event_name": "BXC Outdoor Food Fest", "start_date": "2026-11-01", "end_date": "2026-12-31", "registration_deadline": "2026-10-20"}'::jsonb,
    '{"total_slots": 10, "available_slots": 3}'::jsonb,
    'Bintaro Jaya Xchange Mall Parkir Barat FT-04', 'Pondok Jaya', 'Pondok Aren', 'Tangerang Selatan', 'Banten', 'Indonesia', 'ID', '15220',
    -6.28420000, 106.72810000, 'https://maps.google.com/?q=-6.2842,106.7281', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun KRL Jurangmangu", "distanceKm": 0.1}]'::jsonb,
    '["day", "month"]'::jsonb, 250000.00, 3200000.00, 500000.00,
    7, '["event_day_1", "event_week_1"]'::jsonb,
    'everyday', 'flexible_days', 'deposit_refundable',
    'Daya listrik colokan 2200 VA.',
    '["power", "wifi", "parking", "trash-area", "security"]'::jsonb,
    '["Armada food truck wajib laik jalan."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1526367790999-0150786686a2?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 19, TRUE
);

-- 3.3 Temporary: Pop-Up Booth Surabaya Shopping Festival
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Bazaar Atrium Grand City Surabaya',
    'Booth pameran bazaar pada event Surabaya Shopping Festival di Atrium Utama Grand City.',
    'bazaar-booth', 'temporary', 'indoor',
    1300,
    '{"event_name": "Surabaya Shopping Fest 2026", "start_date": "2026-11-10", "end_date": "2026-11-15", "registration_deadline": "2026-10-30"}'::jsonb,
    '{"total_slots": 40, "available_slots": 9}'::jsonb,
    'Grand City Surabaya, Atrium Utama Stand 18', 'Ketabang', 'Genteng', 'Surabaya', 'Jawa Timur', 'Indonesia', 'ID', '60272',
    -7.26120000, 112.74810000, 'https://maps.google.com/?q=-7.2612,112.7481', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun Surabaya Gubeng", "distanceKm": 0.5}]'::jsonb,
    '["day"]'::jsonb, 350000.00, 300000.00,
    6, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'non_refundable',
    'Listrik 1300 VA & partisi karpet terpasang.',
    '["power", "air-conditioner", "wifi", "parking", "security"]'::jsonb,
    '["Diwajibkan mematuhi jam buka pameran."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 14, TRUE
);

-- 3.4 Temporary: Pop-Up Stand CFD Slamet Riyadi Solo
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Lapak Minggu CFD Jalan Slamet Riyadi Solo',
    'Lapak mingguan Car Free Day Jalan Slamet Riyadi Solo terpanjang di Indonesia. Pengunjung ribuan masyarakat berolahraga.',
    'pop-up-booth', 'temporary', 'outdoor',
    900,
    '{"event_name": "Sunday CFD Slamet Riyadi Fest", "start_date": "2026-10-04", "end_date": "2026-12-27", "registration_deadline": "2026-09-28"}'::jsonb,
    '{"total_slots": 60, "available_slots": 18}'::jsonb,
    'Jl. Slamet Riyadi (Depan Sriwedari) Plot 24', 'Sriwedari', 'Laweyan', 'Surakarta', 'Jawa Tengah', 'Indonesia', 'ID', '57141',
    -7.56820000, 110.81210000, 'https://maps.google.com/?q=-7.5682,110.8121', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Taman Sriwedari Solo", "distanceKm": 0.1}]'::jsonb,
    '["day"]'::jsonb, 100000.00, 50000.00,
    1, '["event_day_1"]'::jsonb,
    'weekends', 'flexible_days', 'deposit_refundable',
    'Pengelolaan sampah dilakukan panitia CFD.',
    '["trash-area", "security"]'::jsonb,
    '["Bongkar tenda maksimal pukul 09:00 WIB."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 12, TRUE
);

-- 3.5 Temporary: Food Truck Court Kampus IPB Dramaga
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, monthly_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Food Truck Park Kampus IPB Dramaga Bogor',
    'Spot outdoor khusus parkir food truck di area Lapangan Gymnasium Kampus IPB Dramaga Bogor.',
    'food-truck-spot', 'temporary', 'outdoor',
    2200,
    '{"event_name": "IPB Student Food Carnival", "start_date": "2026-10-10", "end_date": "2026-11-10", "registration_deadline": "2026-09-30"}'::jsonb,
    '{"total_slots": 15, "available_slots": 4}'::jsonb,
    'Jl. Raya Dramaga Kampus IPB, Area Gym FT-03', 'Dramaga', 'Dramaga', 'Bogor', 'Jawa Barat', 'Indonesia', 'ID', '16680',
    -6.55820000, 106.72810000, 'https://maps.google.com/?q=-6.5582,106.7281', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "IPB University Dramaga", "distanceKm": 0.0}]'::jsonb,
    '["day", "month"]'::jsonb, 150000.00, 2200000.00, 300000.00,
    5, '["event_day_1"]'::jsonb,
    'everyday', 'flexible_days', 'deposit_refundable',
    'Tersedia listrik colokan 2200 VA.',
    '["power", "wifi", "parking", "trash-area", "security"]'::jsonb,
    '["Dilarang menjual minuman beralkohol."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1526367790999-0150786686a2?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 16, TRUE
);

-- 3.6 Temporary: Pop-Up Booth Bali Arts Festival Denpasar
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Bazaar Pesta Kesenian Bali Taman Budaya Art Center',
    'Booth bazaar temporary pada ajang Pesta Kesenian Bali (PKB) di Art Center Denpasar. Ribuan pengunjung lokal & wisatawan.',
    'bazaar-booth', 'temporary', 'semi-outdoor',
    1300,
    '{"event_name": "Pesta Kesenian Bali 2026", "start_date": "2026-11-15", "end_date": "2026-12-15", "registration_deadline": "2026-11-01"}'::jsonb,
    '{"total_slots": 50, "available_slots": 10}'::jsonb,
    'Jl. Nusa Indah, Taman Budaya Art Center Stand 12', 'Sumerta Kelod', 'Denpasar Timur', 'Denpasar', 'Bali', 'Indonesia', 'ID', '80239',
    -8.65820000, 115.23810000, 'https://maps.google.com/?q=-8.6582,115.2381', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Taman Budaya Bali (Art Center)", "distanceKm": 0.0}]'::jsonb,
    '["day"]'::jsonb, 250000.00, 300000.00,
    10, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'deposit_refundable',
    'Listrik & penerangan gratis dari panitia PKB.',
    '["power", "seating", "parking", "trash-area", "security"]'::jsonb,
    '["Diutamakan menjual kerajinan lokal dan kuliner khas Bali."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555939594-58d7cb561ad1?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 27, TRUE
);

-- 3.7 Temporary: Pop-Up Stand Festival Kuliner Nusantara Makassar
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Kuliner Anjungan Pantai Losari Makassar',
    'Spot kuliner temporary di pelataran Anjungan Pantai Losari Makassar saat festival kuliner nusantara.',
    'pop-up-booth', 'temporary', 'outdoor',
    1300,
    '{"event_name": "Makassar Culinary Expo 2026", "start_date": "2026-10-18", "end_date": "2026-10-22", "registration_deadline": "2026-10-05"}'::jsonb,
    '{"total_slots": 40, "available_slots": 8}'::jsonb,
    'Anjungan Pantai Losari, Stand B-05', 'Maddini', 'Ujung Pandang', 'Makassar', 'Sulawesi Selatan', 'Indonesia', 'ID', '90111',
    -5.13820000, 119.40810000, 'https://maps.google.com/?q=-5.1382,119.4081', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Pantai Losari", "distanceKm": 0.0}]'::jsonb,
    '["day"]'::jsonb, 200000.00, 200000.00,
    5, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'non_refundable',
    'Sudah termasuk daya listrik 1300 VA.',
    '["power", "seating", "parking", "trash-area", "security"]'::jsonb,
    '["Menjaga kebersihan area anjungan pantai."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 13, TRUE
);

-- 3.8 Temporary: Pop-Up Stand Night Market Malang
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Night Market Kayutangan Heritage Malang',
    'Stand pop-up pada event Pasar Malam Kayutangan Heritage Malang. Kawasan wisata arsitektur kolonial favorit anak muda.',
    'pop-up-booth', 'temporary', 'outdoor',
    1300,
    '{"event_name": "Kayutangan Heritage Night Market", "start_date": "2026-10-02", "end_date": "2026-11-02", "registration_deadline": "2026-09-22"}'::jsonb,
    '{"total_slots": 30, "available_slots": 6}'::jsonb,
    'Jl. Basuki Rahmat (Kayutangan), Stand 14', 'Kauman', 'Klojen', 'Malang', 'Jawa Timur', 'Indonesia', 'ID', '65119',
    -7.98120000, 112.63210000, 'https://maps.google.com/?q=-7.9812,112.6321', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Alun-Alun Tugu Malang", "distanceKm": 0.3}]'::jsonb,
    '["day"]'::jsonb, 180000.00, 200000.00,
    4, '["event_day_1"]'::jsonb,
    'weekends', 'flexible_days', 'deposit_refundable',
    'Listrik penerangan disiapkan panitia.',
    '["power", "trash-area", "security"]'::jsonb,
    '["Tampilan stand bernuansa retro/vintage disukai."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 18, TRUE
);

-- 3.9 Temporary: Bazaar Booth Mall Kelapa Gading
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Bazaar Kuliner Kampoeng Woenig MKG',
    'Booth bazaar kuliner nusantara pada event tahunan Kuliner Kampoeng Woenig di Mall Kelapa Gading.',
    'bazaar-booth', 'temporary', 'semi-outdoor',
    2200,
    '{"event_name": "Kampoeng Woenig MKG 2026", "start_date": "2026-10-20", "end_date": "2026-11-05", "registration_deadline": "2026-10-05"}'::jsonb,
    '{"total_slots": 45, "available_slots": 11}'::jsonb,
    'Mall Kelapa Gading La Piazza, Stand K-08', 'Kelapa Gading Timur', 'Kelapa Gading', 'Jakarta Utara', 'DKI Jakarta', 'Indonesia', 'ID', '14240',
    -6.15710000, 106.90710000, 'https://maps.google.com/?q=-6.1571,106.9071', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "La Piazza Kelapa Gading", "distanceKm": 0.0}]'::jsonb,
    '["day"]'::jsonb, 400000.00, 400000.00,
    10, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'non_refundable',
    'Listrik 2200 VA & fasiltas cuci alat F&B.',
    '["power", "water", "seating", "parking", "security"]'::jsonb,
    '["Menu makanan wajib lolos kurasi panitia MKG."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565123409695-7b5ef63a2efb?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 30, TRUE
);

-- 3.10 Temporary: Food Truck Court ICE BSD
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Outdoor Food Truck Area ICE BSD Concert',
    'Kavling outdoor food truck di pelataran ICE BSD saat konser musik k-pop & pameran otomotif nasional.',
    'food-truck-spot', 'temporary', 'outdoor',
    3500,
    '{"event_name": "ICE BSD Mega Concerts 2026", "start_date": "2026-11-12", "end_date": "2026-11-15", "registration_deadline": "2026-10-25"}'::jsonb,
    '{"total_slots": 20, "available_slots": 5}'::jsonb,
    'Indonesia Convention Exhibition (ICE) BSD, Courtyard 02', 'Pagedangan', 'Pagedangan', 'Tangerang', 'Banten', 'Indonesia', 'ID', '15339',
    -6.30120000, 106.63810000, 'https://maps.google.com/?q=-6.3012,106.6381', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "AEON Mall BSD City", "distanceKm": 0.5}]'::jsonb,
    '["day"]'::jsonb, 500000.00, 500000.00,
    4, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'non_refundable',
    'Listrik 3500 VA colokan industri.',
    '["power", "high-power", "wifi", "parking", "trash-area", "security"]'::jsonb,
    '["Kebersihan area sekitar armada menjadi tanggung jawab tenant."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1526367790999-0150786686a2?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.9, 28, TRUE
);

-- 3.11 Temporary: Pop-Up Stand CFD Dago Cikapayang
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Lapak Tenda Canopy Taman Cikapayang Dago Bandung',
    'Spot lapak mingguan CFD Dago depan Taman Cikapayang Bandung. Sangat pas untuk produk minuman dingin, roti, atau jus.',
    'pop-up-booth', 'temporary', 'outdoor',
    900,
    '{"event_name": "Sunday CFD Cikapayang Fest", "start_date": "2026-10-04", "end_date": "2026-12-27", "registration_deadline": "2026-09-28"}'::jsonb,
    '{"total_slots": 40, "available_slots": 10}'::jsonb,
    'Taman Cikapayang Dago Plot A-08', 'Lebak Siliwangi', 'Coblong', 'Bandung', 'Jawa Barat', 'Indonesia', 'ID', '40132',
    -6.89910000, 107.61010000, 'https://maps.google.com/?q=-6.8991,107.6101', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "ITB Bandung", "distanceKm": 0.3}]'::jsonb,
    '["day"]'::jsonb, 150000.00, 100000.00,
    1, '["event_day_1"]'::jsonb,
    'weekends', 'flexible_days', 'deposit_refundable',
    'Fasilitas kantong sampah disediakan panitia.',
    '["trash-area", "security"]'::jsonb,
    '["Bongkar tenda tepat jam 10:00 WIB."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 11, TRUE
);

-- 3.12 Temporary: Pop-Up Stand Festival Sekaten Jogja
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Pasar Rakyat Sekaten Alun-Alun Utara Jogja',
    'Spot bazaar festival budaya Sekaten di Alun-Alun Utara Yogyakarta. Sangat dipadati pengunjung malam hari.',
    'pop-up-booth', 'temporary', 'outdoor',
    1300,
    '{"event_name": "Sekaten Festival Jogja 2026", "start_date": "2026-10-10", "end_date": "2026-10-25", "registration_deadline": "2026-09-30"}'::jsonb,
    '{"total_slots": 70, "available_slots": 15}'::jsonb,
    'Alun-Alun Utara Yogyakarta, Plot A-21', 'Prawirodirjan', 'Gondomanan', 'Yogyakarta', 'DI Yogyakarta', 'Indonesia', 'ID', '55121',
    -7.80120000, 110.36510000, 'https://maps.google.com/?q=-7.8012,110.3651', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Kraton Ngayogyakarta Hadiningrat", "distanceKm": 0.2}]'::jsonb,
    '["day"]'::jsonb, 175000.00, 200000.00,
    5, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'deposit_refundable',
    'Kebutuhan daya listrik PLN pasar malam.',
    '["power", "trash-area", "security"]'::jsonb,
    '["Menjaga kebersihan area tumpukan sampah basah."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 15, TRUE
);

-- 3.13 Temporary: Bazaar Booth Mall Taman Anggrek
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Atrium Utama Mall Taman Anggrek Jakarta Barat',
    'Booth bazaar indoor berkelas di Atrium Mall Taman Anggrek. Dekat ice rink, traffic ramai sepanjang hari.',
    'bazaar-booth', 'temporary', 'indoor',
    2200,
    '{"event_name": "Taman Anggrek Culinary Fest 2026", "start_date": "2026-11-01", "end_date": "2026-11-08", "registration_deadline": "2026-10-20"}'::jsonb,
    '{"total_slots": 35, "available_slots": 7}'::jsonb,
    'Mall Taman Anggrek Lt. Ground Atrium Stand 05', 'Tanjung Duren Selatan', 'Grogol Petamburan', 'Jakarta Barat', 'DKI Jakarta', 'Indonesia', 'ID', '11470',
    -6.17820000, 106.79210000, 'https://maps.google.com/?q=-6.1782,106.7921', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Tarumanagara (UNTAR)", "distanceKm": 0.8}]'::jsonb,
    '["day"]'::jsonb, 450000.00, 500000.00,
    7, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'non_refundable',
    'Listrik 2200 VA & partisi karpet terpasang.',
    '["power", "air-conditioner", "wifi", "parking", "security"]'::jsonb,
    '["Wajib memakai ID Card panitia pameran."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565123409695-7b5ef63a2efb?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 22, TRUE
);

-- 3.14 Temporary: Pop-Up Stand Festival Teluk Penyu Cilacap
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Bazaar Wisata Pantai Teluk Penyu Cilacap',
    'Booth bazaar temporer pada acara pesta laut wisata Pantai Teluk Penyu Cilacap.',
    'bazaar-booth', 'temporary', 'outdoor',
    1300,
    '{"event_name": "Pesta Laut Teluk Penyu 2026", "start_date": "2026-10-15", "end_date": "2026-10-18", "registration_deadline": "2026-10-01"}'::jsonb,
    '{"total_slots": 30, "available_slots": 8}'::jsonb,
    'Area Pantai Teluk Penyu, Stand C-04', 'Cilacap', 'Cilacap Selatan', 'Cilacap', 'Jawa Tengah', 'Indonesia', 'ID', '53211',
    -7.73820000, 109.01210000, 'https://maps.google.com/?q=-7.7382,109.0121', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Benteng Pendem Cilacap", "distanceKm": 0.2}]'::jsonb,
    '["day"]'::jsonb, 150000.00, 150000.00,
    4, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'deposit_refundable',
    'Termasuk pencahayaan listrik.',
    '["power", "seating", "parking", "security"]'::jsonb,
    '["Menjaga kebersihan area pasir pantai."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555939594-58d7cb561ad1?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.5, 9, TRUE
);

-- 3.15 Temporary: Pop-Up Stand Night Market Palembang
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, daily_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Booth Night Market Benteng Kuto Besak Palembang',
    'Stand bazaar malam hari di pelataran Benteng Kuto Besak (BKB) dengan pemandangan Jembatan Ampera.',
    'pop-up-booth', 'temporary', 'outdoor',
    1300,
    '{"event_name": "BKB Ampera Night Market 2026", "start_date": "2026-11-01", "end_date": "2026-11-30", "registration_deadline": "2026-10-20"}'::jsonb,
    '{"total_slots": 40, "available_slots": 12}'::jsonb,
    'Pelataran BKB Palembang, Stand A-10', '19 Ilir', 'Bukit Kecil', 'Palembang', 'Sumatera Selatan', 'Indonesia', 'ID', '30113',
    -2.99120000, 104.76120000, 'https://maps.google.com/?q=-2.9912,104.7612', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "tourist-area", "name": "Jembatan Ampera", "distanceKm": 0.2}]'::jsonb,
    '["day"]'::jsonb, 175000.00, 200000.00,
    5, '["event_day_1"]'::jsonb,
    'everyday', 'mandatory_full', 'deposit_refundable',
    'Listrik penerangan malam hari terisi.',
    '["power", "trash-area", "security"]'::jsonb,
    '["Dilarang membuang limbah minyak langsung ke sungai Musi."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.7, 16, TRUE
);