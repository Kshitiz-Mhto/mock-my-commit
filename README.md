<p align="center">
<img src="./assets/logo.png" alt="mock-my-commit logo" width=280>
</p>

**mock-my-commit**: A Git commit message validator that roasts your bad commit messages with sarcastic, passive-aggressive one-liners. Powered by the Mistral text generation model API, this tool ensures you never write another "fixed stuff" commit without feeling properly judged.

---

### Features

- **Validates Git Commit Messages**: Ensures commit messages meet best practices before Git registers the commit.

- **Sarcastic review**: Generates passive-aggressive feedback for poorly written commit messages that embeds in your soul.

- **AI-Powered**: Uses Mistral text generation API to craft scarcastic one-liners that you deserve.

- **Seamless Git Hook Integration**: Works as a `commit-msg` hook to validate messages before they are committed.

---

### Project: mock-my-commit

This project, `mock-my-commit`, is built using the Go programming language. Below is a list of the commands offered by this tool.

<p align="center">
  <img src="./assets/home.png" alt="mock-my-commit Logo" width="800">
</p>

> *Demo*

<p align="center">
  <img src="./assets/working.png" alt="mock-my-commit demo" width="800">
</p>

---

### *Brief Command Screenshot*

> *setup-API-key*

<p align="center">
  <img src="./assets/setup.png" alt="mock-my-commit setup" width="800">
</p>

> *install*

<p align="center">
  <img src="./assets/install.png" alt="mock-my-commit install <scope>" width="800">
</p>

> *run-hook*

<p align="center">
  <img src="./assets/run_hook.png" alt="mock-my-commit rub-hho" width="800">
</p>

---

## Technologies Used

- [cobra](https://github.com/spf13/cobra) – CLI framework for handling commands.
- [term](https://golang.org/x/term) – Terminal handling utilities.
- [emoji](https://github.com/enescakir/emoji) – For adding emoji reactions to roasts.
- [terminfo](https://github.com/xo/terminfo) – Terminal capabilities information.
- [mistral-go](https://github.com/gage-technologies/mistral-go) – API client for interacting with the Mistral text generation model.
- [gookit/color](https://github.com/gookit/color) – Colored terminal output.

---

## Local setup

### Prerequisites

- **Linux OS**  
- **Go (version 1.22 or higher):**  
  Install from the official [Go installation guide](https://go.dev/doc/install).
- **Make:**  
  Install using your package manager (e.g., `sudo apt-get install make` on Debian/Ubuntu).
- **Mistral API Key:**  
  Generate your API key from the [Mistral API Keys console](https://console.mistral.ai/api-keys/).

### Installation

1. Navigate to the project repo and run 

```bash
make build
```

2. Create a symlink using 

```bash
sudo ln -s /path/to/mock-my-commit/bin/mock-my-commit /usr/local/bin/mock-my-commit
```

3. Test the Executable

```bash
mock-my-commit
```

4. View the Manual

```bash
mock-my-commit -h
```

### Flowchart

> *setup-flowdiagram*

<p align="center">
  <img src="./assets/setup_flow.png" alt="mock-my-commit rub-hho" width="800">
</p>



## Makefile Documentation

This `Makefile` provides an easy interface to build, test, install, and clean the mock-my-commit project. It automates common tasks required for the development and deployment of the project.

## Variables

- **BINARY_NAME**: Specifies the name of the output binary file. In this case, it is set to `mock-my-commit`.

- **OUTPUT_DIR**: Specifies the directory where the binary will be output. The default is `bin`.

- **MAIN_FILE**: Specifies the path to the main Go file. This is where the Go code for the project starts (default: `./main.go`).

- **INSTALL_DIR**: Specifies where the binary should be installed. By default, it installs to `$HOME/go/bin` if the `GOBIN` environment variable is not set.

## Targets

The Makefile has the following targets that automate the build process:

#### `build`

- **Description**: Builds the binary from the main Go file and outputs it to the `bin` directory.

- **Commands**:
    1. Creates the `bin` directory if it doesn't already exist.

    2. Runs the `go build` command to compile the Go project into a binary named `mock-my-commit`.

    3. Outputs the path to the built binary.

- **Usage**: 

    ```bash
    make build
    ```

#### `test`

- **Description**: Runs the tests for the project using `go test`.

- **Commands**:

    1. Executes the `go test` command on the entire project (`./...`).

    2. Prints the results of the tests to the terminal.

- **Usage**: 

    ```bash
    make test
    ```

### `install`

- **Description**: Builds the project and installs the binary to the `GOBIN` directory.

- **Commands**:

    1. Runs the `build` target to compile the binary.

    2. Installs the binary to the directory specified by `GOBIN`.

- **Usage**:

    ```bash
    make install
    ```

### `clean`

- **Description**: Cleans up the build directory by removing the `bin` directory and its contents.

- **Commands**:

    1. Deletes the `bin` directory.

    2. Outputs a message when cleanup is complete.

- **Usage**: 

    ```bash
    make clean
    ```
