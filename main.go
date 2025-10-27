package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: margo <output.pdf> <input1.pdf> <input2.pdf> ...")
		fmt.Println("Example: margo merged.pdf file1.pdf file2.pdf file3.pdf")
		os.Exit(1)
	}

	outputFile := os.Args[1]
	inputFiles := os.Args[2:]

	// Validate input files
	for _, file := range inputFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("Error: File '%s' does not exist\n", file)
			os.Exit(1)
		}
		if filepath.Ext(file) != ".pdf" {
			fmt.Printf("Error: File '%s' is not a PDF file\n", file)
			os.Exit(1)
		}
	}

	fmt.Printf("Merging %d PDF files into '%s'...\n", len(inputFiles), outputFile)

	// Merge PDFs using pdfcpu
	err := api.MergeCreateFile(inputFiles, outputFile, true, nil)
	if err != nil {
		fmt.Printf("Error merging PDFs: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ PDFs merged successfully!")

	// Print file info
	if info, err := os.Stat(outputFile); err == nil {
		fmt.Printf("Output file size: %.2f MB\n", float64(info.Size())/(1024*1024))
	}
}
