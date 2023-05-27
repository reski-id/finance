package usecase_test

import (
	"errors"
	"finance/domain"
	"finance/feature/limit/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define mock structs

type MockLimitData struct {
	mock.Mock
}

func (m *MockLimitData) Insert(insertLimit domain.Limit) domain.Limit {
	args := m.Called(insertLimit)
	return args.Get(0).(domain.Limit)
}

func (m *MockLimitData) Update(IDLimit int, updatedLimit domain.Limit) domain.Limit {
	args := m.Called(IDLimit, updatedLimit)
	return args.Get(0).(domain.Limit)
}

func (m *MockLimitData) Delete(IDLimit int) error {
	args := m.Called(IDLimit)
	return args.Error(0)
}

func (m *MockLimitData) GetLimitID(dataID int) []domain.Limit {
	if dataID == -1 {
		return nil
	}
	result := m.Called(dataID)
	if result.Get(0) == nil {
		return nil
	}
	return result.Get(0).([]domain.Limit)
}

// Test cases

func TestAddLimit_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockLimitData)
	mockData.On("Insert", mock.AnythingOfType("domain.Limit")).Return(domain.Limit{
		ID:         1,
		CustomerID: 1,
		Tenor1:     100000,
		Tenor2:     200000,
		Tenor3:     700000,
		Tenor6:     800000,
	})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	newLimit := domain.Limit{ID: 1, CustomerID: 1, Tenor1: 100000, Tenor2: 200000, Tenor3: 700000, Tenor6: 800000}

	// Perform the action
	result, err := useCase.AddLimit(newLimit)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, domain.Limit{
		ID:         1,
		CustomerID: 1,
		Tenor1:     100000,
		Tenor2:     200000,
		Tenor3:     700000,
		Tenor6:     800000,
	}, result)

	// Assert that the mock was called with the expected arguments
	mockData.AssertCalled(t, "Insert", newLimit)
}

func TestAddLimit_Error(t *testing.T) {
	// Prepare mock
	mockData := new(MockLimitData)
	mockData.On("Insert", mock.AnythingOfType("domain.Limit")).Return(domain.Limit{})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	newLimit := domain.Limit{ID: 0, CustomerID: 1, Tenor1: 0.5, Tenor2: 0.6, Tenor3: 0.7, Tenor6: 0.8}

	// Perform the action
	result, err := useCase.AddLimit(newLimit)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, domain.Limit{}, result)

	// Assert that the mock was called with the expected arguments
	mockData.AssertCalled(t, "Insert", newLimit)
}

func TestUpLimit_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockLimitData)
	mockData.On("Update", 1, mock.AnythingOfType("domain.Limit")).Return(domain.Limit{
		ID:         1,
		CustomerID: 1,
		Tenor1:     200000,
		Tenor2:     300000,
		Tenor3:     800000,
		Tenor6:     900000,
	})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	IDLimit := 1
	updateData := domain.Limit{ID: 1, CustomerID: 1, Tenor1: 200000, Tenor2: 300000, Tenor3: 800000, Tenor6: 900000}

	// Perform the action
	result, err := useCase.UpLimit(IDLimit, updateData)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, domain.Limit{
		ID:         1,
		CustomerID: 1,
		Tenor1:     200000,
		Tenor2:     300000,
		Tenor3:     800000,
		Tenor6:     900000,
	}, result)

	// Assert that the mock was called with the expected arguments
	mockData.AssertCalled(t, "Update", IDLimit, updateData)
}

func TestUpLimit_InvalidData(t *testing.T) {
	// Prepare mock
	mockData := new(MockLimitData)

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Prepare test data
	IDLimit := -1
	updateData := domain.Limit{ID: 1, CustomerID: 1, Tenor1: 200000, Tenor2: 300000, Tenor3: 800000, Tenor6: 900000}

	// Perform the action
	result, err := useCase.UpLimit(IDLimit, updateData)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, domain.Limit{}, result)

	// Assert that the mock was not called
	mockData.AssertNotCalled(t, "Update")
}

func TestDelLimit_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockLimitData)
	mockData.On("Delete", 1).Return(nil)

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Perform the action
	result, err := useCase.DelLimit(1)

	// Assert the result
	assert.NoError(t, err)
	assert.True(t, result)

	// Assert that the mock was called with the expected argument
	mockData.AssertCalled(t, "Delete", 1)
}

func TestDelLimit_Error(t *testing.T) {
	// Prepare mock
	mockData := new(MockLimitData)
	mockData.On("Delete", 1).Return(errors.New("delete error"))

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Perform the action
	result, err := useCase.DelLimit(1)

	// Assert the result
	assert.Error(t, err)
	assert.False(t, result)

	// Assert that the mock was called with the expected argument
	mockData.AssertCalled(t, "Delete", 1)
}

func TestGetSpecificLimit_Success(t *testing.T) {
	// Prepare mock
	mockData := new(MockLimitData)
	mockData.On("GetLimitID", 1).Return([]domain.Limit{
		{ID: 1, CustomerID: 1, Tenor1: 100000, Tenor2: 200000, Tenor3: 300000, Tenor6: 400000},
		{ID: 2, CustomerID: 1, Tenor1: 500000, Tenor2: 600000, Tenor3: 700000, Tenor6: 800000},
	})

	// Create use case with mock
	useCase := usecase.New(mockData)

	// Perform the action
	result, err := useCase.GetSpecificLimit(1)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, []domain.Limit{
		{ID: 1, CustomerID: 1, Tenor1: 100000, Tenor2: 200000, Tenor3: 300000, Tenor6: 400000},
		{ID: 2, CustomerID: 1, Tenor1: 500000, Tenor2: 600000, Tenor3: 700000, Tenor6: 800000},
	}, result)

	// Assert that the mock was called with the expected argument
	mockData.AssertCalled(t, "GetLimitID", 1)
}
