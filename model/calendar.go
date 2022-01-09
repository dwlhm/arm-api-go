package model

import "gorm.io/gorm"

type Calendar struct {
	gorm.Model
	Tanggal      string
	DetailJadwal []DetailJadwal `gorm:"foreignKey:CalendarRefer"'`
}

type DetailJadwal struct {
	gorm.Model
	Lokasi          string
	Kota            string
	Alamat          string
	Waktu           string
	PemangkuManaqib string
	Pengisi         string
	CalendarRefer   uint
}

type CalendarRequest struct {
	Tanggal         string   `json:"tanggal" binding:"required"`
	Kota            string   `json:"kota" binding:"required"`
	Lokasi          string   `json:"lokasi" binding:"required"`
	Alamat          string   `json:"alamat" binding:"required"`
	Waktu           string   `json:"waktu" binding:"required"`
	PemangkuManaqib string   `json:"pemangku-manaqib" binding:"required"`
	Pengisi         []string `json:"pengisi" binding:"required"`
}

type CalendarUpdate struct {
	Kota            string   `json="kota"`
	Lokasi          string   `json="lokasi"`
	Alamat          string   `json="alamat"`
	Waktu           string   `json="waktu"`
	PemangkuManaqib string   `json="pemangku-manaqib"`
	Pengisi         []string `json="pengisi"`
}

type CalendarUpdateDb struct {
	Kota            string `json="kota"`
	Lokasi          string `json="lokasi"`
	Alamat          string `json="alamat"`
	Waktu           string `json="waktu"`
	PemangkuManaqib string `json="pemangku-manaqib"`
	Pengisi         string `json="pengisi"`
}

type CalendarResponse struct {
	Tanggal string                 `json:"tanggal"`
	Jadwal  []DetailJadwalResponse `json:"jadwal"`
}

type DetailJadwalResponse struct {
	ID              uint     `json:"id"`
	Kota            string   `json:"kota"`
	Lokasi          string   `json:"lokasi"`
	Alamat          string   `json:"alamat"`
	Waktu           string   `json:"waktu"`
	PemangkuManaqib string   `json:"pemangku manaqib"`
	Pengisi         []string `json:"pengisi"`
}
