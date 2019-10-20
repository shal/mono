package mono

import (
	"testing"
)

func TestNewPersonal(t *testing.T) {
	personal := NewPersonal("token")

	if personal == nil {
		t.Error()
	}
}
