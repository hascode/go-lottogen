package main

import "./lotto"

func main() {
	lottery := lotto.Lottery{Take: 5, Of: 50, Take2: 2, Of2: 10}
	lottery.Run()
}
