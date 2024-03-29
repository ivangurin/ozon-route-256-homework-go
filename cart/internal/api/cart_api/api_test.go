package cartapi

import (
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
)

func TestGetDescription(t *testing.T) {

	sp := suite.NewSuiteProvider(t)
	api := NewAPI(
		sp.GetCartServiceMock(),
	)

	desc := api.GetDescription()

	require.Len(t, desc.Handlers, 5, "Должно быть 5 хендлеров")

}
