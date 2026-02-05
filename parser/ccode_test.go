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
					ID:          parser.SubmitProcess,
					Date:        time.Date(2026, time.February, 3, 23, 28, 45, 0, time.UTC),
					Description: "Submit command issued.",
				},
				{
					ID:            parser.ProcessStarted,
					Date:          time.Date(2026, time.February, 3, 23, 28, 45, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG200I",
				},
				{
					ID:            parser.ProcessStarted,
					Date:          time.Date(2026, time.February, 3, 23, 28, 45, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG200I",
				},
				// TODO(adam): missing XCPK
				{
					ID:            parser.CopyTerminationRecord,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "SCPA000I",
				},
				{
					ID:            parser.CopyTerminationRecord,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "SCPA000I",
				},
				{
					ID:            parser.ProcessEnded,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG252I",
				},
				{
					ID:            parser.ProcessEnded,
					Date:          time.Date(2026, time.February, 3, 23, 28, 52, 0, time.UTC),
					Description:   "sample",
					ProcessNumber: "14",
					Code:          0,
					MessageID:     "XSMG252I",
				},
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
