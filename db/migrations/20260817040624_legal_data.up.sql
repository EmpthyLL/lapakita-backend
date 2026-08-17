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
            "body": "Lapakita is an online venue and operating platform connecting Tenants (Business Operators), Stall Owners, and B2B Suppliers. Lapakita is not a real estate broker, property manager, cleaner, law enforcement agent, or direct seller of products. Lapakita provides digital infrastructure, contract tools, escrow payment facilitation, and business analytics."
          }
        ]
      },
      {
        "id": "user-roles",
        "number": "2",
        "title": "User Roles & Single Account Multi-Persona",
        "subsections": [
          {
            "body": "A user registers a single primary account verified by email and phone number. A single account may operate across three personas (Tenant, Stall Owner, Supplier). Users are responsible for all activities under their credentials."
          }
        ]
      },
      {
        "id": "leasing-contracts",
        "number": "3",
        "title": "Stall Leasing, Contracts & Payment Timelines",
        "subsections": [
          {
            "title": "Digital Lease Agreement & Owner Configurations",
            "body": "Stall Owners configure specific lease rules for their listings, including Start Date options (1st, 15th, End of Month, or custom dates between 1-28), Minimum Lease Durations, and Payment Cycles (Monthly, Quarterly, Semesterly, Yearly). Tenants use search filters to find stalls matching their preferred timeline."
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
            "body": "To protect Owners from application spam, Tenants with two (2) or more unpaid or cancelled approved applications within a 30-day window are flagged. Flagged Tenants are required to submit a temporary 35% commitment deposit when applying. This deposit is 100% REFUNDABLE and non-punitive: if the lease becomes active, 100% of the deposit is applied directly toward the Tenant''s initial rent and security deposit balance. If the application is cancelled or fails to proceed before the Start Date, the commitment deposit is fully refunded back to the Tenant''s account. Account flag status and public star ratings remain the primary disciplinary measures for non-completion."
          }
        ]
      },
      {
        "id": "escrow",
        "number": "4",
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
        "id": "keys-access",
        "number": "5",
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
        "number": "6",
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
        "number": "7",
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
        "number": "8",
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
        "number": "9",
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
    'terms',
    'id',
    'Syarat dan Ketentuan',
    'Syarat dan ketentuan ini mengatur penggunaan Lapakita oleh Penyewa, Pemilik Lapak, dan Supplier. Harap baca dengan cermat sebelum menggunakan platform.',
    '[
      {
        "id": "platform-nature",
        "number": "1",
        "title": "Pendahuluan & Sifat Platform",
        "subsections": [
          {
            "body": "Lapakita adalah platform operasional dan tempat sewa online yang menghubungkan Penyewa (Pelaku Usaha), Pemilik Lapak, dan Supplier B2B. Lapakita bukan agen properti, pengelola gedung, penyedia jasa kebersihan, aparat penegak hukum, atau penjual langsung. Lapakita menyediakan infrastruktur digital, alat pembuatan kontrak, fasilitasi pembayaran escrow, dan analitik bisnis."
          }
        ]
      },
      {
        "id": "user-roles",
        "number": "2",
        "title": "Peran Pengguna & Satu Akun Multi-Persona",
        "subsections": [
          {
            "body": "Pengguna mendaftarkan satu akun utama yang diverifikasi melalui email dan nomor telepon. Satu akun dapat beroperasi dalam tiga persona (Penyewa, Pemilik Lapak, Supplier). Pengguna bertanggung jawab penuh atas seluruh aktivitas di bawah kredensial mereka."
          }
        ]
      },
      {
        "id": "leasing-contracts",
        "number": "3",
        "title": "Sewa Lapak, Kontrak & Tenggat Pembayaran",
        "subsections": [
          {
            "title": "Perjanjian Sewa Digital & Konfigurasi Pemilik",
            "body": "Pemilik Lapak mengonfigurasi aturan sewa spesifik pada listing mereka, termasuk opsi Tanggal Mulai (Tanggal 1, 15, Akhir Bulan, atau tanggal kustom antara 1-28), Durasi Sewa Minimum, dan Siklus Pembayaran (Bulanan, Triwulan, Semesteran, Tahunan). Penyewa menggunakan filter pencarian untuk menemukan lapak yang sesuai dengan preferensi mereka."
          },
          {
            "title": "Kunci Persetujuan & Tenggat Pembayaran",
            "body": "Setelah persetujuan Pemilik, lapak akan dikunci sementara dan dihapus dari pencarian publik. Penyewa wajib menyelesaikan pembayaran sewa awal dan deposit jaminan melalui Payment Gateway pada atau sebelum Tanggal Mulai yang dipilih."
          },
          {
            "title": "Sanksi Pembatalan & Pembatalan Kontrak",
            "body": "Jika Penyewa gagal menyelesaikan pembayaran sesuai tenggat Tanggal Mulai, Pemilik berhak untuk segera membatalkan kontrak dan memberikan ulasan publik bintang 1 atas pelanggaran komitmen."
          },
          {
            "title": "Biaya Komitmen Anti-Spam (Dapat Dikembalikan)",
            "body": "Untuk melindungi Pemilik dari spam pengajuan, Penyewa dengan dua (2) atau lebih pengajuan yang disetujui namun tidak dibayar/dibatalkan dalam kurun waktu 30 hari akan ditandai. Penyewa yang ditandai wajib menyetorkan deposit komitmen sementara sebesar 35% saat mengajukan sewa. Deposit ini 100% DAPAT DIKEMBALIKAN dan tidak bersifat menghukum: jika sewa aktif, 100% deposit langsung dialokasikan untuk mengurangi tagihan sewa awal dan deposit jaminan. Jika pengajuan dibatalkan sebelum Tanggal Mulai, deposit komitmen dikembalikan penuh ke akun Penyewa. Status penandaan akun dan ulasan bintang publik tetap menjadi tindakan disipliner utama."
          }
        ]
      },
      {
        "id": "escrow",
        "number": "4",
        "title": "Deposit Jaminan & Pengelolaan Escrow",
        "subsections": [
          {
            "title": "Penyimpanan Escrow",
            "body": "Deposit jaminan dikumpulkan melalui Escrow Payment Gateway berizin dan dipegang secara netral selama masa sewa. Deposit tidak disimpan di rekening bank pribadi Pemilik selama masa sewa aktif."
          },
          {
            "title": "Cakupan Penggunaan",
            "body": "Deposit jaminan murni berfungsi sebagai jaminan terhadap kerusakan fisik properti atau penggantian kunci yang hilang, bukan sebagai denda harian."
          },
          {
            "title": "Klaim Kerusakan & Proses Banding",
            "body": "Saat penyewa keluar, Pemilik dapat mengajukan klaim kerusakan disertai rincian biaya dan bukti foto berstempel waktu. Penyewa memiliki batas waktu untuk Menerima atau Mengajukan Banding. Jika Diterima, dana dicairkan ke rekening bank Pemilik, dan sisanya dikembalikan ke Penyewa. Jika Dibanding, Dukungan Platform Lapakita bertindak sebagai peninjau administratif netral untuk memeriksa catatan foto awal vs. akhir dan membuat penyesuaian deposit yang bersifat mengikat."
          },
          {
            "title": "Batas Deposit & Kerusakan Mayor Properti",
            "body": "Deposit Jaminan yang ditetapkan Pemilik merupakan batas maksimal jaminan escrow yang dapat dipulihkan langsung melalui platform. Lapakita tidak bertanggung jawab atas biaya perbaikan yang melebihi jumlah deposit. Dalam kasus kerusakan berat atau vandalisme yang melebihi deposit, Lapakita akan mencairkan 100% deposit yang tersedia kepada Pemilik serta memberikan bukti KYC terverifikasi untuk membantu Pemilik dalam proses hukum formal. Akun Penyewa yang melanggar akan diblokir permanen."
          }
        ]
      },
      {
        "id": "keys-access",
        "number": "5",
        "title": "Kunci Fisik, Duplikasi & Tanggung Jawab Silinder Kunci",
        "subsections": [
          {
            "title": "Serah Terima Kunci Awal",
            "body": "Kunci diserahkan secara langsung dari Pemilik kepada Penyewa pada awal masa sewa."
          },
          {
            "title": "Pengembalian Kunci & Keluar (Kebebasan Pengembalian)",
            "body": "Pengembalian kunci fisik saat masa sewa berakhir bersifat opsional. Penyewa tidak dikenakan sanksi hanya karena tidak mengembalikan kunci, dan tidak diwajibkan mengembalikan kunci hasil duplikasi."
          },
          {
            "title": "Duplikasi Kunci",
            "body": "Penyewa bebas menduplikasi kunci secara mandiri di tukang kunci lokal dengan biaya sendiri selama masa sewa."
          },
          {
            "title": "Rekomendasi Keamanan Pemilik (Higiene Silinder Kunci)",
            "body": "Lapakita sangat menyarankan Pemilik Lapak untuk mengganti silinder/gagang kunci di antara periode penyewa yang berbeda. Jika Pemilik memilih menggunakan kembali set kunci lama, Pemilik menanggung seluruh risiko keamanan terkait potensi kunci duplikat. Lapakita tidak bertanggung jawab atas pelanggaran keamanan properti akibat penggunaan kunci lama."
          },
          {
            "title": "Protokol Kunci Hilang — Pemilik Memiliki Kunci Cadangan",
            "body": "Jika Penyewa kehilangan kunci tetapi Pemilik memiliki kunci cadangan, Penyewa hanya membayar biaya duplikasi kunci, yang dapat dipotong dari deposit jaminan atau dibayarkan langsung ke Pemilik."
          },
          {
            "title": "Protokol Kunci Hilang — Kehilangan Kunci Total",
            "body": "Jika seluruh kunci hilang dan tukang kunci harus membongkar, membuat kunci baru dari nol, atau mengganti seluruh silinder kunci: Pemilik bertanggung jawab mengelola proses penggantian dan menanggung biaya perangkat keras silinder kunci sebagai pemilik aset. Penyewa hanya membayar biaya pembuatan kunci individu yang dibuat untuk mereka sebagai sanksi atas kelalaian."
          }
        ]
      },
      {
        "id": "utilities-electricity",
        "number": "6",
        "title": "Utilitas, Listrik & Biaya Operasional",
        "subsections": [
          {
            "title": "Penyediaan oleh Pemilik",
            "body": "Pemilik Lapak bertanggung jawab menyediakan infrastruktur utilitas operasional dasar, termasuk kapasitas daya listrik (kVA), meteran air, atau sambungan pipa sesuai yang diiklankan pada listing."
          },
          {
            "title": "Penggunaan & Tanggung Jawab Tagihan Penyewa",
            "body": "Konsumsi harian listrik, air, internet, pembuangan sampah, atau iuran pemeliharaan pasar lokal selama masa sewa aktif menjadi tanggung jawab penuh Penyewa. Penyewa wajib mengisi ulang token listrik prabayar (PLN) atau membayar tagihan utilitas pascabayar secara langsung."
          },
          {
            "title": "Tunggakan Utilitas Saat Keluar",
            "body": "Jika Penyewa mengosongkan lapak dengan meninggalkan tunggakan tagihan utilitas atau iuran kebersihan lokal, Pemilik berhak memotong tepat sejumlah tunggakan tersebut dari deposit jaminan escrow Penyewa saat keluar."
          }
        ]
      },
      {
        "id": "cleanliness-eviction",
        "number": "7",
        "title": "Kebersihan Lapak, Barang Tertinggal & Pengosongan",
        "subsections": [
          {
            "title": "Kewajiban Kebersihan",
            "body": "Penyewa bertanggung jawab penuh untuk mengosongkan seluruh barang pribadi dan inventaris saat keluar. Pemilik bertanggung jawab menyajikan ruang yang bersih bagi penyewa baru."
          },
          {
            "title": "Reaktivasi Listing Manual",
            "body": "Lapak aktif atau tertunda secara otomatis disembunyikan dari pasar. Setelah penyewa keluar atau kontrak dibatalkan, lapak TIDAK otomatis muncul kembali. Pemilik bertanggung jawab penuh untuk mengaktifkan/menerbitkan kembali listing secara manual setelah ruang fisik bersih dan siap untuk peninjauan baru."
          },
          {
            "title": "Barang yang Ditinggalkan",
            "body": "Barang-barang yang ditinggalkan oleh penyewa yang dihentikan atau keluar setelah masa sewa dapat dibuang, disimpan, atau dibersihkan oleh Pemilik Lapak atas diskresi penuh mereka. Lapakita tidak bertanggung jawab atas barang yang ditinggalkan."
          }
        ]
      },
      {
        "id": "supplier-disputes",
        "number": "8",
        "title": "Pasar Supplier & Sengketa B2B",
        "subsections": [
          {
            "title": "Transaksi Peer-to-Peer",
            "body": "Pasar Supplier B2B menghubungkan Penyewa secara langsung dengan Supplier."
          },
          {
            "title": "Penanganan Sengketa",
            "body": "Lapakita tidak menyediakan arbitrasi admin manual untuk keluhan produk supplier (misalnya salah bahan, keterlambatan pengiriman, cacat stok ringan). Pembeli dan Supplier harus menyelesaikan masalah melalui obrolan langsung. Pembeli mempertahankan hak penuh untuk memberikan rating bintang dan ulasan publik pada katalog produk dan profil supplier."
          }
        ]
      },
      {
        "id": "payouts",
        "number": "9",
        "title": "Pencairan Dana & Persyaratan Rekening Bank",
        "subsections": [
          {
            "body": "Pemilik dan Supplier wajib mendaftarkan rekening bank yang valid untuk pencairan dana otomatis. Penyewa wajib mendaftarkan rekening bank yang valid untuk menerima potensi pengembalian deposit."
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
            "title": "Account Identity",
            "body": "Full name, email address, phone number (WhatsApp), and encrypted password."
          },
          {
            "title": "Verification Data (KYC)",
            "body": "ID card (KTP) photo, OCR data, and a selfie with KTP — collected prior to lease signing or supplier activation for legal identity binding."
          },
          {
            "title": "Financial & Payout Data",
            "body": "Bank account name, bank name, and account number for automated payment routing via payment gateway APIs."
          },
          {
            "title": "Operational & Transactional Data",
            "body": "POS sales entries, stock levels, item prices, rental payment history, chat messages, uploaded property photos, and review ratings."
          }
        ]
      },
      {
        "id": "data-usage",
        "number": "2",
        "title": "How We Use Your Data",
        "subsections": [
          {
            "body": "To facilitate digital lease contracts, billing, and automated payout transfers."
          },
          {
            "body": "To display B2B supplier catalogs to relevant tenant business categories."
          },
          {
            "body": "To generate aggregated, non-personally identifiable Business Intelligence (BI) statistics — e.g. area demand trends, average foot traffic scores."
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
            "body": "Users may request account deactivation and data erasure, provided there are no active binding lease contracts, pending escrow deposit claims, or unfulfilled B2B transactions associated with the account."
          }
        ]
      }
    ]'::jsonb
),
(
    gen_random_uuid(),
    'privacy',
    'id',
    'Kebijakan Privasi',
    'Kebijakan ini menjelaskan data apa yang dikumpulkan Lapakita, bagaimana data tersebut digunakan, dan perlindungan yang diterapkan di seluruh akun Penyewa, Pemilik, dan Supplier.',
    '[
      {
        "id": "data-collected",
        "number": "1",
        "title": "Data yang Kami Kumpulkan",
        "subsections": [
          {
            "title": "Identitas Akun",
            "body": "Nama lengkap, alamat email, nomor telepon (WhatsApp), dan kata sandi terenkripsi."
          },
          {
            "title": "Data Verifikasi (KYC)",
            "body": "Foto kartu identitas (KTP), data OCR, dan foto selfie dengan KTP — dikumpulkan sebelum penandatanganan sewa atau aktivasi supplier untuk pengikatan identitas hukum."
          },
          {
            "title": "Data Keuangan & Pencarian Dana",
            "body": "Nama pemilik rekening bank, nama bank, dan nomor rekening untuk perutean pembayaran otomatis via API payment gateway."
          },
          {
            "title": "Data Operasional & Transaksi",
            "body": "Entri penjualan POS, tingkat stok, harga barang, riwayat pembayaran sewa, pesan obrolan, foto properti yang diunggah, dan rating ulasan."
          }
        ]
      },
      {
        "id": "data-usage",
        "number": "2",
        "title": "Cara Kami Menggunakan Data Anda",
        "subsections": [
          {
            "body": "Untuk memfasilitasi kontrak sewa digital, penagihan, dan transfer pencairan dana otomatis."
          },
          {
            "body": "Untuk menampilkan katalog supplier B2B ke kategori bisnis penyewa yang relevan."
          },
          {
            "body": "Untuk menghasilkan statistik Business Intelligence (BI) anonim yang diagregasi — misalnya tren permintaan area, skor lalu lintas pejalan kaki rata-rata."
          },
          {
            "body": "Untuk memverifikasi identitas dalam kasus sengketa sewa hukum atau banding deposit."
          }
        ]
      },
      {
        "id": "data-protection",
        "number": "3",
        "title": "Perlindungan Data & Kerahasiaan",
        "subsections": [
          {
            "title": "Tidak Ada Penjualan Data",
            "body": "Lapakita secara tegas tidak pernah menjual, menyewakan, atau memperdagangkan data pribadi pengguna, pendapatan bisnis, atau log transaksi pribadi kepada pengiklan pihak ketiga atau broker data."
          },
          {
            "title": "Kerahasiaan Data Pendapatan",
            "body": "Angka pendapatan penyewa individu dan buku kas POS bersifat privat untuk akun bisnis penyewa. Pemilik Lapak tidak dapat melihat pendapatan kotor atau margin keuntungan penyewa secara pasti."
          },
          {
            "title": "Infrastruktur Aman",
            "body": "Semua data muatan sensitif, token API, dan kredensial dienkripsi menggunakan protokol standar industri (TLS/SSL) dan disimpan secara aman."
          }
        ]
      },
      {
        "id": "data-retention",
        "number": "4",
        "title": "Retensi Data & Hak Pengguna",
        "subsections": [
          {
            "body": "Pengguna dapat mengajukan penonaktifan akun dan penghapusan data, dengan syarat tidak ada kontrak sewa mengikat yang aktif, klaim deposit escrow yang tertunda, atau transaksi B2B yang belum terpenuhi yang terkait dengan akun tersebut."
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
            "title": "Active Role Preference",
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
),
(
    gen_random_uuid(),
    'cookies',
    'id',
    'Kebijakan Cookies & Penyimpanan Lokal',
    'Lapakita menggunakan penyimpanan browser minimal — murni untuk fungsionalitas esensial, tidak pernah untuk pelacakan invasif atau retargeting iklan.',
    '[
      {
        "id": "what-we-store",
        "number": "1",
        "title": "Apa yang Kami Simpan",
        "subsections": [
          {
            "title": "Status Sesi",
            "body": "Menjaga Anda tetap masuk secara aman saat memuat ulang halaman."
          },
          {
            "title": "Preferensi Peran Aktif",
            "body": "Mengingat apakah Anda terakhir beroperasi dalam mode Penyewa, Pemilik, atau Supplier."
          },
          {
            "title": "Cache POS",
            "body": "Menyimpan sementara item keranjang POS dan daftar produk secara lokal agar antarmuka kasir Anda tetap cepat dan responsif bahkan saat terjadi gangguan jaringan ringan."
          }
        ]
      },
      {
        "id": "third-party-cookies",
        "number": "2",
        "title": "Cookies Pihak Ketiga",
        "subsections": [
          {
            "body": "Kami tidak menggunakan cookies pelacak pihak ketiga yang invasif, skrip pelacak perilaku lintas situs, atau piksel retargeting iklan. Skrip pihak ketiga murni dibatasi pada iFrame Payment Gateway aman (Midtrans/Xendit) untuk pemrosesan pembayaran."
          }
        ]
      }
    ]'::jsonb
)
ON CONFLICT (doc_type, lang) DO UPDATE 
SET title = EXCLUDED.title, intro = EXCLUDED.intro, sections_json = EXCLUDED.sections_json, updated_at = CURRENT_TIMESTAMP;