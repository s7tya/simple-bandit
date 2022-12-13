package main

import (
	"fmt"
	"math/rand"
	"simple-bandit/arm"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

func randomGreedy(arms []arm.Arm, count int) int {
	reward := 0

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(arms))
		reward += arms[index].Play()
	}

	return reward
}

func epsilonGreedy(arms []arm.Arm, count int, epsilon float64) int {
	reward := 0

	for i := 0; i < count; i++ {
		b := &distuv.Binomial{N: 1, P: epsilon}
		index := 0
		if int(b.Rand()) == 1 {
			index = rand.Intn(len(arms))
		} else {
			index = selectHighestArmIndex(arms)
		}

		reward += arms[index].Play()
	}

	return reward
}

func selectHighestArmIndex(arms []arm.Arm) int {
	if len(arms) == 1 {
		return 0
	} else {
		highestArmIndex := 0

		for i := 0; i < len(arms)-1; i++ {
			if arms[i].CalcSuccess() < arms[i+1].CalcSuccess() {
				highestArmIndex = i + 1
			}
		}

		return highestArmIndex
	}
}

func main() {
	count := 10000

	arm1 := arm.New(0.5)
	arm2 := arm.New(0.2)
	arm3 := arm.New(0.7)

	arms := []arm.Arm{*arm1, *arm2, *arm3}

	fmt.Println("Random: ", randomGreedy(arms, count))

	epsilon := 0.4
	fmt.Println("Epsilon: ", epsilonGreedy(arms, count, epsilon))
}
