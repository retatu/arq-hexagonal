package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/retatu/arq-hexagonal/application"
	"github.com/retatu/arq-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("valid_id")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
