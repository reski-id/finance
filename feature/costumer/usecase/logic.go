package usecase

import (
	"errors"
	"finance/domain"
)

type costumerUseCase struct {
	costumerData domain.CostumerData
}

func New(model domain.CostumerData) domain.CostumerUseCase {
	return &costumerUseCase{
		costumerData: model,
	}
}

func (nu *costumerUseCase) AddCostumer(newData domain.Costumer) (domain.Costumer, error) {

	res := nu.costumerData.Insert(newData)

	if res.ID == 0 {
		return domain.Costumer{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *costumerUseCase) UpCostumer(IDCostumer int, updateData domain.Costumer) (domain.Costumer, error) {
	if IDCostumer == -1 {
		return domain.Costumer{}, errors.New("invalid data")
	}

	res := nu.costumerData.Update(IDCostumer, updateData)
	if res.ID == 0 {
		return domain.Costumer{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *costumerUseCase) DelCostumer(IDCostumer int) (bool, error) {
	err := nu.costumerData.Delete(IDCostumer)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (nu *costumerUseCase) GetAllData() ([]domain.Costumer, error) {
	res := nu.costumerData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (nu *costumerUseCase) GetSpecificCostumer(dataID int) ([]domain.Costumer, error) {
	res := nu.costumerData.GetCostumerID(dataID)
	if dataID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}
