package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// fetchVulnerabilities sends a request to the NVD API to fetch vulnerabilities
// of a specific severity, using an API key for authentication.
func fetchVulnerabilities(ctx context.Context, severity string) {
	// Set the start and end dates for the API query.
	pubStartDate := time.Now().AddDate(0, 0, -5).Format("2006-01-02")
	pubEndDate := time.Now().Format("2006-01-02")

	// Construct the NVD API URL with query parameters.
	nvdURL := fmt.Sprintf("https://services.nvd.nist.gov/rest/json/cves/2.0?cvssV3Severity=%s&pubStartDate=%sT00:00:00.000&pubEndDate=%sT00:00:00.000", severity, pubStartDate, pubEndDate)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", nvdURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Retrieve API key from environment variable
	apiKey := os.Getenv("NVD_API_KEY")
	if apiKey == "" {
		fmt.Println("API key not set in environment variables")
		return
	}
	req.Header.Add("X-Api-Key", apiKey)

	// Make the HTTP request using the created request with the header
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request to NVD:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the function is canceled and exit gracefully if so
	if ctx.Err() != nil {
		fmt.Println("Function canceled. Exiting gracefully.")
		return
	}

	// Check if the HTTP request was not successful.
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusForbidden {
			fmt.Printf("HTTP request failed with status code 403 Forbidden [Likely Rate Throttling] for severity %s. Skipping...\n", severity)
			return
		}

		fmt.Printf("HTTP request failed with status code: %s\n", resp.Status)
		body, _ := io.ReadAll(resp.Body) // Read the response body even in case of an error.
		fmt.Println("Response body:", string(body))
		return
	}

	// Read the response body from the HTTP request.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Print an error message if reading the response body fails.
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal the JSON response into the NVDResponse struct.
	var nvdData NVDResponse
	if err := json.Unmarshal(body, &nvdData); err != nil {
		// Print an error message if unmarshaling fails.
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Iterate through each vulnerability in the response and print details.
	for _, vulnerability := range nvdData.Vulnerabilities {
		// Check if the function is canceled and exit gracefully if so
		if ctx.Err() != nil {
			fmt.Println("Function canceled. Exiting gracefully.")
			return
		}

		// Print the CVE ID.
		fmt.Println("CVE ID:", vulnerability.CVE.ID)

		// Print the English description of the vulnerability.
		for _, description := range vulnerability.CVE.Descriptions {
			if description.Lang == "en" {
				fmt.Println("Description:", description.Value)
			}
		}

		// Print CVSS v3.1 metrics.
		for _, metricV3 := range vulnerability.CVE.Metrics.CVSSMetricV31 {
			fmt.Printf("CVSS v3.1 Metrics: %+v\n", metricV3.CVSSDataV3)
		}

		// Print a line break after each CVE.
		fmt.Println()

		// Pause for 1 second to prevent rate limiting and check cancellation again
		time.Sleep(1 * time.Second)
		if ctx.Err() != nil {
			fmt.Println("Function canceled. Exiting gracefully.")
			return
		}
	}
}
