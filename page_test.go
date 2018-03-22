package page

import (
	"testing"
)

func TestPage(t *testing.T) {
	po, err := NewPageObj(2, 5, 2)
	if err != nil {
		t.Error(err)
	}

	t.Log(po)
}
