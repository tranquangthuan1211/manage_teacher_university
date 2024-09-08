package database

type Teach struct {
	BaseModel
	TeachBaseData
}

type TeachBaseData struct {
	Id          string `json:"MAGV" gorm:"column:MAGV"`
	Name        string `json:"HOTEN" gorm:"column:HOTEN"`
	Salary      string `json:"LUONG" gorm:"column:LUONG"`
	Sex         string `json:"PHAI" gorm:"column:PHAI"`
	DateOfBirth string `json:"NGSINH" gorm:"column:NGSINH"`
	Address     string `json:"DIACHI" gorm:"column:DIACHI"`
	IdSubject   string `json:"GVQLCM" gorm:"column:GVQLCM"`
	IdMajor     string `json:"MABM" gorm:"column:MABM"`
}

func (Teach) TableName() string {
	return DB_TEACHES + ".GIAOVIEN"
}
