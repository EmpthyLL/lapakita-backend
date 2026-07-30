package payment

import (
	"lapakita-backend/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService struct {
	SnapClient snap.Client
}

func NewPaymentService(cfg *config.Config) *PaymentService {
	var client snap.Client
	client.New(cfg.PaymentServerKey, midtrans.Sandbox)

	return &PaymentService{
		SnapClient: client,
	}
}
