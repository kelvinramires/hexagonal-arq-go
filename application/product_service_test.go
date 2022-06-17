package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/kelvinramires/hexagonal-arq-go/application"
	mock_application "github.com/kelvinramires/hexagonal-arq-go/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("testing")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
