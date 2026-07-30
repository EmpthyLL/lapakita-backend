# Lapakita Backend

Backend RESTful API untuk platform manajemen dan penyewaan lapak UMKM berbasis **Go**, **Gin**, **GORM**, dan **Google Wire**. Project ini menerapkan arsitektur **Clean Architecture (Modular/Feature-based)** untuk memastikan keterbacaan kode, kemudahan pengujian, dan skalabilitas tinggi.

---

## 🛠️ Tech Stack & Libraries

- **Language:** Go (Golang) 1.22+
- **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database & ORM:** PostgreSQL & [GORM](https://gorm.io/)
- **Dependency Injection:** [Google Wire](https://github.com/google/wire)
- **Configuration Management:** [Viper](https://github.com/spf13/viper)
- **Logging:** [Uber Zap Logger](https://github.com/uber-go/zap)
- **Database Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **Live Reload Development:** [Air](https://github.com/air-verse/air)
- **Third-Party Integrations:**
  - **Firebase:** Push Notification (FCM) & Analytics
  - **Storage:** ImageKit SDK v2 & MinIO (Object Storage)
  - **Payment Gateway:** Midtrans Snap SDK

---

## 📁 Struktur Project

Aplikasi ini menggunakan pendekatan **Modular / Feature-Based Clean Architecture**:

```text
lapakita-backend/
├── cmd/
│   └── api/
│       ├── main.go            # Entry point aplikasi
│       ├── server.go          # Inisialisasi HTTP Server & Gin Router
│       ├── wire.go            # Dependency Injection Spec (Wire)
│       └── wire_gen.go        # Auto-generated code oleh Wire
├── config/
│   └── config.go              # Config struct & Viper loader
├── internal/
│   ├── entity/                # Domain Entities / GORM Models
│   └── feature/               # Modul / Fitur berbasis folder
│       └── [feature_name]/    # Contoh: product, auth, order
│           ├── dto/           # Data Transfer Objects
│           ├── handler/       # Delivery Layer (Gin Handlers)
│           ├── repository/    # Data Access Layer (GORM)
│           └── usecase/       # Business Logic Layer
├── pkg/                       # Shared Services & External Integrations
│   ├── database/              # DB Connection (PostgreSQL)
│   ├── firebase/              # FCM & Firebase Analytics
│   ├── logger/                # Zap Logger Setup
│   ├── payment/               # Midtrans Integration
│   └── storage/               # ImageKit & MinIO Storage Drivers
├── migrations/                # File Migrasi SQL
├── .env.example               # Template Environment Variable
├── Dockerfile                 # Multi-stage Docker Build
├── docker-compose.yml         # Container Setup (App & Postgres)
├── Makefile                   # Shortcut commands
└── .air.toml                  # Konfigurasi Air (Hot Reload)
```
