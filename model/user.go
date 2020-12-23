package model

type User struct {
	ID    uint   `gorm:"primary_key;unique_index;AUTO_INCREMENT"`
	Name  string `gorm:"type:varchar(180);unique_index"`
	Pass  []byte
	Email string `gorm:"type:varchar(180);unique_index"`
}

// User Model
//
// The User holds information about permission and other stuff.
//
// swagger:model User
type UserRegisterRequest struct {
	// The user id.
	//
	// read only: true
	// required: true
	// example: 25
	ID uint `json:"id"`
	// The user name. For login.
	//
	// required: true
	// example: unicorn
	Name string `binding:"required" json:"username" query:"name" form:"username"`

	// The user's password
	Pass string `binding:"required" json:"password" query:"name" form:"password"`

	// The user's email
	Email string `binding:"required" json:"email" query:"name" form:"email"`
}
