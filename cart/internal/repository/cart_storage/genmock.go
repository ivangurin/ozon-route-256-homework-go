//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package cartstorage

// ClientMock ...
type StorageMock interface {
	Storage
}
