//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package orderservice

// ServiceMock ...
type ServiceMock interface {
	Service
}
