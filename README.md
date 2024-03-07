
# CLI Timer (clitimer)

CLI Timer (clitimer) is a simple, versatile command-line utility written in Go. It offers two primary modes of operation: a stopwatch mode that counts up from zero, and a countdown timer mode that counts down from a specified duration. This tool is designed to be dependency-free, requiring only the Go runtime to build and run, making it highly portable and easy to use.

## Features

- **Stopwatch Mode:** By default, `clitimer` functions as a stopwatch, counting up in seconds from the moment it is started until it is manually stopped with `CTRL-C`.
- **Countdown Timer Mode:** By providing a duration through a command-line flag, `clitimer` operates as a countdown timer, counting down to zero from the specified time.

## Installation

To use `clitimer`, you need to have Go installed on your system. If you don't have Go, you can download and install it from the [official Go website](https://golang.org/dl/).

### Building from Source

1. Clone the repository or download the source code to your local machine.
2. Navigate to the directory containing `clitimer`.
3. Build the binary using the Go compiler:

```sh
go build -o clitimer
```

This command compiles the source code into a binary named `clitimer`.

## Usage

After building `clitimer`, you can start it in either mode by using the following commands:

### Stopwatch Mode

Simply run the binary without any flags:

```sh
./clitimer
```

### Countdown Timer Mode

To use the countdown timer mode, specify the duration with the `-d` flag followed by the duration string. The duration string is a sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300s", "1.5h" or "2h45m". Valid time units are "s", "m", "h".

```sh
./clitimer -d 1m30s
```

This command starts a countdown timer for 1 minute and 30 seconds.

## Stopping the Timer

In both modes, you can stop the timer by pressing `CTRL-C`. Upon stopping, the current time and the total elapsed time (for stopwatch mode) or the remaining time (for countdown mode) will be displayed.

## Contributing

Contributions to `clitimer` are welcome! Feel free to open issues or submit pull requests on our GitHub repository.

## License

`clitimer` is released under the MIT License. See the LICENSE file for more details.
