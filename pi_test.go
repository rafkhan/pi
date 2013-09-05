package pi_test

import (
	"time"
	"testing"
	"github.com/rafkhan/pi"
)

// Tests to make sure Setup returns an error
// if no flags are specified
func TestSetupFailure(t *testing.T) {
	if err := pi.Setup(0); err == nil {
		t.Fail();
	}
}

func TestWiring(t *testing.T) {
	pi.Setup(pi.SETUP_WIRING);
	blink(7, 300);
}

func TestGpio(t *testing.T) {
	pi.Setup(pi.SETUP_GPIO);
	blink(4, 100);
}


func blink(pin int, delay time.Duration) {
	for i := 0; i < 5; i++ {
		pi.Write(pin, pi.HIGH);
		time.Sleep(delay * time.Millisecond);
		pi.Write(pin, pi.LOW);
		time.Sleep(delay * time.Millisecond);
	}
}
