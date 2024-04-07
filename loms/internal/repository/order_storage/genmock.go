//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package orderstorage

// RepositoryMock ...
type RepositoryMock interface {
	Repository
}
