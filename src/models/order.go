package models

type Order struct {
	Model
	TransactionId   string `json:"transaction_id" gorm:"null"`
	UserId          uint   `json:"user_id"`
	Code            string `json:"code"`
	AmbassadorEmail string `json:"ambassador_email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email" `
	Address         string `json:"address" gorm:"null"`
	City            string `json:"city" gorm:"null"`
	Zip             string `json:"zip" gorm:"null"`
	Complete        bool   `json:"complete" gorm:"false"`
}
