package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tyler-smith/go-bip39" // Importing the BIP39 package for generating mnemonic seed phrases.
)

// Constants for the entropy sizes corresponding to the number of words in the seed phrase.
const (
	entropySize12Words = 128 // Entropy size for 12 words seed.
	entropySize24Words = 256 // Entropy size for 24 words seed.
	commandMake12      = "make 12" // Command to generate 12-word seed phrase.
	commandMake24      = "make 24" // Command to generate 24-word seed phrase.
	commandExit        = "exit"    // Command to exit the program.
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Creating a new buffered reader for standard input.
	
	// Mapping commands to their respective entropy sizes for mnemonic generation.
	commandMap := map[string]int{
		commandMake12: entropySize12Words,
		commandMake24: entropySize24Words,
	}

	for {
		fmt.Println(`Enter "make 12" for a 12-word seed phrase, "make 24" for a 24-word seed phrase, or "exit" to quit:`)
		input, err := reader.ReadString('\n') // Reading user input from standard input.
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue // Continue prompting the user for input in case of read errors.
		}
		input = strings.TrimSpace(input) // Trimming newline and spaces from the input.

		if input == commandExit { // Check if the user wants to exit the program.
			fmt.Println("Exiting program.")
			break // Exiting the loop and the program.
		}

		entropySize, ok := commandMap[input] // Retrieve the entropy size for the given command.
		if !ok {
			fmt.Println("Invalid command. Please enter \"make 12\", \"make 24\", or \"exit\".")
			continue // Continue prompting for input if an invalid command is entered.
		}

		mnemonic, err := generateMnemonic(entropySize) // Generating the mnemonic seed phrase.
		if err != nil {
			fmt.Println("Error generating mnemonic: ", err)
			continue // Continue prompting for input if there was an error generating the mnemonic.
		}

		fmt.Printf("Your mnemonic seed phrase is:\n%s\n\n", mnemonic) // Displaying the generated mnemonic seed phrase.
	}
}

// generateMnemonic generates a mnemonic seed phrase given the entropy size.
func generateMnemonic(size int) (string, error) {
	entropy, err := bip39.NewEntropy(size) // Generating new entropy of the given size.
	if err != nil {
		return "", err // Return the error if entropy generation fails.
	}
	return bip39.NewMnemonic(entropy) // Generating and returning the mnemonic seed phrase from the entropy.
}
