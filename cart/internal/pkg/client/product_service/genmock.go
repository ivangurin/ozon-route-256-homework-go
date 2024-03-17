//go:generate mkdir -p ./mocks
//go:generate minimock -i "ClientMock" -o ./mocks/ -s ".go" -g
package productservice

// ClientMock ...
type ClientMock interface {
	IClient
}
