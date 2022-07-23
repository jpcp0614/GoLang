package main

import (
	"log"
	"math"
	"fmt"
)

type norgateMathError struct {
	lat		string
	long	string
	err		error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{lat: "50N", long: "40W", err: nme}
	}
	return math.Sqrt(f), nil
}

// 2022/07/23 13:42:32 a norgate math error occurred: 50N 40W norgate math redux: square root of negative number: -10.23