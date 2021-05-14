# evermos
```
Jawaban saya terkait dari permasalahan yang diberikan adalah
1. Tidak adanya pengecekan stock barang yang dilakukan pada saat sicustomer melakukan 'checkout'.
2. Menampilkan semua produk meskipun ada produk yang stoknya adalah 0


Dari permasalahan tersebut, maka saya menyarankan solusi:
1. Adanya penegcekan stok pada saat customer ingin melakukan checkout/pembayaran. 
   Pada saat sicustomer ingin melakukan checkout dan pembayaran, harus dilakukan kembali pengecekan ke database apakah stock masih tersedia.
   Stoknya harus lebih besar dari 0. Jika ya, maka checkout/payment berhasil. 
   Jika tidak, muncul pemberitahuan "Out of Stock / Stok tidak tersedia". 
   Jadi jika ada customer lain melakukan order dengan produk yang sama, jumlah yang dibeli harus disesuaikan dengan stok yang baru. 
   Hal ini juga dapat mengatasi jumlah stok barang minus.
```
```
Untuk menjalankan API ini harus dipastikan GO, dan PostgreSql sudah terinstall di local. 
Untuk menjalankan pada local anda, anda dapat terlebih dahulu melakukan pengaturan pada folder config/app.yaml sesuai pengaturan postgresql local anda
dsn: "host=localhost port=5432 user=postgres password=12345678 dbname=evermos sslmode=disable"

Kemudia anda dapat melakukan running aplikasi pada local dengan menjalankan 
> go run main.go
# Server will run at localhost:8080

Anda dapat melakukan test API nya pada postman, dengan route ada dalam folder controller.
Contoh: Post Order
localhost:8080/api/orders

dengan body
{
    "cust_id": 1,
    "item_id":1,
    "quantity":10,
    "is_pay":false,
    "dispatched": false
}

Anda juga dapat melakukan functonal testing dengan terlebih dahulu masuk kedalam folde testing
> go test

```
