package database

type WorkingBase struct {
	MaDT   string `json:"MADT" gorm:"column:MADT"`
	STT    string `json:"STT" gorm:"column:STT"`
	TenCV  string `json:"TENCV" gorm:"column:TENCV"`
	NgayBD string `json:"NGAYBD" gorm:"column:NGAYBD"`
	NgayKT string `json:"NGAYKT" gorm:"column:NGAYKT"`
}

func (WorkingBase) TableName() string {
	return DB_TEACHER + ".CONGVIEC"
}
