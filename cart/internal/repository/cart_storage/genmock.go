//go:generate mkdir -p ./mocks
//go:generate minimock -i "StorageMock" -o ./mocks/ -s ".go" -g
package cartstorage

// ClientMock ...
type StorageMock interface {
	IStorage
}
