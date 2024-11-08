package mod1_test

import (
	"testing"

	"example.com/myapp/mod1"
)

func TestFuncs(t *testing.T) {
	t.Run("mod1", func(t *testing.T) {
		t.Parallel()
		mod1.F1()
		mod1.F2()
	})
}
