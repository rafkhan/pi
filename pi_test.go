package pi_test

import (
	"testing"

	"github.com/rafkhan/pi"
)

// Tests to make sure Setup returns an error
// if no flags are specified

/*
func TestSetupFailure(t *testing.T) {
	if err := pi.Setup(0); err == nil {
		t.Fail();
	}
}
*/

func TestBlink(t *testing.T) {
//	pi.Setup(pi.SETUP_GPIO);
	pi.Test()

	/*
	for i := 0; i < 5; i++ {
		pi.Write(8, 1);
		time.Sleep(500 * time.Millisecond);
		pi.Write(8, 0);
		time.Sleep(500 * time.Millisecond);
	}
	*/
}
