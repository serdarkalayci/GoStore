package model

// Order is a model in the "orders" table.
type Order struct {
	ID       int     `json:"id,omitempty"`
	Subtotal float64 `json:"subtotal,string" gorm:"type:decimal(18,2)"`

	Customer   Customer `json:"customer" gorm:"ForeignKey:CustomerID"`
	CustomerID int      `json:"-"`

	Products []Product `json:"products" gorm:"many2many:order_products"`
}
