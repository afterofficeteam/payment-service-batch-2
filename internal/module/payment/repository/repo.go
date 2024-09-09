package repository

import "codebase-app/internal/module/payment/ports"

type repo struct {
}

func NewPaymentRepository() ports.PaymentRepository {
	return &repo{}
}
