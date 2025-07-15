# Days CLI

A simple tool to track how many hybrid days you need to work in a month.
Defaults to a 0.2 ratio, but can be adjusted via flags.

## Overview

This CLI calculates the number of days you need to be in the office based on the number of weekdays in a month, adjusted by a work ratio and any days off taken. The result is rounded up to the nearest 0.5.

## Installation

### Prebuilt binaries (Linux & Windows)

You can get a pre-built binary for your system from the GitHub Releases page.

1. **Go to Releases:** Visit the [Releases page](https://github.com/DeltaScratchpad/days-cli/releases).
2. **Download:** Find the latest release and download the appropriate binary for your operating system and architecture.
    * **Linux:** `days-linux-amd64` or `days-linux-arm64`
    * **Windows:** `days-windows-amd64.exe` or `days-windows-arm64.exe`
3. **Make Executable (Linux/macOS):** If you're on Linux or macOS, you might need to make the downloaded binary executable:

    ```bash
    chmod +x ./days-linux-amd64 # or the name of your downloaded binary
    ```

4. **Place in PATH (Optional but Recommended):** For easier access, move the binary to a directory included in your system's `PATH` (e.g., `/usr/local/bin` or `~/bin`).

    ```bash
    mv ./days-linux-amd64 /usr/local/bin/days # Rename for convenience
    ```

### Building from Source (macOS)

If you prefer to build the `days` CLI from its source code on macOS, follow these steps:

1. **Install Go:** Ensure you have Go installed on your system. You can download it from the [official Go website](https://go.dev/doc/install).
2. **Clone the Repository:**

    ```bash
    git clone https://github.com/DeltaScratchpad/days-cli.git
    cd days
    ```

3. **Build the Binary:**

    ```bash
    go build -o days .
    ```

    This command will compile the `days` application and create an executable named `days` in the current directory.
4. **Install to /usr/local/bin:**
    To make the `days` command available globally from any directory in your terminal, move the compiled binary to `/usr/local/bin`:

    ```bash
    sudo mv days /usr/local/bin/
    ```

## Usage

Run the tool from the command line:

```bash
ᐅ  ./days

Calculate hybrid days needed (rounded to nearest 0.5).
Flags: -year  2025, -month 7, -ratio  0.2, -off 0

Year: 2025, Month: 7 has 23 weekdays
You need to be in the office: 5 days
```

### Flags

* **year:** Specify the year (default: current year).
* **month:** Specify the month (default: current month).
* **ratio:** Define the fraction of weekdays required in the office (default: 0.2).
* **off:** Define the number of days off already taken (default: 0).

## Examples

To calculate for July 2025:

```bash
ᐅ ./days -year 2025 -month 7
```

For help:

```bash
ᐅ ./days -h
```

## Bash Alias Setup

To set up a bash alias with the ratio fixed to 0.5, add the following to your shell config (e.g. ~/.bashrc or ~/.zshrc):

```bash
alias days="[saved location]/days -ratio 0.5"
```

Reload your shell with:

```bash
source ~/.bashrc # or: source ~/.zshrc
```
