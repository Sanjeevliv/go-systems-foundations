package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type IPEntry struct {
	Raw string
	IP net.IP
	IsValid bool
	LineNum int
}

type ParseResult struct {
	Entries []IPEntry
	ValidCount int
}

func parseIPFile(filePath string) (ParseResult, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return ParseResult{}, fmt.Errorf("failed to open file %w", err)
	}
	defer file.Close()

	var result ParseResult
	result.Entries = make([]IPEntry, 0, 100)

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
	
		if len(line) == 0 {
			continue
		}

		parsedIP := net.ParseIP(line)
		isValid := parsedIP != nil

		entry := IPEntry{
			Raw: line,
			IP: parsedIP,
			IsValid: isValid,
			LineNum: lineNumber,
		}

		result.Entries = append(result.Entries, entry)
		if isValid {
			result.ValidCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return ParseResult{}, fmt.Errorf("error reading file stream: %w", err)
	}
	return result, nil
}

func main() {
	filePath := "/Users/sanjeev/Desktop/MicroProjects/Project_1/ips.txt"

	result, err := parseIPFile(filePath)
	if err != nil {
		fmt.Printf("Execution error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Processing Complete.\n")
	fmt.Printf("Total Records Processed: %d\n", len(result.Entries))
	fmt.Printf("Valid IP Addresses Found: %d\n\n", result.ValidCount)

	for _, entry := range result.Entries {
		status := "Invalid"
		if entry.IsValid {
			status = "Valid"
		}
		fmt.Printf("Line %d: %-15s [%s]\n", entry.LineNum, entry.Raw, status)
	}
}