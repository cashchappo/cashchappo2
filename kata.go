package main

import (
	"fmt"
)

func main() {

	dadYearsOld := 24
	sonYearsOld := 5

	a1 := TwiceAsOld(dadYearsOld, sonYearsOld)
	a2 := -a1
	fmt.Println(a2)
	fmt.Println(a1)

}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {

	years := sonYearsOld * 2
	years1 := dadYearsOld - years

	return years1
}
