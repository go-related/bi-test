package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unidoc/unipdf/v3/model"
	"io/ioutil"
	"net/http"
)

type AIResponse struct {
	Summary string `json:"summary"`
}

func extractTextFromPDF(pdfBytes []byte) (string, error) {
	reader, err := model.NewPdfReader(bytes.NewReader(pdfBytes))
	if err != nil {
		return "", err
	}

	numPages, err := reader.GetNumPages()
	if err != nil {
		return "", err
	}

	var text string
	for i := 1; i <= numPages; i++ {
		page, err := reader.GetPage(i)
		if err != nil {
			return "", err
		}

		content, err := page.GetAllContentStreams()
		if err != nil {
			return "", err
		}
		text += content + " "
	}
	return text, nil
}

func summarizeTextUsingAI(text string) (string, error) {
	apiURL := "https://api-inference.huggingface.co/models/facebook/bart-large-cnn"
	client := &http.Client{}

	requestBody, err := json.Marshal(map[string]string{"inputs": text})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer YOUR_HUGGINGFACE_API_KEY")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var aiResponse []AIResponse
	err = json.NewDecoder(resp.Body).Decode(&aiResponse)
	if err != nil || len(aiResponse) == 0 {
		return "", fmt.Errorf("Failed to parse AI response")
	}

	return aiResponse[0].Summary, nil
}

func summarizePDFHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	pdfBytes, err := ioutil.ReadFile(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	text, err := extractTextFromPDF(pdfBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract text"})
		return
	}

	summary, err := summarizeTextUsingAI(text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI summarization failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}

func main() {
	r := gin.Default()
	r.POST("/summarize", summarizePDFHandler)
	r.Run(":8080")
}
