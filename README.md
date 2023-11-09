# Stamp - Watermark Command Line Tool

Stamp is a command-line tool for adding watermarks to images. With Stamp, you can easily overlay text or images on your images to protect your content or add branding.

## Installation

You can download the latest version of Stamp from the [releases page](https://github.com/KhashayarKhm/stamp/releases/latest). Choose the binary that matches your operating system and architecture, and then follow these installation instructions.

### Linux/macOS

```shell
# Download the binary
$ curl -Lo stamp https://github.com/KhashayarKhm/stamp/releases/latest/download/stamp-linux-amd64

# Make it executable
$ chmod +x stamp

# Move it to a directory in your PATH (optional)
$ sudo mv stamp /usr/local/bin/
```

### Windows

Download the binary from the [releases page](https://github.com/KhashayarKhm/stamp/releases/latest) and add it to your system's PATH.

## Usage

Stamp provides the following commands:

- `watermark` :  Add watermarks to images.
- `completion`: Generate autocompletion scripts for various shells.
- `help`      : Get help on any specific command or usage.

### Adding Watermarks

The `watermark` command allows you to add watermarks to images.

Example:

```shell
# Add a watermark to an image
$ stamp watermark input.jpg -o output.jpg -w watermark.png
```

### Autocompletion

You can generate autocompletion scripts for your preferred shell using the `completion` command. Supported shells include Bash, Zsh, and Fish.

Example (Bash):

```shell
# Generate Bash autocompletion script
$ stamp completion bash > stamp-completion.sh

# Source the autocompletion script
$ source stamp-completion.sh
```

## Contributing

Feel free to contribute to this project by opening issues, making feature requests, or submitting pull requests. Your feedback and contributions are welcome.

## License

This project is licensed under the GPLv3 or later License. See the [LICENSE](https://github.com/KhashayarKhm/stamp/blob/main/LICENSE) file for more details.
