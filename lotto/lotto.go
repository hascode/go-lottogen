package lotto

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Roll is a specific area of numbers and part of a lottery
type Roll struct {
	Take  int
	Lower int
	Upper int
	Title string
}

// Roll rolls a roll :P
func (r *Roll) Roll() {
	fmt.Printf("# %v - %v aus %v bis %v\n", r.Title, r.Take, r.Lower, r.Upper)
	nums := singleRoll(r.Take, r.Lower, r.Upper)
	printRolled(nums)
}

// Lottery specifies a set of specific rolls
type Lottery struct {
	Name  string
	Rolls []Roll
}

// NewLottery creates a new lottery
func newLottery(name string, rolls ...Roll) Lottery {
	return Lottery{Name: name, Rolls: rolls}
}

// Run runs the lottery
func (l *Lottery) run() {
	fmt.Printf("###### %v\n", l.Name)

	for _, roll := range l.Rolls {
		roll.Roll()
	}
}

func singleRoll(take int, lower int, upper int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	nums := make([]int, take)
	numRange := upper - lower

	var i int
	for i < take {
		num := rand.Intn(numRange-1) + 1
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

// RunCli starts the lottery selection in the command line
func RunCli() {
	lotteries := []Lottery{
		newLottery("Lotto 6 aus 49", Roll{Title: "Hauptzahlen", Take: 6, Lower: 1, Upper: 49}, Roll{Title: "Zusatzzahlen", Take: 1, Lower: 0, Upper: 9}),
		newLottery("Euro-Lotto", Roll{Title: "Hauptzahlen", Take: 5, Lower: 1, Upper: 50}, Roll{Title: "Zusatzzahlen", Take: 2, Lower: 1, Upper: 10}),
	}
	fmt.Println("Please select your lottery")
	for i, lottery := range lotteries {
		fmt.Printf("%v) %v\n", i+1, lottery.Name)
	}

	var selection int
	fmt.Scanf("%d", &selection)

	fmt.Println("How many runs?")
	var runs int
	fmt.Scanf("%d", &runs)

	l := lotteries[selection-1]

	for i := 0; i < runs; i++ {
		l.run()
	}

}
