package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type SummaryStats struct {
	Stats []SummaryStat
}

type SummaryStat struct {
	ID            RecordID
	Date          time.Time
	Description   string
	ProcessNumber string
	Code          int
	MessageID     string
}

func ParseCCode(input string) (SummaryStats, error) {
	var out SummaryStats

	var err error
	lines := strings.Split(input, "\n")

	// Find the row with a bunch of hyphens
	var shouldParseLine bool
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "------") {
			shouldParseLine = true
		}
		// Stop when we see a line of equals again
		if shouldParseLine && strings.Contains(line, "======") {
			break
		}

		// Skip lines when we aren't actively parsing
		if !shouldParseLine {
			continue
		}

		// Parse the line
		// We've seen lines like the following:
		//   E SUBP  02/03/2026 23:28:45 Submit command issued.
		//   P PSTR  02/03/2026 23:28:45 sample            14                0      XSMG200I

		cols := strings.Fields(line)
		if len(cols) < 2 {
			continue // invalid line
		}

		switch strings.ToUpper(cols[1]) {
		case SubmitProcess.ID:
			if len(cols) < 4 {
				continue
			}

			// Parse a submit process line
			rec := SummaryStat{
				ID:          SubmitProcess,
				Description: strings.Join(cols[4:], " "),
			}
			rec.Date, err = parseSummaryDate(cols[2:4])
			if err != nil {
				return out, fmt.Errorf("parsing %s date: %v", SubmitProcess.ID, err)
			}
			out.Stats = append(out.Stats, rec)

		default:
			// Parse a line which looks like
			//   P PSTR  02/03/2026 23:28:45 sample            14                0      XSMG200I
			if len(cols) < 8 {
				continue
			}

			// Parse an update line
			ccode := LookupRecordID(cols[1])
			if ccode == nil {
				continue
			}

			rec := SummaryStat{
				ID: *ccode,
			}
			rec.Date, err = parseSummaryDate(cols[2:4])
			if err != nil {
				return out, fmt.Errorf("parsing %s date: %v", rec.ID, err)
			}
			// Start adding columns right to left
			idx := len(cols) - 1
			if len(cols) > idx {
				rec.MessageID = cols[idx]
			}
			idx--
			if len(cols) > idx {
				cc, err := strconv.ParseInt(cols[idx], 10, 16)
				if err != nil {
					return out, err
				}
				rec.Code = int(cc)
			}
			idx--

			if len(cols) > 8 {
				idx-- // skip NODENAME
			}

			if len(cols) > idx {
				rec.ProcessNumber = cols[idx]
			}

			if len(cols) > idx {
				rec.Description = strings.Join(cols[4:idx], " ")
			}

			out.Stats = append(out.Stats, rec)
		}
	}

	return out, nil
}

func parseSummaryDate(fields []string) (time.Time, error) {
	return time.Parse("01/02/2006 15:04:05", strings.Join(fields, " "))
}
