CREATE EXTENSION IF NOT EXISTS "pgcrypto";

INSERT INTO business_types (
    id,
    label_lang,
    group_name_lang,
    default_bep_months,
    default_capital,
    avg_gross_margin_ratio,
    industry_rent_to_revenue_ratio,
    permanence_presets,
    recommended_landmarks
) VALUES

-- =============================================================================
-- F&B (FOOD & BEVERAGES)
-- =============================================================================

(
    gen_random_uuid(),
    '{"en": "Full-Service Restaurant", "id": "Restoran Layanan Penuh"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    12,
    60000000.00,
    0.5000,
    0.1500,
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
    }'::jsonb,
    '["office", "market", "residential", "culinary-center"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Coffee Shop & Cafe", "id": "Kedai Kopi & Kafe"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    9,
    35000000.00,
    0.7000,
    0.1800,
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
    }'::jsonb,
    '["office", "campus", "residential", "transit-station"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Bakery & Pastry Shop", "id": "Toko Roti & Kue"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    9,
    30000000.00,
    0.5500,
    0.1500,
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
    }'::jsonb,
    '["residential", "school", "market", "transit-station"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Quick-Service / Fast Food", "id": "Makanan Cepat Saji"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    8,
    25000000.00,
    0.5000,
    0.1500,
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
    }'::jsonb,
    '["campus", "office", "market", "culinary-center"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Beverage & Snack Kiosk", "id": "Kios Minuman & Makanan Ringan"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    4,
    10000000.00,
    0.6500,
    0.1600,
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
    }'::jsonb,
    '["campus", "school", "residential", "transit-station", "transit-bus"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Street Food & Hawker Stall", "id": "Jajanan Kaki Lima & Lapak Kuliner"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    3,
    6000000.00,
    0.6000,
    0.1200,
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
    }'::jsonb,
    '["school", "campus", "residential", "transit-bus"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Meat, Poultry & Seafood Retail", "id": "Kios Daging, Unggas & Seafood"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    5,
    12000000.00,
    0.4500,
    0.1200,
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
    }'::jsonb,
    '["market", "residential", "culinary-center"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Food Truck & Mobile Unit", "id": "Food Truck & Unit Keliling"}'::jsonb,
    '{"en": "F&B (Food & Beverages)", "id": "Makanan & Minuman"}'::jsonb,
    6,
    20000000.00,
    0.5500,
    0.1500,
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
    }'::jsonb,
    '["culinary-center", "campus", "office"]'::jsonb
),

-- =============================================================================
-- RETAIL & COMMERCE
-- =============================================================================

(
    gen_random_uuid(),
    '{"en": "Mini Market & Convenience Store", "id": "Minimarket & Toko Kelontong Modern"}'::jsonb,
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    12,
    60000000.00,
    0.2000,
    0.0800,
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
    }'::jsonb,
    '["residential", "office", "gas-station", "healthcare"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Fresh Fruits, Vegetables & Spices", "id": "Toko Buah, Sayur & Bumbu Segar"}'::jsonb,
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    4,
    8000000.00,
    0.2500,
    0.0800,
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
    }'::jsonb,
    '["market", "residential"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Fashion, Apparel & Accessory Boutique", "id": "Batik, Pakaian & Aksesori"}'::jsonb,
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    10,
    40000000.00,
    0.4500,
    0.1400,
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
    }'::jsonb,
    '["market", "office", "campus"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "General Retail & Hobby Store", "id": "Toko Kelontong Umum & Hobi"}'::jsonb,
    '{"en": "Retail & Commerce", "id": "Ritel & Perdagangan"}'::jsonb,
    9,
    25000000.00,
    0.3500,
    0.1200,
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
    }'::jsonb,
    '["residential", "market", "school"]'::jsonb
),

-- =============================================================================
-- SERVICES
-- =============================================================================

(
    gen_random_uuid(),
    '{"en": "Beauty Salon, Barbershop & Spa", "id": "Salon Kecantikan, Barbershop & Spa"}'::jsonb,
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    10,
    30000000.00,
    0.6000,
    0.1500,
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
    }'::jsonb,
    '["residential", "office", "campus"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Service, Repair Shop & Laundry", "id": "Jasa Reparasi, Servis & Laundry"}'::jsonb,
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    8,
    20000000.00,
    0.5000,
    0.1200,
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
    }'::jsonb,
    '["residential", "office", "gas-station"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Professional Office & Agency", "id": "Kantor Profesional & Agensi"}'::jsonb,
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    12,
    45000000.00,
    0.5500,
    0.1500,
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
    }'::jsonb,
    '["office", "market", "government"]'::jsonb
),
(
    gen_random_uuid(),
    '{"en": "Education & Studio Space", "id": "Studio Foto, Seni & Tempat Kursus"}'::jsonb,
    '{"en": "Services", "id": "Jasa & Layanan"}'::jsonb,
    12,
    35000000.00,
    0.5000,
    0.1400,
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
    }'::jsonb,
    '["school", "campus", "residential"]'::jsonb
);