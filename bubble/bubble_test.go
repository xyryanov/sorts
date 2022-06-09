package bubble

import (
	"testing"
	"sort"
	"time"
	"math/rand"
)

func TestSort(t *testing.T) {
	unsorted := randArray()
	Sort(unsorted)
	
	if !sort.IntsAreSorted(unsorted) {
		t.Error("array isnt sorted")
	}
}

func randArray() []int {
	rand.Seed(time.Now().UnixNano())
	var arr []int

	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		arr = append(arr, n)
	}

	return arr
}