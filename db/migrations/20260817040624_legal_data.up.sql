CREATE EXTENSION IF NOT EXISTS "pgcrypto";

INSERT INTO cms_legal_documents (id, doc_type, lang, title, intro, sections_json)
VALUES 
(
    gen_random_uuid(),
    'terms',
    'en',
    'Terms and Conditions',
    'These terms govern your use of Lapakita as a Tenant, Stall Owner, or Supplier. Please read them carefully before using the platform.',
    '[
      {
        "id": "platform-nature",
        "number": "1",
        "title": "Introduction & Platform Nature",
        "subsections": [
          {
            "body": "Lapakita is an online venue and operating platform connecting Tenants (Business Operators), Stall Owners, and B2B Suppliers. Lapakita is not a real estate broker, property manager, cleaner, law enforcement agent, or direct seller of physical products. Lapakita provides digital infrastructure, lease contract tools, escrow payment facilitation, and business intelligence analytics."
          }
        ]
      },
      {
        "id": "user-roles",
        "number": "2",
        "title": "User Roles & Multi-Persona Profiles",
        "subsections": [
          {
            "body": "A user registers a primary account verified by email and phone number. A single account may operate across three personas (Tenant, Stall Owner, Supplier) with customizable role-specific display names and avatars. Users remain responsible for all activities under their credentials."
          }
        ]
      },
      {
        "id": "stall-permanence-bazaar",
        "number": "3",
        "title": "Stall Classification & Short-Term Bazaar Events",
        "subsections": [
          {
            "title": "Operational Permanence Levels",
            "body": "Lapakita classifies properties into three distinct operational types: (a) Permanent (Independent properties with 24/7 access, physical sqm specifications, and no parent entity restrictions), (b) Semi-Permanent (Managed complex stalls such as mall shops, food court counters, and traditional market stalls bound by parent entity operating hours and parent complex names), and (c) Temporary (Short-term pop-up bazaar booths, street vendor spots, and food truck bays)."
          },
          {
            "title": "Pop-Up & Bazaar Event Booking & Lease Rules",
            "body": "Temporary bazaar booths are bound by specific event schedules, registration deadlines, slot availability, and event-specific lease rules. Event lease rules configure operating days (''everyday'', ''weekends'', ''weekdays'', ''flexible''), attendance requirements (''mandatory_full'' or ''flexible_days''), and cancellation policies (''pro_rata'', ''deposit_refundable'', or ''non_refundable''). Tenants applying for bazaar booths agree to abide by these event terms."
          }
        ]
      },
      {
        "id": "leasing-contracts",
        "number": "4",
        "title": "Stall Leasing, Contracts & Payment Timelines",
        "subsections": [
          {
            "title": "Digital Lease Agreement & Owner Configurations",
            "body": "Stall Owners configure specific lease rules for their listings. For Permanent & Semi-Permanent stalls, rules include Start Date options (1st, 15th, End of Month, or custom dates between 1-28), Minimum Lease Months, and Payment Cycles (Monthly, Quarterly, Semesterly, Yearly). For Temporary stalls, rules include Minimum Lease Days, Start Day options (e.g. Event Day 1, Event Day 2, Event Week 1), and Daily/Monthly event pricing."
          },
          {
            "title": "Approval Lock & Payment Deadline",
            "body": "Upon Owner approval, the stall is temporarily locked and removed from public search. The Tenant must complete the initial rent and security deposit payment via the Payment Gateway on or before the selected Start Date."
          },
          {
            "title": "Ghosting Penalty & Contract Cancellation",
            "body": "If a Tenant fails to complete payment by the Start Date deadline, the Owner holds the right to immediately cancel the contract and issue a 1-star public review for breach of commitment."
          },
          {
            "title": "Refundable Anti-Spam Commitment Fee",
            "body": "To protect Owners from application spam, Tenants with two (2) or more unpaid or cancelled approved applications within a 30-day window are flagged. Flagged Tenants are required to submit a temporary 35% commitment deposit when applying. This deposit is 100% REFUNDABLE and non-punitive: if the lease becomes active, 100% of the deposit is applied directly toward the Tenant''s initial rent and security deposit balance. If the application is cancelled or fails to proceed before the Start Date, the commitment deposit is fully refunded back to the Tenant''s account."
          }
        ]
      },
      {
        "id": "escrow",
        "number": "5",
        "title": "Security Deposit & Escrow Handling",
        "subsections": [
          {
            "title": "Escrow Storage",
            "body": "Security deposits are collected via a licensed Payment Gateway Escrow and held neutrally during the lease term. Deposits do not reside in the Owner''s personal bank account during active tenancy."
          },
          {
            "title": "Usage Scope",
            "body": "Security deposits exist strictly as a guarantee against physical property damage or lost key reproduction, not as daily punitive fines."
          },
          {
            "title": "Damage Claim & Appeal Process",
            "body": "Upon tenant exit, the Owner may submit a damage claim with itemized costs and timestamped photo evidence. The Tenant has a designated window to Accept or Appeal the claim. If Accepted, funds are disbursed to the Owner''s payout bank account, and the remainder is returned to the Tenant''s registered bank account. If Appealed, Lapakita Platform Support acts as a neutral administrative reviewer to inspect initial vs. final photo records and make a final binding deposit adjustment."
          },
          {
            "title": "Deposit Limits & Major Property Damage",
            "body": "The Security Deposit set by the Owner represents the maximum escrow guarantee recoverable directly through the platform. Lapakita is not liable for repair costs exceeding the deposited amount. In cases of severe property destruction or vandalism exceeding the deposit, Lapakita will disburse 100% of the available deposit to the Owner and provide verified KYC evidence to assist the Owner in formal legal proceedings. The offending Tenant''s account will be permanently blacklisted."
          }
        ]
      },
      {
        "id": "analytics-reports",
        "number": "6",
        "title": "Generated Reports & Data Analysis History",
        "subsections": [
          {
            "title": "Report Generation & Execution",
            "body": "Subscribed users (Premium or Active Tier) may execute automated data analysis reports, including Tenant Multi-Timeline Business Forecasts, Owner Vacancy Loss Analyses, and Supplier Market Opportunity Gap Analyses."
          },
          {
            "title": "Historical Data Archive Rights",
            "body": "All generated reports are compiled into structured JSON payloads and saved permanently in the user''s Report History archive. Users retain full rights to view, export, and download previously generated reports at any time, even if their subscription plan reverts to the Free tier. Regenerating new analysis reports requires an active subscription."
          }
        ]
      },
      {
        "id": "keys-access",
        "number": "7",
        "title": "Physical Keys, Duplication & Lock Cylinder Responsibility",
        "subsections": [
          {
            "title": "Initial Key Handover",
            "body": "Keys are handed over directly from the Owner to the Tenant at the start of the lease."
          },
          {
            "title": "Key Returns & Exit (Freedom of Return)",
            "body": "Returning physical keys upon lease termination is optional. Tenants are not penalized solely for unreturned keys, nor are they required to return duplicated sets."
          },
          {
            "title": "Key Duplication",
            "body": "Tenants are free to duplicate keys independently at local locksmiths at their own expense during the lease term."
          },
          {
            "title": "Owner Security Recommendation (Lock Cylinder Hygiene)",
            "body": "Lapakita strongly urges Stall Owners to replace the lock cylinder/knob set between different tenancies. If an Owner chooses to reuse an old lock set with spare keys, the Owner accepts all inherent security risks regarding potential duplicate keys. Lapakita bears no liability for property security breaches resulting from reused locks."
          },
          {
            "title": "Lost Key Protocol — Owner Has a Spare Key",
            "body": "If the Tenant loses their keys but the Owner has a master or spare key, the Tenant pays strictly for the key duplication fee, which can be deducted from the security deposit or paid directly to the Owner."
          },
          {
            "title": "Lost Key Protocol — Total Key Loss",
            "body": "If all keys are lost and a locksmith must pick the lock, forge a new key from scratch, or replace the entire lock cylinder: the Owner is responsible for managing the lock replacement process and covering any structural/lock cylinder hardware costs, as the underlying asset owner. The Tenant pays only for the cost of the individual key(s) created for them, as penalty for their negligence."
          }
        ]
      },
      {
        "id": "utilities-electricity",
        "number": "8",
        "title": "Utilities, Electricity & Operational Expenses",
        "subsections": [
          {
            "title": "Owner Provision",
            "body": "Stall Owners are responsible for providing basic operational utility infrastructure, including electrical power capacity (kVA), water meters, or plumbing connections as advertised in the listing."
          },
          {
            "title": "Tenant Usage & Billing Responsibility",
            "body": "Ongoing consumption of electricity, water, internet, trash disposal, or local market maintenance fees during the active lease term is the sole responsibility of the Tenant. Tenants must top up prepaid electricity tokens (PLN) or pay monthly utility bills directly."
          },
          {
            "title": "Utility Arrears Upon Exit",
            "body": "If a Tenant vacates a stall with unpaid post-paid utility bills or unpaid local maintenance fees, the Owner is entitled to deduct the exact outstanding arrears amount from the Tenant''s escrow security deposit upon exit."
          }
        ]
      },
      {
        "id": "cleanliness-eviction",
        "number": "9",
        "title": "Stall Cleanliness, Abandoned Items & Evictions",
        "subsections": [
          {
            "title": "Cleanliness Duty",
            "body": "Tenants are fully responsible for removing all personal items and inventory upon exit. Owners are responsible for presenting a clean space to incoming tenants."
          },
          {
            "title": "Manual Listing Reactivation",
            "body": "Active or pending stalls are automatically hidden from the marketplace. Upon a tenant''s exit or contract cancellation, the stall does NOT automatically reappear. It is the Owner''s sole responsibility to manually reactivate/publish the listing once the physical space is clean and ready for new viewings."
          },
          {
            "title": "Abandoned Goods",
            "body": "Items left behind by an evicted or departed tenant after lease termination may be disposed of, kept, or cleared by the Stall Owner at their sole discretion. Lapakita bears no liability for abandoned property."
          }
        ]
      },
      {
        "id": "supplier-disputes",
        "number": "10",
        "title": "Supplier Marketplace & B2B Disputes",
        "subsections": [
          {
            "title": "Peer-to-Peer Transactions",
            "body": "The B2B Supplier Marketplace connects Tenants directly with Suppliers."
          },
          {
            "title": "Dispute Handling",
            "body": "Lapakita does not provide manual admin arbitration for supplier product complaints (e.g. wrong ingredients, delayed deliveries, minor stock defects). Buyers and Suppliers must resolve issues via direct chat. Buyers retain full rights to leave public star ratings and reviews on product catalogs and supplier profiles."
          }
        ]
      },
      {
        "id": "payouts",
        "number": "11",
        "title": "Payouts & Bank Account Requirements",
        "subsections": [
          {
            "body": "Owners and Suppliers must register a valid bank account for automated payout disbursements. Tenants must register a valid bank account to receive potential deposit refunds."
          }
        ]
      }
    ]'::jsonb
),
(
    gen_random_uuid(),
    'privacy',
    'en',
    'Privacy Policy',
    'This policy explains what data Lapakita collects, how it''s used, and the protections in place across Tenant, Owner, and Supplier accounts.',
    '[
      {
        "id": "data-collected",
        "number": "1",
        "title": "Data We Collect",
        "subsections": [
          {
            "title": "Account Identity & Role Profiles",
            "body": "Full name, email address, multi-phone contact numbers (WhatsApp), role-specific avatars, display names, and encrypted password credentials."
          },
          {
            "title": "Verification Data (KYC)",
            "body": "ID card (KTP) photo, NIK, OCR data, and official business document photos collected prior to lease signing, stall publishing, or supplier activation."
          },
          {
            "title": "Financial & Payout Data",
            "body": "Bank account holder name, bank code, and account number for automated payment routing and escrow payouts."
          },
          {
            "title": "Operational & Generated Analysis Data",
            "body": "POS sales entries, stock levels, item prices, rental payment history, chat messages, uploaded property media, and saved historical analysis report payloads."
          }
        ]
      },
      {
        "id": "data-usage",
        "number": "2",
        "title": "How We Use Your Data",
        "subsections": [
          {
            "body": "To facilitate digital lease contracts, short-term bazaar booth bookings, billing, and automated payout transfers."
          },
          {
            "body": "To display B2B supplier catalogs to relevant tenant business categories."
          },
          {
            "body": "To compile historical business forecast analysis and store structured report history accessible via user dashboards."
          },
          {
            "body": "To verify identity in cases of legal lease disputes or deposit appeals."
          }
        ]
      },
      {
        "id": "data-protection",
        "number": "3",
        "title": "Data Protection & Non-Disclosure",
        "subsections": [
          {
            "title": "No Data Selling",
            "body": "Lapakita strictly never sells, rents, or trades user personal data, business revenues, or private transaction logs to third-party advertisers or data brokers."
          },
          {
            "title": "Privacy of Revenue Data",
            "body": "Individual tenant revenue figures and POS ledgers are strictly private to the tenant''s business account. Stall Owners cannot view a tenant''s exact gross revenue or profit margins."
          },
          {
            "title": "Secure Infrastructure",
            "body": "All sensitive payload data, API tokens, and credentials are encrypted using industry-standard protocols (TLS/SSL) and stored securely."
          }
        ]
      },
      {
        "id": "data-retention",
        "number": "4",
        "title": "Data Retention & User Rights",
        "subsections": [
          {
            "body": "Users may request account deactivation and data erasure, provided there are no active binding lease contracts, pending escrow deposit claims, unfulfilled B2B orders, or active event bookings associated with the account."
          }
        ]
      }
    ]'::jsonb
),
(
    gen_random_uuid(),
    'cookies',
    'en',
    'Cookies & Local Storage Policy',
    'Lapakita uses minimal browser storage — strictly for essential functionality, never for invasive tracking or ad retargeting.',
    '[
      {
        "id": "what-we-store",
        "number": "1",
        "title": "What We Store",
        "subsections": [
          {
            "title": "Session State",
            "body": "Keeping you logged in securely across page reloads."
          },
          {
            "title": "Active Role & Profile Preference",
            "body": "Remembering whether you last operated in Tenant, Owner, or Supplier mode."
          },
          {
            "title": "POS Cache",
            "body": "Temporarily caching POS cart items and product lists locally so your cashier interface remains fast and responsive even during minor network drops."
          }
        ]
      },
      {
        "id": "third-party-cookies",
        "number": "2",
        "title": "Third-Party Cookies",
        "subsections": [
          {
            "body": "We do not use invasive third-party tracking cookies, cross-site behavioral tracking scripts, or ad-retargeting pixels. Third-party scripts are strictly limited to secure Payment Gateway iFrames (Midtrans/Xendit) for payment processing."
          }
        ]
      }
    ]'::jsonb
)
ON CONFLICT (doc_type, lang) DO UPDATE 
SET title = EXCLUDED.title, intro = EXCLUDED.intro, sections_json = EXCLUDED.sections_json, updated_at = CURRENT_TIMESTAMP;