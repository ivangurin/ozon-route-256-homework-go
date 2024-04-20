//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package redis

// ClientMock ...
type ClientMock interface {
	Client
}
