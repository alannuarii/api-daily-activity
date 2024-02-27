package models

type Activity struct {
    ID          int    `db:"id" json:"id"`
    Tanggal     string `db:"tanggal" json:"tanggal"`
    Jenis       string `db:"jenis" json:"jenis"`
    Pekerjaan   string `db:"pekerjaan" json:"pekerjaan"`
    Kode        string `db:"kode" json:"kode"`
    Perusahaan  string `db:"perusahaan" json:"perusahaan"`
    Foto        string `db:"foto" json:"foto"`
}

