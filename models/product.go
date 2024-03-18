package models

type Product struct {
	Id          int  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi   string `gorm:"type:varchar(300)" json:"deskripsi"`
}
