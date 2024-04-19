//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package kafka

// ServiceMock ...
type ProducerMock interface {
	Producer
}
