package models

type User struct {
	ID       uint64 `json:"id" gorm:"not null; primary_key;"`
	Name     string `json:"name" gorm:"not null; default:null;"`
	Surname  string `json:"surname" gorm:"not null; default:null;"`
	Username string `json:"username" gorm:"not null; default:null;"`
	Email    string `json:"email" gorm:"not null; default:null;"`
	RoleID   int    `json:"role_id" gorm:"not null; default:null;"`
}
