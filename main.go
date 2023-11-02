package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func main() {
	// Replace with the path to your JSON key file
	credentialsFile := "credentials.json"

	// Initialize a context and create a client with your API credentials
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Text to be translated from English to German
	inputText := "Hello, world!"

	// Translate the text
	translatedText, err := translateText(ctx, client, inputText, language.English, language.German)
	if err != nil {
		log.Fatalf("Translation error: %v", err)
	}

	// Print the translated text
	fmt.Printf("Input Text: %s\n", inputText)
	fmt.Printf("Translated Text: %s\n", translatedText)
}

func translateText(ctx context.Context, client *translate.Client, text string, sourceLang, targetLang language.Tag) (string, error) {
	translations, err := client.Translate(ctx, []string{text}, targetLang, &translate.Options{
		Source: sourceLang,
	})
	if err != nil {
		return "", err
	}
	if len(translations) > 0 {
		return translations[0].Text, nil
	}
	return "", nil
}
