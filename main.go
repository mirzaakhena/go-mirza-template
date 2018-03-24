package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	melody "gopkg.in/olahol/melody.v1"
)

type Application struct {
	DB *gorm.DB
}

type Transaksi struct {
	ID        string    `json:"id"`
	Deskripsi string    `json:"deskripsi"`
	Tanggal   time.Time `json:"tanggal"`
	Nilai     float64   `json:"nilai"`
}

type StandartResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (app Application) simpanData(c *gin.Context) {

	// siapkan requestnya
	var request struct {
		Deskripsi string  `json:"deskripsi"`
		Nilai     float64 `json:"nilai"`
	}

	// bind objectnya
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, StandartResponse{"Gagal Binding Object", nil})
		return
	}

	// validasi input
	if request.Nilai == 0 {
		c.JSON(http.StatusBadRequest, StandartResponse{"Nilai tidak boleh 0", nil})
		return
	}

	// siapkan objeknya
	trx := Transaksi{
		ID:        uuid.NewV4().String(),
		Deskripsi: request.Deskripsi,
		Tanggal:   time.Now(),
		Nilai:     request.Nilai,
	}

	// simpan objeknya
	app.DB.Create(&trx)

	c.JSON(http.StatusOK, StandartResponse{"Sukses Insert", trx})
}

func (app Application) ambilSemuaData(c *gin.Context) {

	// ambil semua transaksi
	var listTrx []Transaksi
	app.DB.Find(&listTrx)

	c.JSON(http.StatusOK, StandartResponse{"Sukses Ambil Data", listTrx})
}

func (app Application) ubahData(c *gin.Context) {

	// ambil transaksi_id nya
	transaksiID := c.Param("transaksiID")

	// siapkan requestnya
	var request struct {
		Deskripsi string  `json:"deskripsi"`
		Nilai     float64 `json:"nilai"`
	}

	// bind objectnya
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, StandartResponse{"Gagal Binding Object", nil})
		return
	}

	// validasi input
	if request.Nilai == 0 {
		c.JSON(http.StatusBadRequest, StandartResponse{"Nilai tidak boleh 0", nil})
		return
	}

	// ambil transaksi by id
	var trx Transaksi
	app.DB.Where("ID = ?", transaksiID).First(&trx)

	// ubah nilai2 nya
	trx.Nilai = request.Nilai
	trx.Deskripsi = request.Deskripsi

	// simpan kembali
	app.DB.Save(&trx)

	c.JSON(http.StatusOK, StandartResponse{"Sukses Ubah Data", trx})
}

func (app Application) hapusData(c *gin.Context) {

	// ambil transaksi_id nya
	transaksiID := c.Param("transaksiID")

	// ambil transaksi by id
	app.DB.Delete(Transaksi{}, "ID = ?", transaksiID)

	c.JSON(http.StatusOK, StandartResponse{"Sukses Hapus Data", nil})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	m := melody.New()

	// buka koneksi
	db, err := gorm.Open("sqlite3", "transaksi.db")
	if err != nil {
		panic("failed to connect database")
	}

	// nanti ditutup
	defer db.Close()

	// bikin tabel
	db.AutoMigrate(&Transaksi{})

	app := Application{
		DB: db,
	}

	r.POST("/transaksi", app.simpanData)
	r.GET("/transaksi", app.ambilSemuaData)
	r.PUT("/transaksi/:transaksiID", app.ubahData)
	r.DELETE("/transaksi/:transaksiID", app.hapusData)

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Printf("%+v", s.Request.URL)
		m.Broadcast(msg)
	})

	r.Run(":8081")
}
