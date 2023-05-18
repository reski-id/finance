package usecase

import (
	"errors"
	"finance/domain"
)

type transUseCase struct {
	transactionData domain.TransactionData
}

func New(model domain.TransactionData) domain.TransactionUseCase {
	return &transUseCase{
		transactionData: model,
	}
}

func (nu *transUseCase) AddTransaction(newData domain.Transaction) (domain.Transaction, error) {

	res := nu.transactionData.Insert(newData)

	if res.ID == 0 {
		return domain.Transaction{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *transUseCase) GetSpecificTransaction(dataID int) ([]domain.Transaction, error) {
	res := nu.transactionData.GetTransactionID(dataID)
	if dataID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}
