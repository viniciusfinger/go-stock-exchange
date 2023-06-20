package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transaction
}

func NewOrder(id string, investor *Investor, asset *Asset, shares int, pendingShares int, price float64, orderType string, status string) *Order {
	return &Order{
		ID:            id,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: pendingShares,
		Price:         price,
		OrderType:     orderType,
		Status:        "OPEN",
		Transactions:  []*Transaction{},
	}
}
