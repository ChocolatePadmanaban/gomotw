# String Functions Goglang 

- First Letter of the Words Capital 


```
package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Printing all words in Title letters
	c := cases.Title(language.Und, cases.NoLower)

	s := "The quick brown fox jumped over the lazy dog."
	fmt.Println(c.String(s))

	// If your string is in a language with the special casing rules, for example, Turkish or Azeri, use the strings.ToUpperSpecial() function.
	c = cases.Title(language.Turkish, cases.NoLower)
	s = "en iyi web sitesi"

	fmt.Println(c.String(s))

}
```

- Upper case title can be given to other languages too

## OutPut :

```
String_Functions % go run .    
The Quick Brown Fox Jumped Over The Lazy Dog.
En İyi Web Sitesi
```