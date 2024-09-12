package database

type TopicBaseData struct {
	MaDT    string `json:"MADT" gorm:"column:MADT"`
	TenDT   string `json:"TENDT" gorm:"column:TENDT"`
	CapQL   string `json:"CAPQL" gorm:"column:CAPQL"`
	KinhPhi string `json:"KINHPHI" gorm:"column:KINHPHI"`
	NgayBD  string `json:"NGAYBD" gorm:"column:NGAYBD"`
	NgayKT  string `json:"NGAYKT" gorm:"column:NGAYKT"`
	GVCNDT  string `json:"GVCNDT" gorm:"column:GVCNDT"`
}

func (TopicBaseData) TableName() string {
	return DB_TEACHER + ".DETAI"
}
