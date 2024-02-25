package models

import "time"

type Activity struct {
    ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Tanggal    time.Time `gorm:"type:date" json:"tanggal"`
    Jenis      string    `gorm:"type:varchar(10)" json:"jenis"`
    Perusahaan string    `gorm:"type:varchar(30)" json:"perusahaan"`
    Pekerjaan  string    `gorm:"type:varchar(75)" json:"pekerjaan"`
    Kode       string    `gorm:"type:varchar(6)" json:"kode"`
}

type Photo struct {
    ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    Tanggal time.Time `gorm:"type:date" json:"tanggal"`
    Kode    string `gorm:"type:varchar(6)" json:"kode"`
    Foto    string `gorm:"type:varchar(100)" json:"foto"`
}
