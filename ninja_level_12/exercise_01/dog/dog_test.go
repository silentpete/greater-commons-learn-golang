package dog

import "testing"

func TestYears(t *testing.T) {
	got := Years(1)
	expect := 7
	if got != expect {
		t.Error("got:", got, "expect", expect)
	}
}
