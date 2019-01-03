package luckykids

import "testing"

func Test_Sample1(t *testing.T) {
	behaviors := []int{3, 3, 4, 5, 2, 1, 5, 5}

	got := luckyKids(behaviors)
	if got != 2 {
		t.Errorf("luckyKids() = %v, want %v", got, 2)
	}
}

func Test_Sample2(t *testing.T) {
	behaviors := []int{5, 5, 5, 5}

	got := luckyKids(behaviors)
	want := 1
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample3(t *testing.T) {
	behaviors := []int{10, 100, 1}

	got := luckyKids(behaviors)
	want := 3
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample4(t *testing.T) {
	behaviors := []int{50, 10000, 50, 7}

	got := luckyKids(behaviors)
	want := 3
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample5(t *testing.T) {
	behaviors := []int{7, 8, 9, 8, 7, 6, 5}

	got := luckyKids(behaviors)
	want := 6
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample6(t *testing.T) {
	behaviors := []int{7, 101, 100, 101, 101, 70, 70, 60, 5, 50, 40, 4, 40, 101, 100}

	got := luckyKids(behaviors)
	want := 10
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample7(t *testing.T) {
	behaviors := []int{1, 1, 1, 1, 1, 4, 4, 4, 4, 4, 1, 1, 1, 1, 1, 4, 4, 4, 4, 4, 1, 1, 1, 1, 1}

	got := luckyKids(behaviors)
	want := 11
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample8(t *testing.T) {
	behaviors := []int{8, 7, 6, 5, 4, 3, 2, 1}

	got := luckyKids(behaviors)
	want := 8
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample9(t *testing.T) {
	behaviors := []int{5, 8, 8, 17, 20, 2, 6, 4, 4, 15, 7, 9, 9, 9, 1, 6, 7, 2, 2, 8, 4, 1, 100, 33, 20, 8}

	got := luckyKids(behaviors)
	want := 12
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample10(t *testing.T) {
	behaviors := []int{2}

	got := luckyKids(behaviors)
	want := 1
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample11(t *testing.T) {
	behaviors := []int{100000, 99999, 100000, 99999, 99998}

	got := luckyKids(behaviors)
	want := 4
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}

func Test_Sample12(t *testing.T) {
	behaviors := []int{7, 11, 10, 9, 5, 7, 5, 7, 6, 6, 6, 2}

	got := luckyKids(behaviors)
	want := 9
	if got != want {
		t.Errorf("luckyKids() = %v, want %v", got, want)
	}
}
