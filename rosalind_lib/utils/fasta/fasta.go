package fasta

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type FastaEntry struct {
	Label   string
	Content string
}

type FastaContent []FastaEntry

func ReadFastaContent(f *os.File) FastaContent {
	scanner := bufio.NewScanner(f)
	labelRegex := regexp.MustCompile(`>Rosalind_\d{4}`) // Rosalind-specific label format

	var entry FastaEntry
	var content FastaContent
	var sb strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if labelRegex.MatchString(line) {
			if sb.Len() > 0 {
				entry.Content = sb.String()
				content = append(content, entry)
				sb.Reset()
			}
			entry.Label = line
		} else {
			sb.WriteString(line)
		}
	}

	// TODO: Handle more gracefully
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if sb.Len() > 0 {
		entry.Content = sb.String()
		content = append(content, entry)
		sb.Reset()
	}

	return content
}
