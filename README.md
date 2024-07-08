# Combigo

Combigo is a Go module designed to generate combinations of characters based on a specified pattern from given alphabets. The pattern uses capital letters to signify positions where characters from the alphabets should be inserted. Each capital letter corresponds to an alphabet provided in the input.

## Features

- Generate all possible combinations based on the pattern and alphabets.
- Handle patterns with static characters alongside dynamic characters from alphabets.
- Simple and intuitive API.

## Installation

To install the module, use the following command:

```bash
go get github.com/brevis/combigo
```

## Usage

### Importing the Module

First, import the module in your Go code:

```go
import (
    "github.com/brevis/combigo"
)
```

### Generating Combinations

Use the `GenerateCombinations` function to generate combinations based on the pattern and alphabets.

#### Example

```go
package main

import (
    "fmt"
    "github.com/brevis/combigo"
)

func main() {
    pattern := "ABlix.com"
    alphabets := []string{"aeiouy", "bcdfghjklmnpqrstvwxz"}

    combinations := combigo.GenerateCombinations(pattern, alphabets)
    
    for _, combo := range combinations {
        fmt.Println(combo)
    }
}
```

#### Output

The above code will generate and print combinations such as:

```
ablix.com
aclix.com
...
yblix.com
yclix.com
...
yzlix.com
```

### API Reference

#### `func GenerateCombinations(pattern string, alphabets []string) []string`

Generates all possible combinations based on the given pattern and alphabets.

- **pattern**: A string where capital letters indicate positions for characters from alphabets.
- **alphabets**: An array of strings, where each string is an alphabet to be used for corresponding capital letters in the pattern.

#### `func GetCombinationsCount(pattern string, alphabets []string) int`

Returns the total number of possible combinations based on the pattern and alphabets.

- **pattern**: A string where capital letters indicate positions for characters from alphabets.
- **alphabets**: An array of strings, where each string is an alphabet to be used for corresponding capital letters in the pattern.

#### `func GetFirstPatternCombination(pattern string, alphabets []string) string`

Returns the first possible combination based on the pattern and alphabets.

- **pattern**: A string where capital letters indicate positions for characters from alphabets.
- **alphabets**: An array of strings, where each string is an alphabet to be used for corresponding capital letters in the pattern.

#### `func GetLastPatternCombination(pattern string, alphabets []string) string`

Returns the last possible combination based on the pattern and alphabets.

- **pattern**: A string where capital letters indicate positions for characters from alphabets.
- **alphabets**: An array of strings, where each string is an alphabet to be used for corresponding capital letters in the pattern.

#### `func GetNextCombination(combination string, pattern string, alphabets []string) (string, error)`

Generates the next combination based on the current combination, pattern, and alphabets.

- **combination**: The current combination string.
- **pattern**: A string where capital letters indicate positions for characters from alphabets.
- **alphabets**: An array of strings, where each string is an alphabet to be used for corresponding capital letters in the pattern.

#### `func GetNextCharacter(character string, alphabet string) (string, error)`

Returns the next character in the given alphabet.

- **character**: The current character.
- **alphabet**: The alphabet string.

## Contributing

If you'd like to contribute to Combigo, please fork the repository and use a feature branch. Pull requests are warmly welcome.

## License

Combigo is released under the MIT License. See the LICENSE file for more details.