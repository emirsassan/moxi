# Moxi

Moxi is a lightweight build tool designed to be a simple alternative to CMake. It simplifies the build process for projects, providing an easy-to-use solution for compiling.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Moxi is built with simplicity in mind, offering a straightforward approach to building projects without the complexities often associated with larger build systems. Whether you're working on a small application or a larger software project, Moxi aims to streamline the build process, allowing you to focus on your code rather than build configurations.

Key features of Moxi include:

- **Ease of Use:** Moxi is designed to be easy to understand and use. With a minimal learning curve, developers can quickly adopt Moxi for their projects.

- **Fast Builds:** Moxi prioritizes efficient build times, making it well-suited for projects of various sizes. Build times are optimized to enhance developer productivity.

## Getting Started

### Installation

As of now, Moxi does not provide precompiled releases. To use Moxi, you'll need to compile it from the source code. Moxi is written in Go, so make sure you have Go installed on your system.

This version specifies the installation steps for Windows, assumes the use of PowerShell, and refers to the executable as `moxi.exe`.

1. Clone the Moxi repository:
   
   ```PowerShell
   git clone https://github.com/emirsassan/moxi.git
   ```

2. Change into the Moxi directory:

   ```bash
   cd moxi
   ```

3. Build Moxi:

   ```bash
   go build -o moxi main.go
   ```

4. Optionally, move the moxi.exe binary to a directory in your system's PATH.

Now, you should be able to use the `moxi` command in your terminal.

## Usage

Moxi uses a simple syntax for defining tasks in a `Moxifile`. Here's an example `Moxifile`:

```mo
task: build => go build -o myapp main.go
task: test => go test ./...
```

In this example:

- The `build` task compiles the main Go file into an executable named `myapp`.
- The `test` task runs tests for the entire project.

To execute a task, run:

```bash
moxi build
```

Replace `build` with the desired task name. If no task is specified, Moxi will run all tasks defined in the `Moxifile`.

## Contributing

We welcome contributions from the community to enhance Moxi. If you'd like to contribute, please follow these steps:

1. **Fork the Repository:** Click the "Fork" button on the top-right corner of the [GitHub repository](https://github.com/emirsassan/moxi) to create your own copy.

2. **Create a New Branch:** Create a new branch for your feature or bug fix. Naming conventions may vary, but use something descriptive like `feature/new-feature` or `fix/bug-fix`.

3. **Make Changes:** Make your changes in the new branch. Please adhere to our coding standards and conventions.

4. **Commit Changes:** Commit your changes with clear and concise commit messages.

5. **Push Changes:** Push your changes to your fork on GitHub.


6. **Open a Pull Request (PR):** Open a pull request on the [main repository](https://github.com/emirsassan/moxi) and describe your changes. Make sure to reference any related issues.
   Our maintainers will review your contribution and provide feedback. Thank you for helping improve Moxi!


## License

Moxi is licensed under the [Apache-2.0 license](LICENSE).
