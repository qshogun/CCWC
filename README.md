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

Sure, here's the completed README.md file:

```markdown
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

This command will count the number of bytes, lines, words, and characters in `test.txt`.

## Installation

To install CCWC, you need to have Go installed on your machine. Then, you can clone this repository and build the tool with the following commands:

```sh
git clone https://github.com/qshogun/ccwc.git
cd ccwc
go build
```

This will create an executable named `ccwc` in the current directory.

## Testing

CCWC comes with a suite of tests that you can run with the following command:

```sh
go test
```

## Contributing

Contributions to CCWC are welcome. Please open an issue or a pull request if you want to contribute.

## License

CCWC is licensed under the MIT License. See the `LICENSE` file for more details.

This README.md file provides a comprehensive overview of the CCWC tool, including its usage, installation instructions, testing procedures, contribution guidelines, and licensing information.
