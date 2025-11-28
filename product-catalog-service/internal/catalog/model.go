package catalog

import (
	"time"

	"github.com/shopspring/decimal"
)

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID                  string          `json:"id"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	Price               decimal.Decimal `json:"price"`
	QuantityInInventory int             `json:"quantity_in_inventory"`
	CategoryID          Category        `json:"category_id"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
}
