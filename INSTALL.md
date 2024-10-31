# Installation Guide

This document provides detailed instructions for installing the project using various methods.

## Prerequisites

- `Go` 1.22.7 or later

## Quick Install

To quickly install the latest version using `Go`:

```bash
go install github.com/krisvih32/simplyVerse@latest
```

This will download, compile, and install the project in your `Go` workspace.

## Installation Methods

### 1. Using `go install` without cloning the repo

1. Run the following command:
   ```bash
   go install github.com/krisvih32/simplyVerse@latest
   ```
2. Ensure your Go bin directory is in your `PATH`.

### 2. Using `go install` with cloning the repo

1. Clone the repository:
   ```bash
   git clone https://github.com/krisvih32/simplyVerse.git
   ```
2. Navigate to the project directory:
   ```bash
   cd simplyVerse
   ```
3. Run the install command:
   ```bash
   go install
   ```

### 3. Using Make

<!-- trunk-ignore-begin(markdownlint/MD024) -->

#### Prerequisites

<!-- trunk-ignore-end(markdownlint/MD024) -->

`GNU Make`>=3.81

1. Clone the repository:
   ```bash
   git clone https://github.com/krisvih32/simplyVerse.git
   ```
2. Navigate to the project directory:
   ```bash
   cd simplyVerse
   ```
3. Run the install command:
   ```bash
   make install
   ```

### 4. Using `go build`

1. Clone the repository:
   ```bash
   git clone https://github.com/krisvih32/simplyVerse.git
   ```
2. Navigate to the project directory:
   ```bash
   cd simplyVerse
   ```
3. Build the project:
   ```bash
   go build
   ```
4. Move the project to bin
   ```bash
   mv simplyVerse $(go env GOPATH)/bin
   ```

## Verifying the Installation

After installation, you can verify it by running:

```bash
simplyVerse
```

This should display the prompt as if enter was pressed.

## Troubleshooting

### Common Issues

1. `command not found` error

   - Ensure that your Go bin directory is in your PATH.
   - Adhere to the prerequisites

2. Compilation errors
   - Make sure you have an acceptable version of Go installed.

### Getting Help

If you encounter any issues not covered here, please:

- Check the FAQ in our documentation.
- Open an issue on our GitHub repository.

## Uninstallation

To uninstall the project:

1. If installed with

   ```bash
   go install github.com/krisvih32/simplyVerse
   ```

   Then:

   ```bash
   rm $(go env GOPATH)/bin/simplyVerse
   ```

2. If installed with
   ```bash
   go install
   ```
   Then:
   1. Navigate to the directory you installed the program in
   2. Run:
      ```bash
      go clean -i
      ```
3. If installed with Make, then:
   1. Navigate to the directory you installed the program in
   2. Run:
   ```bash
   make uninstall
   ```

## Updating

To update to the latest version:

1. If installed with

   ```bash
   go install github.com/krisvih32/simplyVerse
   ```

   Then:

   ```bash
   go install github.com/krisvih32/simplyVerse
   ```

2. If installed with
   ```bash
   go install
   ```
   Then:
   1. Navigate to the directory you installed the program in
   2. Run:
      ```bash
      go install
      ```
3. If installed with Make, then:
   1. Navigate to the directory you installed the program in
   2. Run:
   ```bash
   make install
   ```
