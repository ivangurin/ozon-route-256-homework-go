//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package productservice

// ServiceMock ...
type ClientMock interface {
	IClient
}
