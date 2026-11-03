# Golang DB

Project latihan belajar **database di Golang** memakai paket `database/sql` dan driver [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql). Contohnya dijalankan sebagai fungsi test dan terhubung ke MySQL atau MariaDB (misalnya dari XAMPP).

## Struktur Project

```
golang-db/
├── go.mod
├── go.sum
├── database.go        # Fungsi GetConnection() dengan pengaturan connection pool
├── database_test.go   # Test koneksi dasar ke database
└── sql_test.go        # Contoh INSERT (Exec) dan SELECT (Query)
```

## Prasyarat

- [Go](https://go.dev/dl/) versi 1.21 atau lebih baru
- MySQL atau MariaDB yang sedang berjalan di `localhost:3306`

## Menyiapkan Database

Nyalakan server database, lalu masuk lewat Terminal. Contoh untuk XAMPP di macOS:

```bash
/Applications/XAMPP/xamppfiles/bin/mysql -u root
```

Buat database dan tabel yang dipakai project ini:

```sql
CREATE DATABASE golang_db;
USE golang_db;

CREATE TABLE customer (
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;
```

Setiap perintah SQL harus diakhiri tanda titik koma (`;`).

## Konfigurasi Koneksi

String koneksi ada di `database.go` dan `database_test.go`:

```
root:@tcp(localhost:3306)/golang_db
```

Formatnya `user:password@tcp(host:port)/nama_database`. Pada contoh ini user `root` tanpa password. Sesuaikan jika pengaturan database kamu berbeda.

## Instalasi

```bash
git clone <url-repository-kamu>
cd golang-db
go mod tidy
```

`go mod tidy` mengunduh driver MySQL dan merapikan `go.mod`.

## Menjalankan Test

Semua test:

```bash
go test -v
```

Satu test tertentu:

```bash
go test -v -run TestConnect
go test -v -run TestExcecSql
go test -v -run TestQuerySql
```

## Isi Materi

| Berkas | Test | Penjelasan |
| --- | --- | --- |
| `database_test.go` | `TestConnect` | Membuka koneksi ke database dengan `sql.Open`. |
| `database.go` | `GetConnection` | Membuat koneksi dan mengatur connection pool: maksimum 100 koneksi terbuka, 10 koneksi idle, idle maksimal 5 menit, umur koneksi maksimal 60 menit. |
| `sql_test.go` | `TestExceSql` | Menyisipkan data ke tabel `customer` dengan `ExecContext`. |
| `sql_test.go` | `TestQuerySql` | Mengambil semua data dari tabel `customer` dengan `QueryContext`, lalu membacanya baris demi baris dengan `rows.Next()` dan `rows.Scan()`. |

## Catatan

- `TestExcecSql` menyisipkan data dengan `id` bernilai `eko`. Karena `id` adalah primary key, menjalankannya dua kali akan menghasilkan error *Duplicate entry*. Hapus datanya dulu dengan `DELETE FROM customer WHERE id = 'eko';` atau ganti nilai `id`-nya.
- Jalankan `TestExcecSql` sebelum `TestQuerySql` supaya ada data yang dibaca.
- Pastikan server MySQL atau MariaDB sudah berjalan sebelum menjalankan test. Kalau tidak, akan muncul error `connection refused`.
- Driver MySQL dipilih `v1.8.1` agar cocok dengan Go 1.21.

## Dependensi

- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) v1.8.1