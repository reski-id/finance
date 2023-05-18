package data

import (
	"finance/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type transactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.TransactionData {
	return &transactionData{
		db: db,
	}
}

func (nd *transactionData) Insert(newData domain.Transaction) domain.Transaction {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Transaction{}
	}
	return cnv.ToDomain()
}

func (nd *transactionData) GetTransactionID(dataID int) []domain.Transaction {
	var data []Transaction
	err := nd.db.Where("ID = ?", dataID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
