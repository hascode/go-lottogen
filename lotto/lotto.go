package lotto

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Lottery creates a lottery
type Lottery struct {
	Take  int
	Of    int
	Take2 int
	Of2   int
}

// Run runs the lottery
func (l *Lottery) Run() {
	fmt.Printf("## Lotto %v aus %v ##\n", l.Take, l.Of)
	fmt.Println("- Hauptzahlen")
	nums := roll(l.Take, l.Of)
	printRolled(nums)

	fmt.Println("- Zusatzzahlen")
	nums = roll(l.Take2, l.Of2)
	printRolled(nums)
}

func roll(take int, of int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	nums := make([]int, take)

	var i int
	for i < take {
		num := rand.Intn(of-1) + 1
		if !alreadyExists(num, nums) {
			nums[i] = num
			i++
		}
	}

	sort.Sort(sort.IntSlice(nums))
	return nums
}

func printRolled(nums []int) {
	for _, val := range nums {
		fmt.Printf("  %2d\n", val)
	}
}

func alreadyExists(num int, numbers []int) (exists bool) {
	for _, x := range numbers {
		if num == x {
			exists = true
			return
		}
	}
	return
}
