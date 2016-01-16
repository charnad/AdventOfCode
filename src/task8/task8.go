package task8

import (
	"bufio"
	"os"
	//	"strings"
	//	"strconv"
	"fmt"
	"regexp"
	"strings"
)

func Solve() {
	file, err := os.Open("res/task8.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	totalChars := 0
	totalData := 0
	totalExtended := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		originalLine := scanner.Text()
		totalChars += len(originalLine)

		//fmt.Println(line, "is", len(line), "long")

		backslash, _ := regexp.Compile(`\\\\|\\"`)
		line := backslash.ReplaceAllString(originalLine, ",")

		quotes, _ := regexp.Compile(`^"|"$`)
		line = strings.ToLower(quotes.ReplaceAllString(line, ""))

		hex, _ := regexp.Compile(`\\x[0-9a-f]{2}`)
		line = hex.ReplaceAllString(line, ".")

		totalData += len(line)

		slash, _ := regexp.Compile(`\\`)
		extendedLine := slash.ReplaceAllString(originalLine, "\\\\")

		quotes2, _ := regexp.Compile(`"`)
		extendedLine = quotes2.ReplaceAllString(extendedLine, "\\\"")

		extendedLine = "\"" + extendedLine + "\""

		totalExtended += len(extendedLine)
		fmt.Println(extendedLine, "is", len(extendedLine), "long")
	}

	fmt.Println("Total characters ", totalChars, "total data", totalData, "Difference: ", totalChars-totalData)
	fmt.Println("Total characters ", totalChars, "total extended", totalExtended, "Difference: ", totalExtended-totalChars)
}
