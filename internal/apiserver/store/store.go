package store

//go:generate mockgen -self_package=github.com/marmotedu/iam/internal/apiserver/store -destination mock_store.go -package store github.com/daz2yy/go-base/internal/apiserver/store Factory

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	// Users() UserStore
	// Secrets() SecretStore
	// Policies() PolicyStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
