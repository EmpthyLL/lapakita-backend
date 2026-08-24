CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 1. PERMANENT: Ruko Dago Commercial Space (Bandung)
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    size_sqm, length_meters, width_meters, floor_level, electricity_capacity_va,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, quarterly_rate, semesterly_rate, yearly_rate, security_deposit,
    minimum_lease_months, start_date_options, utility_terms, facility_values, house_rules,
    display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Ruko Dago Commercial Space - Ground Floor Main Road',
    'Ruko komersial mandiri lokasi sangat strategis di koridor utama Jalan Ir. H. Juanda (Dago). Bebas jam operasional (akses 24/7), cocok untuk Cafe, Boutique, Office, atau Clinic. Dilengkapi halaman parkir sendiri, listrik 5500 VA, dan saluran air PDAM lancar.',
    'shophouse', 'permanent', 'indoor',
    45.00, 9.00, 5.00, 1, 5500,
    'Jl. Ir. H. Juanda No. 102', 'Lebak Siliwangi', 'Coblong', 'Bandung', 'Jawa Barat', 'Indonesia', 'ID', '40132',
    -6.89150000, 107.61060000, 'https://maps.google.com/?q=-6.8915,107.6106', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Institut Teknologi Bandung (ITB)", "distanceKm": 0.4}, {"categoryValue": "office", "name": "Dago Plaza & Business Hub", "distanceKm": 0.2}]'::jsonb,
    '["month", "quarter", "semester", "year"]'::jsonb, 7500000.00, 21000000.00, 40000000.00, 75000000.00, 5000000.00,
    3, '["1", "15", "eom"]'::jsonb, 'Listrik PLN & PDAM dibayar mandiri sesuai pemakaian meteran.',
    '["power", "high-power", "water", "drainage", "air-conditioner", "wifi", "parking", "toilet", "trash-area", "security", "cctv"]'::jsonb,
    '["Penyewa memegang kunci mandiri dan bertanggung jawab penuh atas keamanan internal.", "Renovasi interior diperbolehkan dengan konfirmasi sebelum pengerjaan."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1517248135467-4c7edcad34c4?w=1200&h=800&fit=crop", "facilityImages": [{"url": "https://images.unsplash.com/photo-1497366216548-37526070297c?w=1200&h=800&fit=crop", "caption": "Area Interior Lantai 1"}, {"url": "https://images.unsplash.com/photo-1578916171728-46686eac8d58?w=1200&h=800&fit=crop", "caption": "Area Parkir Depan Ruko"}], "virtualTour360Url": "https://pannellum.org/images/alma.jpg"}'::jsonb,
    4.9, 22, TRUE
);

-- 2. PERMANENT: Kios Garasi Komersial Jakal UGM (Sleman)
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
    'Kios Garasi Komersial Jalan Kaliurang KM 5',
    'Garasi rumah diubah menjadi kios komersial hook strategis. Dekat kawasan kos mahasiswa UGM. Cocok untuk warung kopi, laundry ekspres, atau usaha aksesoris.',
    'garage-driveway', 'permanent', 'indoor',
    24.00, 6.00, 4.00, 1, 2200,
    'Jl. Kaliurang KM 5.5 No. 12', 'Caturtunggal', 'Depok', 'Sleman', 'DI Yogyakarta', 'Indonesia', 'ID', '55281',
    -7.75810000, 110.38120000, 'https://maps.google.com/?q=-7.7581,110.3812', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Gadjah Mada (UGM)", "distanceKm": 0.8}]'::jsonb,
    '["month"]'::jsonb, 2500000.00, 1000000.00,
    1, '["1", "15"]'::jsonb, 'Listrik kWh meteran terpisah.',
    '["power", "water", "wifi", "parking"]'::jsonb,
    '["Menjaga kebersihan area garasi depan."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1521017432531-fbd92d768814?w=1200&h=800&fit=crop", "facilityImages": [{"url": "https://images.unsplash.com/photo-1554118811-1e0d58224f24?w=1200&h=800&fit=crop", "caption": "Tampak Depan Kios"}]}'::jsonb,
    4.9, 19, TRUE
);

-- 3. PERMANENT: Container Kiosk Pasar Minggu (Jakarta Selatan)
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
    'Container Kiosk Hook Jalan Raya Pasar Minggu',
    'Kios kontainer moderen di pinggir jalan raya utama. Visibilitas tinggi untuk brand F&B kekinian atau usaha minuman boba.',
    'street-kiosk', 'permanent', 'semi-outdoor',
    8.00, 4.00, 2.00, 1, 1300,
    'Jl. Raya Pasar Minggu No. 45', 'Pejaten Timur', 'Pasar Minggu', 'Jakarta Selatan', 'DKI Jakarta', 'Indonesia', 'ID', '12510',
    -6.28120000, 106.84210000, 'https://maps.google.com/?q=-6.2812,106.8421', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "transit-station", "name": "Stasiun Pasar Minggu", "distanceKm": 0.5}]'::jsonb,
    '["month"]'::jsonb, 1800000.00, 1000000.00,
    1, '["1", "15", "eom"]'::jsonb, 'Listrik flat 150rb/bulan.',
    '["power", "trash-area", "security"]'::jsonb,
    '["Dilarang membuang minyak F&B sembarangan."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1578916171728-46686eac8d58?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.6, 15, TRUE
);

-- 4. SEMI-PERMANENT: Kios Plaza Margonda (Depok)
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
    'Kios Ground Floor Plaza Margonda - Main Corridor',
    'Kios komersial strategis di koridor utama ground floor Plaza Margonda. Posisi hook dengan visibilitas tinggi dari eskalator utama. Sangat cocok untuk F&B Grab-and-Go atau Beverage Kiosk.',
    'mall-island', 'semi-permanent', 'indoor',
    3500, 'Plaza Margonda', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. Margonda Raya No. 188, GF Blok G-05', 'Pondok Cina', 'Beji', 'Depok', 'Jawa Barat', 'Indonesia', 'ID', '16424',
    -6.37320000, 106.83290000, 'https://maps.google.com/?q=-6.3732,106.8329', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "campus", "name": "Universitas Indonesia (UI)", "distanceKm": 0.6}]'::jsonb,
    '["month", "quarter"]'::jsonb, 3500000.00, 9900000.00, 2500000.00,
    1, '["1", "15", "eom"]'::jsonb, 'Listrik isi token mandiri.',
    '["power", "high-power", "water", "drainage", "grease-trap", "air-conditioner", "wifi", "toilet", "cleaning-service", "security"]'::jsonb,
    '["Wajib buka dan tutup sesuai jam operasional mall (10:00 - 22:00 WIB)."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "facilityImages": [{"url": "https://images.unsplash.com/photo-1441986300917-64674bd600d8?w=1200&h=800&fit=crop", "caption": "Instalasi Air & Grease-Trap"}]}'::jsonb,
    4.8, 14, TRUE
);

-- 5. SEMI-PERMANENT: Food Court Tunjungan Plaza (Surabaya)
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
    'Counter Food Court Area Utama Tunjungan Plaza',
    'Counter F&B siap pakai di area food court Tunjungan Plaza 3. Termasuk piring, meja bersama, dan fasilitas cuci piring dari pengelola.',
    'food-court-counter', 'semi-permanent', 'indoor',
    4400, 'Tunjungan Plaza 3', '{"opening_time": "10:00", "closing_time": "22:00", "is_24_hours": false}'::jsonb,
    'Jl. Jend. Basuki Rachmat No. 8-12, Lt. 5', 'Kedungdoro', 'Tegalsari', 'Surabaya', 'Jawa Timur', 'Indonesia', 'ID', '60261',
    -7.26210000, 112.73920000, 'https://maps.google.com/?q=-7.2621,112.7392', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Bumi Mandiri Tower", "distanceKm": 0.3}]'::jsonb,
    '["month"]'::jsonb, 5500000.00, 3000000.00,
    1, '["1", "15"]'::jsonb, 'Bagi hasil utilitas dikelola mall.',
    '["power", "water", "air-conditioner", "wifi", "toilet", "cleaning-service"]'::jsonb,
    '["Mengikuti standar higienitas F&B mall."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    5.0, 31, TRUE
);

-- 6. TEMPORARY: Pop-Up Booth Festival Ramadan Senayan (Jakarta Pusat)
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
    'Pop-Up Booth A-12 Kuliner Festival Ramadan Senayan',
    'Spot booth bazaar temporary pada event festival kuliner Ramadan di pelataran Parkir Timur Senayan. Paket sewa sudah mencakup meja, kursi, listrik 1300 VA, dan kebersihan event.',
    'bazaar-booth', 'temporary', 'outdoor',
    1300,
    '{"event_name": "Ramadan Culinary Fest 2026", "start_date": "2026-03-20", "end_date": "2026-03-23", "registration_deadline_days": 5}'::jsonb,
    '{"total_slots": 20, "available_slots": 6}'::jsonb,
    'GBK Parkir Timur Senayan, Booth A-12', 'Gelora', 'Tanah Abang', 'Jakarta Pusat', 'DKI Jakarta', 'Indonesia', 'ID', '10270',
    -6.21830000, 106.80220000, 'https://maps.google.com/?q=-6.2183,106.8022', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Kawasan Bisnis SCBD", "distanceKm": 0.8}]'::jsonb,
    '["day", "month"]'::jsonb, 250000.00, 2500000.00, 300000.00,
    1, '["event_day_1", "event_day_2", "event_week_1"]'::jsonb,
    'everyday', 'mandatory_full', 'deposit_refundable',
    'Listrik 1300 VA ter-include.',
    '["power", "wifi", "seating", "parking", "trash-area", "security"]'::jsonb,
    '["Loading barang H-1 sebelum event dimulai."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565123409695-7b5ef63a2efb?w=1200&h=800&fit=crop", "facilityImages": [{"url": "https://images.unsplash.com/photo-1555396273-367ea4eb4db5?w=1200&h=800&fit=crop", "caption": "Denah Layout Outdoor"}]}'::jsonb,
    4.7, 9, TRUE
);

-- 7. TEMPORARY: Food Truck Spot Breeze BSD (Tangerang Selatan)
INSERT INTO stalls (
    id, stall_owner_id, title, description, property_type, permanence_type, placement,
    electricity_capacity_va, event_schedule, slot_info,
    street_address, suburb, district, city, province, country, country_code, postal_code,
    latitude, longitude, map_url, embedded_map_url, nearby_landmarks,
    allowed_payment_cycles, monthly_rate, security_deposit,
    minimum_lease_days, start_date_options, event_operating_days, event_attendance_requirement, event_cancellation_policy,
    utility_terms, facility_values, house_rules, display_media, rating_avg, review_count, is_published
) VALUES (
    gen_random_uuid(), (SELECT id FROM users ORDER BY random() LIMIT 1),
    'Outdoor Courtyard Spot Food Truck Breeze BSD',
    'Lapak outdoor khusus parkir armada food truck / VW Combi F&B di pelataran Mall The Breeze BSD. Ramai pengunjung anak muda dan keluarga saat akhir pekan.',
    'food-truck-spot', 'temporary', 'outdoor',
    2200,
    '{"event_name": "BSD Food Truck Park", "start_date": "2026-04-01", "end_date": "2026-06-30", "registration_deadline_days": 5}'::jsonb,
    '{"total_slots": 8, "available_slots": 2}'::jsonb,
    'The Breeze BSD City, Outdoor Parkir Barat', 'Sampora', 'Cisauk', 'Tangerang Selatan', 'Banten', 'Indonesia', 'ID', '15345',
    -6.30210000, 106.65420000, 'https://maps.google.com/?q=-6.3021,106.6542', 'https://www.google.com/maps/embed?pb=!1m18...',
    '[{"categoryValue": "office", "name": "Unilever Head Office BSD", "distanceKm": 0.4}]'::jsonb,
    '["month"]'::jsonb, 3000000.00, 1000000.00,
    30, '["event_day_1", "event_week_1"]'::jsonb,
    'weekends', 'flexible_days', 'pro_rata',
    'Listrik colok sambungan parkir 2200 VA.',
    '["power", "wifi", "parking", "trash-area", "security"]'::jsonb,
    '["Armada wajib memiliki izin sertifikasi kelaikan F&B."]'::jsonb,
    '{"mainImage": "https://images.unsplash.com/photo-1565123409695-7b5ef63a2efb?w=1200&h=800&fit=crop", "facilityImages": []}'::jsonb,
    4.8, 15, TRUE
);