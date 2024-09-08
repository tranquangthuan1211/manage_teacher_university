package database

type UserBaseData struct {
	Email    string `json:"email" gorm:"column:email"`
	Name     string `json:"name" gorm:"column:name"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Role     string `json:"role" gorm:"column:role"`
	Password string `json:"password" gorm:"column:password"`
}
type UserResponse struct {
	Id string `json:"id" gorm:"column:id"`
	UserBaseData
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Code   int          `json:"code" example:"200"`
	Token  string       `json:"token" example:"iuniu32neui3rn38fh784e5yn78f5r57R&FGU*^TU?;.'grteuiHIUN98"`
	Expire string       `json:"expire" example:"2005-08-15T15:52:01+00:00"`
	Data   UserResponse `json:"data"`
}
type Register struct {
	ID        string `json:"Id" gorm:"column:Id" example:"1"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Name      string `json:"Name" gorm:"not null" example:"Tran Quang Thuan"`
	Username  string `json:"Username" example:"0379823839"`
	Password  string `json:"Password" gorm:"not null" example:"123"`
	Role      string `json:"Role" example:"officer"`
	Email     string `json:"Email" example:"ndan.itus@gmail.com"`
}

func (UserBaseData) TableName() string {
	return DB_TEACHES + ".USERS"
}
