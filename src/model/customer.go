package model

// Customer is a model in the "customers" table.
type Customer struct {
	ID   int     `json:"id,omitempty"`
	Name *string `json:"name" gorm:"not null"`
}
