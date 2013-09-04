package pi_test

import (
	"github.com/rafkhan/pi"
	"testing"
)

// Tests to make sure Setup returns an error
// if no flags are specified
func TestSetupFailure(t *testing.T) {
	if err := pi.Setup(0); err == nil {
		t.Fail();
	}
}


