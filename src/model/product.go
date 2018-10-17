package model

// Product is a model in the "products" table.
type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  *string `json:"name"  gorm:"not null;unique"`
	Price float64 `json:"price,string" gorm:"type:decimal(18,2)"`
}
