package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	errors "github.com/Ksualboy/Globant-bootcamp/2ndPart/errorsPackage/errors"
)

func main() {
	clear()
	fmt.Println("==========================================")
	fmt.Println("==========================================\n\n")
	fmt.Println("                 WELCOME\n\n")
	fmt.Println("==========================================")
	fmt.Println("==========================================")

	time.Sleep(time.Second * 2)
	clear()

	fmt.Printf("You are here to create errors, so feel free to break everything")
	fmt.Printf("\nPress enter to continue ")
	fmt.Scanf("%s")

	var customError error
	var uinput string

	for i := 1; customError == nil; i++ {
		clear()
		fmt.Printf("Enter a number from 1 to 10\n\n")

		_, err := fmt.Scanf("%s", &uinput)
		fmt.Println("")

		if err != nil {
			fmt.Println("We got an error! Nicely done c:")
			customError = errors.NewInternal()
			continue
		}

		inputInt, err := strconv.ParseInt(uinput, 10, 32)
		if err != nil {
			fmt.Println("That's not a number... Error.")
			fmt.Println("Nice C:!!")
			customError = errors.NewOther()
			continue
		}

		if inputInt < 1 || inputInt > 10 {
			fmt.Println("Number out of allowed bounds... Error.")
			fmt.Println("Nice work!!")
			customError = errors.NewThirdParty()
			continue
		}

		fmt.Printf("Eeeh, that's right, %d is inside range\n", inputInt)
		fmt.Printf("Aren't you a smart individual...\n")
		fmt.Printf("\n\nPress enter to try again")
		fmt.Scanf("%s")
	}
	fmt.Printf("\n\nExiting")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 1333)
		fmt.Printf(".")
	}

	msg := "Ok now it's time to guess which type of error you got"

	errorTypes := map[int]string{1: "Internal", 2: "ThirdParty", 3: "Other"}

	for {
		clear()
		fmt.Println(msg)
		fmt.Printf("\nTypes of errors:\n\n")
		for i := 1; i < len(errorTypes)+1; i++ {
			fmt.Printf("%d  - %s\n", i, errorTypes[i])
		}
		fmt.Printf("\nPut the number of the error you think you got: ")

		if _, err := fmt.Scanf("%s", &uinput); err != nil {
			msg = "Stop pressing enter and do the guess"
			continue
		}

		num, err := strconv.ParseInt(uinput, 10, 32)
		if err != nil {
			msg = "That's not a number. Guess the error >:["
			continue
		}

		if num < 1 || num > 3 {
			msg = "You know that isn't an option. Just guess the error"
			continue
		}

		if errors.ErrorChecker(customError, errorTypes[int(num)]) {
			clear()
			fmt.Printf("You did it! It is an %s error\n\n", customError)
			time.Sleep(time.Millisecond * 1500)
			fmt.Printf("Closing")
			for i := 0; i < 3; i++ {
				time.Sleep(time.Millisecond * 800)
				fmt.Printf(".")
			}
			fmt.Println("")
			return
		}

		msg = "Wrong! You can try as many times as you need"
	}

}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
