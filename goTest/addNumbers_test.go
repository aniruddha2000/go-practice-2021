package main

import "testing"

type AddNumbersTest struct {
	a   int
	b   int
	res int
}

var testData = []AddNumbersTest{
	{1, 2, 3},
	{5, 5, 10},
}

func TestAddNumbers(t *testing.T) {
	for _, test := range testData {
		res := AddNumbers(test.a, test.b)
		if res != test.res {
			t.Fatal("Not passed")
		}
	}
}
