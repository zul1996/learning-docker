# My Go App

## Deskripsi

`My Go App` adalah contoh aplikasi Go sederhana yang dikemas dalam container Docker. Aplikasi ini menggunakan Docker untuk mempermudah distribusi dan penerapan (deployment).

## Struktur Proyek

- **Dockerfile**: File konfigurasi untuk membangun image Docker.
- **go.mod**: File modul Go yang berisi informasi tentang dependensi.
- **go.sum**: File yang melacak versi yang digunakan oleh dependensi.
- **main.go**: Kode sumber utama aplikasi Go.

## Prasyarat

- [Docker](https://www.docker.com/) harus diinstal di sistem Anda.
- [Go](https://golang.org/dl/) (untuk pengembangan lokal, tidak wajib jika hanya ingin membangun image Docker).

## Langkah-langkah

### 1. Kloning Repositori

Kloning repositori `my-go-app` ke direktori lokal Anda:

```bash
git clone https://github.com/username/my-go-app.git
cd my-go-app
```

### 2. Buat Dockerfile image dengan Nama "my-go-app-v2"

```bash
docker build -t my-go-app:v2 .
```

![Build Dockerfile Image](././imageSS/docker%20build.png)

### 3. Jalankan Container

Setelah image Docker berhasil dibangun, Anda dapat menjalankan container dari image tersebut di latar belakang dan memetakan port 8080:

```bash
docker run -d --name my-go-app-container -p 5555:5555 my-go-app:v2

```

![Build Dockerfile Image](././imageSS/terminal%20runing%20container.png)

### 4. Cek Running Program n logs

Buka browser dan akses localhost:55555

![Running FIle Golang](././imageSS/localhost%205555.png)

### 5. Cek logs and akses shell didalam container

melihat atau edit file di container sangat tidak direkomendasikan,, kita berikan contoh memakai logs dan cat, logs untuk melihat kerja container golang kita, cat untuk mencetak file yang berada di golang kita

```bash
 docker logs go-app
```

```bash
 docker exec -it go-app cat ./main.go
```

![Running FIle Golang](././imageSS/see%20log%20and%20cats.png)

### 6. Final file di docker desktop

![File Image](././imageSS/my-go-app.png)

![Container Running](././imageSS/running.png)
