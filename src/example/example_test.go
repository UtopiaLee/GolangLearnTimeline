package example

import "testing"

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add(2,4) error,expect:%d,actual:%d", 6, r)
	}
	t.Logf("test add succ")
}

func TestSub(t *testing.T) {
	s := sub(6, 2)
	if s != 4 {
		t.Fatalf("sub(6,2) error ,expected:%d,actual:%d", 4, s)
	}
	t.Logf("test sub success")
}
