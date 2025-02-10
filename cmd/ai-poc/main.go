package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")

	apiKey := os.Getenv("AI_API_KEY")
	folderPath := os.Getenv("AI_DOUCMENTS_PATH")
	resultsPath := os.Getenv("AI_RESULTS_PATH")
	instance := 1
	client := NewAIClient(apiKey)

	err := CreateSummaryForFiles(folderPath, resultsPath, fmt.Sprintf("%d", instance), client)
	if err != nil {
		panic(err)
	}
	fmt.Println("finished running")
}
