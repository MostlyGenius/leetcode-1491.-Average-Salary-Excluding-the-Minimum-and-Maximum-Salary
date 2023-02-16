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
		} else if maxS < salary[i] {
			maxS = salary[i]
		}
		i++
	}
	//find mean
	averageSalary := 0

	//add up all items
	for i < len(salary) {
		averageSalary = averageSalary + salary[i]
	}
	//subtrack and low
	averageSalary = averageSalary - (minS + maxS)
	//find mean
	averageSalary = averageSalary / 2
	fmt.Println(averageSalary)

}

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	average(salary)

}
