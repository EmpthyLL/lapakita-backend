CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- SEED DATA: 16 MASTER BUSINESS TYPES (I18n, Financial Benchmarks & Permanence Presets)
INSERT INTO business_types (
    id,
    group_name_lang,
    label_lang,
    default_bep_months,
    default_capital,
    avg_gross_margin_ratio,
    industry_rent_to_revenue_ratio,
    recommended_landmarks,
    permanence_presets
) VALUES
-- 1. Full-Service Restaurant
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Full-Service Restaurant", "id": "Restoran Layanan Penuh"}'::jsonb,
    12,
    60000000.00,
    0.5000,
    0.1500,
    '["office", "market", "residential", "culinary-center"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway", "street-kiosk"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 40, "max": 120},
            "recommendedFloors": {"min": 1, "max": 2},
            "facilities": ["power", "high-power", "water", "drainage", "grease-trap", "ventilation", "air-conditioner", "seating", "toilet", "parking", "trash-area"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop", "food-court-counter"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "22:00",
            "facilities": ["power", "high-power", "water", "drainage", "grease-trap", "ventilation", "air-conditioner", "seating", "toilet", "parking", "trash-area", "cleaning-service"]
        }
    }'::jsonb
),

-- 2. Coffee Shop & Cafe
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Coffee Shop & Cafe", "id": "Kedai Kopi & Kafe"}'::jsonb,
    9,
    35000000.00,
    0.7000,
    0.1800,
    '["office", "campus", "residential", "transit-station"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway", "street-kiosk"],
            "allowedPlacements": ["indoor", "semi-outdoor", "outdoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 20, "max": 60},
            "recommendedFloors": {"min": 1, "max": 2},
            "facilities": ["power", "high-power", "water", "drainage", "ventilation", "wifi", "seating", "toilet", "parking"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop", "mall-island", "food-court-counter"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "22:00",
            "facilities": ["power", "high-power", "water", "drainage", "wifi", "seating", "trash-area"]
        },
        "temporary": {
            "allowedPropertyTypes": ["bazaar-booth", "food-truck-spot", "street-vendor-spot"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "outdoor",
            "registrationWindowDaysBefore": 7,
            "typicalDurationDays": 3,
            "facilities": ["power", "water", "trash-area", "seating"]
        }
    }'::jsonb
),

-- 3. Bakery & Pastry Shop
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Bakery & Pastry Shop", "id": "Toko Roti & Kue"}'::jsonb,
    9,
    30000000.00,
    0.5500,
    0.1500,
    '["residential", "school", "market", "transit-station"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway", "street-kiosk"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 15, "max": 40},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "high-power", "water", "ventilation", "air-conditioner", "storage", "display-case", "toilet"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop", "mall-island", "traditional-market-shop"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "08:00",
            "defaultClosingTime": "21:00",
            "facilities": ["power", "display-case", "trash-area"]
        },
        "temporary": {
            "allowedPropertyTypes": ["bazaar-booth"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "registrationWindowDaysBefore": 5,
            "typicalDurationDays": 3,
            "facilities": ["power", "display-case", "trash-area"]
        }
    }'::jsonb
),

-- 4. Quick-Service / Fast Food
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Quick-Service / Fast Food", "id": "Makanan Cepat Saji"}'::jsonb,
    8,
    25000000.00,
    0.5000,
    0.1500,
    '["campus", "office", "market", "culinary-center"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "street-kiosk", "garage-driveway"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "semi-outdoor",
            "recommendedSizeSqm": {"min": 15, "max": 35},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "water", "drainage", "grease-trap", "ventilation", "trash-area", "parking"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["food-court-counter", "mall-island", "traditional-market-shop"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "semi-outdoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "22:00",
            "facilities": ["power", "water", "drainage", "trash-area"]
        },
        "temporary": {
            "allowedPropertyTypes": ["bazaar-booth", "street-vendor-spot"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "outdoor",
            "registrationWindowDaysBefore": 5,
            "typicalDurationDays": 3,
            "facilities": ["power", "water", "trash-area"]
        }
    }'::jsonb
),

-- 5. Beverage & Snack Kiosk
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Beverage & Snack Kiosk", "id": "Kios Minuman & Camilan"}'::jsonb,
    4,
    10000000.00,
    0.6500,
    0.1600,
    '["campus", "school", "residential", "transit-station", "transit-bus"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["street-kiosk", "garage-driveway"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "semi-outdoor",
            "recommendedSizeSqm": {"min": 4, "max": 12},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "water", "drainage", "trash-area", "display-case"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-island", "food-court-counter", "open-market-stall"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "22:00",
            "facilities": ["power", "water", "drainage", "trash-area", "display-case"]
        },
        "temporary": {
            "allowedPropertyTypes": ["bazaar-booth", "street-vendor-spot"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "outdoor",
            "registrationWindowDaysBefore": 3,
            "typicalDurationDays": 3,
            "facilities": ["power", "water", "trash-area"]
        }
    }'::jsonb
),

-- 6. Street Food & Hawker Stall
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Street Food & Hawker Stall", "id": "Jajanan Kakilima & Lapak Kakilima"}'::jsonb,
    3,
    6000000.00,
    0.6000,
    0.1200,
    '["school", "campus", "residential", "transit-bus"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["street-kiosk", "garage-driveway"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "outdoor",
            "recommendedSizeSqm": {"min": 4, "max": 12},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "water", "trash-area"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["open-market-stall"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "outdoor",
            "defaultOpeningTime": "16:00",
            "defaultClosingTime": "23:00",
            "facilities": ["power", "water", "trash-area"]
        },
        "temporary": {
            "allowedPropertyTypes": ["street-vendor-spot", "bazaar-booth"],
            "allowedPlacements": ["outdoor"],
            "defaultPlacement": "outdoor",
            "registrationWindowDaysBefore": 2,
            "typicalDurationDays": 2,
            "facilities": ["power", "trash-area"]
        }
    }'::jsonb
),

-- 7. Meat, Poultry & Seafood Retail
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Meat, Poultry & Seafood Retail", "id": "Toko Daging, Ayam & Makanan Laut"}'::jsonb,
    5,
    12000000.00,
    0.4500,
    0.1200,
    '["market", "residential", "culinary-center"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "street-kiosk"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "semi-outdoor",
            "recommendedSizeSqm": {"min": 10, "max": 25},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "water", "drainage", "trash-area"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["traditional-market-shop", "open-market-stall"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "semi-outdoor",
            "defaultOpeningTime": "04:00",
            "defaultClosingTime": "12:00",
            "facilities": ["water", "drainage", "trash-area"]
        }
    }'::jsonb
),

-- 8. Food Truck & Mobile Unit
(
    gen_random_uuid(),
    '{"en": "F&B (Food & Beverages)", "id": "F&B (Makanan & Minuman)"}'::jsonb,
    '{"en": "Food Truck & Mobile Unit", "id": "Food Truck & Unit Seluler"}'::jsonb,
    6,
    20000000.00,
    0.5500,
    0.1500,
    '["culinary-center", "campus", "office"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["garage-driveway"],
            "allowedPlacements": ["outdoor"],
            "defaultPlacement": "outdoor",
            "recommendedSizeSqm": {"min": 12, "max": 25},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "water", "trash-area", "parking"]
        },
        "temporary": {
            "allowedPropertyTypes": ["food-truck-spot"],
            "allowedPlacements": ["outdoor"],
            "defaultPlacement": "outdoor",
            "registrationWindowDaysBefore": 5,
            "typicalDurationDays": 3,
            "facilities": ["power", "water", "trash-area", "parking"]
        }
    }'::jsonb
),

-- 9. Mini Market & Convenience Store
(
    gen_random_uuid(),
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    '{"en": "Mini Market & Convenience Store", "id": "Toko Kelontong & Minimarket"}'::jsonb,
    12,
    60000000.00,
    0.2000,
    0.0800,
    '["residential", "office", "gas-station", "healthcare"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 40, "max": 100},
            "recommendedFloors": {"min": 1, "max": 2},
            "facilities": ["power", "high-power", "water", "air-conditioner", "storage", "security", "cctv", "parking"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "22:00",
            "facilities": ["power", "high-power", "water", "air-conditioner", "storage", "security", "cctv"]
        }
    }'::jsonb
),

-- 10. Fresh Fruits, Vegetables & Spices
(
    gen_random_uuid(),
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    '{"en": "Fresh Fruits, Vegetables & Spices", "id": "Buah, Sayur & Bumbu Segar"}'::jsonb,
    4,
    8000000.00,
    0.2500,
    0.0800,
    '["market", "residential"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "street-kiosk", "garage-driveway"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "semi-outdoor",
            "recommendedSizeSqm": {"min": 8, "max": 20},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["water", "drainage", "trash-area"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["traditional-market-shop", "open-market-stall"],
            "allowedPlacements": ["semi-outdoor", "outdoor"],
            "defaultPlacement": "semi-outdoor",
            "defaultOpeningTime": "04:00",
            "defaultClosingTime": "14:00",
            "facilities": ["water", "drainage", "trash-area"]
        },
        "temporary": {
            "allowedPropertyTypes": ["street-vendor-spot"],
            "allowedPlacements": ["outdoor"],
            "defaultPlacement": "outdoor",
            "registrationWindowDaysBefore": 1,
            "typicalDurationDays": 1,
            "facilities": ["trash-area"]
        }
    }'::jsonb
),

-- 11. Fashion, Apparel & Accessory Boutique
(
    gen_random_uuid(),
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    '{"en": "Fashion, Apparel & Accessory Boutique", "id": "Butik Pakaian & Aksesori"}'::jsonb,
    10,
    40000000.00,
    0.4500,
    0.1400,
    '["market", "office", "campus"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 15, "max": 40},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "air-conditioner", "display-case", "storage", "security", "wifi"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop", "mall-island", "traditional-market-shop", "open-market-stall"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "21:00",
            "facilities": ["power", "display-case", "security"]
        },
        "temporary": {
            "allowedPropertyTypes": ["bazaar-booth"],
            "allowedPlacements": ["indoor", "semi-outdoor", "outdoor"],
            "defaultPlacement": "indoor",
            "registrationWindowDaysBefore": 5,
            "typicalDurationDays": 3,
            "facilities": ["power", "display-case"]
        }
    }'::jsonb
),

-- 12. General Retail & Hobby Store
(
    gen_random_uuid(),
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    '{"en": "General Retail & Hobby Store", "id": "Toko Ritel Umum & Hobi"}'::jsonb,
    9,
    25000000.00,
    0.3500,
    0.1200,
    '["residential", "market", "school"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "street-kiosk", "garage-driveway"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 10, "max": 30},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "display-case", "storage", "security", "cctv"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop", "mall-island", "traditional-market-shop"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "21:00",
            "facilities": ["power", "display-case", "security"]
        },
        "temporary": {
            "allowedPropertyTypes": ["bazaar-booth"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "indoor",
            "registrationWindowDaysBefore": 5,
            "typicalDurationDays": 3,
            "facilities": ["power", "display-case"]
        }
    }'::jsonb
),

-- 13. Beauty Salon, Barbershop & Spa
(
    gen_random_uuid(),
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    '{"en": "Beauty Salon, Barbershop & Spa", "id": "Salon Kecantikan, Pangkas Rambut & Spa"}'::jsonb,
    10,
    30000000.00,
    0.6000,
    0.1500,
    '["residential", "office", "campus"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 15, "max": 45},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "water", "drainage", "air-conditioner", "seating", "toilet", "parking"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "21:00",
            "facilities": ["power", "water", "drainage", "air-conditioner", "seating", "toilet"]
        }
    }'::jsonb
),

-- 14. Service, Repair Shop & Laundry
(
    gen_random_uuid(),
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    '{"en": "Service, Repair Shop & Laundry", "id": "Bengkel Servis, Reparasi & Laundry"}'::jsonb,
    8,
    20000000.00,
    0.5000,
    0.1200,
    '["residential", "office", "gas-station"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway", "street-kiosk"],
            "allowedPlacements": ["indoor", "semi-outdoor"],
            "defaultPlacement": "semi-outdoor",
            "recommendedSizeSqm": {"min": 15, "max": 40},
            "recommendedFloors": {"min": 1, "max": 1},
            "facilities": ["power", "high-power", "water", "drainage", "storage", "parking"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["traditional-market-shop"],
            "allowedPlacements": ["semi-outdoor"],
            "defaultPlacement": "semi-outdoor",
            "defaultOpeningTime": "08:00",
            "defaultClosingTime": "17:00",
            "facilities": ["power", "water", "storage"]
        }
    }'::jsonb
),

-- 15. Professional Office & Agency
(
    gen_random_uuid(),
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    '{"en": "Professional Office & Agency", "id": "Kantor Profesional & Agensi"}'::jsonb,
    12,
    45000000.00,
    0.5500,
    0.1500,
    '["office", "market", "government"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 20, "max": 60},
            "recommendedFloors": {"min": 1, "max": 2},
            "facilities": ["power", "high-power", "wifi", "air-conditioner", "security", "toilet", "parking"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "09:00",
            "defaultClosingTime": "18:00",
            "facilities": ["power", "high-power", "wifi", "air-conditioner", "security", "toilet"]
        }
    }'::jsonb
),

-- 16. Education & Studio Space
(
    gen_random_uuid(),
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    '{"en": "Education & Studio Space", "id": "Ruang Edukasi, Kursus & Studio"}'::jsonb,
    12,
    35000000.00,
    0.5000,
    0.1400,
    '["school", "campus", "residential"]'::jsonb,
    '{
        "permanent": {
            "allowedPropertyTypes": ["shophouse", "garage-driveway"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "recommendedSizeSqm": {"min": 25, "max": 60},
            "recommendedFloors": {"min": 1, "max": 2},
            "facilities": ["power", "water", "air-conditioner", "wifi", "seating", "toilet", "security", "parking"]
        },
        "semi-permanent": {
            "allowedPropertyTypes": ["mall-shop"],
            "allowedPlacements": ["indoor"],
            "defaultPlacement": "indoor",
            "defaultOpeningTime": "10:00",
            "defaultClosingTime": "20:00",
            "facilities": ["power", "water", "air-conditioner", "wifi", "seating", "toilet"]
        }
    }'::jsonb
);