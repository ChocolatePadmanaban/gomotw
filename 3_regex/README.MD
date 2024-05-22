---
layout: default
---

# Golang Regular Expressions :

## Simple Regex in golang:

- Import the regexp package: This package provides functions to work with regular expressions.
- Compile the regular expression: The MustCompile function is used to create a regular expression object.
- Find all matches: The FindAllString method returns a slice of all occurrences of the regular expression in the input string. The second argument -1 means to find all matches.
- Print the matches: Use fmt.Println to output the matches.

```
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
	fmt.Println(matches)
}
```
**Output:**
```
$ go run workout_regex.go 
[ts ts]
```

## Alpha numeric regex in golang

for getting alphanumeric regular expression we use `\w` for it

note: `\w` matches with `_` underscore too 

```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the input string
	input := "h32rb17"

	// Compile the regular expression
	re := regexp.MustCompile(`\w`)

	// Find all matches
	matches := re.FindAllString(input, -1)

	// Print the matches
	fmt.Println(matches)
}
```

**Output:**
```
[h 3 2 r b 1 7]
```


## Additional Symbols to match sepcific types of Characters:

- `.` matches to all characters including symbols 
- `\d` matches to all single digits [0-9]
- `\s` matches to all single spaces
- `\.` matches to the period character 

## symbols to qualify occurences 

- `+` symbol represents one or more occurences of a specific character 
- In the following example the pattern places it after the `\d` symbol to find matches to one or more occurences of a single digit 

```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the input string
	input := "h32rb17"

	// Compile the regular expression
	re := regexp.MustCompile(`\d+`)

	// Find all matches
	matches := re.FindAllString(input, -1)

	// Print the matches
	fmt.Println(matches)
}
```

**Output:**

```
[32 17]
```

- `*` symbol represents zero, one, or more occurences of specific characters.
- if you want to indicate specific number of repetitions to allow, you can place this number in curly brackets `{ }`
- In the following example, the regular expression pattern `\d{2}` instructs go to return all matches of exactly two single digits in a row from string of multiple device IDs:

```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the input string
	input := "h32rb17 k825t0m c2994eh"

	// Compile the regular expression
	re := regexp.MustCompile(`\d{2}`)

	// Find all matches
	matches := re.FindAllString(input, -1)

	// Print the matches
	fmt.Println(matches)
}
```

**Output:**

```
[32 17 82 29 94]
```

- you can specify range within curly brakets by separating two numbers with a comma.
- example `\d{1,3}` will get all digits from 1 to 3 

## Contructing a pattern:

- Contructing a regular expression requires you to breakdown the pattern you are searching for into smaller chunks and represent those chunks using the symbols you have learned 
- Example For each Employee the following string contains their employee ID, their 
username followed by a colon `:`, their attempted logins for the day, and their deplartment:

```
employee_logins_string := "1001 bmoreno: 12 Marketing 1002 tshah: 7 Human Resources 1003 sgilmore: 5 Finance"
```

- Our task is to extract the username and the login attempts, without the employee's ID number or department.

```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the pattern and input string
	pattern := `\w+:\s\d+`
	employeeLoginsString := "1001 bmoreno: 12 Marketing 1002 tshah: 7 Human Resources 1003 sgilmore: 5 Finance"

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(employeeLoginsString, -1)

	// Print the matches
	fmt.Println(matches)
}
```

**Output:**

```
[bmoreno: 12 tshah: 7 sgilmore: 5]
```

## Construct a pattern 2:

- Construct a regular expression to get the email-id of the email logs

```
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the input string
	emailLog := `Email received June 2 from user1@email.com
            Email received June 2 from user2@email.com
            Email rejected June 2 from invalid_email@email.com`

	// Define the pattern
	pattern := `\w+@\w+\.\w+`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(emailLog, -1)

	// Print the matches
	fmt.Println(matches)
}
```

**Output:**

```
[user1@email.com user2@email.com invalid_email@email.com]
```