# go-mirza-template


Beberapa library yg dipake disini adalah

- gin-gonic
- gorm
- go.uuid
- cors


## CRUD
Create Read Update Delete atau disingkat CRUD adalah 4 fungsi dasar pada aplikasi yang biasanya melibatkan database (bisa menyimpan data). 

### Create
Create adalah fungsi untuk melakukan penyimpanan data. Pada proses ini aplikasi akan meminta input data yang diperlukan dan mewakili model yang ada di aplikasi utama di tabel database. 

Misal kita punya model spt ini: 

```
type Transaksi struct {
	ID        string    
	Deskripsi string    
	Tanggal   time.Time 
	Nilai     float64
	Deleted   bool  
}
```

Bentuk request akan tergantung pada **bussiness process** / logic / requirement yang dibutuhkan. Misal dalam hal ini
- ID akan di generate otomatis oleh aplikasi
- Tanggal (tanggal pembuatan data) juga akan di buat oleh aplikasi

Maka bentuk input yang dimungkinkan adalah. Misal dalam hal ini Tanggal adalah tanggal input aplikasi (created date).

```
type CreateTransaksiRequest struct {
	Deskripsi string
	Nilai     float64   
}
```

Jika requirement yang diminta bahwa tanggal boleh ditentukan oleh user, maka bentuk requestnya bisa begini

```
type CreateTransaksiRequest struct {
	Deskripsi string    
	Tanggal   time.Time 
	Nilai     float64   
}
```

Jika proses ini berhasil, maka aplikasi akan mengembalikan response berupa

```
type CreateTransaksiResponse struct {
	ID        string    
	Deskripsi string    
	Tanggal   time.Time 
	Nilai     float64   
}
```

Atau bisa juga hanya mengembalikan ID saja

```
type CreateTransaksiResponse struct {
	ID        string 
}
```

Semua tergantung pada requirement

Jika gagal, maka aplikasi akan mengembalikan pesan error. Dan sistem akan mengembalikan parameter apa yang error

### Read
Read adalah proses membaca kembali data yang sudah pernah diinput ke aplikasi. Ada beberapa macam proses read. Yaitu :
- GetAll : mengambil semua data. Untuk data yang sangat banyak, biasanya ada parameter untuk paging. Untuk paging kita membutuhkan 2 parameter yaitu halaman berapa dan berapa data perhalaman. Untuk determinate paging, kita membutuhkan informasi tentang jumlah total dari item model yang ada di aplikasi. Untuk undeterminate paging, maka kita perlu suatu flag khusus yang akan memerikan informasi ketersediakan data selanjutnya.

- GetByX : mengambil beberapa data secara spesifik tergantung pada kondisi tertentu. Untuk model read semacam ini variasinya sangat banyak tergantung dari kompleksitas model yang dimiliki, dan spesifikas kebutuhan yang diminta.

- GetOne : mengambil hanya satu data saja. Biasanya akan diminta id dari salah satu model yang sudah ada.

Data yang akan dikembalikan biasanya berupa List. Biasanya data yang dikembalikan adalah data yang hanya diperlukan oleh user misal

```
type GetTransaksiResponse struct {
	ID        string    
	Deskripsi string    
	Tanggal   time.Time 
	Nilai     float64
}
```

Perhatikan bahwa disini field Deleted tidak kita kembalikan ke user karena hanya data yang available saja yang kita sajikan. Bukan data yang sudah kita hapus

### Update
Adalah proses mengubah / mengedit data yang sudah pernah diinput aplikasi. Biasanya hanya satu data yang diupdate. Selain menyertakan data yang baru (data yang akan diubah) pada proses mengupdate ini biasanya aplikasi juga akan meminta parameter berupa id model yang valid (ada didatabase).

Jika objeknya tersedia (id yang diberikan valid), maka data akan berhasil diubah. Namun sebaliknya jika idnya tidak dikenal, maka akan mengamblikan pesan error.

### Delete
Adalah proses menghapus data yang sudah pernah diinput. Hampir mirip seperti proses Update, proses delete juga akan meminta input id model yang akan dihapus. Proses delete ada dua macam. Yaitu, 
- Real delete. Yaitu menghapus data langsung ke database. Data akan hilang dari database dan tidak bisa dikembalikan lagi
- Flagging delete. Yaitu memberikan tanda / flag pada item yang akan dihapus bahwasanya data tersebut sudah dihapus. Jadi ini adalah proses menghapus secara logical saja. Data aslinya masih ada dan masih bisa dikembalikan lagi sekiranya diperlukan dan itupun harus menyesuaikan requirement aplikasi. Jangan sampai proses recover data ini malah merusak data yang lain. Perhatikan bahwa contoh model yang kita gunakan pada penjelasan sebelumnya adalah yang menggunakan `Deleted` field.


## Validasi

Setiap CRUD selalu ada proses validasi atau pengecekan kelengkapan data sebelum data masuk kedalam database. Hal ini adalah untuk mencegah data yang tidak valid masuk kedalam database dan merusak logic aplikasi.



https://hackernoon.com/restful-api-designing-guidelines-the-best-practices-60e1d954e7c9
https://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api
https://medium.com/datafire-io/designing-a-restful-api-81792139a56
http://www.restapitutorial.com/lessons/httpmethods.html

