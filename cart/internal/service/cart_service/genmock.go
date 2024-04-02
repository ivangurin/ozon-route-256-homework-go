//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package cartservice

// ServiceMock ...
type ServiceMock interface {
	Service
}
