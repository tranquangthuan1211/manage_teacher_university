package database

type SubjectBase struct {
	IdMonHoc     string `json:"MAMH" gorm:"column:MAMH"`
	NameMonHoc   string `json:"TENMH" gorm:"column:TENMH"`
	PhongHoc     string `json:"PHONGHOC" gorm:"column:PHONGHOC"`
	DienThoai    string `json:"DIENTHOAI" gorm:"column:DIENTHOAI"`
	TruongBoMon  string `json:"TRUONGBM" gorm:"column:TRUONGBM"`
	MaKhoa       string `json:"MAKHOA" gorm:"column:MAKHOA"`
	NgayNhanChuc string `json:"NGAYNHANCHUC" gorm:"column:NGAYNHANCHUC"`
}

func (SubjectBase) TableName() string {
	return DB_TEACHER + ".MONHOC"
}
