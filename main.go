package main

import "fmt"

func average(salary []int) {

	i := 0
	minS := salary[0]
	maxS := salary[0]

	//to find max and min
	for i < len(salary) {
		if minS > salary[i] {
			minS = salary[i]
		}
		if maxS < salary[i] {
			maxS = salary[i]
		}
		i++
	}
	fmt.Println(minS)
	fmt.Println(maxS)

}

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	average(salary)

}
