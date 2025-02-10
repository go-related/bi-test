package main

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

func CreateSummaryForFiles(folderPath, resultsPath, runningInstance string, client *AIClient) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}
	resultFile, err := os.Create(resultsPath + "/" + runningInstance + ".txt")
	if err != nil || resultFile == nil {
		return errors.Wrap(err, "failed to to write results file")
	}
	for _, file := range files {
		if !file.IsDir() {
			if strings.Contains(file.Name(), ".pdf") {
				filePath := folderPath + "/" + file.Name()
				text, err := extractTextFromPDF(filePath)
				if err != nil {
					return err
				}
				summary, err := client.Summarize(text)
				if err != nil {
					return errors.Wrap(err, "failed to summarize")
				}

				writeResult(resultFile, file.Name(), text, summary)
				time.Sleep(time.Minute)
			}
		}
	}
	return nil
}

func writeResult(file *os.File, name, text, summary string) {

	file.WriteString("*******************\n")
	file.WriteString(name + ": \n")

	file.WriteString(text + "\n")
	file.WriteString("overview:" + "\n")
	file.WriteString(summary)
	file.WriteString("\n*******************")

	fmt.Println("*******************")
	fmt.Println(name + ": ")
	fmt.Println(summary)
	fmt.Println("*******************")
	fmt.Println()
	fmt.Println()
}

func extractTextFromPDF(pdfFile string) (string, error) {
	f, r, err := pdf.Open(pdfFile)
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	_, err = buf.ReadFrom(b)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
