package translate

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	// "google.golang.org/api/option"
)

func main() {
	// import "cloud.google.com/go/translate"
	// import "google.golang.org/api/option"
	// import "golang.org/x/text/language"

	// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:/projects/src/hyperdic/server/hyperdic-d02647eebf1c.json")

	ctx := context.Background()

	// const apiKey = "AIzaSyBYKupQPuoSdS6cysM5u864N1vBSf9OKcA"
	// client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatal("Failed to create client: %v", err)
	}
	// defer client.Close()

	resp, err := client.Translate(ctx, []string{"Hello, world!"}, language.Russian, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", resp)
}

// [START translate_translate_text]
func TranslateText(targetLanguage, text string) (string, error) {
	// translate works only with the GOOGLE_APPLICATION_CREDENTIALS env variable.
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:/projects/src/hyperdic/server/hyperdic-d02647eebf1c.json")
	fmt.Println("FOO:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}

// [END translate_translate_text]
// [START translate_detect_language]

func detectLanguage(text string) (*translate.Detection, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lang, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return nil, err
	}
	return &lang[0][0], nil
}

// [END translate_detect_language]
// [START translate_list_codes]
// [START translate_list_language_names]

func listSupportedLanguages(w io.Writer, targetLanguage string) error {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	langs, err := client.SupportedLanguages(ctx, lang)
	if err != nil {
		return err
	}

	for _, lang := range langs {
		fmt.Fprintf(w, "%q: %s\n", lang.Tag, lang.Name)
	}

	return nil
}

// [END translate_list_language_names]
// [END translate_list_codes]

// [START translate_text_with_model]

func translateTextWithModel(targetLanguage, text, model string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, &translate.Options{
		Model: model, // Either "mnt" or "base".
	})
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}

// [END translate_text_with_model]
