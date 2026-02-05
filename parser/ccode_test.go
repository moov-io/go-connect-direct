package parser_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/go-connect-direct/parser"

	"github.com/stretchr/testify/require"
)

func TestParseCCode(t *testing.T) {
	cases := []struct {
		inputFilepath string
		expected      []parser.SummaryStat
	}{
		{
			inputFilepath: filepath.Join("testdata", "ccode_stats.txt"),
			expected: []parser.SummaryStat{
				{
					Type:        "E",
					ID:          parser.SubmitProcess,
					Date:        time.Date(2026, time.February, 3, 23, 28, 45, 0, time.UTC),
					Description: "Submit command issued.",
				},
				{
					Type:          "P",
					ID:            parser.ProcessStarted,
					Date:          time.Date(2026, time.February, 3, 23, 28, 45, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG200I",
				},
				{
					Type:          "P",
					ID:            parser.ProcessStarted,
					Date:          time.Date(2026, time.February, 3, 23, 28, 45, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG200I",
				},
				{
					Type:          "P",
					ID:            parser.CheckpointingDisabled,
					Date:          time.Date(2026, time.February, 3, 23, 28, 46, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          4,
					MessageID:     "XCPK005W",
				},
				{
					Type:          "P",
					ID:            parser.CopyTerminationRecord,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "SCPA000I",
				},
				{
					Type:          "P",
					ID:            parser.CopyTerminationRecord,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "SCPA000I",
				},
				{
					Type:          "P",
					ID:            parser.ProcessEnded,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG252I",
				},
				{
					Type:          "P",
					ID:            parser.ProcessEnded,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG252I",
				},
			},
		},
		{
			inputFilepath: filepath.Join("testdata", "ccode_error.txt"),
			expected: []parser.SummaryStat{
				{Type: "P", ID: parser.RecordID{ID: "XIPT"}, Date: time.Date(2026, time.February, 5, 22, 45, 40, 0, time.UTC), Description: "SENDFILE", ProcessNumber: "21", Code: 8, MessageID: "XIPT004I"},
				{Type: "E", ID: parser.RecordID{ID: "RNCF"}, Date: time.Date(2026, time.February, 5, 22, 45, 40, 0, time.UTC), Description: "Attempt to connect to remote node frbpajcd02 failed"},
				{Type: "P", ID: parser.RecordID{ID: "XIPT"}, Date: time.Date(2026, time.February, 5, 22, 46, 10, 0, time.UTC), Description: "SENDFILE", ProcessNumber: "21", Code: 8, MessageID: "XIPT004I"},
				{Type: "E", ID: parser.RecordID{ID: "RNCF"}, Date: time.Date(2026, time.February, 5, 22, 46, 10, 0, time.UTC), Description: "Attempt to connect to remote node frbpajcd02 failed"},
				{Type: "P", ID: parser.RecordID{ID: "XIPT"}, Date: time.Date(2026, time.February, 5, 22, 46, 40, 0, time.UTC), Description: "SENDFILE", ProcessNumber: "21", Code: 8, MessageID: "XIPT004I"},
				{Type: "E", ID: parser.RecordID{ID: "RNCF"}, Date: time.Date(2026, time.February, 5, 22, 46, 40, 0, time.UTC), Description: "Attempt to connect to remote node frbpajcd02 failed"},
				{Type: "P", ID: parser.RecordID{ID: "XIPT"}, Date: time.Date(2026, time.February, 5, 22, 47, 10, 0, time.UTC), Description: "SENDFILE", ProcessNumber: "21", Code: 8, MessageID: "XIPT004I"},
				{Type: "E", ID: parser.RecordID{ID: "RNCF"}, Date: time.Date(2026, time.February, 5, 22, 47, 10, 0, time.UTC), Description: "Attempt to connect to remote node frbpajcd02 failed"},
				{Type: "E", ID: parser.SubmitProcess, Date: time.Date(2026, time.February, 5, 22, 50, 40, 0, time.UTC), Description: "Submit command issued."},
				{Type: "P", ID: parser.RecordID{ID: "XIPT"}, Date: time.Date(2026, time.February, 5, 22, 57, 10, 0, time.UTC), Description: "SENDFILE", ProcessNumber: "21", Code: 8, MessageID: "XIPT004I"},
				{Type: "E", ID: parser.RecordID{ID: "RNCF"}, Date: time.Date(2026, time.February, 5, 22, 57, 10, 0, time.UTC), Description: "Attempt to connect to remote node frbpajcd02 failed"},
				{Type: "P", ID: parser.RecordID{ID: "XIPT"}, Date: time.Date(2026, time.February, 5, 23, 7, 11, 0, time.UTC), Description: "SENDFILE", ProcessNumber: "21", Code: 8, MessageID: "XIPT004I"},
				{Type: "E", ID: parser.RecordID{ID: "RNCF"}, Date: time.Date(2026, time.February, 5, 23, 7, 11, 0, time.UTC), Description: "Attempt to connect to remote node frbpajcd02 failed"},
				{Type: "P", ID: parser.RecordID{ID: "XIPT"}, Date: time.Date(2026, time.February, 5, 23, 17, 11, 0, time.UTC), Description: "SENDFILE", ProcessNumber: "21", Code: 8, MessageID: "XIPT004I"},
				{Type: "E", ID: parser.RecordID{ID: "RNCF"}, Date: time.Date(2026, time.February, 5, 23, 17, 11, 0, time.UTC), Description: "Attempt to connect to remote node frbpajcd02 failed"},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.inputFilepath, func(t *testing.T) {
			bs, err := os.ReadFile(tc.inputFilepath)
			require.NoError(t, err)

			got, err := parser.ParseCCode(string(bs))
			require.NoError(t, err)
			require.Equal(t, len(tc.expected), len(got.Stats))

			for idx := range tc.expected {
				require.Equal(t, tc.expected[idx], got.Stats[idx], fmt.Sprintf("stats[%d]", idx))
			}
		})
	}
}

func TestParseCCode_ByCode(t *testing.T) {
	cases := []struct {
		inputFilepath string

		successful   int
		warnings     int
		errors       int
		catastrophic int
	}{
		{
			inputFilepath: filepath.Join("testdata", "ccode_stats.txt"),
			successful:    6,
			warnings:      1,
			errors:        0,
			catastrophic:  0,
		},
		{
			inputFilepath: filepath.Join("testdata", "ccode_error.txt"),
			successful:    0,
			warnings:      0,
			errors:        7,
			catastrophic:  0,
		},
	}
	for _, tc := range cases {
		t.Run(tc.inputFilepath, func(t *testing.T) {
			bs, err := os.ReadFile(tc.inputFilepath)
			require.NoError(t, err)

			got, err := parser.ParseCCode(string(bs))
			require.NoError(t, err)

			filtered := got.ByCodes(parser.CompletionCodeSuccess)
			require.Equal(t, tc.successful, len(filtered))

			filtered = got.ByCodes(parser.CompletionCodeWarning)
			require.Equal(t, tc.warnings, len(filtered))

			filtered = got.ByCodes(parser.CompletionCodeError)
			require.Equal(t, tc.errors, len(filtered))

			filtered = got.ByCodes(parser.CompletionCodeCatastrophicError)
			require.Equal(t, tc.catastrophic, len(filtered))

			// combine errors and catastrophic
			filtered = got.ByCodes(parser.CompletionCodeWarning, parser.CompletionCodeError, parser.CompletionCodeCatastrophicError)
			require.Equal(t, tc.warnings+tc.errors+tc.catastrophic, len(filtered))
		})
	}
}
