Phrase Generator

![Phrase_Generator_SS](https://github.com/Julian-1001/Phrase_Generator/assets/162458677/caf10986-0486-4d9b-a53e-95ef0564277d)

Overview

The Phrase Generator is an open-source command-line tool that generates mnemonic seed phrases based on the BIP39 standard. It supports creating both 12-word and 24-word seed phrases. This tool is designed to be freely downloadable, modifiable, and usable by anyone interested in generating cryptographic seed phrases for cryptocurrency wallets or other cryptographic applications while offline.

Prerequisites

Go (Golang) installed on your machine. You can download and install Go from the official Go website: https://go.dev/dl/

Before running the Phrase Generator, ensure that all external dependencies are properly installed. The program relies on the github.com/tyler-smith/go-bip39 package, which must be explicitly listed in your project's go.mod file. If you encounter an error regarding missing packages, you will need to add this package to your project's dependencies.
Adding Missing Dependencies
If you receive an error indicating that the github.com/tyler-smith/go-bip39 package is missing, follow these steps to resolve the issue:
Navigate to Your Project's Root Directory: Ensure you are in the directory where the go.mod file is located. This is typically the root directory of your project.
Add the Dependency: Run the following command in PowerShell (or your terminal of choice) to add the github.com/tyler-smith/go-bip39 package to your project's dependencies:
shell
go get github.com/tyler-smith/go-bip39
This command fetches the go-bip39 package and adds it to your go.mod file, while also updating the go.sum file which tracks the checksums of your dependencies for integrity verification.
After completing these steps, your project should have all the necessary dependencies installed, and you can proceed with building and running the Phrase Generator.


Installation

To set up the Phrase Generator on your local machine, follow these steps:

Clone the repository or download the source code to your local machine.
Navigate to the project directory where the Phrase_Generator.go file is located.
Build the program by running the following command in your terminal:

    
    go build Phrase_Generator.go

This command creates an executable named Phrase_Generator (or Phrase_Generator.exe on Windows).

Usage

To generate mnemonic seed phrases, follow these steps:

Open a terminal or command prompt.
Navigate to the directory containing the Phrase Generator executable.
Run the program:
    On Linux/macOS:

    ./Phrase_Generator

On Windows:

    .\Phrase_Generator.exe

Follow the on-screen instructions. You can generate a 12-word or 24-word seed phrase by entering make 12 or make 24, respectively. To exit the program, type exit.

Contributing

Contributions to the Phrase Generator are welcome! Please feel free to submit pull requests or open issues on the project's GitHub page to suggest improvements or add new features.

MIT License

Contact

For questions contact Julian Melendez at https://www.linkedin.com/in/julian-melendez-0a9ba22b8/
