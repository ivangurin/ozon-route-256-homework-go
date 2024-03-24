//go:generate mkdir -p ./mocks
//go:generate minimock -i "RepositoryMock" -o ./mocks/ -s ".go" -g
package orderstorage

// RepositoryMock ...
type RepositoryMock interface {
	Repository
}
