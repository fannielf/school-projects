The program will receive as arguments the name of a file containing a text that needs some modifications and the name of the file the modified text should be placed in.

The program executes following modifications:

Keyword (hex) replaces the word before with the decimal version of the word (assuming that the word will always be a hexadecimal number).

Keyword (bin) replaces the word before with the decimal version of the word (assuming that the word will always be a binary number).

Keyword (up) converts the word before with the UPPERCASE version of it.

Keyword (low) converts the word before with the lowercase version of it.

Keyword (cap) converts the word before with the Capitalized version of it.

For keywords cap|up|low, if a number appears after it within same quotation: (low, <number>) it handles the previously specified number of words as per keyword.

Every instance of the punctuations ., ,, !, ?, : and ; shall be attached to the previous word and with space apart from the next one, if the following character is not a punctuation itself.

The quotation mark ' will always be found with another instance of it and they will be placed to the right and left of the word/sentence in the middle of them, without any spaces.

Every instance of a is turned into an if the next word begins with a vowel (a, e, i, o, u).
