package pixabay

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const pixabayBaseURL = "https://pixabay.com/api/"

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Get the GITHUB_USERNAME environment variable
	pixabayAPIKey, exists := os.LookupEnv("PIXABAY_API_KEY")

	if exists {
		fmt.Println(pixabayAPIKey)
	}
}

// Lala prints some message
func Lala() {
	fmt.Println("HEELO IMAGES")
}

// C:\projects\src\hyperdic\server\internal\app\hyperdic\pixabay.go
