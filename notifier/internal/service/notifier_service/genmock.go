//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package notifierservice

// ServiceMock ...
type ServiceMock interface {
	Service
}
