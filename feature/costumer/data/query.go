package data

import (
	"finance/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type costumerData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.CostumerData {
	return &costumerData{
		db: db,
	}
}

func (nd *costumerData) Insert(newData domain.Costumer) domain.Costumer {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Costumer{}
	}
	return cnv.ToDomain()
}

func (bd *costumerData) Update(dataID int, updatedData domain.Costumer) domain.Costumer {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", dataID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Costumer{}
	}
	cnv.ID = uint(dataID)
	return cnv.ToDomain()
}

func (nd *costumerData) Delete(dataID int) error {
	result := nd.db.Where("ID = ?", dataID).Delete(&Costumer{})
	if result.Error != nil {
		log.Println("Cannot delete data", result.Error.Error())
		return result.Error
	}
	if result.RowsAffected == 0 {
		log.Println("No data deleted")
		return fmt.Errorf("data with ID %d does not exist", dataID)
	}
	return nil
}

func (nd *costumerData) GetAll() []domain.Costumer {
	var data []Costumer
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *costumerData) GetCostumerID(dataID int) []domain.Costumer {
	var data []Costumer
	err := nd.db.Where("ID = ?", dataID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
