package models

type Vendor struct {
	VendorID string `gorm:"primaryKey" json:"vendorId"`
	Name     string `json:"name"`
	Email    string `json:"emailAddress"`
	Phone    string `json:"phoneNumber"`
	Address  string `json:"address"`
	Contact  string `json:"contact"`
}
