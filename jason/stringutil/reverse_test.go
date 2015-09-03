package stringutil

import "testing"
import "fmt"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, iworld", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		fmt.Printf("\n")
		fmt.Printf(c.want)
		fmt.Printf("\n")
		fmt.Printf(c.in)
		fmt.Printf("\n")
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
