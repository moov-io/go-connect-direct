## moov-io/go-connect-direct

[![GoDoc](https://pkg.go.dev/badge/github.com/moov-io/go-connect-direct?status.svg)](https://pkg.go.dev/github.com/moov-io/go-connect-direct)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/go-connect-direct)](https://goreportcard.com/report/github.com/moov-io/go-connect-direct)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

A library for parsing IBM Connect:Direct statistics. The `parser` package provides utilities for parsing summary statistics and completion codes from IBM Sterling Connect:Direct process logs. It extracts structured data such as record IDs, dates, descriptions, process numbers, completion codes, and message IDs from log input strings.

This package is useful for analyzing Connect:Direct process outputs, identifying success/warning/error states, and categorizing events based on predefined record IDs (e.g., copy terminations, session starts/ends, process starts/ends).

Key features:
- Parses log lines into a `SummaryStats` struct.
- Supports lookup of record IDs for various Connect:Direct events and processes.
- Handles completion codes: Success (0), Warning (4), Error (8), Catastrophic Error (16).
- Categorizes records into Event, Process, or External Source.

## Installation

To install the package, use Go modules:

```bash
go get github.com/moov-io/go-connect-direct
```

Replace `moov-io/go-connect-direct` with the actual repository owner if this is hosted on GitHub.

## Usage

Import the package in your Go code:

```go
import "github.com/moov-io/go-connect-direct/parser"
```

### Parsing Log Input

The primary function is `ParseCCode(input string) (SummaryStats, error)`, which takes a raw log string and returns parsed statistics.

Example:

```go
package main

import (
	"fmt"
	"github.com/moov-io/go-connect-direct/parser"
)

func main() {
	input := `
	------ ... ------
	E SUBP  02/03/2026 23:28:45 Submit command issued.
	P PSTR  02/03/2026 23:28:45 sample            14                0      XSMG200I
	======
	`

	stats, err := parser.ParseCCode(input)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		return
	}

	for _, stat := range stats.Stats {
		fmt.Printf("ID: %s, Date: %s, Description: %s, Code: %d\n",
			stat.ID.ID, stat.Date.Format("2006-01-02 15:04:05"), stat.Description, stat.Code)
	}
}
```

### Looking Up Record IDs

Use `LookupRecordID(code string) *RecordID` to retrieve details for a given record code.

Example:

```go
rid := parser.LookupRecordID("CTRC")
if rid != nil {
	fmt.Printf("Record: %s - %s (%s)\n", rid.ID, rid.Description, rid.Category)
}
```

### Completion Codes

Predefined constants for completion codes:
- `CompletionCodeSuccess = 0`
- `CompletionCodeWarning = 4`
- `CompletionCodeError = 8`
- `CompletionCodeCatastrophicError = 16`

You can filter stats by code using `SummaryStats.ByCode(code int)`.

### Supported Record IDs

The package defines numerous `RecordID` constants based on IBM Connect:Direct documentation. Examples include:

- `CopyTerminationRecord` ("CTRC"): Copy Termination Record (Process category)
- `ProcessStarted` ("PSTR"): Process started (Event category)
- `SessionStarted` ("SSTR"): Session started (Event category)
- `CheckpointingDisabled` ("XCPK"): Checkpointing was disabled for the copy step (Process category)

For a full list, refer to the source code or GoDoc.

## Data Structures

### SummaryStats
```go
type SummaryStats struct {
	Stats []SummaryStat
}
```
- `ByCode(code int) []SummaryStat`: Filters stats by completion code.

### SummaryStat
```go
type SummaryStat struct {
	ID            RecordID
	Date          time.Time
	Description   string
	ProcessNumber string
	Code          int
	MessageID     string
}
```

### RecordID
```go
type RecordID struct {
	ID          string
	Category    RecordCategoy
	Description string
}
```

### RecordCategoy
```go
type RecordCategoy string
```
- `CategoryEvent = "Event"`
- `CategoryProcess = "Process"`
- `CategoryExternalSource = "External Source"`

## Error Handling

The parser returns errors for invalid date formats, malformed lines, or integer parsing issues. Ensure input logs follow the expected format (e.g., lines with fields separated by spaces, starting after a hyphen row).

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for bugs, features, or improvements.

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

## License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Based on IBM Connect:Direct documentation (e.g., [Process Submission](https://www.ibm.com/docs/en/connect-direct/6.3.0?topic=processes-submitting-process)).
- Includes fixes referenced from APARs like IT41867.

For questions or support, open an issue on the GitHub repository.
