CREATE EXTENSION IF NOT EXISTS "pgcrypto";

INSERT INTO
    cms_public_faqs (
        id,
        lang,
        category_id,
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
        'all',
        'Platform Fundamentals',
        'What is Lapakita?',
        'Lapakita is an all-in-one digital operating platform designed specifically for micro, small, and medium enterprises (SMEs/UMKM). It unifies physical stall rentals, Point of Sale (POS) operations, financial business analytics, and a B2B supplier marketplace into a single ecosystem.',
        'general',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Dasar-Dasar Platform',
        'Apa itu Lapakita?',
        'Lapakita adalah platform operasional digital serba ada yang dirancang khusus untuk usaha mikro, kecil, dan menengah (UMKM). Lapakita menyatukan sewa lapak fisik, operasional Point of Sale (POS), analitik keuangan bisnis, dan pasar supplier B2B dalam satu ekosistem.',
        'general',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'all',
        'Platform Fundamentals',
        'How does the single account multi-role system work?',
        'You only need one email and phone number to register. From your profile menu, you can toggle between Tenant, Stall Owner, and Supplier modes. Each role has its own isolated dashboard, settings, and workflows, eliminating the need for multiple accounts.',
        'general',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Dasar-Dasar Platform',
        'Bagaimana cara kerja sistem satu akun multi-peran?',
        'Anda hanya memerlukan satu email dan nomor telepon untuk mendaftar. Dari menu profil, Anda dapat beralih antara mode Penyewa, Pemilik Lapak, dan Supplier. Setiap peran memiliki dasbor, pengaturan, dan alur kerja terisolasi sendiri, sehingga tidak perlu membuat banyak akun.',
        'general',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'all',
        'Platform Fundamentals',
        'Is Lapakita a mobile app or web platform?',
        'Lapakita is built as a responsive web application optimized for smartphones, tablets, and desktop computers. Cashiers can easily process transactions on mobile phones while property owners manage portfolios on desktops.',
        'general',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Dasar-Dasar Platform',
        'Apakah Lapakita berupa aplikasi seluler atau platform web?',
        'Lapakita dibangun sebagai aplikasi web responsif yang dioptimalkan untuk ponsel pintar, tablet, dan komputer desktop. Kasir dapat dengan mudah memproses transaksi di ponsel sementara pemilik properti mengelola portofolio di desktop.',
        'general',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'all',
        'Pricing, Plans & Subscriptions',
        'Is Lapakita free to use?',
        'Yes! Basic stall discovery, lease applications, POS cashier operations, inventory management, and diagnostic health overviews are completely free for all users.',
        'general',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Harga, Paket & Langganan',
        'Apakah Lapakita gratis untuk digunakan?',
        'Ya! Pencarian lapak dasar, pengajuan sewa, operasional kasir POS, manajemen inventaris, dan tinjauan diagnostik kesehatan bisnis sepenuhnya gratis untuk semua pengguna.',
        'general',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'all',
        'Pricing, Plans & Subscriptions',
        'What is the difference between Single-Role and All-Access Plans?',
        'Our Single-Role Plan is Rp 55,000/month (or Rp 495,000/year) and unlocks premium features for just one role (e.g. Tenant only). The All-Access Ecosystem Bundle is Rp 125,000/month (or Rp 1,125,000/year) and unlocks premium analytics across all three roles simultaneously.',
        'general',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Harga, Paket & Langganan',
        'Apa perbedaan antara Paket Single-Role dan All-Access?',
        'Paket Single-Role kami seharga Rp 55.000/bulan (atau Rp 495.000/tahun) dan membuka fitur premium untuk satu peran saja (misal: Penyewa saja). Paket Bundel Ekosistem All-Access seharga Rp 125.000/bulan (atau Rp 1.125.000/tahun) dan membuka analitik premium di ketiga peran secara bersamaan.',
        'general',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'all',
        'Pricing, Plans & Subscriptions',
        'How does Lapakita earn money if base features are free?',
        'We maintain a transparent business model. Revenue comes from optional Premium Subscriptions, a small percentage transaction fee on active stall rent collection, and lightweight administrative processing fees on supplier orders.',
        'general',
        6
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Harga, Paket & Langganan',
        'Bagaimana Lapakita menghasilkan uang jika fitur dasar gratis?',
        'Kami mempertahankan model bisnis yang transparan. Pendapatan berasal dari Langganan Premium opsional, persentase kecil biaya transaksi pada penarikan sewa lapak aktif, dan biaya pemrosesan administratif ringan pada pesanan supplier.',
        'general',
        6
    ),
    (
        gen_random_uuid (),
        'en',
        'all',
        'Account Security & Verification',
        'Do I need to upload an ID card (KTP) during registration?',
        'No. Registration requires only your name, email, phone number, and password. ID verification (KYC) is requested gradually — only when you are ready to sign a binding lease contract, list a stall, or register as a supplier.',
        'general',
        7
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Keamanan Akun & Verifikasi',
        'Apakah saya perlu mengunggah KTP saat pendaftaran?',
        'Tidak. Pendaftaran hanya memerlukan nama, email, nomor telepon, dan kata sandi Anda. Verifikasi identitas (KYC) diminta secara bertahap — hanya saat Anda siap menandatangani kontrak sewa yang mengikat, mendaftarkan lapak, atau mendaftar sebagai supplier.',
        'general',
        7
    ),
    (
        gen_random_uuid (),
        'en',
        'all',
        'Account Security & Verification',
        'Is my business revenue data kept private?',
        'We enforce strict data privacy. Your POS sales ledgers and financial profits are strictly confidential to your business account. Stall owners cannot view your net revenue, and data is never sold to third-party advertisers.',
        'general',
        8
    ),
    (
        gen_random_uuid (),
        'id',
        'all',
        'Keamanan Akun & Verifikasi',
        'Apakah data pendapatan bisnis saya dijaga kerahasiaannya?',
        'Kami menerapkan privasi data yang ketat. Buku kas penjualan POS dan keuntungan finansial Anda bersifat sangat rahasia untuk akun bisnis Anda. Pemilik lapak tidak dapat melihat pendapatan bersih Anda, dan data tidak pernah dijual kepada pengiklan pihak ketiga.',
        'general',
        8
    );

INSERT INTO
    cms_public_faqs (
        id,
        lang,
        category_id,
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
        'tenant',
        'POS & Non-Stall Business Operations',
        'Can I use Lapakita if I don''t rent a physical stall from the platform?',
        'Yes! Home-based businesses, cloud kitchens, online stores, or businesses renting spaces elsewhere can fully use our POS cashier, product/inventory management, staff access, and supplier marketplace.',
        'tenant',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'POS & Operasional Bisnis Non-Lapak',
        'Bisakah saya menggunakan Lapakita jika saya tidak menyewa lapak fisik dari platform?',
        'Bisa! Bisnis rumahan, cloud kitchen, toko online, atau bisnis yang menyewa tempat di lokasi lain dapat sepenuhnya menggunakan kasir POS, manajemen produk/inventaris, akses staf, dan pasar supplier kami.',
        'tenant',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'POS & Non-Stall Business Operations',
        'How do restricted cashier accounts work?',
        'You can create staff credentials for your cashiers. Staff members can process customer orders, issue receipts, and log cash/QRIS payments, but they cannot access financial profit reports or cost breakdowns.',
        'tenant',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'POS & Operasional Bisnis Non-Lapak',
        'Bagaimana cara kerja akun kasir dengan akses terbatas?',
        'Anda dapat membuat kredensial staf untuk kasir Anda. Anggota staf dapat memproses pesanan pelanggan, menerbitkan resi, dan mencatat pembayaran tunai/QRIS, tetapi mereka tidak dapat mengakses laporan keuntungan finansial atau rincian biaya.',
        'tenant',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Stall Search, Landmarks & ROI Filter',
        'How does the Landmark & Radius search work?',
        'You can search by city or street, or pair a specific landmark (such as a university, school, or office complex) with a custom radius distance (e.g., within 3 km) to find nearby available stalls.',
        'tenant',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Pencarian Lapak, Landmark & Filter ROI',
        'Bagaimana cara kerja pencarian Landmark & Radius?',
        'Anda dapat mencari berdasarkan kota atau jalan, atau memasangkan landmark tertentu (seperti universitas, sekolah, atau kompleks perkantoran) dengan jarak radius kustom (misal: dalam radius 3 km) untuk menemukan lapak terdekat yang tersedia.',
        'tenant',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Stall Search, Landmarks & ROI Filter',
        'What is the Budget & Target ROI Match filter?',
        'Instead of guessing, you can input your total capital, business preset, and target break-even period (e.g., 6 months). Our search engine filters and highlights stalls with monthly rent prices that mathematically align with your budget goals.',
        'tenant',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Pencarian Lapak, Landmark & Filter ROI',
        'Apa itu filter Pencocokan Anggaran & Target ROI?',
        'Daripada mereka-reka, Anda dapat memasukkan total modal, preset bisnis, dan target periode impas/BEP (misal: 6 bulan). Mesin pencari kami menyaring dan menyoroti lapak dengan harga sewa bulanan yang secara matematis selaras dengan target anggaran Anda.',
        'tenant',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Stall Search, Landmarks & ROI Filter',
        'Can I inspect the stall before signing a lease?',
        'We highly recommend visiting the location in person to verify physical facilities, street access, and neighborhood conditions before submitting an application or signing the digital contract.',
        'tenant',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Pencarian Lapak, Landmark & Filter ROI',
        'Bisakah saya mensurvei lapak sebelum menandatangani sewa?',
        'Kami sangat menyarankan untuk mengunjungi lokasi secara langsung untuk memverifikasi fasilitas fisik, akses jalan, dan kondisi lingkungan sebelum mengajukan sewa atau menandatangani kontrak digital.',
        'tenant',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Lease Contracts, Keys & Deposits',
        'What happens once the owner approves my lease request?',
        'A digital contract is created with price-locked terms. Once you make the first month''s rent payment and security deposit through the Payment Gateway, the lease becomes active.',
        'tenant',
        6
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Kontrak Sewa, Kunci & Deposit',
        'Apa yang terjadi setelah pemilik menyetujui permintaan sewa saya?',
        'Kontrak digital dibuat dengan syarat harga yang dikunci. Setelah Anda melakukan pembayaran sewa bulan pertama dan deposit jaminan melalui Payment Gateway, masa sewa menjadi aktif.',
        'tenant',
        6
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Lease Contracts, Keys & Deposits',
        'Where does my security deposit go?',
        'Your deposit is stored safely in a neutral Escrow Payment Gateway account — not in the owner''s personal bank account. It is fully refunded to your registered bank account upon lease completion, provided there are no unreturned key fees or physical property damages.',
        'tenant',
        7
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Kontrak Sewa, Kunci & Deposit',
        'Ke mana dana deposit jaminan saya disimpan?',
        'Deposit Anda disimpan secara aman di rekening Escrow Payment Gateway yang netral — bukan di rekening bank pribadi pemilik. Deposit dikembalikan penuh ke rekening bank terdaftar Anda setelah masa sewa selesai, selama tidak ada biaya kunci yang tidak dikembalikan atau kerusakan fisik properti.',
        'tenant',
        7
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Lease Contracts, Keys & Deposits',
        'What if I lose my physical keys during the lease?',
        'You are free to duplicate keys independently at local locksmiths at your own expense. If all keys are lost, you can submit an in-app key replacement request to the owner; the key reproduction fee will be deducted from your deposit.',
        'tenant',
        8
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Kontrak Sewa, Kunci & Deposit',
        'Bagaimana jika saya kehilangan kunci fisik selama masa sewa?',
        'Anda bebas menduplikasi kunci secara mandiri di tukang kunci lokal dengan biaya sendiri. Jika seluruh kunci hilang, Anda dapat mengajukan permintaan penggantian kunci dalam aplikasi ke pemilik; biaya pembuatan kunci akan dipotong dari deposit Anda.',
        'tenant',
        8
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Financial Forecasts & Analytics',
        'How does the Multi-Timeline Business Forecast work?',
        'Our forecast engine projects your margins, cash flow, and break-even targets across 1-week, 1-month, 6-month, and 1-year timelines. It uses your POS sales history or custom financial presets across Conservative, Balanced, and Optimistic market scenarios.',
        'tenant',
        9
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Prakiraan Keuangan & Analitik',
        'Bagaimana cara kerja Proyeksi Bisnis Multi-Linimasa?',
        'Mesin proyeksi kami memproyeksikan margin, arus kas, dan target impas Anda dalam skala waktu 1 minggu, 1 bulan, 6 bulan, dan 1 tahun. Ini menggunakan riwayat penjualan POS Anda atau preset keuangan kustom dalam skenario pasar Konservatif, Seimbang, dan Optimis.',
        'tenant',
        9
    ),
    (
        gen_random_uuid (),
        'en',
        'tenant',
        'Financial Forecasts & Analytics',
        'What is the Prescriptive Operational Co-Pilot?',
        'It provides actionable recommendations based on your actual POS data — such as suggesting opening hour expansions, discount strategies for high-margin items, or automated restock reminders.',
        'tenant',
        10
    ),
    (
        gen_random_uuid (),
        'id',
        'tenant',
        'Prakiraan Keuangan & Analitik',
        'Apa itu Co-Pilot Operasional Preskriptif?',
        'Co-Pilot memberikan rekomendasi yang dapat ditindaklanjuti berdasarkan data POS aktual Anda — seperti menyarankan perpanjangan jam buka, strategi diskon untuk barang bermargin tinggi, atau pengingat stok ulang otomatis.',
        'tenant',
        10
    );

INSERT INTO
    cms_public_faqs (
        id,
        lang,
        category_id,
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
        'owner',
        'Listing Stalls & Tenant Vetting',
        'How do I list a stall on Lapakita?',
        'From your Owner Dashboard, click ''Add Stall'', upload clear photos of the space, select available facilities (electrical kVA, water, seating), set landmark locations, and define monthly rent and deposit amounts.',
        'owner',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Mendaftarkan Lapak & Seleksi Penyewa',
        'Bagaimana cara mendaftarkan lapak di Lapakita?',
        'Dari Dasbor Pemilik Anda, klik ''Tambah Lapak'', unggah foto ruangan yang jelas, pilih fasilitas yang tersedia (kVA listrik, air, tempat duduk), tentukan lokasi landmark, dan tetapkan biaya sewa bulanan serta jumlah deposit.',
        'owner',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Listing Stalls & Tenant Vetting',
        'Can I review applicants before accepting them?',
        'Yes. When a tenant applies, you can inspect their profile rating, previous rental reviews, and proposed business category before clicking ''Approve'' or ''Reject''.',
        'owner',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Mendaftarkan Lapak & Seleksi Penyewa',
        'Bisakah saya meninjau calon penyewa sebelum menerimanya?',
        'Bisa. Ketika penyewa mengajukan sewa, Anda dapat memeriksa rating profil mereka, ulasan sewa sebelumnya, dan kategori bisnis yang diajukan sebelum mengklik ''Setujui'' atau ''Tolak''.',
        'owner',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Listing Stalls & Tenant Vetting',
        'Can I change the rent price during an active lease?',
        'No. Once a lease is approved and signed, the rent and deposit terms are locked for the duration of that contract. You may update listing prices for future tenancies once the current lease expires.',
        'owner',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Mendaftarkan Lapak & Seleksi Penyewa',
        'Bisakah saya mengubah harga sewa selama masa sewa aktif?',
        'Tidak bisa. Setelah sewa disetujui dan ditandatangani, ketentuan sewa dan deposit dikunci selama durasi kontrak tersebut. Anda dapat memperbarui harga listing untuk penyewa di masa mendatang setelah sewa saat ini berakhir.',
        'owner',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Payments, Overdues & Evictions',
        'How do I receive monthly rent payouts?',
        'Rent payments made by tenants via the Payment Gateway are disbursed automatically to your registered bank account after platform fee deduction.',
        'owner',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Pembayaran, Keterlambatan & Pengosongan',
        'Bagaimana cara saya menerima pencairan sewa bulanan?',
        'Pembayaran sewa yang dilakukan oleh penyewa melalui Payment Gateway dicairkan secara otomatis ke rekening bank terdaftar Anda setelah dipotong biaya platform.',
        'owner',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Payments, Overdues & Evictions',
        'What happens if a tenant is late on rent?',
        'The system tracks payment due dates. Once a deadline passes, a red Overdue badge appears on your dashboard. You retain full manual discretion to grant a grace period, issue reminders, or terminate the contract.',
        'owner',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Pembayaran, Keterlambatan & Pengosongan',
        'Apa yang terjadi jika penyewa terlambat membayar sewa?',
        'Sistem melacak tenggat waktu pembayaran. Setelah tenggat lewat, lencana Terlambat berwarna merah muncul di dasbor Anda. Anda memiliki diskresi manual penuh untuk memberikan masa tenggang, menerbitkan pengingat, atau menghentikan kontrak.',
        'owner',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Payments, Overdues & Evictions',
        'How are non-paying tenants handled?',
        'Lapakita does not dispatch physical eviction teams. If a tenant fails to pay after grace periods, you can terminate the contract in-app and approach the stall directly to request move-out and reclaim your property.',
        'owner',
        6
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Pembayaran, Keterlambatan & Pengosongan',
        'Bagaimana penanganan penyewa yang tidak membayar?',
        'Lapakita tidak mengirimkan tim pengosongan fisik. Jika penyewa gagal membayar setelah masa tenggang, Anda dapat mengakhiri kontrak dalam aplikasi dan mendatangi lapak secara langsung untuk meminta pengosongan dan mengambil kembali properti Anda.',
        'owner',
        6
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Key Control & Property Damage Claims',
        'Do I need to install smart locks or QR access hardware?',
        'No expensive hardware is required. You can hand over physical keys directly to the tenant at the start of the lease.',
        'owner',
        7
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Kontrol Kunci & Klaim Kerusakan Properti',
        'Apakah saya perlu memasang kunci pintar atau perangkat akses QR?',
        'Tidak diperlukan perangkat keras mahal. Anda dapat menyerahkan kunci fisik secara langsung kepada penyewa pada awal masa sewa.',
        'owner',
        7
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Key Control & Property Damage Claims',
        'Should I replace the door lock between different tenants?',
        'We strongly recommend replacing the lock cylinder/knob set between tenancies for security hygiene. If you choose to reuse old lock sets, you accept inherent security risks regarding potential duplicate keys.',
        'owner',
        8
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Kontrol Kunci & Klaim Kerusakan Properti',
        'Haruskah saya mengganti kunci pintu di antara penyewa yang berbeda?',
        'Kami sangat menyarankan untuk mengganti silinder/gagang kunci di antara periode penyewa demi kebersihan keamanan. Jika Anda memilih untuk menggunakan kembali set kunci lama, Anda menanggung risiko keamanan terkait potensi kunci duplikat.',
        'owner',
        8
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Key Control & Property Damage Claims',
        'How do I claim funds from the security deposit for damages?',
        'Upon tenant exit, submit a damage claim via your dashboard with itemized repair costs and timestamped photo evidence. Once verified or agreed upon by the tenant, the claim amount is disbursed from escrow to your bank account.',
        'owner',
        9
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Kontrol Kunci & Klaim Kerusakan Properti',
        'Bagaimana cara saya mengklaim dana dari deposit jaminan atas kerusakan?',
        'Saat penyewa keluar, ajukan klaim kerusakan melalui dasbor Anda disertai rincian biaya perbaikan dan bukti foto berstempel waktu. Setelah diverifikasi atau disetujui oleh penyewa, jumlah klaim dicairkan dari escrow ke rekening bank Anda.',
        'owner',
        9
    ),
    (
        gen_random_uuid (),
        'en',
        'owner',
        'Property Portfolio Analytics',
        'What is the Vacancy Loss Tracker?',
        'It calculates the exact financial loss accumulating during vacant periods day by day, giving you clear visibility into overall portfolio performance.',
        'owner',
        10
    ),
    (
        gen_random_uuid (),
        'id',
        'owner',
        'Analitik Portofolio Properti',
        'Apa itu Pelacak Kerugian Kekosongan?',
        'Fitur ini menghitung kerugian finansial secara tepat yang terakumulasi selama periode lapak kosong hari demi hari, memberikan Anda visibilitas yang jelas tentang kinerja portofolio secara keseluruhan.',
        'owner',
        10
    );

INSERT INTO
    cms_public_faqs (
        id,
        lang,
        category_id,
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
        'supplier',
        'Marketplace Matchmaking & Catalog Rules',
        'How do tenant businesses find my products?',
        'Lapakita automatically showcases your product catalog directly inside the procurement dashboard of active SME tenants whose business category matches your wholesale goods. No ad spend is required.',
        'supplier',
        1
    ),
    (
        gen_random_uuid (),
        'id',
        'supplier',
        'Pencocokan Pasar & Aturan Katalog',
        'Bagaimana bisnis penyewa menemukan produk saya?',
        'Lapakita secara otomatis menampilkan katalog produk Anda langsung di dalam dasbor pengadaan penyewa UMKM aktif yang kategori bisnisnya cocok dengan barang grosir Anda. Tidak diperlukan biaya iklan.',
        'supplier',
        1
    ),
    (
        gen_random_uuid (),
        'en',
        'supplier',
        'Marketplace Matchmaking & Catalog Rules',
        'Can I set Minimum Order Quantities (MOQ) and tiered pricing?',
        'Yes. You can configure MOQ requirements per product and set volume-based discount tiers (e.g., 1-10 units at normal price, 10+ units at a discounted wholesale rate).',
        'supplier',
        2
    ),
    (
        gen_random_uuid (),
        'id',
        'supplier',
        'Pencocokan Pasar & Aturan Katalog',
        'Bisakah saya menetapkan Jumlah Pesanan Minimum (MOQ) dan harga berjenjang?',
        'Bisa. Anda dapat mengonfigurasi persyaratan MOQ per produk dan menetapkan tingkatan diskon berbasis volume (misal: 1-10 unit pada harga normal, 10+ unit pada harga grosir berdiskon).',
        'supplier',
        2
    ),
    (
        gen_random_uuid (),
        'en',
        'supplier',
        'Orders, Reorders & Delivery Disputes',
        'How does the 1-Click Reorder feature work?',
        'Tenants who tag you as their ''Primary Supplier'' can trigger instant reorder requests directly from their POS inventory alerts, giving you predictable recurring orders.',
        'supplier',
        3
    ),
    (
        gen_random_uuid (),
        'id',
        'supplier',
        'Pesanan, Pemesanan Ulang & Sengketa Pengiriman',
        'Bagaimana cara kerja fitur Pemesanan Ulang 1-Klik?',
        'Penyewa yang menandai Anda sebagai ''Supplier Utama'' dapat memicu permintaan pemesanan ulang instan langsung dari peringatan inventaris POS mereka, memberi Anda pesanan berulang yang terprediksi.',
        'supplier',
        3
    ),
    (
        gen_random_uuid (),
        'en',
        'supplier',
        'Orders, Reorders & Delivery Disputes',
        'How are delivery disputes or product defect issues handled?',
        'Supplier transactions follow a direct peer-to-peer marketplace model. Buyers and suppliers resolve order discrepancies via direct chat. Buyers retain the right to leave public star ratings and product reviews on your catalog.',
        'supplier',
        4
    ),
    (
        gen_random_uuid (),
        'id',
        'supplier',
        'Pesanan, Pemesanan Ulang & Sengketa Pengiriman',
        'Bagaimana penanganan sengketa pengiriman atau masalah cacat produk?',
        'Transaksi supplier mengikuti model pasar peer-to-peer langsung. Pembeli dan supplier menyelesaikan ketidaksesuaian pesanan melalui obrolan langsung. Pembeli mempertahankan hak untuk memberikan rating bintang dan ulasan produk secara publik pada katalog Anda.',
        'supplier',
        4
    ),
    (
        gen_random_uuid (),
        'en',
        'supplier',
        'Orders, Reorders & Delivery Disputes',
        'How do I receive payments for fulfilled orders?',
        'Payments processed through the checkout gateway are disbursed directly to your registered bank account upon order fulfillment confirmation.',
        'supplier',
        5
    ),
    (
        gen_random_uuid (),
        'id',
        'supplier',
        'Pesanan, Pemesanan Ulang & Sengketa Pengiriman',
        'Bagaimana cara saya menerima pembayaran untuk pesanan yang telah dipenuhi?',
        'Pembayaran yang diproses melalui checkout gateway dicairkan langsung ke rekening bank terdaftar Anda setelah konfirmasi pemenuhan pesanan.',
        'supplier',
        5
    ),
    (
        gen_random_uuid (),
        'en',
        'supplier',
        'Demand & Supply Analytics',
        'What are Subscriber Demand Signals?',
        'It aggregates privacy-safe demand patterns from your subscriber base (e.g., predicted weekly coffee bean consumption) so you can optimize warehouse inventory before peak demand.',
        'supplier',
        6
    ),
    (
        gen_random_uuid (),
        'id',
        'supplier',
        'Analitik Permintaan & Penawaran',
        'Apa itu Sinyal Permintaan Pelanggan?',
        'Fitur ini mengagregasi pola permintaan yang aman secara privasi dari basis pelanggan Anda (misal: prediksi konsumsi biji kopi mingguan) sehingga Anda dapat mengoptimalkan stok gudang sebelum permintaan puncak.',
        'supplier',
        6
    );