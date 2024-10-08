package report_test

import (
	"bytes"
	"testing"

	"github.com/ayoisaiah/f2/internal/config"
	"github.com/ayoisaiah/f2/internal/file"
	"github.com/ayoisaiah/f2/internal/status"
	"github.com/ayoisaiah/f2/internal/testutil"
	"github.com/ayoisaiah/f2/report"
)

func TestReport(t *testing.T) {
	testCases := []testutil.TestCase{
		{
			Name: "report unchanged file names",
			Changes: file.Changes{
				{
					SourcePath: "macos_update_notes_2023.txt",
					TargetPath: "macos_update_notes_2023.txt",
					Status:     status.Unchanged,
				},
				{
					SourcePath: "macos_user_guide_macos_sierra.pdf",
					TargetPath: "macos_user_guide_macos_sierra.pdf",
					Status:     status.Unchanged,
				},
			},
			Args: []string{"-r"},
		},
	}

	reportTest(t, testCases)
}

func reportTest(t *testing.T, cases []testutil.TestCase) {
	t.Helper()

	for i := range cases {
		tc := cases[i]

		for i := range tc.Changes {
			tc.Changes[i].Position = i
		}

		t.Run(tc.Name, func(t *testing.T) {
			if tc.SetupFunc != nil {
				t.Cleanup(tc.SetupFunc(t, ""))
			}

			conf := testutil.GetConfig(t, &tc, ".")

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			config.Stdout = &stdout
			config.Stderr = &stderr

			report.Report(conf, tc.Changes, false)

			testutil.CompareGoldenFile(t, &tc, stdout.Bytes())
		})
	}
}
