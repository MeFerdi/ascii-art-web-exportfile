package ascii

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	hashing "ascii-art/hashing"
)

// bannerMap is a map that stores the ASCII art for different banner files.
// It maps the banner file name to the corresponding ASCII art string.
var bannerMap map[string]string

// init initializes the bannerMap and loads the ASCII art from the banner files.
// It reads the contents of the standard, shadow, and thinkertoy banner files
// and stores them in the bannerMap.
func init() {
	bannerMap = make(map[string]string)
	loadBanner(filepath.Join("banners", "shadow.txt"))
	loadBanner(filepath.Join("banners", "standard.txt"))
	loadBanner(filepath.Join("banners", "thinkertoy.txt"))
}

// loadBanner reads the contents of a banner file and stores it in the bannerMap.
// It takes a file path as input and appends the ASCII art lines to a slice.
// The joined lines are then stored in the bannerMap using the file name as the key.
func loadBanner(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.New("file modified")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return nil
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	bannerMap[filepath.Base(filename)] = strings.Join(lines, "\n")
	return nil
}

// GetLetterArray retrieves the ASCII art representation for a given character
// from the specified banner file. It takes a character and a banner style as input,
// and returns a slice of strings representing the ASCII art for that character.
// If the banner style is not found or the character is out of range, it returns an error.
func GetLetterArray(char rune, bannerStyle string) ([]string, error) {
	banner, ok := bannerMap[bannerStyle+".txt"]
	if !ok {
		return nil, errors.New("banner style not found")
	}

	fileHash := hashing.ValidFile([]byte(banner))

	// Hash validation for each banner style
	switch bannerStyle {
	case "standard":
		if fileHash != "c3ec7584fb7ecfbd739e6b3f6f63fd1fe557d2ae3e24f870730d9cf8b2559e94" {
			return nil, errors.New("standard file corrupted")
		}
	case "shadow":
		if fileHash != "78ccd616680eb9068fe1465db1c852ceaffd8c0f318e3aa0414e1635508e85bf" {
			return nil, errors.New("shadow file corrupted")
		}
	case "thinkertoy":
		if fileHash != "e3c7a11f41a473d9b0d3bf2132a8f6dabb754bd16efa3897fa835a432d3b9caa" {
			return nil, errors.New("ThinkerToy file corrupted")
		}
	default:
		return nil, errors.New("unsupported banner style")
	}

	alphabet := strings.Split(banner, "\n")
	start := (int(char) - 32) * 9

	if start < 0 || start+9 > len(alphabet) {
		return nil, errors.New("character out of range")
	}

	arr := alphabet[start : start+9]
	return arr, nil
}

// PrintAscii returns the ASCII art representation of a given string.
// It takes a string and a banner style as input, and returns the ASCII art
// representation of the string. If the input string contains non-ASCII characters
// or the banner file is corrupted, it returns an error.
func PrintAscii(str, bannerStyle string) (string, error) {
	replaceNewLine := strings.ReplaceAll(str, "\r\n", "\n")
	lines := strings.Split(replaceNewLine, "\n")
	var asciiArt strings.Builder
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			asciiArt.WriteString("\n")
			continue
		}
		letterArrays := [][]string{}
		for _, letter := range line {
			if letter < 32 || letter > 126 {
				return "", fmt.Errorf("non-ASCII character '%c' encountered", letter)
			}
			arr, err := GetLetterArray(letter, bannerStyle)
			if err != nil {
				return " ", errors.New("file corrupted")
			}

			letterArrays = append(letterArrays, arr)
		}
		for i := 0; i < 9; i++ {
			for _, letter := range letterArrays {
				asciiArt.WriteString(letter[i])
			}
			asciiArt.WriteString("\n")
		}
	}
	return asciiArt.String(), nil
}
