# CCWC - Custom Character Word Counter

CCWC is a command-line tool written in Go that provides various statistics about a text file. It can count the number of bytes, lines, words, and characters in a file.

## Usage

You can use CCWC with the following command-line options:

- `-c`: Count the number of bytes in a file
- `-l`: Count the number of lines in a file
- `-w`: Count the number of words in a file
- `-m`: Count the number of characters in a file

If no option is provided, CCWC will count all four statistics.

Here is an example of how to use CCWC:

```sh
ccwc -c -l -w -m test.txt
```
