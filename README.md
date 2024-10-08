# Sistem Pemesanan Restoran CLI

Ini adalah aplikasi Command-Line Interface (CLI) sederhana berbasis Go untuk sistem pemesanan di restoran. Pengguna dapat menambahkan item dari menu yang sudah ditentukan, menghitung total pesanan, dan menampilkan ringkasan pesanan akhir.

## Fitur

- Menu yang Tersedia: `nasi goreng`, `mie goreng`, `ayam bakar`, `teh manis`, `jus jeruk`.
- Tambahkan beberapa item ke dalam pesanan, dan perbarui jumlah item yang sudah ada.
- Manajemen pesanan aman menggunakan `sync.Mutex` pada Go.
- Hitung dan tampilkan total pesanan secara real-time.
- Encode detail pesanan menjadi Base64 untuk penyimpanan yang lebih aman.
- Proses pesanan secara asinkron dengan mekanisme timeout untuk mensimulasikan keterlambatan persiapan.

## Persyaratan

- Go versi 1.18 atau lebih tinggi.

## Cara Menjalankan Aplikasi

1. Clone repository ke mesin lokal Anda:
   ```bash
   git clone https://github.com/adtbch/ManajemenPemesanan_Golang.git
   cd <repository-directory>

2. Bangun dan jalankan program Go:
   ```bash
    go run main.go

3. Ikuti instruksi yang ditampilkan di layar untuk menambahkan item ke pesanan.

## Contoh Penggunaan
    Setelah menjalankan program, Anda akan disambut dengan pesan selamat datang dan daftar menu yang tersedia:
    
    Selamat Datang di Restoran Sederhana!
    Menu yang Tersedia:
    - Nasi Goreng: 20000.00
    - Mie Goreng: 15000.00
    - Ayam Bakar: 25000.00
    - Teh Manis: 5000.00
    - Jus Jeruk: 10000.00
    Masukkan nama menu yang ingin dipesan, diikuti dengan jumlahnya:

    Masukkan nama menu (ketik 'selesai' untuk menyelesaikan pesanan): nasi goreng
    Masukkan jumlah pesanan: 2
    Total sementara: 40000.00
    Ketika Anda selesai memesan, ketik selesai untuk menampilkan ringkasan akhir.

    Program akan menampilkan versi terenkripsi dari detail pesanan dan status pemrosesan setiap item:

    Informasi pesanan terenkripsi (base64): RGV0YWlsIFBlc2FuYW46IG5hc2kgZ29yZW5nOiAyLCA=
    Pesanan nasi goreng telah selesai diproses.
    Gambaran Kode

## Gambaran Kode
- Struct MenuItem: Mewakili item dalam menu restoran, berisi name, price, dan quantity.
- Struct Order: Berisi daftar item yang dipesan dan sync.Mutex untuk operasi yang aman dari thread.
- Interface MenuOperations: Mendefinisikan metode untuk menambahkan item dan menampilkan pesanan.

## Metode Utama
- addItem: Menambahkan item ke dalam pesanan, memperbarui jumlah jika item sudah ada.
- displayOrder: Menampilkan ringkasan lengkap dari pesanan.
- getMenuInput & getQuantityInput: Fungsi utilitas untuk mengambil input pengguna terkait menu dan jumlah.
- calculateOrderTotal: Menghitung total harga dari pesanan saat ini.
- displayEncodedOrderInfo: Menampilkan detail pesanan dalam format Base64.
- processOrdersWithTimeout: Mensimulasikan pemrosesan pesanan secara asinkron, dengan timeout untuk setiap item.

## Kontak
Untuk pertanyaan atau masukan, silakan hubungi adit.bachtiar091@gmail.com.