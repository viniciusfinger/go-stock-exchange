package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Price        float64
	Total        float64
	DateTime     time.Time
	Shares       int
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price

	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Total:        total,
		Price:        price,
		Shares:       shares,
	}
}

func (t *Transaction) CalculateTotal(shares int, price float64) {
	t.Total = float64(t.Shares) * t.Price
}

func (t *Transaction) CloseBuyOrderTransaction() {
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) CloseSellOrderTransaction() {
	if t.SellingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) AddBuyOrderPendingShares(shares int) {
	t.BuyingOrder.PendingShares += shares
}

func (t *Transaction) AddSellOrderPendingShares(shares int) {
	t.SellingOrder.PendingShares += shares
}
