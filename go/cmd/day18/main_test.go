package main

import (
	"testing"
)

func TestEval(t *testing.T) {
	cases := []struct{
		name string
		inp string
		start int
		want int
	}{
		{
			"test-1", "1 + 2 * 3 + 4 * 5 + 6",
			0, 71,
		},
		{
			"test-2", "1 + (2 * 3) + (4 * (5 + 6))",
			0, 51,
		},
		{
			"test-3", "2 * 3 + (4 * 5",
			0, 26,
		},
		{
			"test-4", "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			0, 437,
		},
		{
			"test-5", "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			0, 12240,
		},
		{
			"test-6", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			0, 13632,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T){
			res, _ := eval(test.inp, test.start)
			if res != test.want {
				t.Fatalf("eval returned %d, but expected %d\n", res, test.want)
			}
		})
	}
}

func BenchmarkPart1(b *testing.B) {
	//file := "../../../inputs/day18/input.txt"
	for i := 0; i < b.N; i++ {
	}
}
