//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package cartservice

// ClientMock ...
type ClientMock interface {
	Client
}
