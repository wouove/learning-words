package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wouove/learning-words/internal/client"
	"github.com/wouove/learning-words/internal/extractor"
	"github.com/wouove/learning-words/internal/loader"
	"github.com/wouove/learning-words/internal/transformer"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "words-learner",
		Short: "Tool to prepare words from input file and stores them in the input format for Anki",
	}

	translateWords := &cobra.Command{
		Use:   "translate",
		Short: "Translates words",
		Run:   handle,
	}
	rootCmd.AddCommand(translateWords)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing the command:", err)
	}
}

func handle(cmd *cobra.Command, args []string) {
	inputPath := os.Getenv("INPUT_PATH")
	outputPath := os.Getenv("OUTPUT_PATH")
	extractorPort, err := extractor.NewSqlLiteExtractorAdapter(inputPath)
	if err != nil {
		fmt.Printf("setting up extractor: %w", err)
	}
	stringFormatter := loader.NewMultipleChoiceStringFormatter(3)
	loaderPort := loader.NewCSVLoaderAdapter(outputPath, stringFormatter)
	deepLClient := client.NewDeepLClient()
	transf := transformer.NewTransformer(
		extractorPort,
		loaderPort,
		deepLClient,
	)
	err = transf.Transform()
	if err != nil {
		fmt.Printf("transforming words: %w", err)
	}
}
