//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package productservice

// ClientMock ...
type ClientMock interface {
	Client
}
