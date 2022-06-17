package application_test

import (
	"github.com/kelvinramires/hexagonal-arq-go/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Papi"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price must be greater than 0 to enable", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Sopa"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "price must be 0 to disable", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Vanilla"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "status must be enabled/disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "price must be greater or equal to 0", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	var testId = uuid.NewV4().String()
	product.ID = testId
	require.Equal(t, testId, product.GetId())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Papi"

	require.Equal(t, "Papi", product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}

	product.Status = application.DISABLED
	require.Equal(t, application.DISABLED, product.GetStatus())

	product.Status = application.ENABLED
	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}

	product.Price = 10.0
	require.Equal(t, 10.0, product.GetPrice())
}
