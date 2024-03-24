//go:generate mkdir -p ./mocks
//go:generate minimock -i "ServiceMock" -o ./mocks/ -s ".go" -g
package orderservice

// ServiceMock ...
type ServiceMock interface {
	Service
}
