package epochdate

import (
	"testing"
	"time"
)

func TestFromSeconds(t *testing.T) {
	testCases := map[string]struct {
		Input string
		Want  int
	}{
		"1970-01-01": {
			Input: "1970-01-01",
			Want:  0,
		},
		"1971-01-01": {
			Input: "1971-01-01",
			Want:  365,
		},
		"1972-01-01": {
			Input: "1972-01-01",
			Want:  730,
		},
		"1973-01-01": {
			Input: "1973-01-01",
			Want:  1096,
		},
		"2020-02-02": {
			Input: "2020-02-02",
			Want:  18294,
		},
		"2000-02-02": {
			Input: "2000-02-02",
			Want:  10989,
		},
		"2001-02-02": {
			Input: "2001-02-02",
			Want:  11355,
		},
		"2100-02-02": {
			Input: "2100-02-02",
			Want:  47514,
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			when, err := time.ParseInLocation("2006-01-02", tc.Input, time.UTC)
			if err != nil {
				t.Fatalf("got %v; want nil", err)
			}

			n := FromSeconds(when.Unix())
			if got, want := n, tc.Want; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}

			v := Time(n)
			if got, want := v.Format("2006-01-02"), tc.Input; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}

func TestZero(t *testing.T) {
	if got, want := FromSeconds(0), 0; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
