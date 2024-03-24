//go:generate mkdir -p ./mocks
//go:generate minimock -i "ServiceMock" -o ./mocks/ -s ".go" -g
package stockservice

// ServiceMock ...
type ServiceMock interface {
	Service
}
