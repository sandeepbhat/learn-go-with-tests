package sum

import (
	"reflect"
	"testing"
)

// TestSum - Tests for Sum()
func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})
}

// CheckSums - Helper to compare sums
func checkSums(t *testing.T, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d but expected %d", got, want)
	}
}

// TestSumAll - Run tests for SumAll()
func TestSumAll(t *testing.T) {

	t.Run("Test sum of numbers in all sets", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2, 3})
		want := []int{3, 5}

		checkSums(t, got, want)
	})
}

// TestSumOfAllTails - Run tests for SumOfAllTails
func TestSumOfAllTails(t *testing.T) {

	// Run test with non-empty slice
	t.Run("sum of non-empty tails", func(t *testing.T) {
		got := SumOfAllTails([]int{1, 2}, []int{2, 3})
		want := []int{2, 3}

		checkSums(t, got, want)
	})

	// Run test with empty slice
	t.Run("sum of tails in empty slices", func(t *testing.T) {
		got := SumOfAllTails([]int{}, []int{2, 3})
		want := []int{0, 3}

		checkSums(t, got, want)
	})
}

// BenchmarkSumOfAllTails - Run a benchmark for SumOfAllTails
func BenchmarkSumOfAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfAllTails([]int{1, 2}, []int{2, 3})
	}
}
