# Phrase Generator (BIP39)

![Phrase_Generator_SS](https://github.com/Julian-1001/Phrase_Generator/assets/162458677/caf10986-0486-4d9b-a53e-95ef0564277d)

## Overview

Phrase Generator is a lightweight command-line tool written in Go that generates mnemonic seed phrases compliant with the **BIP39 standard**. It supports both 12-word and 24-word phrases and is designed to run **entirely offline**, minimizing exposure during seed generation.

This project was built to explore practical cryptography concepts, secure randomness, and dependency management in a small, auditable codebase. It is intended for educational and experimental use rather than production wallet generation.

---

## Purpose

The goal of this project was not to create a full-featured wallet tool, but to understand how mnemonic seed phrases are generated, how entropy is handled, and how cryptographic dependencies should be selected and managed responsibly.

By keeping the implementation minimal and offline, the project emphasizes correctness, transparency, and risk awareness over convenience.

---

## What This Project Demonstrates

This project demonstrates:

- Understanding of the BIP39 mnemonic standard  
- Appropriate use of vetted cryptographic libraries instead of custom implementations  
- Offline-first design to reduce attack surface  
- Secure handling of sensitive material in a local execution context  
- Dependency integrity and reproducible builds using Go modules  

---

## Implementation Notes

The application is implemented in **Go** and leverages the widely adopted  
`github.com/tyler-smith/go-bip39` library to handle entropy generation and mnemonic construction. Using an established library avoids reimplementing cryptographic primitives and reduces the likelihood of subtle security flaws.

Dependency integrity is managed through `go.mod` and `go.sum`, ensuring reproducible builds and checksum verification.

---

## Build and Run

This project requires **Go (Golang)**.

From the project root directory, build the executable:

    go build Phrase_Generator.go

This produces a standalone executable.

Run the program from the terminal:

Linux / macOS:
    ./Phrase_Generator

Windows:
    .\Phrase_Generator.exe

Follow the on-screen prompts to generate a **12-word** or **24-word** mnemonic phrase.

---

## Limitations

This tool is intentionally minimal. It does not provide secure storage, key derivation visualization, or hardware-backed protection, and it is not intended to replace audited wallet software or hardware devices.

The project focuses on understanding correctness and cryptographic primitives rather than providing end-user safety guarantees.

---

## Security Considerations

Mnemonic seed phrases are highly sensitive secrets. Users are responsible for ensuring the security of their execution environment and for handling generated phrases appropriately.

Offline operation reduces exposure risk, but does not eliminate it.

---

## License

MIT License
