package model

// MUserDetail struct
type MUserDetail struct {
	ID      int64  `gorm:"primary_key;auto_increment" json:"id"`
	Address string `gorm:"size 255" json:"address"`
	DOB     string `gorm:"size 8" json:"dob"`
	POB     string `gorm:"size 255" json:"pob"`
	Phone   string `gorm:"size 255" json:"phone"`
	Email   string `gorm:"size 255" json:"email"`
	// UserID  int64  `sql:"type:bigint(20) REFERENCES m_user(id)" json:"userId"`
	UserID int64 `json:"userId"`
	MUser  MUser `gorm:"ForeignKey:UserID;AssociationForeignKey:ID" json:"user"`
}

// MUserDetails array of MUserDetail
type MUserDetails []MUserDetail
