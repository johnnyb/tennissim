package main

import "fmt"

func main() {
	p1Stats := PlayerStats{
		FirstServeRatio:  0.8,
		SecondServeRatio: 0.8,
		HitRatio:         0.8,
	}
	p2Stats := PlayerStats{
		FirstServeRatio:  0.8,
		SecondServeRatio: 0.8,
		HitRatio:         0.7,
	}

	for i := 0; i < 100; i++ {
		m := CreateMatch(&p1Stats, &p2Stats)
		m.Play()
		fmt.Printf("%s\n", m.ScoreString())
	}
}
