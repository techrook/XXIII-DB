package main

import (
	"fmt"
	"log"
	"os"
	"github.com/techrook/xxiii-db/write"
)

func main() {
	path := "xx3_database.db"
	
	// // Test 1: Initial Write
	fmt.Println("--- Test 1: Initializing Database ---")
	initialData := []byte("Version 1: Original Data")
	err := write.SaveData2(path, initialData)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	printFileContent(path)

	// Test 2: Atomic Update
	fmt.Println("\n--- Test 2: Updating Atomically ---")
	newData := []byte("Version 2: Updated Data with Atomic Rename")
	err = write.SaveData2(path, newData)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	printFileContent(path)
}

// Helper to see what's actually in the file
func printFileContent(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not read file: %v\n", err)
		return
	}
	fmt.Printf("Current File Content: [%s]\n", string(content))
}