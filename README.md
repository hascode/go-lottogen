# Go Lotto Generator

A simple lotto number generator written in Go.

## Building


```
go build -o lotto

```

## Using

```
package main

import "lotto"

func main() {
	lottery := lotto.Lottery{Take: 5, Of: 50, Take2: 2, Of2: 10}
	lottery.Run()
}
```

-----

**2016 Micha Kops / hasCode.com**