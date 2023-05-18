package usecase

import (
	"errors"
	"finance/domain"
)

type limitUseCase struct {
	costumerData domain.LimitData
}

func New(model domain.LimitData) domain.LimitUseCase {
	return &limitUseCase{
		costumerData: model,
	}
}

func (nu *limitUseCase) AddLimit(newData domain.Limit) (domain.Limit, error) {

	res := nu.costumerData.Insert(newData)

	if res.ID == 0 {
		return domain.Limit{}, errors.New("error insert data")
	}
	return res, nil
}

func (nu *limitUseCase) UpLimit(IDLimit int, updateData domain.Limit) (domain.Limit, error) {
	if IDLimit == -1 {
		return domain.Limit{}, errors.New("invalid data")
	}

	res := nu.costumerData.Update(IDLimit, updateData)
	if res.ID == 0 {
		return domain.Limit{}, errors.New("error update data")
	}

	return res, nil
}

func (nu *limitUseCase) DelLimit(IDLimit int) (bool, error) {
	err := nu.costumerData.Delete(IDLimit)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (nu *limitUseCase) GetSpecificLimit(dataID int) ([]domain.Limit, error) {
	res := nu.costumerData.GetLimitID(dataID)
	if dataID == -1 {
		return nil, errors.New("error update data")
	}

	return res, nil
}
