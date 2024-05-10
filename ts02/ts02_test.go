package main

import (
	"reflect"
	"testing"
)

func TestNOD(t *testing.T) {
	someStruct := []struct {
		want  []int
		array []int
	}{
		{
			want:  []int{2, 3, 6},
			array: []int{42, 12, 18},
		},
		{
			want:  []int{2, 3, 6},
			array: []int{36, 60, 42},
		},
	}

	for _, v := range someStruct {
		got := NOD(v.array)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("want %v got %v", v.want, got)
		}
	}
}
