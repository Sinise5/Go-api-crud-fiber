Fiber PostgreSQL Authentication & CRUD API
===
## 📌 Deskripsi Proyek
Proyek ini adalah RESTful API berbasis **Golang dengan Fiber** sebagai framework web yang ringan dan cepat. API ini terhubung dengan **PostgreSQL** sebagai database utama dan menerapkan **JWT Authentication** untuk keamanan. API ini dibuat untuk kebutuhan aplikasi yang memerlukan fitur **registrasi, login, dan manajemen pengguna (CRUD)** dengan standar keamanan yang baik.

## 📌 Fitur Utama
✅ **Autentikasi dengan JWT** → Login & register dengan hashing password menggunakan **bcrypt**.  
✅ **Proteksi Endpoint** → Middleware JWT untuk membatasi akses hanya kepada pengguna yang telah login.  
✅ **CRUD User** → Mendukung **Create, Read, Update, Delete** untuk data pengguna.  
✅ **Struktur Proyek yang Terorganisir** → Menggunakan **Object-Oriented Programming (OOP)** dan pemisahan concern.  

## 📌 Teknologi yang Digunakan
- **Golang** → Bahasa pemrograman utama.
- **Fiber** → Framework web yang ringan dan cepat.
- **PostgreSQL** → Database relational untuk menyimpan data user.
- **JWT (JSON Web Token)** → Untuk autentikasi dan otorisasi pengguna.
- **bcrypt** → Untuk hashing password agar lebih aman.

## Install & Dependence

- go mod init myapp
- go get github.com/gofiber/fiber/v2
- go get github.com/gofiber/jwt/v3
- go get github.com/gofiber/basicauth
- go get github.com/jmoiron/sqlx
- go get github.com/lib/pq
- go get golang.org/x/crypto/bcrypt
- go get github.com/golang-jwt/jwt/v5

## Use
- for test
  ```
  go run main.go
  ```
## 📌 Endpoint API

## API Endpoints

| Method | Endpoint       | Deskripsi              | Authentication |
|--------|--------------|----------------------|---------------|
| **POST**   | `/auth/register` | Register User        | ❌ Tidak perlu |
| **POST**   | `/auth/login`    | Login & Dapatkan Token | ❌ Tidak perlu |
| **GET**    | `/users`         | Get List User        | ✅ Perlu Token |
| **PUT**    | `/users/:id`     | Update User         | ✅ Perlu Token |
| **DELETE** | `/users/:id`     | Hapus User         | ✅ Perlu Token |


## Directory Hierarchy
```
|—— config
|    |—— database.go
|—— controllers
|    |—— authController.go
|    |—— userController.go
|—— go.mod
|—— go.sum
|—— main.go
|—— middlewares
|    |—— jwt.go
|—— models
|    |—— user.go
|—— routes
|    |—— auth.go
|    |—— user.go
```

## 📌 Keamanan API
✔ Menggunakan bcrypt untuk hashing password, sehingga data tidak tersimpan dalam bentuk teks biasa.
✔ Token JWT memiliki expiration time untuk membatasi waktu sesi pengguna.
✔ Endpoint CRUD dilindungi oleh middleware JWT, sehingga hanya pengguna yang memiliki token valid yang bisa mengaksesnya.



## References
- [paper-1]()
- [paper-2]()
- [code-1](https://github.com)
- [code-2](https://github.com)
  
## License

## Citing
If you use xxx,please use the following BibTeX entry.
```
```
