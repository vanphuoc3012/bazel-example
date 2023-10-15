package go_hello_world

import "testing"

func TestGreeter(t *testing.T) {
	expected := "Hello World!! Phuoc"
	actual := HelloWorld("Phuoc")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
