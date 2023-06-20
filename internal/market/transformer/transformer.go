package transformer

import (
	"github.com/viniciusfinger/bolsa-valores/internal/market/dto"
	"github.com/viniciusfinger/bolsa-valores/internal/market/entity"
)

//Parecido com o modelMapper do Java, converte DTO em entity e vice-versa.

func TransformInput(input dto.TradeInput) *entity.Order {
	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID)
	order := entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, input.OrderType)

	if input.CurrentShares > 0 {
		investorAssetPosion := entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(investorAssetPosion)
	}

	return order
}

func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderID:    order.ID,
		InvestorID: order.Investor.ID,
		AssetID:    order.Asset.ID,
		OrderType:  order.OrderType,
		Shares:     order.Shares,
		Partial:    order.PendingShares,
		Status:     order.Status,
	}

	var transactionsOutput []*dto.TransactionOutput

	for _, t := range order.Transactions {
		transactionOutput := &dto.TransactionOutput{
			TransactionID: t.ID,
			BuyerID:       t.BuyingOrder.ID,
			SellerID:      t.SellingOrder.ID,
			AssetID:       t.SellingOrder.Asset.ID,
			Shares:        t.SellingOrder.Shares - t.SellingOrder.PendingShares,
			Price:         t.Price,
		}

		transactionsOutput = append(transactionsOutput, transactionOutput)
	}

	output.TransactionsOutput = transactionsOutput
	return output
}
