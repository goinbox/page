package page

import (
	"testing"
)

func TestPage(t *testing.T) {
	po, err := NewPageObj(2, 15, 2)
	if err != nil {
		t.Error(err)
	}

	t.Log(po)

	po.Rewind()
	t.Log(po)
	for po.Next() == nil {
		t.Log(po)
	}
}
