CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Clean existing data
TRUNCATE TABLE cms_public_faqs;

-- =============================================================================
-- 1. GENERAL & PLATFORM (role_type: 'all')
-- =============================================================================
-- English (en)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'en',
        'Platform Fundamentals',
        'What is Lapakita?',
        'Lapakita is an all-in-one digital operating platform designed specifically for micro, small, and medium enterprises (SMEs/UMKM). It unifies physical stall rentals across permanent, semi-permanent, and temporary bazaar spaces, Point of Sale (POS) operations, financial business analytics, and a B2B supplier marketplace into a single ecosystem.',
        'all',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'Platform Fundamentals',
        'How does the single account multi-role system work?',
        'You only need one email and phone number to register. From your profile menu, you can toggle between Tenant, Stall Owner, and Supplier modes. Each role has its own isolated dashboard, settings, and workflows, eliminating the need for multiple accounts.',
        'all',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'Platform Fundamentals',
        'Is Lapakita a mobile app or web platform?',
        'Lapakita is built as a responsive web application optimized for smartphones, tablets, and desktop computers. Cashiers can easily process transactions on mobile phones while property owners manage portfolios on desktops.',
        'all',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'Pricing, Plans & Subscriptions',
        'Is Lapakita free to use?',
        'Yes! Basic stall discovery, lease applications, POS cashier operations, inventory management, and diagnostic health overviews are completely free for all users.',
        'all',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'Pricing, Plans & Subscriptions',
        'What is the difference between Single-Role and All-Access Plans?',
        'Our Single-Role Plan is Rp 55,000/month (or Rp 495,000/year) and unlocks premium features for just one role (e.g. Tenant only). The All-Access Ecosystem Bundle is Rp 125,000/month (or Rp 1,125,000/year) and unlocks premium analytics across all three roles simultaneously.',
        'all',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'Pricing, Plans & Subscriptions',
        'How does Lapakita earn money if base features are free?',
        'We maintain a transparent business model. Revenue comes from optional Premium Subscriptions, a small percentage transaction fee on active stall rent collection, and lightweight administrative processing fees on supplier orders.',
        'all',
        6
    ),
    (
        gen_random_uuid (),
        'en',
        'Account Security & Verification',
        'Do I need to upload an ID card (KTP) during registration?',
        'No. Registration requires only your name, email, phone number, and password. ID verification (KYC) is requested gradually — only when you are ready to sign a binding lease contract, list a stall, or register as a supplier.',
        'all',
        7
    ),
    (
        gen_random_uuid (),
        'en',
        'Account Security & Verification',
        'Is my business revenue data kept private?',
        'We enforce strict data privacy. Your POS sales ledgers and financial profits are strictly confidential to your business account. Stall owners cannot view your net revenue, and data is never sold to third-party advertisers.',
        'all',
        8
    );

-- Indonesian (id)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'id',
        'Dasar Platform',
        'Apa itu Lapakita?',
        'Lapakita adalah platform operasional digital serba ada yang dirancang khusus untuk Usaha Mikro, Kecil, dan Menengah (UMKM). Platform ini menyatukan penyewaan lapak fisik (permanen, semi-permanen, dan bazaar temporer), sistem Kasir Point of Sale (POS), analitik bisnis keuangan, dan marketplace supplier B2B ke dalam satu ekosistem terpadu.',
        'all',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'Dasar Platform',
        'Bagaimana cara kerja sistem satu akun multi-peran?',
        'Anda hanya membutuhkan satu email dan nomor telepon untuk mendaftar. Dari menu profil, Anda dapat beralih antara mode Penyewa, Pemilik Lapak, dan Supplier. Setiap peran memiliki dashboard, pengaturan, dan alur kerja mandiri, tanpa perlu membuat banyak akun.',
        'all',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'Dasar Platform',
        'Apakah Lapakita berupa aplikasi mobile atau platform web?',
        'Lapakita dibangun sebagai aplikasi web responsif yang dioptimalkan untuk ponsel pintar, tablet, dan komputer desktop. Kasir dapat dengan mudah memproses transaksi melalui ponsel, sementara pemilik properti mengelola portofolio di desktop.',
        'all',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'Harga, Paket & Berlangganan',
        'Apakah Lapakita gratis digunakan?',
        'Ya! Fitur pencarian lapak dasar, pengajuan sewa, operasi kasir POS, manajemen inventaris, dan ringkasan kesehatan bisnis diagnosis dapat digunakan 100% gratis oleh semua pengguna.',
        'all',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'Harga, Paket & Berlangganan',
        'Apa perbedaan antara Paket Single-Role dan All-Access?',
        'Paket Single-Role seharga Rp 55.000/bulan (atau Rp 495.000/tahun) membuka fitur premium untuk satu peran saja (misal: Penyewa saja). Paket All-Access Bundle seharga Rp 125.000/bulan (atau Rp 1.125.000/tahun) membuka analitik premium di ketiga peran sekaligus secara bersamaan.',
        'all',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'Harga, Paket & Berlangganan',
        'Bagaimana Lapakita menghasilkan uang jika fitur dasar gratis?',
        'Kami menerapkan model bisnis yang transparan. Pendapatan berasal dari Langganan Premium opsional, persentase biaya transaksi ringan pada penagihan sewa lapak aktif, serta biaya pemrosesan administratif pada pesanan supplier.',
        'all',
        6
    ),
    (
        gen_random_uuid (),
        'id',
        'Keamanan Akun & Verifikasi',
        'Apakah saya wajib mengunggah KTP saat pendaftaran?',
        'Tidak. Pendaftaran hanya membutuhkan nama, email, nomor telepon, dan kata sandi. Verifikasi identitas (KYC) diminta secara bertahap — hanya saat Anda siap menandatangani kontrak sewa, memublikasikan lapak, atau mendaftar sebagai supplier.',
        'all',
        7
    ),
    (
        gen_random_uuid (),
        'id',
        'Keamanan Akun & Verifikasi',
        'Apakah data pendapatan bisnis saya dijaga kerahasiaannya?',
        'Kami menerapkan privasi data yang ketat. Pembukuan penjualan POS dan keuntungan finansial Anda bersifat rahasia untuk akun bisnis Anda. Pemilik lapak tidak dapat melihat pendapatan bersih Anda, dan data tidak pernah dijual ke pengiklan pihak ketiga.',
        'all',
        8
    );

-- =============================================================================
-- 2. TENANT & BUSINESS OPERATIONS (role_type: 'tenant')
-- =============================================================================
-- English (en)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'en',
        'POS & Non-Stall Business Operations',
        'Can I use Lapakita if I don''t rent a physical stall from the platform?',
        'Yes! Home-based businesses, cloud kitchens, online stores, or businesses renting spaces elsewhere can fully use our POS cashier, product/inventory management, staff access, and supplier marketplace.',
        'tenant',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'POS & Non-Stall Business Operations',
        'How do restricted cashier accounts work?',
        'You can create staff credentials for your cashiers. Staff members can process customer orders, issue receipts, and log cash/QRIS payments, but they cannot access financial profit reports or cost breakdowns.',
        'tenant',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'Stall Search, Permanence & Target ROI Filter',
        'How do the Operational Permanence Search Tabs work?',
        'You can filter spaces by 3 operational levels: Permanent (standalone shophouses with 24/7 access & sqm physical specs), Semi-Permanent (mall shops, food courts, and traditional market stalls bound by parent complex operating hours), and Temporary (short-term bazaar booths, food truck spots, and street vendor spots).',
        'tenant',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'Stall Search, Permanence & Target ROI Filter',
        'How does the Landmark & Radius search work?',
        'You can search by city or street, or pair a specific landmark (such as a university, school, or office complex) with a custom radius distance (e.g., within 3 km) to find nearby available stalls.',
        'tenant',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'Stall Search, Permanence & Target ROI Filter',
        'What is the Budget & Target ROI Match filter?',
        'For Permanent & Semi-Permanent stalls, input your total capital, business preset, and target BEP months to filter mathematically viable rents. For Temporary Bazaar spots, input your Target Daily Revenue to surface daily or monthly event rates that match your sales goals.',
        'tenant',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'Stall Search, Permanence & Target ROI Filter',
        'Can I inspect the stall before signing a lease?',
        'We highly recommend visiting the location in person to verify physical facilities, street access, and neighborhood conditions before submitting an application or signing the digital contract.',
        'tenant',
        6
    ),
    (
        gen_random_uuid (),
        'en',
        'Lease Contracts, Keys & Deposits',
        'What happens once the owner approves my lease request?',
        'A digital contract is created with price-locked terms. Once you make the initial rent payment and security deposit through the Payment Gateway on or before your selected start date, the lease becomes active.',
        'tenant',
        7
    ),
    (
        gen_random_uuid (),
        'en',
        'Lease Contracts, Keys & Deposits',
        'Where does my security deposit go?',
        'Your deposit is stored safely in a neutral Escrow Payment Gateway account — not in the owner''s personal bank account. It is fully refunded to your registered bank account upon lease completion, provided there are no unreturned key fees or physical property damages.',
        'tenant',
        8
    ),
    (
        gen_random_uuid (),
        'en',
        'Lease Contracts, Keys & Deposits',
        'How do lease rules work for Temporary Bazaar Events?',
        'Temporary spots do not use monthly lease terms. Instead, they define minimum lease days, event operating days (e.g., Everyday vs. Weekends Only), attendance requirements (Mandatory Full vs. Flexible), and clear cancellation policies (Pro-Rata, Deposit Refundable, or Strict Non-Refundable).',
        'tenant',
        9
    ),
    (
        gen_random_uuid (),
        'en',
        'Lease Contracts, Keys & Deposits',
        'What if I lose my physical keys during the lease?',
        'You are free to duplicate keys independently at local locksmiths. If all keys are lost, the owner replaces the lock cylinder; you pay strictly for the cost of the new key duplicated for your use.',
        'tenant',
        10
    ),
    (
        gen_random_uuid (),
        'en',
        'Financial Forecasts & Analytics',
        'How does the Multi-Timeline Business Forecast work?',
        'Our forecast engine projects your margins, cash flow, and break-even targets across 1-week, 1-month, 6-month, and 1-year timelines. It uses your POS sales history or custom financial presets across Conservative, Balanced, and Optimistic market scenarios.',
        'tenant',
        11
    ),
    (
        gen_random_uuid (),
        'en',
        'Financial Forecasts & Analytics',
        'What is the Prescriptive Operational Co-Pilot?',
        'It provides actionable recommendations based on your actual POS data — such as suggesting opening hour expansions, discount strategies for high-margin items, or automated restock reminders.',
        'tenant',
        12
    );

-- Indonesian (id)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'id',
        'Operasional POS & Bisnis Non-Lapak',
        'Bisakah saya menggunakan Lapakita jika tidak menyewa lapak fisik dari platform?',
        'Ya! Bisnis rumahan, cloud kitchen, toko online, atau bisnis yang menyewa ruang di tempat lain dapat menggunakan penuh kasir POS, manajemen produk/inventaris, akun akses staf, dan marketplace supplier kami.',
        'tenant',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'Operasional POS & Bisnis Non-Lapak',
        'Bagaimana cara kerja akun akses kasir terbatas?',
        'Anda dapat membuat kredensial staf khusus untuk kasir Anda. Anggota staf dapat memproses pesanan pelanggan, mencetak resi, dan mencatat pembayaran tunai/QRIS, tetapi tidak dapat mengakses laporan keuntungan keuangan atau rincian biaya modal.',
        'tenant',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'Pencarian Lapak, Permanensi & Filter Target ROI',
        'Bagaimana cara kerja Tab Pencarian Permanensi Operasional?',
        'Anda dapat memfilter ruang berdasarkan 3 tingkat operasional: Permanen (ruko independen dengan akses 24/7 & spesifikasi fisik m2), Semi-Permanen (toko mall, food court, dan lapak pasar tradisional yang terikat jam operasional komplek induk), dan Temporer (booth bazaar jangka pendek, spot food truck, dan lapak PKL).',
        'tenant',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'Pencarian Lapak, Permanensi & Filter Target ROI',
        'Bagaimana cara kerja pencarian Landmark & Radius?',
        'Anda dapat mencari berdasarkan kota atau jalan, atau memasangkan landmark tertentu (seperti universitas, sekolah, atau komplek perkantoran) dengan jarak radius kustom (misal: dalam radius 3 km) untuk menemukan lapak terdekat.',
        'tenant',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'Pencarian Lapak, Permanensi & Filter Target ROI',
        'Apa itu filter Anggaran & Cocokkan Target ROI?',
        'Untuk lapak Permanen & Semi-Permanen, masukkan total modal, preset jenis bisnis, dan target bulan BEP untuk memfilter harga sewa yang layak secara matematis. Untuk spot Bazaar Temporer, masukkan Target Pendapatan Harian untuk menampilkan tarif harian/bulanan yang cocok dengan target penjualan Anda.',
        'tenant',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'Pencarian Lapak, Permanensi & Filter Target ROI',
        'Bisakah saya mengecek lapak secara langsung sebelum menandatangani sewa?',
        'Kami sangat menyarankan untuk mengunjungi lokasi secara langsung untuk memverifikasi fasilitas fisik, akses jalan, dan kondisi lingkungan sebelum mengajukan sewa atau menandatangani kontrak digital.',
        'tenant',
        6
    ),
    (
        gen_random_uuid (),
        'id',
        'Kontrak Sewa, Kunci & Deposit Jaminan',
        'Apa yang terjadi setelah pemilik menyetujui pengajuan sewa saya?',
        'Kontrak digital akan dibuat dengan tarif yang terkunci. Setelah Anda melakukan pembayaran sewa awal dan deposit jaminan melalui Payment Gateway pada atau sebelum tanggal mulai yang dipilih, masa sewa resmi aktif.',
        'tenant',
        7
    ),
    (
        gen_random_uuid (),
        'id',
        'Kontrak Sewa, Kunci & Deposit Jaminan',
        'Ke mana dana deposit jaminan saya disimpan?',
        'Deposit Anda disimpan secara aman di rekening Escrow Payment Gateway yang netral — bukan di rekening pribadi pemilik lapak. Deposit akan dikembalikan penuh ke rekening bank Anda saat masa sewa berakhir, selama tidak ada denda kunci atau kerusakan fisik properti.',
        'tenant',
        8
    ),
    (
        gen_random_uuid (),
        'id',
        'Kontrak Sewa, Kunci & Deposit Jaminan',
        'Bagaimana aturan sewa berlaku untuk Event Bazaar Temporer?',
        'Spot temporer tidak menggunakan jangka waktu sewa bulanan. Sebagai gantinya, event menentukan minimal hari sewa, hari operasional (misal: Setiap Hari vs Akhir Pekan Saja), persyaratan kehadiran (Wajib Penuh vs Fleksibel), dan kebijakan pembatalan yang jelas (Pro-Rata, Deposit Refundable, atau Non-Refundable).',
        'tenant',
        9
    ),
    (
        gen_random_uuid (),
        'id',
        'Kontrak Sewa, Kunci & Deposit Jaminan',
        'Bagaimana jika kunci fisik saya hilang selama masa sewa?',
        'Anda bebas menduplikasi kunci secara mandiri di tukang kunci lokal. Jika seluruh kunci hilang, pemilik akan mengganti silinder kunci; Anda hanya membayar biaya pembuatan kunci baru yang diduplikasi untuk penggunaan Anda.',
        'tenant',
        10
    ),
    (
        gen_random_uuid (),
        'id',
        'Proyeksi Keuangan & Analitik',
        'Bagaimana cara kerja Proyeksi Bisnis Multi-Timeline?',
        'Mesin proyeksi kami memprediksi margin, arus kas, dan target BEP Anda dalam timeline 1 minggu, 1 bulan, 6 bulan, dan 1 tahun. Mesin ini menggunakan riwayat penjualan POS atau preset keuangan kustom dalam skenario pasar Konservatif, Seimbang, dan Optimis.',
        'tenant',
        11
    ),
    (
        gen_random_uuid (),
        'id',
        'Proyeksi Keuangan & Analitik',
        'Apa itu Co-Pilot Operasional Preskriptif?',
        'Fitur ini memberikan rekomendasi aksi berdasarkan data transaksi POS nyata Anda — seperti menyarankan perluasan jam buka, strategi diskon untuk item ber-margin tinggi, atau pengingat otomatis untuk stok ulang barang.',
        'tenant',
        12
    );

-- =============================================================================
-- 3. STALL OWNER OPERATIONS (role_type: 'owner')
-- =============================================================================
-- English (en)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'en',
        'Listing Stalls & Tenant Vetting',
        'How do I list a stall or event spot on Lapakita?',
        'From your Owner Dashboard, click ''Add Stall'', select the permanence type (Permanent, Semi-Permanent, or Temporary), fill in specific attributes (such as sqm size for Permanent, opening hours for Semi-Permanent, or event schedule & slots for Temporary), upload clear photos, and define rent rates.',
        'owner',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'Listing Stalls & Tenant Vetting',
        'Can I review applicants before accepting them?',
        'Yes. When a tenant applies, you can inspect their profile rating, previous rental reviews, and proposed business category before clicking ''Approve'' or ''Reject''.',
        'owner',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'Listing Stalls & Tenant Vetting',
        'Can I change the rent price during an active lease?',
        'No. Once a lease is approved and signed, the rent and deposit terms are locked for the duration of that contract. You may update listing prices for future tenancies once the current lease expires.',
        'owner',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'Payments, Overdues & Evictions',
        'How do I receive rent payouts?',
        'Rent payments made by tenants via the Payment Gateway are disbursed automatically to your registered bank account after platform fee deduction.',
        'owner',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'Payments, Overdues & Evictions',
        'What happens if a tenant is late on rent?',
        'The system tracks payment due dates. Once a deadline passes, a red Overdue badge appears on your dashboard. You retain full manual discretion to grant a grace period, issue reminders, or terminate the contract.',
        'owner',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'Payments, Overdues & Evictions',
        'How are non-paying tenants handled?',
        'Lapakita does not dispatch physical eviction teams. If a tenant fails to pay after grace periods, you can terminate the contract in-app and approach the stall directly to request move-out and reclaim your property.',
        'owner',
        6
    ),
    (
        gen_random_uuid (),
        'en',
        'Key Control & Property Damage Claims',
        'Do I need to install smart locks or QR access hardware?',
        'No expensive hardware is required. You can hand over physical keys directly to the tenant at the start of the lease.',
        'owner',
        7
    ),
    (
        gen_random_uuid (),
        'en',
        'Key Control & Property Damage Claims',
        'Should I replace the door lock between different tenants?',
        'We strongly recommend replacing the lock cylinder/knob set between tenancies for security hygiene. Stall Owners accept all security risks regarding potential duplicate keys if they choose to reuse old lock sets.',
        'owner',
        8
    ),
    (
        gen_random_uuid (),
        'en',
        'Key Control & Property Damage Claims',
        'How do I claim funds from the security deposit for damages?',
        'Upon tenant exit, submit a damage claim via your dashboard with itemized repair costs and timestamped photo evidence. Once verified or agreed upon by the tenant, the claim amount is disbursed from escrow to your bank account.',
        'owner',
        9
    ),
    (
        gen_random_uuid (),
        'en',
        'Property Portfolio Analytics',
        'What is the Vacancy Loss Tracker?',
        'It calculates the exact financial loss accumulating during vacant periods day by day, giving you clear visibility into overall portfolio performance.',
        'owner',
        10
    );

-- Indonesian (id)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'id',
        'Pendaftaran Lapak & Seleksi Penyewa',
        'Bagaimana cara memublikasikan lapak atau spot acara di Lapakita?',
        'Dari Dashboard Pemilik Anda, klik ''Tambah Lapak'', pilih jenis permanensi (Permanen, Semi-Permanen, atau Temporer), isi atribut spesifik (seperti ukuran m2 untuk Permanen, jam operasional untuk Semi-Permanen, atau jadwal & slot acara untuk Temporer), unggah foto yang jelas, dan tentukan tarif sewa.',
        'owner',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'Pendaftaran Lapak & Seleksi Penyewa',
        'Bisakah saya meninjau profil pendaftar sebelum menyetujuinya?',
        'Ya. Ketika penyewa mengajukan sewa, Anda dapat memeriksa rating profil mereka, ulasan sewa sebelumnya, dan kategori bisnis yang diajukan sebelum mengklik ''Setujui'' atau ''Tolak''.',
        'owner',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'Pendaftaran Lapak & Seleksi Penyewa',
        'Bisakah saya mengubah harga sewa selama masa sewa berlangsung?',
        'Tidak. Setelah kontrak sewa disetujui dan ditandatangani, ketentuan sewa dan deposit terkunci selama durasi kontrak tersebut. Anda dapat memperbarui harga listing untuk penyewaan di masa mendatang setelah kontrak saat ini berakhir.',
        'owner',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'Pembayaran, Tunggakan & Pengosongan',
        'Bagaimana cara saya menerima pencairan dana sewa?',
        'Pembayaran sewa yang dilakukan oleh penyewa melalui Payment Gateway dicairkan secara otomatis ke rekening bank terdaftar Anda setelah dipotong biaya platform.',
        'owner',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'Pembayaran, Tunggakan & Pengosongan',
        'Apa yang terjadi jika penyewa terlambat membayar sewa?',
        'Sistem melacak tanggal jatuh tempo pembayaran. Begitu batas waktu lewat, lencana Merah Terlambat muncul di dashboard Anda. Anda memiliki hak diskresi penuh untuk memberikan masa tenggang, mengirimkan pengingat, atau mengakhiri kontrak.',
        'owner',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'Pembayaran, Tunggakan & Pengosongan',
        'Bagaimana penanganan penyewa yang tidak membayar?',
        'Lapakita tidak menurunkan tim pengosongan fisik. Jika penyewa gagal membayar setelah masa tenggang, Anda dapat mengakhiri kontrak di aplikasi dan mendatangi lapak secara langsung untuk meminta pengosongan dan mengambil kembali properti Anda.',
        'owner',
        6
    ),
    (
        gen_random_uuid (),
        'id',
        'Kontrol Kunci & Klaim Kerusakan Properti',
        'Apakah saya perlu memasang kunci pintar (smart lock) atau perangkat QR?',
        'Tidak diperlukan perangkat keras yang mahal. Anda dapat menyerahkan kunci fisik secara langsung kepada penyewa pada awal masa sewa.',
        'owner',
        7
    ),
    (
        gen_random_uuid (),
        'id',
        'Kontrol Kunci & Klaim Kerusakan Properti',
        'Haruskah saya mengganti kunci pintu di antara penyewa yang berbeda?',
        'Kami sangat menyarankan untuk mengganti set silinder/knob kunci di antara penyewa yang berbeda demi kebersihan keamanan. Pemilik Lapak menanggung seluruh risiko keamanan terkait potensi kunci duplikat jika memilih untuk menggunakan kembali set kunci lama.',
        'owner',
        8
    ),
    (
        gen_random_uuid (),
        'id',
        'Kontrol Kunci & Klaim Kerusakan Properti',
        'Bagaimana cara mengklaim dana dari deposit jaminan untuk kerusakan?',
        'Saat penyewa keluar, ajukan klaim kerusakan melalui dashboard Anda beserta rincian biaya perbaikan dan bukti foto berstempel waktu. Setelah diverifikasi atau disetujui oleh penyewa, nominal klaim dicairkan dari escrow ke rekening bank Anda.',
        'owner',
        9
    ),
    (
        gen_random_uuid (),
        'id',
        'Analitik Portofolio Properti',
        'Apa itu Pelacak Kerugian Kekosongan (Vacancy Loss Tracker)?',
        'Fitur ini menghitung kerugian finansial tepat yang terakumulasi selama periode lapak kosong hari demi hari, memberikan transparansi jelas atas kinerja portofolio Anda.',
        'owner',
        10
    );

-- =============================================================================
-- 4. SUPPLIER & B2B MARKETPLACE (role_type: 'supplier')
-- =============================================================================
-- English (en)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'en',
        'Marketplace Matchmaking & Catalog Rules',
        'How do tenant businesses find my products?',
        'Lapakita automatically showcases your product catalog directly inside the procurement dashboard of active SME tenants whose business category matches your wholesale goods. No ad spend is required.',
        'supplier',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'Marketplace Matchmaking & Catalog Rules',
        'Can I set Minimum Order Quantities (MOQ) and tiered pricing?',
        'Yes. You can configure MOQ requirements per product and set volume-based discount tiers (e.g., 1-10 units at normal price, 10+ units at a discounted wholesale rate).',
        'supplier',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'Orders, Reorders & Delivery Disputes',
        'How does the 1-Click Reorder feature work?',
        'Tenants who tag you as their ''Primary Supplier'' can trigger instant reorder requests directly from their POS inventory alerts, giving you predictable recurring orders.',
        'supplier',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'Orders, Reorders & Delivery Disputes',
        'How are delivery disputes or product defect issues handled?',
        'Supplier transactions follow a direct peer-to-peer marketplace model. Buyers and suppliers resolve order discrepancies via direct chat. Buyers retain the right to leave public star ratings and product reviews on your catalog.',
        'supplier',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'Orders, Reorders & Delivery Disputes',
        'How do I receive payments for fulfilled orders?',
        'Payments processed through the checkout gateway are disbursed directly to your registered bank account upon order fulfillment confirmation.',
        'supplier',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'Demand & Supply Analytics',
        'What are Subscriber Demand Signals?',
        'It aggregates privacy-safe demand patterns from your subscriber base (e.g., predicted weekly coffee bean consumption) so you can optimize warehouse inventory before peak demand.',
        'supplier',
        6
    );

-- Indonesian (id)
INSERT INTO
    cms_public_faqs (
        id,
        lang,
        sub_topic_title,
        question,
        answer,
        role_type,
        sort_order
    )
VALUES
    (
        gen_random_uuid (),
        'id',
        'Pencocokan Marketplace & Aturan Katalog',
        'Bagaimana bisnis penyewa menemukan produk saya?',
        'Lapakita secara otomatis menampilkan katalog produk Anda secara langsung di dalam dashboard pengadaan milik penyewa UMKM aktif yang kategori bisnisnya cocok dengan barang grosir Anda. Tanpa perlu biaya iklan.',
        'supplier',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'Pencocokan Marketplace & Aturan Katalog',
        'Bisakah saya menentukan Minimal Pemesanan (MOQ) dan harga bertingkat?',
        'Ya. Anda dapat mengonfigurasi persyaratan MOQ per produk dan mengatur tingkatan diskon berbasis volume (misal: 1-10 unit harga normal, 10+ unit harga grosir terdiskon).',
        'supplier',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'Pesanan, Pemesanan Ulang & Sengketa Pengiriman',
        'Bagaimana cara kerja fitur Pemesanan Ulang 1-Klik?',
        'Penyewa yang menandai Anda sebagai ''Supplier Utama'' dapat memicu permintaan pemesanan ulang instan langsung dari peringatan inventaris POS mereka, memberikan Anda pesanan berulang yang terprediksi.',
        'supplier',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'Pesanan, Pemesanan Ulang & Sengketa Pengiriman',
        'Bagaimana sengketa pengiriman atau masalah cacat produk ditangani?',
        'Transaksi supplier mengikuti model marketplace peer-to-peer langsung. Pembeli dan supplier menyelesaikan ketidaksesuaian pesanan melalui chat langsung. Pembeli tetap berhak memberikan rating bintang publik dan ulasan pada katalog produk Anda.',
        'supplier',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'Pesanan, Pemesanan Ulang & Sengketa Pengiriman',
        'Bagaimana cara saya menerima pembayaran untuk pesanan yang telah dipenuhi?',
        'Pembayaran yang diproses melalui gateway checkout dicairkan langsung ke rekening bank terdaftar Anda setelah konfirmasi pemenuhan pesanan.',
        'supplier',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'Analitik Permintaan & Penawaran',
        'Apa itu Sinyal Permintaan Pelanggan (Subscriber Demand Signals)?',
        'Fitur ini mengagregasi pola permintaan yang aman secara privasi dari basis pelanggan Anda (misal: prediksi konsumsi biji kopi mingguan) sehingga Anda dapat mengoptimalkan stok gudang sebelum puncak permintaan.',
        'supplier',
        6
    );