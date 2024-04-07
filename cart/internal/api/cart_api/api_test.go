package cartapi

import (
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
)

func TestGetDescription(t *testing.T) {

	t.Parallel()

	sp := suite.NewSuiteProvider()
	api := NewAPI(
		sp.GetCartServiceMock(),
	)

	desc := api.GetDescription()

	require.Len(t, desc.Handlers, 5, "Должно быть 5 хендлеров")

}
