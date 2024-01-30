package main

// main is the entry point of the application.
func main() {
	severities := []string{"CRITICAL", "HIGH", "MEDIUM", "LOW"}
	for _, severity := range severities {
		fetchVulnerabilities(severity)
	}
}
