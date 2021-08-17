package utils

import "testing"

func TestQuickParsePanics(t *testing.T) {
	// No need to check whether `recover()` is nil. Just turn off the panic.
	defer func() { recover() }()

	QuickParse("abc")

	// Never reaches here if `utils.QuickParse` panics.
	t.Errorf("did not panic")
}
