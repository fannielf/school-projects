package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// finding a number value from a string eg "2)"" and converts it into an integer
func getInt(s string) int {

	num := 1

	re := regexp.MustCompile(`\d+`)

	match := re.FindString(s)
	if match != "" {
		number, _ := strconv.Atoi(match)

		num = number
	}

	return num
}

// finding an index of a slice that equals (hex) or (bin) and converts the previous index accordingly. Removes even empty indexes in addition to identifier indexes.
func fixHexBin(s string) string {
	slice := strings.Split(s, " ")

	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		} else if slice[i] == "(hex)" || slice[i] == "(bin)" {
			if i > 0 {
				indexToConvert := slice[i-1]
				var decimalValue int64
				var err error

				if slice[i] == "(hex)" {
					decimalValue, err = strconv.ParseInt(indexToConvert, 16, 64)
				} else if slice[i] == "(bin)" {
					decimalValue, err = strconv.ParseInt(indexToConvert, 2, 64)
				}
				if err != nil {
					fmt.Println("Error converting value:", err)
					os.Exit(1)
				}
				slice[i-1] = strconv.FormatInt(decimalValue, 10)
				slice = append(slice[:i], slice[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(slice, " ")
}

// finding an index of a slice that equals "a" or "A" and checks if the following index starts with a vowel. If yes, adding "n" to the end of the index.
func fixArticle(s string) string {
	slice := strings.Split(s, " ")
	var isVowel = map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	for i := 0; i < len(slice)-1; i++ {
		if strings.ToLower(slice[i]) == "a" {
			nextWord := slice[i+1]
			if isVowel[rune(nextWord[0])] {
				slice[i] += "n"
			}
		}
	}
	return strings.Join(slice, " ")
}

// finding an index of a slice that equals equals case specific commands (cap) (up) (low) for one or various words and converts the previous index(es) accordingly. Removes even empty indexes as well as possible value indexes.
func fixCase(s string) string {
	regex := regexp.MustCompile(`\([^()]*\)|[^()\s]+`)
	slice := regex.FindAllString(s, -1)

	re := regexp.MustCompile(`\((cap|up|low)`)

	for i := 0; i < len(slice); i++ {
		if re.MatchString(strings.ToLower(slice[i])) {
			keyWord := re.FindString(strings.ToLower(slice[i]))
			num := getInt(slice[i])
			if i > 0 {
				for j := num; j > 0; j-- {
					wordToMod := slice[i-j]
					switch keyWord {
					case "(cap":
						wordToMod = strings.ToUpper(wordToMod[:1]) + wordToMod[1:]
					case "(up":
						wordToMod = strings.ToUpper(wordToMod)
					case "(low":
						wordToMod = strings.ToLower(wordToMod)
					}
					slice[i-j] = wordToMod
				}
			}
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return strings.Join(slice, " ")
}

// Checking if index is a punctuation or if an index includes a punctuation and applies punctuation rules where all punctuation is attached to the previous word and are followed by a space.
func fixPunctuation(s string) string {
	slice := strings.Split(s, " ")

	var isPunctuationMark = map[string]bool{
		".": true,
		",": true,
		"!": true,
		"?": true,
		":": true,
		";": true,
	}

	for i := 0; i < len(slice); i++ {
		if i > 0 && isPunctuationMark[slice[i]] {
			slice[i-1] = slice[i-1] + slice[i]
			slice = append(slice[:i], slice[i+1:]...)
			i--
			continue
		} else {
			newWord := ""
			for j, char := range slice[i] {
				if j == 0 {
					newWord = string(char)
				} else {
					if !isPunctuationMark[string(char)] && isPunctuationMark[string(newWord[j-1])] {
						newWord += " "
					}
					newWord += string(char)
				}
			}
			if isPunctuationMark[string(newWord[0])] && i > 0 {
				slice[i-1] = slice[i-1] + newWord
				slice = append(slice[:i], slice[i+1:]...)
				i--
			} else {
				slice[i] = newWord
			}
		}
	}

	return strings.Join(slice, " ")
}

// With the help of boolean value, checking if "'" is the opening or closing quote and builds the new string according to quotation rules. Adding a new line and returning the finalized text to be saved to output.
// Rules: opening quote preceded by a space and attached to the following word; closing quote followed by a space and attached to the previous word
func fixQuotation(s string) string {
	slice := strings.Split(s, " ")
	newString := ""

	inQuote := false

	for _, word := range slice {
		if len(newString) == 0 && word != "'" {
			newString = word
		} else if word == "'" {
			if inQuote {
				newString = strings.TrimSpace(newString) + word
			} else {
				newString += " " + word
			}
			inQuote = !inQuote
		} else if inQuote && newString[len(newString)-1] == '\'' {
			newString += word
		} else {
			newString += " " + word
		}
	}
	return newString
}

// opens the file location, reads the contents of the file and casts into a string
func getOriginalText(text string) string {

	fileLocation, err := os.Open(text)
	if err != nil {
		fmt.Printf("Error: Couldn't open %s file location.\n", text)
	}
	defer fileLocation.Close()

	fileContent, err := io.ReadAll(fileLocation)
	if err != nil {
		fmt.Println("Error: Couldn't read the file contents.")
	}
	return (string(fileContent))
}

// Checking if the output file exists and either create or open the file. Writes the modified content to the output file.
func writeToFile(file string, content string) {

	textToFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // WRONGLY opening the file for writing only (not reading), if not existing yet CREATE (self explanatory), TRUNC wipes the contents of the file
	if err != nil {
		fmt.Printf("Error: Couldn't open nor create %s .\n", file)
	}
	defer textToFile.Close()

	_, err = textToFile.Write([]byte(content))
	if err != nil {
		fmt.Printf("Error: Failed to write content to %s .\n", file)
	}
}

// validating the input (amount of arguments, both files need to be .txt files and making sure the input file exists)
func isValid(arg []string) {

	if len(arg) != 3 {
		fmt.Println("Error: Invalid amount of arguments.")
		os.Exit(1) // General error
	}
	for _, file := range arg[1:] {
		ext := filepath.Ext(file)
		if strings.ToLower(ext) != ".txt" {
			fmt.Printf("Error: Input %s has to be a .txt file.\n", file)
			os.Exit(4) // Invalid input or arguments
		}
	}
	_, err := os.Stat(arg[1])
	if os.IsNotExist(err) {
		fmt.Printf("Error: File %s not found.\n", arg[1])
		os.Exit(3) // File-related error
	}
}

func main() {

	isValid(os.Args[0:])

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	originalContent := getOriginalText(inputFile)

	readyContent := handlingContent(originalContent)

	// modifiedContent := fixHexBin(originalContent)
	// modifiedContent = fixArticle(modifiedContent)
	// modifiedContent = fixCase(modifiedContent)
	// modifiedContent = fixPunctuation(modifiedContent)
	// readyContent := fixQuotation(modifiedContent)

	writeToFile(outputFile, readyContent)
}

func handlingContent(s string) string {
	contentLines := strings.Split(s, "\n")

	for i, line := range contentLines {
		textToModify := line
		textToModify = fixHexBin(textToModify)
		textToModify = fixArticle(textToModify)
		textToModify = fixCase(textToModify)
		textToModify = fixPunctuation(textToModify)
		contentLines[i] = fixQuotation(textToModify)
	}

	return strings.Join(contentLines, "\n")
}
