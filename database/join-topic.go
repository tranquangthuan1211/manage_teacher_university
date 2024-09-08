package database

type TopicBaseData struct {
	Id_Teacher string `json:"MAGV" gorm:"column:MAGV"`
	Id_Topic   string `json:"MADT" gorm:"column:MADT"`
	STT        string `json:"STT" gorm:"column:STT"`
	Allowance  string `json:"PHUCAP" gorm:"column:PHUCAP"`
	Result     string `json:"KETQUA" gorm:"column:KETQUA"`
}

func (TopicBaseData) TableName() string {
	return DB_TEACHES + ".THAMGIADT"
}
