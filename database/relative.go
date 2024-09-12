package database

type RelativeBaseData struct {
	MAGV     string `json:"MAGV" gorm:"column:MAGV"`
	Ten      string `json:"TEN" gorm:"column:TEN"`
	NgaySinh string `json:"NGAYSINH" gorm:"column:NGAYSINH"`
	PHAI     string `json:"PHAI" gorm:"column:PHAI"`
}

func (RelativeBaseData) TableName() string {
	return DB_TEACHER + ".NGUOITHAN"
}
