package main

import "fmt"
import "container/ring"

func main() {
	fmt.Println(play(10, 1618))
	fmt.Println(play(13, 7999))
	fmt.Println(play(17, 1104))
	fmt.Println(play(21, 6111))
	fmt.Println(play(30, 5807))
	fmt.Println(play(448, 71628))
	fmt.Println(play(448, 71628*100))
}

//Thanks to fharding1's solution:
//https://github.com/fharding1/adventofcode-2018/blob/master/day9/main.go
// I learned about container/ring!
func play(players, lastValue int) int {
	scores := make([]int, players)
	circle := ring.New(1)
	circle.Value = 0

	for marble := 1; marble <= lastValue; marble++ {
		if marble%23 == 0 {
			circle = circle.Move(-7)
			scores[marble%players] += marble + circle.Value.(int)

			circle = circle.Prev()
			circle.Unlink(1)
			circle = circle.Next()

			continue
		}

		new := ring.New(1)
		new.Value = marble

		circle = new.Link(circle.Next().Next())
	}

	var highscore int
	for _, v := range scores {
		if v > highscore {
			highscore = v
		}
	}

	return highscore
}
