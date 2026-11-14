package main


import "testing"

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	data := generateRandomElements(0)

	if len(data) != 0 {
		t.Errorf("ожидалась длина 0, получено %d", len(data))
	}
}

func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	size := 1000
	data := generateRandomElements(size)

	if len(data) != size {
		t.Errorf("ожидалась длина %d, получено %d", size, len(data))
	}

	for _, value := range data {
		if value <= 0 {
			t.Errorf("ожидалось положительное число, получено %d", value)
		}
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	got := maximum([]int{})
	want := 0

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaximumSingleElement(t *testing.T) {
	got := maximum([]int{42})
	want := 42

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaximumSeveralElements(t *testing.T) {
	got := maximum([]int{3, 8, 1, 7})
	want := 8

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaximumNilSlice(t *testing.T) {
	got := maximum(nil)
	want := 0

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaxChunksEmptySlice(t *testing.T) {
	got := maxChunks([]int{})
	want := 0

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaxChunksNormalSize(t *testing.T) {
	data := generateRandomElements(10_000)
	got := maxChunks(data)
	want := maximum(data)

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaxChunksLessThanChunks(t *testing.T) {
	got := maxChunks([]int{5, 2, 9, 1})
	want := 9

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}

func TestMaxChunksNotDivisibleSize(t *testing.T) {
	got := maxChunks([]int{4, 6, 2, 9, 1, 7, 3, 5, 8})
	want := 9

	if got != want {
		t.Errorf("ожидалось %d, получено %d", want, got)
	}
}