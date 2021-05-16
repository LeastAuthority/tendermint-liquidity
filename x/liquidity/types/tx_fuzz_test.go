//+build gofuzz

package types

import (
	"testing"

	fleece "github.com/leastauthority/fleece/fuzzing"
	"github.com/stretchr/testify/require"
)

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

func TestFuzzMsgWithdrawWithinBatch_raw(t *testing.T) {
	_, panics, _ := fleece.
		MustNewCrasherIterator(env, FuzzMsgWithdrawWithinBatch_raw, filters...).
		TestFailingLimit(t, crashLimit)

	require.Zero(t, panics)
}

func TestFuzzMsgWithdrawWithinBatch_structured(t *testing.T) {
	_, panics, _ := fleece.
		MustNewCrasherIterator(env, FuzzMsgWithdrawWithinBatch_structured, filters...).
		TestFailingLimit(t, crashLimit)

	require.Zero(t, panics)
}
