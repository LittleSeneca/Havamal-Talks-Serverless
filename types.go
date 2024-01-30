package main

// Description struct represents the language and text of a vulnerability description.
type Description struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

// CVSSDataV3 struct represents various metrics in the CVSS version 3 scoring system.
type CVSSDataV3 struct {
	AttackComplexity      string  `json:"attackComplexity"`
	AttackVector          string  `json:"attackVector"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore"`
	BaseSeverity          string  `json:"baseSeverity"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	PrivilegesRequired    string  `json:"privilegesRequired"`
	Scope                 string  `json:"scope"`
	UserInteraction       string  `json:"userInteraction"`
	VectorString          string  `json:"vectorString"`
	Version               string  `json:"version"`
}

// CVSSMetricV3 struct wraps the CVSSDataV3 struct for CVSS version 3 metrics.
type CVSSMetricV3 struct {
	CVSSDataV3 CVSSDataV3 `json:"cvssData"`
}

// CVSSDataV3 struct represents various metrics in the CVSS version 2 scoring system.
type CVSSDataV2 struct {
	AccessComplexity      string  `json:"accessComplexity"`
	AccessVector          string  `json:"accessVector"`
	Authentication        string  `json:"authentication"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	BaseScore             float64 `json:"baseScore"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	VectorString          string  `json:"vectorString"`
	Version               string  `json:"version"`
}

// CVSSMetricV3 struct wraps the CVSSDataV3 struct for CVSS version 3 metrics.
type CVSSMetricV2 struct {
	CVSSDataV2 CVSSDataV2 `json:"cvssData"`
}

// Metrics struct wraps the CVSSMetricV2 and CVSSMetricV3 structs into one package for display.
type Metrics struct {
	CVSSMetricV31 []CVSSMetricV3 `json:"cvssMetricV31"`
	CVSSMetricV2  []CVSSMetricV2 `json:"cvssMetricV2"`
}

// CVE struct combines the structs built from component parts into one unified deliverable
type CVE struct {
	ID           string        `json:"id"`
	Published    string        `json:"published"`
	LastModified string        `json:"lastModified"`
	VulnStatus   string        `json:"vulnStatus"`
	Descriptions []Description `json:"descriptions"`
	Metrics      Metrics       `json:"metrics"`
}

// Vulnerability struct presents the CVE struct to be printed in the main function of this project
type Vulnerability struct {
	CVE CVE `json:"cve"`
}

// NVDResponse struct represents the response structure from the NVD API.
type NVDResponse struct {
	ResultsPerPage  int             `json:"resultsPerPage"`
	StartIndex      int             `json:"startIndex"`
	TotalResults    int             `json:"totalResults"`
	Format          string          `json:"format"`
	Version         string          `json:"version"`
	Timestamp       string          `json:"timestamp"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}
