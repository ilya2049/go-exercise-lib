package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6}

	p, _ := stats.PercentileNearestRank(data, 50)

	fmt.Print(p)
}
