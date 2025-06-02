# Smart Copy

Smart Copy is a tool that allows you to copy the contents of files or directories to the clipboard, with intelligent filtering options. It consists of two components:

1. A command-line tool written in Go.
2. A VSCode extension that provides commands to copy files or open editors.

## Command-Line Tool

### Installation

To install the command-line tool, ensure you have Go (version 1.24.3 or later) installed on your system. Then, build the binary with:

```bash
go build -o smartcopy
```

This creates an executable named `smartcopy` in your current directory.

### Usage

Basic usage:

```bash
smartcopy <path> [--mode <mode>]
```

- `<path>`: The file or directory to copy.
- `--mode <mode>`: An optional flag to specify the copying mode (defaults to "minimal").

#### Modes

- **minimal**: Copies only files not ignored by `.gitignore` that are deemed important (e.g., source code, excluding binaries and documentation).
- **gitignore**: Copies all files not ignored by `.gitignore`.
- **all**: Copies all files, ignoring `.gitignore` and other filters.
- **gitignore-list**: Lists files that would be copied in `gitignore` mode without copying them.

### Examples

1. Copy a directory's contents in minimal mode:
   ```bash
   smartcopy /path/to/directory
   ```

2. Copy all files, ignoring `.gitignore`:
   ```bash
   smartcopy /path/to/directory --mode all
   ```

3. List files that would be copied in gitignore mode:
   ```bash
   smartcopy /path/to/directory --mode gitignore-list
   ```

## VSCode Extension

The VSCode extension enhances the command-line tool by integrating Smart Copy functionality into Visual Studio Code.

### Features

- Copy contents of all open editors to the clipboard.
- Copy contents of a selected folder using different modes (minimal, gitignore, all).

### Installation

To build and install the extension from source:

1. Navigate to the `vscode-smartcopy` directory:
   ```bash
   cd vscode-smartcopy
   ```
2. Install dependencies and package the extension:
   ```bash
   npm install
   vsce package
   ```
3. Install the generated `.vsix` file in VSCode via the Extensions view (`Ctrl+Shift+X` or `Cmd+Shift+X` on Mac) by selecting "Install from VSIX."

Alternatively, install it from the VSCode Marketplace (if published).

### Usage

Access commands via the Command Palette (`Ctrl+Shift+P`):

- `Smart Copy: Copy Open Editors`
- `Smart Copy: Run Minimal`
- `Smart Copy: Run Gitignore List`
- `Smart Copy: Run Gitignore`
- `Smart Copy: Run All`

Or right-click a folder in the Explorer and select a Smart Copy option.

## How It Works

The command-line tool traverses the specified directory, collecting file contents based on the mode:
- Uses `.gitignore` files to filter ignored files (if enabled).
- In "minimal" mode, excludes non-essential files (e.g., binaries, images, documentation) using language detection.
- Prepends each file's content with a header like:
  ```
  // ---- relative/path/to/file ----
  file contents here
  ```

The VSCode extension either:
- Copies open editors' contents with similar headers.
- Invokes the command-line tool for folder-based copying.

## Building for Multiple Platforms

Build binaries for Windows, macOS, and Linux using the provided script:

```bash
node scripts/build-go.mts
```

Binaries are output to `vscode-smartcopy/bin` for each platform and architecture, used automatically by the VSCode extension.

## Dependencies

### Command-Line Tool
- [github.com/sabhiram/go-gitignore](https://github.com/sabhiram/go-gitignore)
- [github.com/atotto/clipboard](https://github.com/atotto/clipboard)
- [github.com/go-enry/go-enry](https://github.com/go-enry/go-enry)

### VSCode Extension
- Node.js and VSCode APIs

See `go.mod` and `package.json` for full dependency lists.

## Contributing

Contributions are welcome! Submit issues or pull requests on the [GitHub repository](https://github.com/mango/smart-copy).

## License

This project is licensed under the MIT License. See the `LICENSE` file for details
