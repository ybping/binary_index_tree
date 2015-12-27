package tree

import "testing"

func TestBinaryIndexTree(t *testing.T) {
	bit := NewBinaryTreeArrary(10)
	sum := bit.Sum(5)
	if sum != 0 {
		t.Error(t, "Sum[0:5] Error! Expected ", 0, " got ", sum)
	}

	for total, i := 0, 1; i <= 5; i++ {
		total += i
		bit.Update(i, i)
		sum = bit.Sum(6)
		bit.Print()
		if sum != total {
			t.Error(t, "Sum[0:5] Error! Expected ", total, " got ", sum)
		}
	}
}
