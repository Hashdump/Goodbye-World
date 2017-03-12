/******************************************************************************/
/*This code uses a logistic map for population growth and some other trickery */
/*because I understand math and code better than simplicity. By they time this*/
/*finishes running, it should print out "Hello World"                         */
/******************************************************************************/

/*
to run this program either compile it:
	$go build population_model.go
then run it
OR run it without compiling:
	$go run population_model.go
*/

package main

import (
	"fmt"
)

/* This where the actual math calculations happen for the iterated map */
func iterate_population(pop float64, u float64, a float64) float64 {
	pop = (u / a) * pop * (a - pop)
	return pop
}

/* This converts the float value that is input to the string it most closely represents */
/* for example 32.004234 is a space because space is 0x20 in hex */
func convert_text(pop float64) string {
	var i int64 = int64(pop)
	var text string = ""
	i = i >> 8                /* Need to get rid of small populations (as in too small for a character) */
	for ; i > 0; i = i >> 8 { //get as many characters out of the float number as we can
		text = text + string(byte(i))
	}
	return text
}

func main() {
	/* Starting population for the bacteria */
	var population1 float64 = 2 //start with a population of 2
	var population2 float64 = 2 //start with a population of 2

	/* Mu is the rate for growth */
	/* rates between 0 and 1 will die out */
	/* rates between 1 and 2 will quickly approach (mu-1)/mu*carrying capacity */
	/* rates between 2 and 3 will will oscilate around and eventually converge to (mu-1)/mu * carrying capacity */
	/* The closer to three you get, the slower it will converge */
	/* rates between 3 and 4 are weird... look them up because they are cool, but will never converge */
	var mu float64 = 2.99999 //More nines = more fun

	/* The output of the logistic map is between 0 and 1, so this alpha is used to shift these values to realistic populations
	   The following calculations will generate two carrying capacities for the two populations */
	var alpha1 float64 = 0x6F6C6C654820 / ((mu - 1) / mu)
	var alpha2 float64 = 0x646C726F5720 / ((mu - 1) / mu)
	for { //forever
		population1 = iterate_population(population1, mu, alpha1)
		population2 = iterate_population(population2, mu, alpha2)
		//fmt.Printf("Bacterial Population1: %v  Bacterial Population2: %v\n", int(population1), int(population2))
		fmt.Println(convert_text(population1) + " " + convert_text(population2))
	}
}
