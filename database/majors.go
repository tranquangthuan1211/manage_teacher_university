package database

type Major struct {
	BaseModel
	MajorBaseData
}

type MajorBaseData struct {
	Id      string `json:"MABM" gorm:"column:MABM"`
	Name    string `json:"TENBM" gorm:"column:TENBM"`
	Rom     string `json:"PHONG" gorm:"column:PHONG"`
	Phone   string `json:"SDT" gorm:"column:DIENTHOAI"`
	Leader  string `json:"TRUONGBM" gorm:"column:TRUONGBM"`
	IdMajor string `json:"MAKHOA" gorm:"column:MAKHOA"`
	Date    string `json:"NGAYNHANCHUC" gorm:"column:NGAYNHANCHUC"`
}

func (Major) TableName() string {
	return DB_TEACHES + ".BOMON"
}
