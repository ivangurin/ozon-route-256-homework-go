//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package kafka_storage

// RepositoryMock ...
type RepositoryMock interface {
	Repository
}
