package models

import (
	"time"

	"github.com/shopspring/decimal"
	// "golang.org/x/text/date"
)

type Order struct {
	ID                  string `gorm:"size:36; not null; uniqueIndex; primary_key"`
	UserID              string `gorm:"size:36; index"`
	Code                string `gorm:"size:100; index"`
	Status              string `gorm:"size:100"`
	OrderDate           time.Time
	PaymentDue          time.Time
	PaymentStatus       string          `gorm:"size:255"`
	PaymentToken        string          `gorm:"size:255"`
	PaymentURL          string          `gorm:"size:255"`
	TotalPrice          decimal.Decimal `gorm:"type:decimal(16,2);"`
	TaxAmount           decimal.Decimal `gorm:"type:decimal(16,2);"`
	TaxPercent          decimal.Decimal `gorm:"type:decimal(16,2);"`
	DiscountAmount      decimal.Decimal `gorm:"type:decimal(16,2);"`
	DiscountPercent     decimal.Decimal `gorm:"type:decimal(16,2);"`
	ShippingCost        decimal.Decimal `gorm:"type:decimal(16,2);"`
	GrandTotal          decimal.Decimal `gorm:"type:decimal(16,2);"`
	Note                string          `gorm:"size:255"`
	ShippingCourier     string          `gorm:"size:255"`
	ShippingServiceName string          `gorm:"size:255"`
	ApprovedBy          string          `gorm:"size:100"`
	ApprovedAt          time.Time
	CancelledBy         string `gorm:"size:100"`
	CancelledAt         time.Time
	CancelledNote       string `gorm:"size:255"`
	DeletedAt           time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
