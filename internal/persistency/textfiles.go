/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package persistency

import (
	"bufio"
	"log"
	"os"
)

// TODO create a test for ReadTextLines that accesses a specific test file.
// ReadTextLines accesses a file and returns all text lines from that file.
func ReadTextLines(path string) ([]string, error) {
	lines := make([]string, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Could not close file: %s, resulted in error: %s", path, err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// TODO create a test for WriteTextLines that accesses a specific test file.
// WriteTextLines writes text lins to a file and overwrites the existing lines.
func WriteTextLines(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed to access file: %s", err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			log.Fatalf(
				"failed to write to file: %s, resulted in error: %s", path, err)
		}
	}
	err = w.Flush()
	if err != nil {
		log.Println("Error flushing buffer: ", err)
		return err
	}
	return nil
}
