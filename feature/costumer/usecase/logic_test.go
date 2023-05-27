package usecase_test

import (
	"database/sql"
	"errors"
	"finance/domain"
	"finance/feature/costumer/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define mock structs

type MockCostumerData struct {
	mock.Mock
}

func (m *MockCostumerData) Insert(newData domain.Costumer) domain.Costumer {
	args := m.Called(newData)
	return args.Get(0).(domain.Costumer)
}

func (m *MockCostumerData) Update(IDCostumer int, updateData domain.Costumer) domain.Costumer {
	args := m.Called(IDCostumer, updateData)
	return args.Get(0).(domain.Costumer)
}

func (m *MockCostumerData) Delete(IDCostumer int) error {
	args := m.Called(IDCostumer)
	return args.Error(0)
}

func (m *MockCostumerData) GetAll() []domain.Costumer {
	args := m.Called()
	return args.Get(0).([]domain.Costumer)
}

func (m *MockCostumerData) GetCostumerID(dataID int) []domain.Costumer {
	if dataID == -1 {
		return nil
	}
	result := m.Called(dataID)
	if result.Get(0) == nil {
		return nil
	}
	return result.Get(0).([]domain.Costumer)
}

// Test cases

func TestAddCostumer_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockCostumerData)
	mockData.On("Insert", mock.AnythingOfType("domain.Costumer")).Return(domain.Costumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       5000.0,
		KTPPhoto:     "ktp_photo.jpg",
		SelfiePhoto:  "selfie_photo.jpg",
	})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	newCostumer := domain.Costumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       5000.0,
		KTPPhoto:     "ktp_photo.jpg",
		SelfiePhoto:  "selfie_photo.jpg",
	}

	// Perform the action
	result, err := useCase.AddCostumer(newCostumer)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, domain.Costumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       5000.0,
		KTPPhoto:     "ktp_photo.jpg",
		SelfiePhoto:  "selfie_photo.jpg",
	}, result)

	// Assert that the mock was called with the expected arguments
	mockData.AssertCalled(t, "Insert", newCostumer)
}

func TestAddCostumer_Error(t *testing.T) {
	// Prepare mock
	mockData := new(MockCostumerData)
	mockData.On("Insert", mock.AnythingOfType("domain.Costumer")).Return(domain.Costumer{})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	newCostumer := domain.Costumer{
		ID:           0,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       0.0,
		KTPPhoto:     "",
		SelfiePhoto:  "",
	}

	// Perform the action
	result, err := useCase.AddCostumer(newCostumer)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, domain.Costumer{}, result)

	// Assert that the mock was called with the expected arguments
	mockData.AssertCalled(t, "Insert", newCostumer)
}

func TestUpCostumer_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockCostumerData)
	mockData.On("Update", 1, mock.AnythingOfType("domain.Costumer")).Return(domain.Costumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       6000.0,
		KTPPhoto:     "ktp_photo.jpg",
		SelfiePhoto:  "selfie_photo.jpg",
	})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	IDCostumer := 1
	updateData := domain.Costumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       6000.0,
		KTPPhoto:     "ktp_photo.jpg",
		SelfiePhoto:  "selfie_photo.jpg",
	}

	// Perform the action
	result, err := useCase.UpCostumer(IDCostumer, updateData)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, domain.Costumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       6000.0,
		KTPPhoto:     "ktp_photo.jpg",
		SelfiePhoto:  "selfie_photo.jpg",
	}, result)

	// Assert that the mock was called with the expected arguments
	mockData.AssertCalled(t, "Update", IDCostumer, updateData)
}

func TestUpCostumer_InvalidData(t *testing.T) {
	// Prepare mock
	mockData := new(MockCostumerData)

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	IDCostumer := -1
	updateData := domain.Costumer{
		ID:           1,
		NIK:          "1234567890",
		FullName:     "John Doe",
		LegalName:    "John",
		PlaceOfBirth: "New York",
		DateOfBirth:  sql.NullTime{},
		Salary:       6000.0,
		KTPPhoto:     "ktp_photo.jpg",
		SelfiePhoto:  "selfie_photo.jpg",
	}

	// Perform the action
	result, err := useCase.UpCostumer(IDCostumer, updateData)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, domain.Costumer{}, result)

	// Assert that the mock was not called
	mockData.AssertNotCalled(t, "Update")
}

func TestDelCostumer_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockCostumerData)
	mockData.On("Delete", 1).Return(nil)

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Perform the action
	result, err := useCase.DelCostumer(1)

	// Assert the result
	assert.NoError(t, err)
	assert.True(t, result)

	// Assert that the mock was called with the expected argument
	mockData.AssertCalled(t, "Delete", 1)
}

func TestDelCostumer_Error(t *testing.T) {
	// Prepare mock
	mockData := new(MockCostumerData)
	mockData.On("Delete", 1).Return(errors.New("delete error"))

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Perform the action
	result, err := useCase.DelCostumer(1)

	// Assert the result
	assert.Error(t, err)
	assert.False(t, result)

	// Assert that the mock was called with the expected argument
	mockData.AssertCalled(t, "Delete", 1)
}

func TestGetSpecificCostumer_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockCostumerData)
	mockData.On("GetCostumerID", 1).Return([]domain.Costumer{
		{
			ID:           1,
			NIK:          "1234567890",
			FullName:     "John Doe",
			LegalName:    "John",
			PlaceOfBirth: "New York",
			DateOfBirth:  sql.NullTime{},
			Salary:       6000.0,
			KTPPhoto:     "ktp_photo.jpg",
			SelfiePhoto:  "selfie_photo.jpg",
		},
	})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Perform the action
	result, err := useCase.GetSpecificCostumer(1)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, []domain.Costumer{
		{
			ID:           1,
			NIK:          "1234567890",
			FullName:     "John Doe",
			LegalName:    "John",
			PlaceOfBirth: "New York",
			DateOfBirth:  sql.NullTime{},
			Salary:       6000.0,
			KTPPhoto:     "ktp_photo.jpg",
			SelfiePhoto:  "selfie_photo.jpg",
		},
	}, result)

	// Assert that the mock was called with the expected argument
	mockData.AssertCalled(t, "GetCostumerID", 1)
}
