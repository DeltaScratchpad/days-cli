# Days CLI

A simple tool to track how many hybrid days you need to work in a month.
Defaults to a 0.2 ratio, but can be adjusted via flags.

## Overview
This CLI calculates the number of days you need to be in the office based on the number of weekdays in a month, adjusted by a work ratio and any days off taken. The result is rounded up to the nearest 0.5.

## Usage

Run the tool from the command line:

```
ᐅ  ./days                  
Calculate hybrid days needed (rounded to nearest 0.5).
Flags: -year  2025, -month 7, -ratio  0.2, -off 0

Year: 2025, Month: 7 has 23 weekdays
You need to be in the office: 5 days
```

### Flags
- **-year:** Specify the year (default: current year).
- **-month:** Specify the month (default: current month).
- **-ratio:** Define the fraction of weekdays required in the office (default: 0.2).
- **-off:** Define the number of days off already taken (default: 0).

## Examples

To calculate for July 2025:

```bash
ᐅ ./days -year 2025 -month 7
```

For help:
```
ᐅ ./days -h
```

## Bash Alias Setup

To set up a bash alias with the ratio fixed to 0.5, add the following to your shell config (e.g. ~/.bashrc or ~/.zshrc):

```bash
alias days="[saved location]/days -ratio 0.5"
```

Reload your shell with:

```bash
source ~/.bashrc   # or: source ~/.zshrc
```
