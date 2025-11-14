# go-reload

A small Golang utility that reads an input text file, applies a series of transformation rules, and writes the processed result to an output file.

## Description

The program takes two arguments:

1. Input file – the path to the original text

2. Output file – the path where the modified text will be written

It processes the text according to a set of keywords and formatting rules described below.

## Transformation Rules

### Keyword-Based Modifiers

These keywords act on the word directly before them:

- `(hex)` – Converts the previous word from hexadecimal to decimal.
- `(bin)` – Converts the previous word from binary to decimal.
- `(up)` – Converts the previous word to UPPERCASE.
- `(low)` – Converts the previous word to lowercase.
- `(cap)` – Converts the previous word to Capitalized form.

### Multi-Word Modifiers

For `cap`, `up`, and `low`, you can optionally specify how many preceding words to modify:

- `(low, 3)`
- `(up, 2)`
- `(cap, 5)`

This applies the keyword’s effect to the given number of words before the keyword.

### Punctuation Rules

The characters:

```
. , ! ? : ;
```

**must always**:

* Attach directly to the previous word,

* Have a space before the next word, unless the next character is also punctuation.

### Quotation Rules

Every `'` will appear as a pair (opening and closing). **They must:**

* Wrap the enclosed content **without spaces** inside the quotes.

Example:
`' hello world ' → 'hello world'`

### Article Correction

If the word a is followed by a word starting with a vowel (`a, e, i, o, u`), it must be replaced with an.

Example:
`a apple → an apple`

## Prerequisites

- Go 1.23.2 or higher
- An input `.txt` file in the project root 
- Go modules enabled (`go mod tidy`)

## Usage
```bash
go run main.go input.txt output.txt
```

## Examples

Copy these examples to the input.txt file and run the program.

**Example 1:**

Testing punctuation, quotation and article rules

***Input:***

`He said ' hello world ! ' This is a apple, a banana, and a orange .`

***Expected output:***

`He said 'hello world!' This is an apple, a banana, and an orange.`

**Example 2:**

Testing hex, bin and casing modifiers.

***Input:***

`1a (hex) 1101 (bin) golang (up) IMPORTANT (low) programming languages (cap, 2)`

***Expected output:***

`26 13 GOLANG important Programming Languages`