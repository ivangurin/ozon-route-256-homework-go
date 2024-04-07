//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package stockstorage

// RepositoryMock ...
type RepositoryMock interface {
	Repository
}
