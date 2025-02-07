# brwords
## DISCLAIMER
### Please DO NOT flood dicio website with requests. Using it from CLI is fine, but you must use some cache if you are going to use it in a project that processes many words.
## Overview

The `brwords` package is a Go package that provides an easy way to extract word meanings, synonyms, etymologies, and usage examples from the (Dicio)[https://dicio.com.br] website (a popular Brazilian Portuguese dictionary). The package includes a Command Line Interface (CLI) tool that enables users to access these features directly from the terminal.

## Features

- **Extract Word Meanings**: Retrieve the definition of any word in Portuguese from Dicio.
- **Synonyms**: Get a list of synonyms for the word you're querying.
- **Etymology**: Retrieve the origin and history of the word.
- **Examples**: Fetch real-life example sentences showing the word in context.
- **Command-Line Interface**: Access the functionality directly from the terminal.
- **Simple Installation**: Install the CLI with a single command.

## Usage
Note: the output is completely in Brazilian Portuguese, but the code and exposed interfaces are in English.

First, add the package using the following command:

```bash
go get github.com/bruceneco/brwords
```

Next, import and use it in your code:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/bruceneco/brwords"
)

func main() {
	s := brwords.NewScrap()
	word, _ := s.Word("sandice")
	res, _ := json.MarshalIndent(word.Meanings, "", "\t")
	fmt.Println(string(res))
}
/*  go run main.go
[
	{
		"conteudo": "Característica, condição, particularidade de quem se comporta ou se comunica através de tolices, de modo tolo ou simplório."
	},
	{
		"conteudo": "Característica, ação ou discurso que denota tolice, ignorância ou ausência de inteligência; tolice, idiotice, parvoíce."
	},
	{
		"conteudo": "Discurso, comportamento, ação que demonstra ausência de lógica; loucura."
	},
	{
		"conteudo": "Abatimento demonstrado com o passar dos anos; fraqueza mental relacionada à idade; senilidade."
	}
]
*/
```
## Installation (CLI)

To use `brwords` as CLI tool, you can easily install it via Go using the following command:

```bash
go install github.com/bruceneco/brwords/cmd/br@latest
```

This will install the `br` command into your `$GOPATH/bin` (or `$HOME/go/bin` if using Go modules).

Make sure you have Go installed and set up on your system before running the command.

## Usage (CLI)

Once installed, you can use the `br` command to query word meanings, synonyms, etymologies, etymologies, and examples from Dicio.

### Basic Command Syntax

```bash
br <word>
```

Where `<word>` is the word you want to look up.

### Example Usage

1. **Get the meaning of a word:**

```bash
br sandice
```

This will return the definition of the word "sandice" (love) from Dicio.

2. **Get synonyms of a word:**

```bash
br sandice --sy
```

This will list synonyms for the word "sandice".

3. **Get etymologies of a word:**

```bash
br sandice -et
```

This will list etymologies for the word "sandice".

4. **Get example usage:**

```bash
br sandice -ex
```

This will display example sentences showing the usage of "sandice".

### Flags

The CLI supports the following flags to narrow down the results:

- `-sy`: Fetch synonyms for the word.
- `-et`: Fetch antonyms for the word.
- `-ex`: Fetch example usage sentences.
- `-m | -d`: Fetch meanings.
- `-json | -j`: Output the result in JSON format.
- `-h`: Display help information.

For example, to get both the meaning and synonyms:

```bash
br sandice -sy
```

## Contributing

Feel free to open issues or contribute to the project by creating pull requests. If you encounter any bugs or have feature requests, please raise an issue on the GitHub repository.

## Contact

For any issues or inquiries, feel free to reach out to the maintainer via the GitHub repository or email.