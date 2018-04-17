package mystring

import "testing"

type Sample struct {
	input string
	want  string
}

func Test(t *testing.T) {
	var tests = []Sample{
		{input: "Hello", want: "olleH"},
		{input: "Hi Baskar!", want: "!raksaB iH"},
		{input: "This will fail!", want: "!liaf lliw siht"},
		{input: "", want: ""},
		{input: "Hello, 世界", want: "界世 ,olleH"},
	}

	for _, v := range tests {
		got := Reverse(v.input)
		if got == v.want {
			t.Logf("Pass for input: %v", v.input)
		} else {
			t.Errorf("Failed for input: %v | Expected: %v | Actual: %v", v.input, v.want, got)
		}
	}

}
