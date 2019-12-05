package main

import (
	"reflect"
	"testing"
)

func Test_numIslands2(t *testing.T) {
	type args struct {
		n         int
		m         int
		operators []*Point
	}
	tests := []struct {
		name string
		args args
		want []int
	}{{
		"test1",
		args{0, 0, []*Point{}},
		[]int{},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands2(tt.args.n, tt.args.m, tt.args.operators); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numIslands2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numIslands2_30(t *testing.T) {
	type args struct {
		n         int
		m         int
		operators []*Point
	}
	tests := []struct {
		name string
		args args
		want []int
	}{{
		"test1",
		args{0, 0, []*Point{&Point{0, 9}, &Point{5, 4}, &Point{0, 12}, &Point{6, 9}, &Point{6, 5}, &Point{0, 4}, &Point{4, 11}, &Point{0, 0}, &Point{3, 5}, &Point{6, 7}, &Point{3, 12}, &Point{0, 5}, &Point{6, 13}, &Point{7, 5}, &Point{3, 6}, &Point{4, 4}, &Point{0, 8}, &Point{3, 1}, &Point{4, 6}, &Point{6, 1}, &Point{5, 12}, &Point{3, 8}, &Point{7, 0}, &Point{2, 9}, &Point{1, 4}, &Point{3, 0}, &Point{1, 13}, &Point{2, 13}, &Point{6, 0}, &Point{6, 4}, &Point{0, 13}, &Point{0, 3}, &Point{7, 4}, &Point{1, 8}, &Point{5, 5}, &Point{5, 7}, &Point{5, 10}, &Point{5, 3}, &Point{6, 10}, &Point{6, 2}, &Point{3, 10}, &Point{2, 7}, &Point{1, 12}, &Point{5, 0}, &Point{4, 5}, &Point{7, 13}, &Point{3, 2}}},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11, 12, 12, 12, 12, 12, 13, 13, 14, 15, 16, 17, 18, 18, 18, 19, 19, 18, 17, 16, 16, 16, 16, 16, 16, 17, 17, 16, 16, 17, 18, 18, 18, 17, 17, 17},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands2(tt.args.n, tt.args.m, tt.args.operators); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numIslands2() = \n%v\n, want\n%v", got, tt.want)
			}
		})
	}
}
