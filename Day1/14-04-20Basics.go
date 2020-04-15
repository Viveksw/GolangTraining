package main

import "fmt"

func main() {

	/*
		2.
			*
			**
			***
			****
			*****
	*/
	numOfIter := 5
	var i int
	for i <= numOfIter {
		var j int
		for j <= i {
			fmt.Print("*")
			j++
		}
		fmt.Println()
		i++
	}

	/*
		3.
			1
			2 3
			4 5 6
			7 8 9 10
			11 12 13 14
	*/
	i = 0
	number := 1
	for i <= numOfIter {
		var j int
		for j <= i {
			fmt.Printf("%d ", number)
			number++
			j++
		}
		fmt.Println()
		i++
	}

	/*
		4. factorial
	*/
	givenNum := 8
	i = 1
	factorialResult := 1
	for i <= givenNum {
		factorialResult *= i
		i++
	}
	fmt.Println(factorialResult)

	/*
		4. Prime number
	*/
	givenNum = 9
	i = 2
	givenNumHalf := givenNum / 2
	fmt.Println(givenNumHalf)
	isPrime := true
	for i < givenNumHalf {
		fmt.Println(givenNumHalf % i)
		if givenNumHalf%i == 0 {
			isPrime = false
			break
		}
		i++
	}
	if isPrime {
		fmt.Printf("%d is prime number", givenNum)
	} else {
		fmt.Printf("%d is not prime number", givenNum)
	}

}
