package closer

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCloser(t *testing.T) {

	t.Parallel()

	closer := NewCloser(os.Interrupt)

	closed := false

	closer.Add(
		func() error {
			closed = true
			return nil
		},
		func() error {
			return errors.New("test")
		})

	closer.Signal()

	closer.Wait()

	require.True(t, closed, "Должно быть закрыто")

}
