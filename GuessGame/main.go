package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction

	playTheGame(firstNumber, secondNumber, substraction, answer)

}

func playTheGame(firstNumber, secondNumber, substraction, answer int) {
	reader := bufio.NewReader(os.Stdin)
	// display a welcome /instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Mutiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)
}
