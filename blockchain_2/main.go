package main

import "fmt"

func main() {
	blockchain := generateBlockchain(1)

	blockchain.addBlock("Alice", "Bob", 100)
	blockchain.addBlock("Bob", "Charlie", 50)

	fmt.Println(blockchain.isValid())
}
