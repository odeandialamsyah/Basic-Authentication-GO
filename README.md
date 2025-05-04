# Basic Authentication API (Go + Gin + GORM)

Proyek ini adalah implementasi API autentikasi sederhana menggunakan **Go**, **Gin Framework**, **GORM**, dan **MySQL**. Proyek mencakup fitur registrasi pengguna, login, manajemen role ("admin" dan "user"), serta autentikasi berbasis token JWT.

---

## Fitur Utama

1. **Registrasi Pengguna**
   - Pengguna dapat mendaftar dengan username, email, dan password.
   - Role default untuk pengguna baru adalah `"user"`.

2. **Login Pengguna**
   - Pengguna dapat login menggunakan email dan password.
   - Setelah login berhasil, pengguna akan mendapatkan token JWT untuk autentikasi.

3. **Manajemen Role**
   - Dukungan untuk dua role: `"admin"` dan `"user"`.
   - Role ditetapkan secara otomatis saat registrasi (default: `"user"`).

4. **Autentikasi Berbasis Token (JWT)**
   - Token JWT digunakan untuk mengamankan endpoint yang memerlukan autentikasi.
   - Middleware memastikan hanya pengguna dengan token valid yang dapat mengakses endpoint tertentu.

5. **Role-Based Access Control (RBAC)**
   - Beberapa endpoint hanya dapat diakses oleh role tertentu (misalnya, admin-only routes).

---

## Teknologi yang Digunakan

- **Bahasa Pemrograman**: Go
- **Framework**: Gin Framework
- **ORM**: GORM
- **Database**: MySQL
- **Autentikasi**: JWT (JSON Web Token)
- **Tools Lain**: bcrypt (password hashing), godotenv (environment variables)
