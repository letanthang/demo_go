package main

import "testing"

func TestAsChan(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, 6}
	ch := AsChan(intSlice...)
	seen := map[int]struct{}{}

	// test duplicate
	for v := range ch {
		if _, ok := seen[v]; ok {
			t.Errorf("duplicate value %d", v)
		}
		seen[v] = struct{}{}
	}

	// test seen all
	for _, v := range intSlice {
		if _, ok := seen[v]; !ok {
			t.Errorf("miss value %d", v)
		}
	}
}

func TestMergeChannel(t *testing.T) {
	ch1 := AsChan(1, 2, 3)
	ch2 := AsChan(4, 5, 6)
	ch3 := AsChan(7, 8, 9)

	mCh := MergeChannel(ch1, ch2, ch3)

	seen := map[int]struct{}{}
	// test duplicate
	for v := range mCh {
		if _, ok := seen[v]; ok {
			t.Errorf("duplicate value %d", v)
		}
		seen[v] = struct{}{}
	}

	// test seen all
	for v := 0; v < 10; v++ {
		if _, ok := seen[v]; !ok {
			t.Errorf("miss value %d", v)
		}
	}

}
