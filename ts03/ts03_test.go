package main

import (
	"slices"
	"testing"
)

func TestSearchSimpleNumber(t *testing.T) {
	someStruct := []struct {
		want []int
		min  int
		max  int
	}{
		{
			want: []int{11, 13, 17, 19},
			min:  11,
			max:  20,
		},
		{
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31},
			min:  2,
			max:  35,
		},
	}
	for _, v := range someStruct {
		got := SearchSimpleNumber(v.min, v.max)
		if slices.Compare(got, v.want) != 0 {
			t.Errorf("want %v but got %v", v.want, got)
		}
	}
}
