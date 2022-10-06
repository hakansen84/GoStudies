package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your Fav Number?")
	user.OwnsADog = readDog("DO you have a dog? (y/n")

	fmt.Printf("Your name is %s, and you are %d years old. Your Fav Number is %.2f. and owns a dog = %v\n",
		user.UserName, user.Age, user.FavouriteNumber, user.OwnsADog)

}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a  number")
		} else {
			return num
		}
	}
}

func readDog(s string) bool {
	for {
		fmt.Println(s)
		prompt()

		char, _, _ := keyboard.GetSingleKey()

		if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'N' {
			return false
		} else {
			fmt.Println("Please enter Y or N")
		}
	}
}
