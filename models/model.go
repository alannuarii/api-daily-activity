package models

type Activity struct {
    ID         uint      `json:"id"`
    Tanggal    string    `json:"tanggal"`
    Jenis      string    `json:"jenis"`
    Perusahaan string    `json:"perusahaan"`
    Pekerjaan  string    `json:"pekerjaan"`
    Kode       string    `json:"kode"`
}

type Photo struct {
    ID      uint      `json:"id"`
    Tanggal string   `json:"tanggal"`
    Kode    string    `json:"kode"`
    Foto    string    `json:"foto"`
}
