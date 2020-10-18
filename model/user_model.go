package model

// MUser struct
type MUser struct {
	ID                 int64  `gorm:"primary_key;auto_increment" json:"id"`
	UserName           string `gorm:"size:255;not null;unique" json:"userName"`
	Password           string `gorm:"size:255;not null" json:"password"`
	AccountExpired     bool   `gorm:"default:false" json:"accountExpired"`
	AccountLocked      bool   `gorm:"default:false" json:"accountLocked"`
	CredentialsExpired bool   `gorm:"default:false" json:"credentialsExpired"`
	Enabled            bool   `gorm:"default:true" json:"enabled"`
}

// MUsers array of MUser type
type MUsers []MUser
