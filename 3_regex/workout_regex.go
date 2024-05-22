package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the input string
	input := "tsnow, tshah, bmoreno"

	// Compile the regular expression
	re := regexp.MustCompile("ts")

	// Find all matches
	matches := re.FindAllString(input, -1)

	// Print the matches
	fmt.Println("Output of ts regex in input:", input, "output:", matches)

	// Define the input string
	input = "h32rb17"

	// Compile the regular expression
	re = regexp.MustCompile(`\w`)

	// Find all matches
	matches = re.FindAllString(input, -1)

	// Print the matches
	fmt.Println("Output of \\w regex in input:", input, "output:", matches)

	// Define the input string
	input = "h32rb17"

	// Compile the regular expression
	re = regexp.MustCompile(`\d+`)

	// Find all matches
	matches = re.FindAllString(input, -1)

	// Print the matches
	fmt.Println("Output of \\d+ regex in input:", input, "output:", matches)

	// Define the input string
	input = "h32rb17 k825t0m c2994eh"

	// Compile the regular expression
	re = regexp.MustCompile(`\d{2}`)

	// Find all matches
	matches = re.FindAllString(input, -1)

	// Print the matches
	fmt.Println("Output of \\d{2} regex in input:", input, "output:", matches)

	// Define the pattern and input string
	pattern := `\w+:\s\d+`
	employeeLoginsString := "1001 bmoreno: 12 Marketing 1002 tshah: 7 Human Resources 1003 sgilmore: 5 Finance"

	// Compile the regular expression
	re = regexp.MustCompile(pattern)

	// Find all matches
	matches = re.FindAllString(employeeLoginsString, -1)

	// Print the matches
	fmt.Println("Output of custom regex in input:", employeeLoginsString, "output:", matches)

	// Define the input string
	emailLog := `Email received June 2 from user1@email.com
            Email received June 2 from user2@email.com
            Email rejected June 2 from invalid_email@email.com`

	// Define the pattern
	pattern = `\w+@\w+\.\w+`

	// Compile the regular expression
	re = regexp.MustCompile(pattern)

	// Find all matches
	matches = re.FindAllString(emailLog, -1)

	// Print the matches
	fmt.Println("Output of custom regex in input:", emailLog, "output:", matches)
}
