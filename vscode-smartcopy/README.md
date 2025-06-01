# Smart Copy

Smart Copy is a Visual Studio Code extension that enhances your workflow by providing intelligent copying of file contents to the clipboard. It offers two primary features:

- **Copy Open Editors**: Copy the contents of all currently open editors, each prefixed with a file path header.
- **Smart Copy Folder**: Copy files from a selected folder or workspace using customizable filtering modes.

## Features

### Copy Open Editors
- **Command**: `Smart Copy: Copy Open Editors`
- **Description**: Copies the contents of all open files in your VS Code workspace to the clipboard. Non-file documents (e.g., virtual or git-related files) are excluded. Each file’s content is prefixed with a header like `// ---- relative/path/to/file ----` for easy identification.
- **How to Use**: Open the Command Palette (`Ctrl+Shift+C`), or right click editor. 

![Screenshot showing the confirmation message after copying open editors.]

### Smart Copy Folder
- **Commands**:
  - `Smart Copy: Run Minimal`
  - `Smart Copy: Run Gitignore List`
  - `Smart Copy: Run Gitignore`
  - `Smart Copy: Run All`
- **Description**: Copies files from a selected folder (or the workspace root) based on different filtering modes:
  - **Minimal**: Copies only essential source files, excluding binaries, images, and documentation.
  - **Gitignore**: Copies files not excluded by `.gitignore` rules.
  - **All**: Copies all files without applying any filters.
  - **Gitignore List**: Displays a list of files that would be copied in "Gitignore" mode without copying them to the clipboard.
- **How to Use**: Right-click a folder in the Explorer pane and select a Smart Copy option, or run the commands via the Command Palette.

**Tip**: Enhance productivity by assigning keyboard shortcuts to these commands. Learn more in the [Key Bindings documentation](https://code.visualstudio.com/docs/getstarted/keybindings).

## Requirements

- The extension comes with pre-built binaries for Windows, macOS, and Linux, ensuring it works out of the box for most users.
- **Linux Users**: You may need to install an additional dependency (e.g., `xclip` or `wl-clipboard`) to enable clipboard functionality. Refer to your distribution’s package manager for installation instructions.

## Extension Settings

This extension does not currently add any custom settings via `contributes.configuration`. All features are accessible directly through commands.

## Known Issues

No known issues at this time. If you encounter any problems, please report them on the [GitHub repository](https://github.com/your-repo/vscode-smartcopy).

## Release Notes

### 1.0.0
Initial release with support for copying open editors and smart copying folders with multiple filtering modes.

## Following Extension Guidelines

This extension complies with the [VSCode Extension Guidelines](https://code.visualstudio.com/api/references/extension-guidelines) to ensure a reliable and user-friendly experience.

Enjoy using Smart Copy!