//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package stockservice

// ServiceMock ...
type ServiceMock interface {
	Service
}
