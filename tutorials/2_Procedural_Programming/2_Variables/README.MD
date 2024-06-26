---
layout: default
---

# Variables

## What is a Variable?
Variables are possibly the most fundamental element of programming. There
really isn’t much you can do without variables. You’re probably familiar with vari-
ables from your days learning algebra. Variables aren’t really any different here.
A ­variable is a name that holds a value. The name stays the same while the value
can change. In algebra, that variable was often x.

### Examples of Variables
A variable is a name that holds some value. It’s like a question with an answer. The
question always exists, but the answer to the question might change. Let’s think of
some examples:

- How many kids do David have? The variable would be the count of David’s kids, and
the value would be one. That might change, though.
- What’s the current stock price of Microsoft? The variable would be Microsoft
stock price, and the current value would be 53.74.
- What color shirt am I wearing? The variable would be David’s shirt color, and
the current value would be blue.

Notice how in each case, there’s a variable and a value. The variable can take on
different values. Notice also how unlike in math, variables don’t have to represent
numbers. You could have variables that represent any number of things, such as
names, colors, and dates.

Variables are the heart of programming. Nearly everything we do involves
manipulating variables. We’ll use variables to represent the information in which
we’re interested, like stock prices or usernames. We’ll use variables to control how
our programs run, like counting repeated actions or checking if something has been
found.

## Variables in Go
```
//1_variables_in_go.go
package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
}

```
Output
```
$ go run 1_variables_in_go.go 
5
$
```
### Different Kinds of Variables
Just as we talked about lots of different kinds of variables previously
```
//2_variabels_in_go.go
package main

import (
	"fmt"
	"time"
)

func main() {
	var aNumber int = -2
	var aDecimal float64 = 7.1
	var aString string = "Hello world"
	var aBoolean bool = true
	var aDate time.Time = time.Now()
	fmt.Println(aNumber)
	fmt.Println(aDecimal)
	fmt.Println(aString)
	fmt.Println(aBoolean)
	fmt.Println(aDate)
}

```
OutPut:
```
$ go run 2_variables_in_go.go 
-2
7.1
Hello world
true
2021-09-30 05:28:22.297022693 +0530 IST m=+0.000104275
$ 
```
Remember, we could have named these anything we wanted,
but we gave them names that made it clear what they are. Think of this like giving
names to a pet: you could name a cat “Fido” and a dog “Tiger,” but you wouldn’t
be surprised if these names confused people. The same way, you could name your
variables phineas, celery, and sovngarde, but these names wouldn’t tell the
reader anything about what they are.

### Naming Rules and Convention 

- Variable names can contain only letters, numbers, and underscores. No spaces
or special characters.
- Variable names must start with letters. Underscores are technically allowed, too,
but we usually only use these in certain situations.
- Variable names must not duplicate certain reserved words. We’ll talk about the
reserved keywords later.
When you violate these rules, you’ll usually get a syntax error
```
//3_naming_rules_and_conventions.go
package main
import "fmt"
func main() {
	my$variable :=5
	my variable := 5
	$variable := 5
	for := 5
}
```
Output:
```
$ go run 3_naming_rules_and_conventions.go 
# command-line-arguments
./3_naming_rules_and_conventions.go:5:4: invalid character U+0024 '$'
./3_naming_rules_and_conventions.go:5:5: syntax error: unexpected variable at end of statement
./3_naming_rules_and_conventions.go:6:5: syntax error: unexpected variable at end of statement
./3_naming_rules_and_conventions.go:7:2: invalid character U+0024 '$'
./3_naming_rules_and_conventions.go:8:6: syntax error: unexpected :=, expecting expression
./3_naming_rules_and_conventions.go:9:1: syntax error: unexpected }, expecting expression
$ 
```
However, beyond these rules, there are certain conventions we want to follow.
The computer isn’t going to yell at us if we don’t follow these, but other people that
have to work with our code might. The main convention is that our code should be
self-documenting. What that means is we should be able to read the variable name
and know what kind of data it’s holding
```
// 4_naming_rules_and_conventions.go
package main

import "fmt"

func main() {
	// Bad Variables
	a := 30
	foo := "David"
	qwerty := 10
	fmt.Println(a, foo, qwerty)
	// Good Variables
	var age int = 30
	var firstName string = "David"
	var catCount int = 10
	fmt.Println(age, firstName, catCount)
}

```
Output:
```
$ go run 4_naming_rules_and_conventions.go 
30 David 10
30 David 10
```

We’ve said that variables are like questions, and values are like their answers. The
variable sticks around while the value changes. There are a few principles that go
along with this, though.

## Data Types

Every variable has a name, and it should have a value as well if we’re using it.
There’s a third thing that variables have: types. Every variable has a type of infor-
mation it’s storing.

### Basic Data Types
There are some basic types you’ll see very often:
- Integers, or whole numbers.
- Real numbers, or numbers that can have decimals.
- Characters, like individual letters and numbers.
- Strings, which are collections of characters in a row.
- Booleans, which just hold either true or false.

However, not all variables hold these basic types. Variables can hold nearly any
type of information you can imagine. You could have variables that hold colors or
pictures. Later in the class, we’ll even talk about how you can create your own data
types. For example, you could create a Person type that includes a person’s name,
birthday, and favorite food.
```
//5_common_types.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInteger int = 5
	//first of the three methods to determine type
	fmt.Printf("%T \n", myInteger)
	var myFloat float64 = 5.0
	//second method to determine type
	fmt.Println(reflect.TypeOf(myFloat))
	var myString string = "5.0"
	fmt.Println(reflect.ValueOf(myString).Kind())
	var myBoolean bool = true
	fmt.Println(reflect.ValueOf(myBoolean).Kind())
}

```
Output 
```
$ go run 5_common_types.go 
int 
float64
string
bool
$ 
```
## Mixing Types

Operators are the most basic way we act
on variables. Some of the basic operators are mathematical, things like addition and
subtraction, so let’s briefly talk about the interaction between types and operators.
For example, what happens when we try to multiply two variables with different
types?
```
// 6_mixing_types.go
package main

import "fmt"

func main() {
	var myNumber1 int = 5
	var myNumber2 int = 3
	var myString string = "Hello world"
	var myFloat float64 = 5.1
	fmt.Println(myNumber1 * myNumber2)
	fmt.Println(myFloat * float64(myNumber2))
	fmt.Println(myNumber1 * myString)

}
```
Output:
```
$ go run 6_mixing_types.go 
# command-line-arguments
./6_mixing_types.go:13:24: invalid operation: myNumber1 * myString (mismatched types int and string)
$ 

```
## Type Conversions in Golang 

There will be times when we need to convert among the different data types. For
example, imagine writing a program that reads some text in from a file. It encounters
the text “5554321.” Should this be stored as the string of text “5554321” or as the
number 5,554,321? Or, imagine it encounters the worse “False”. Should this remain
as the text “False” or be stored as the boolean value False? The correct answer
will depend on the purpose of the code, and so we need a way to convert between
the different types.

### Converting to Strings

Fortunately, golang provides a very simple way of accomplishing this. First, let’s
look at converting numbers into strings. It will become more clear as we go on in the
material why this is useful, but for now we’ll just focus on how to do it.

```
// 7_converting_to_strings.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myNumber int = 5
	fmt.Printf("%T \n", myNumber)
	var myNumberAsString string = string(myNumber)
	fmt.Println(reflect.TypeOf(myNumberAsString))
}
```
OutPut:
```
$ go run 7_converting_to_strings.go 
int 
string
$ 
```
The string() function will work to convert nearly any kind of variable to some
kind of string.
```
package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var myDate = time.Now()
	var myDateAsString string = myDate.String()
	fmt.Println(reflect.TypeOf(myDateAsString), myDateAsString)
}

```
Output:
```
$ go run 8_converting_to_string.go 
string 2021-10-02 14:51:38.832126584 +0530 IST m=+0.000133644
$ 
```
In practice, you may not always need to do this
```
package main

import (
	"fmt"
	"time"
)

func main() {
	var myDate time.Time = time.Now()
	fmt.Println(myDate)
}

```
Output
```
$ go run 9_converting_to_string.go 
2021-10-02 15:03:38.454841859 +0530 IST m=+0.000116320
```

### Converting From Strings

Practically speaking, converting from strings to other values is just as easy

```
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var myIntAsString string = "5"
	myInt, err := strconv.Atoi(myIntAsString)
	if err == nil {
		fmt.Println(reflect.TypeOf(myInt))
	}

	var myFloatAsString string = "5.1"
	myFloat, err := strconv.ParseFloat(myFloatAsString, 64)
	if err == nil {
		fmt.Println(reflect.TypeOf(myFloat))
	}

	var myBoolAsString string = "true"
	myBool, err := strconv.ParseBool(myBoolAsString)
	if err == nil {
		fmt.Println(reflect.TypeOf(myBool))
	}
}

```
Output:
```
$ go run 10_converting_from_string.go 
int
float64
bool
```
There is an important note here compared to the string conversion function.
Whereas almost anything can be converted to a string, not everything can be converted to integers, booleans, and floats. For example, if a string holds an integer, we can still convert it to a float because golang knows what to to with int given as float. However, if a string holds a float,
we cannot convert it to an int because golang does not know what to do with a float given as int
```
//11_converting_from_string.go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var myIntAsString string = "5"
	myIntAsFloat, err := strconv.ParseFloat(myIntAsString, 64)
	// this will work
	if err == nil {
		fmt.Println(reflect.TypeOf(myIntAsFloat))
	}

	var myFloatAsString string = "5.1"
	myFloatAsInt, err := strconv.Atoi(myFloatAsString)
	// this will cause an error
	if err == nil {
		fmt.Println(reflect.TypeOf(myFloatAsInt))
	} else {
		fmt.Println(err)
	}
}
```
Output:
```
$ go run 11_converting_from_string.go 
float64
strconv.Atoi: parsing "5.1": invalid syntax
```

### User Input

Now is a good time to talk about getting input from the user in console applications. Imagine you’re writing a program that returns the square of a number that the user puts in. Without letting the user edit the code, how do you write this?

Here's How in Golang we have scanln. Whatever text you supply to input is given to the user as a prompt. The program then waits for the user to type something into the console and press enter; whatever the user types is stored in the variable.
```
//12_user_input.go
package main

import "fmt"

func main() {
	fmt.Println("Enter an Integer: ")
	var myInt int
	fmt.Scanln(&myInt)
	fmt.Println("The Entered integer is: ", myInt)
}

```
Output:
```
$ go run 12_user_input.go 
Enter an Integer: 
13
The Entered integer is:  13
```

In Golang the type of the input deponds on the **variable assignation**
```
//13_user_input.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Enter an Integer: ")
	var myInt int
	fmt.Scanln(&myInt)
	fmt.Println("The Type of the variable entered before is: ", reflect.TypeOf(myInt))

	fmt.Println("Enter a Float: ")
	var myFloat float64
	fmt.Scanln(&myFloat)
	fmt.Println("The Type of the variable entered before is: ", reflect.TypeOf(myFloat))

	fmt.Println("Enter a string: ")
	var myString string
	fmt.Scanln(&myString)
	fmt.Println("The Type of the variable entered before is: ", reflect.TypeOf(myString))
}

```
Output:
```
$ go run 13_user_input.go 
Enter an Integer: 
12
The Type of the variable entered before is:  int
Enter a Float: 
3.14
The Type of the variable entered before is:  float64
Enter a string: 
gdhdgfgdg
The Type of the variable entered before is:  string
```

## Reserved Keywords in Go

In an earlier lesson, we noted that one of the rules around what names you can use for variables is that you can’t use some of the reserved keywords. What this means is that there are certain words to which go attaches special
meaning and uses to understand what you’re telling it to do. If you try to use these words to mean something different, go gets confused. 

This occurs in the real world in natural conversation, too. For example, if you named your dog “for” or “some,” people would get confused when you would say things like “Can you take for for a walk?” or “Can you pour some some food?”

### Go's Reserved Words:

break, default, func, interface, select, case, defer, go, map, struct, chan, else, goto, package, switch, const, fallthrough, if, range, type, continue, for, import, return, var. 

misusing keywords will lead to error :-
```
//14_misusing_keywords.go
package main

import "fmt"

func main() {
	var default int = 1
	fmt.Println(default)
}
```
Output:
```
# command-line-arguments
./14_misusing_keywords.go:7:6: syntax error: unexpected default, expecting name
./14_misusing_keywords.go:8:14: syntax error: unexpected default, expecting expression
```
