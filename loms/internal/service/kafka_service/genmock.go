//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package kafka_service

// ServiceMock ...
type ServiceMock interface {
	Service
}
