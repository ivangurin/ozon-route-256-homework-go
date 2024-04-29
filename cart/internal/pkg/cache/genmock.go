//go:generate mockery --name=(.+)Mock --case=underscore  --with-expecter
package cache

// CacheMock ...
type CacheMock interface {
	Cache
}
