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
	// if err := godotenv.Load(); err != nil {
	// 	log.Print("No .env file found")
	// }

	pixabayAPIKey, exists := os.LookupEnv("PIXABAY_API_KEY")
	fmt.Print("PIXABAY");
	fmt.Print(exists);

	if exists {
		fmt.Print(pixabayAPIKey)
	}
}

func main() {
	// Get the GITHUB_USERNAME environment variable
	pixabayAPIKey, exists := os.LookupEnv("PIXABAY_API_KEY")
	fmt.Println("PIXABAY");
	fmt.Println(exists);

	if exists {
		fmt.Println(pixabayAPIKey)
	}
}

// Lala prints some message
func Lala() {
	fmt.Println("HEELO IMAGES")
}
