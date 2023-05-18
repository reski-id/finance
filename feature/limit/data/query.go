package data

import (
	"finance/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type limitData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.LimitData {
	return &limitData{
		db: db,
	}
}

func (nd *limitData) Insert(newData domain.Limit) domain.Limit {
	cnv := ToLocal(newData)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Limit{}
	}
	return cnv.ToDomain()
}

func (bd *limitData) Update(dataID int, updatedData domain.Limit) domain.Limit {
	cnv := ToLocal(updatedData)
	err := bd.db.Model(cnv).Where("ID = ?", dataID).Updates(updatedData)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Limit{}
	}
	cnv.ID = uint(dataID)
	return cnv.ToDomain()
}

func (nd *limitData) Delete(dataID int) error {
	result := nd.db.Where("ID = ?", dataID).Delete(&Limit{})
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

func (nd *limitData) GetLimitID(dataID int) []domain.Limit {
	var data []Limit
	err := nd.db.Where("ID = ?", dataID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
