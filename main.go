package main

import (
	"fmt"
	"math/rand"
	"os"
)

var AmoN int
var PlayerHealth int
var HostHealth int
var Amo [8]bool
var Chance bool

func main() {
	fmt.Printf("LIFE OR DEATH SHOTGUN\n")
	AmoN = 8
	PlayerHealth = 3
	HostHealth = 3
	for i := 0; i < AmoN; i++ {
		Amo[i] = rand.Intn(2) != 0
	}
	//fmt.Print(Amo)

	for {
		DoFire()
	}
}

func DoFire() {
	ShowAndCheck()
rechoose:
	fmt.Printf("your turn\nchoose to shot you or bot (y or b)\n>")
	var inp string
	fmt.Scan(&inp)
	if inp == "y" {
		if Amo[AmoN-1] {
			fmt.Printf("kaboom!")
			PlayerHealth--
			AmoN--
			ShowAndCheck()
			if HostHealth != 0 {
				HostDoFire()
			}
		} else {
			fmt.Printf("...")
			AmoN--
			ShowAndCheck()
		}
	} else if inp == "b" {
		if Amo[AmoN-1] {
			fmt.Printf("kaboom!")
			HostHealth--
			AmoN--
			ShowAndCheck()
		} else {
			fmt.Printf("...")
			AmoN--
			ShowAndCheck()
		}
		if HostHealth != 0 {
			HostDoFire()
		}
	} else {
		goto rechoose
	}
}

func HostDoFire() {
	fmt.Printf("\nbot turn\n")
	if rand.Intn(2) != 0 {
		fmt.Printf("\nbot choose you\n")
		if Amo[AmoN-1] {
			fmt.Printf("kaboom!")
			PlayerHealth--
			AmoN--
			ShowAndCheck()
		} else {
			fmt.Printf("...")
			AmoN--
			ShowAndCheck()
		}
	} else {
		fmt.Printf("\nbot choose him\n")
		if Amo[AmoN-1] {
			fmt.Printf("kaboom!")
			HostHealth--
			AmoN--
			ShowAndCheck()
		} else {
			fmt.Printf("...")
			AmoN--
			ShowAndCheck()

			HostDoFire()
		}
	}
}

func ShowAndCheck() bool {
	fmt.Printf("\n| your life | bots life |\n|     %d     |     %d     |\n\n", PlayerHealth, HostHealth)

	if HostHealth <= 0 {
		fmt.Printf("YOU WON!\n")
		Reverse()
		fmt.Scan()
		os.Exit(0)
	} else if PlayerHealth <= 0 {
		fmt.Printf("YOU DIED...\n")
		Reverse()
		fmt.Scan()
		os.Exit(0)
	} else if AmoN <= 0 {
		fmt.Printf("IT IS A TIE...\n")
		Reverse()
		fmt.Scan()
		os.Exit(0)
	}
	return true
}

func Reverse() {
	left := 0
	right := len(Amo) - 1

	for left < right {
		Amo[left], Amo[right] = Amo[right], Amo[left]
		left++
		right--
	}

	fmt.Print(Amo)
}
