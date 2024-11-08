package mod1_test

import (
	"testing"

	"example.com/myapp/mod1"
)

func TestIT(t *testing.T) {
	t.Run("mod1", func(t *testing.T) {
		t.Parallel()
		mod1.F3()
	})
}
