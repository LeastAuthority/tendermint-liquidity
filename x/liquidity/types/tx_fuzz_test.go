//+build gofuzz

package types

import (
	"flag"
	"os"
	"testing"

	fleece "github.com/leastauthority/fleece/fuzzing"
	"github.com/stretchr/testify/require"
)

var (
	crashLimit           int
	fleeceDir            string
	skipPattern          string
	skipPatternDelimiter string
	safe, verbose        bool

	env     *fleece.Env
	filters []fleece.IterFilter
)

func init() {
	flag.IntVar(&crashLimit, "crash-limit", 1000, "number of crashing inputs to test before stopping")
	flag.StringVar(&fleeceDir, "fleece-dir", "fleece", "path to fleece dir relative to repo/module root")
	flag.StringVar(&skipPattern, "skip", "", "if provided, crashers with recorded outputs which match the pattern will be skipped")
	flag.StringVar(&skipPatternDelimiter, "skip-delimiter", "", "delimiter used to split skip pattern")
	flag.BoolVar(&safe, "safe", true, "\"if true, skips crashers with recorded outputs that timed-out or ran out of memory\"")
	flag.BoolVar(&verbose, "verbose", false, "if true, logs each skip")
}

func TestMain(m *testing.M) {
	flag.Parse()
	env = fleece.NewEnv(fleeceDir)

	skipFilter := fleece.SkipFilter(skipPattern, skipPatternDelimiter, verbose)
	filters = []fleece.IterFilter{skipFilter}
	if safe {
		filters = append(filters,
			fleece.SkipTimedOut(verbose),
			fleece.SkipOutOfMemory(verbose))
	}

	os.Exit(m.Run())
}

func TestFuzzMsgSwapWithinBatch_raw(t *testing.T) {
	_, panics, _ := fleece.
		MustNewCrasherIterator(env, FuzzMsgSwapWithinBatch_raw, filters...).
		TestFailingLimit(t, crashLimit)

	require.Zero(t, panics)
}

func TestFuzzMsgSwapWithinBatch_structured(t *testing.T) {
	_, panics, _ := fleece.
		MustNewCrasherIterator(env, FuzzMsgSwapWithinBatch_structured, filters...).
		TestFailingLimit(t, crashLimit)

	require.Zero(t, panics)
}
