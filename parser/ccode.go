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

func (ss SummaryStats) ByCodes(codes ...int) []SummaryStat {
	var out []SummaryStat
	for _, stat := range ss.Stats {
		if stat.Type != "P" {
			continue
		}

		for _, c := range codes {
			if stat.Code == c {
				out = append(out, stat)
			}
		}
	}
	return out
}

type SummaryStat struct {
	// Type is the record type
	//
	// P: Process records (related to process activities, e.g., CAPR category)
	// E: Event records (related to system events, e.g., CAEV category). Note that these are not necessarily errors—many are informational or lifecycle events (e.g., process started or session ended).
	// X: External records (related to external component statistics, e.g., CAEX category, such as stats from an Integrated File Agent).
	Type string

	ID            RecordID
	Date          time.Time
	Description   string
	ProcessNumber string
	Code          int
	MessageID     string
}

var (
	// CompletionCodeSuccess is defined to be a "numeric code returned from a
	// completed Process that indicates failure or success."
	//
	// The following are the standard return codes:
	//  - 0 indicates successful completion
	//  - 4 indicates a warning
	//  - 8 indicates an error
	//  - 16 indicates a catastrophic error
	//
	// Source: https://public.dhe.ibm.com/software/commerce/doc/mft/scc/53/ReportsGuide.pdf
	CompletionCodeSuccess           = 0
	CompletionCodeWarning           = 4
	CompletionCodeError             = 8
	CompletionCodeCatastrophicError = 16
)

// ParseCCode parses the summary output from an IBM Connect:Direct "select statistics" command into structured SummaryStats.
//
// It processes the input string, identifying lines after a hyphen separator (------) and before an equals separator (======),
// parsing each valid line into a SummaryStat. Lines are expected in formats like:
//
//	E SUBP  02/03/2026 23:28:45 Submit command issued.
//	P PSTR  02/03/2026 23:28:45 sample            14                0      XSMG200I
//
// The function handles special cases for submit processes (SUBP) and general records using LookupRecordID.
// Dates are parsed in the format "01/02/2006 15:04:05".
//
// Statistics can be viewed in Connect:Direct with commands like:
//
//	sel stat ccode(ge,0) pnumber=18;
//
// If the input is malformed (e.g., invalid date, insufficient columns, or unparseable codes), an error is returned.
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
				Type:        cols[0],
				ID:          SubmitProcess,
				Description: strings.Join(cols[4:], " "),
			}
			rec.Date, err = parseSummaryDate(cols[2:4])
			if err != nil {
				return out, fmt.Errorf("parsing %s date: %v", SubmitProcess.ID, err)
			}
			out.Stats = append(out.Stats, rec)

		default:
			// Parse a line which looks like:
			// P RECID LOG TIME            PNAME        PNUMBER  STEPNAME   CCOD FDBK MSGID
			// E RECID LOG TIME            MESSAGE TEXT
			// X RECID LOG TIME            APP DESC     USID     NODENAME   CCOD MSGID

			switch cols[0] {
			case "P": // process
				rec, err := parseSummaryProcessRecord(cols)
				if err != nil {
					return out, fmt.Errorf("parsing process record: %v", err)
				}
				if rec != nil {
					out.Stats = append(out.Stats, *rec)
				}

			case "E": // error
				rec, err := parseSummaryErrorRecord(cols)
				if err != nil {
					return out, fmt.Errorf("parsing error record: %v", err)
				}
				if rec != nil {
					out.Stats = append(out.Stats, *rec)
				}

			case "X": // xtra records
				rec, err := parseSummaryExtraRecord(cols)
				if err != nil {
					return out, fmt.Errorf("parsing extra record: %v", err)
				}
				if rec != nil {
					out.Stats = append(out.Stats, *rec)
				}

			}

		}
	}

	return out, nil
}

func parseSummaryProcessRecord(cols []string) (*SummaryStat, error) {
	// example records
	//
	//   P PSTR  02/03/2026 23:28:45 sample            14                0      XSMG200I
	//
	//   P XIPT  02/05/2026 22:45:40 SENDFILE          21                8      XIPT004I

	rec := &SummaryStat{}

	if len(cols) < 8 {
		return nil, nil
	}
	rec.Type = cols[0]

	ccode := LookupRecordID(cols[1])
	if ccode == nil {
		ccode = &RecordID{
			ID: strings.ToUpper(cols[1]),
		}
	}
	rec.ID = *ccode

	date, err := parseSummaryDate(cols[2:4])
	if err != nil {
		return rec, fmt.Errorf("parsing %s date: %v", rec.ID, err)
	}
	rec.Date = date

	// Start adding columns right to left
	idx := len(cols) - 1
	if len(cols) > idx {
		rec.MessageID = cols[idx]
	}
	idx--

	if len(cols) > idx {
		cc, err := strconv.ParseInt(cols[idx], 10, 16)
		if err != nil {
			return rec, err
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

	return rec, nil
}

func parseSummaryErrorRecord(cols []string) (*SummaryStat, error) {
	// example records
	//
	//    E RNCF  02/05/2026 22:45:40 Attempt to connect to remote node frbpajcd02 failed

	rec := &SummaryStat{}

	if len(cols) < 4 {
		return nil, nil
	}
	rec.Type = cols[0]

	ccode := LookupRecordID(cols[1])
	if ccode == nil {
		ccode = &RecordID{
			ID: strings.ToUpper(cols[1]),
		}
	}
	rec.ID = *ccode

	date, err := parseSummaryDate(cols[2:4])
	if err != nil {
		return rec, fmt.Errorf("parsing %s date: %v", rec.ID, err)
	}
	rec.Date = date

	if len(cols) > 4 {
		rec.Description = strings.Join(cols[4:], " ")
	}

	return rec, nil
}

func parseSummaryExtraRecord(cols []string) (*SummaryStat, error) {
	return nil, nil // TODO(adam):
}

func parseSummaryDate(fields []string) (time.Time, error) {
	return time.Parse("01/02/2006 15:04:05", strings.Join(fields, " "))
}
