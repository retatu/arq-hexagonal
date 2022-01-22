package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/retatu/arq-hexagonal/adapters/cli"
	"github.com/retatu/arq-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc"
	productName := "Teste"
	productPrice := 123.2
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus)

	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(serviceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(serviceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID: %s\n Name: %s\n Price: %f\n Status: %s\n",
		productId, productName, productPrice, productStatus)
	result, err = cli.Run(serviceMock, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

}
