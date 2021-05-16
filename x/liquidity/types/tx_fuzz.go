// +build gofuzz

package types

import (
	"bytes"
	"errors"
	"reflect"

	gofuzz "github.com/google/gofuzz"
	fleece "github.com/leastauthority/fleece/fuzzing"
)

func FuzzMsgSwapWithinBatch_raw(data []byte) int {
	msg1 := MsgSwapWithinBatch{}
	if err := msg1.Unmarshal(data); err != nil {
		return fleece.FuzzNormal
	}

	msg1Data, err := msg1.Marshal()
	if err != nil {
		panic(err)
	}

	msg2 := MsgSwapWithinBatch{}
	if err := msg2.Unmarshal(msg1Data); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(msg1, msg2) {
		// TODO: only run during triage
		//diff, _ := messagediff.PrettyDiff(msg1, msg2)
		//fmt.Println(diff)

		panic(errors.New("deserialized messages didn't' match"))
	}

	return fleece.FuzzInteresting
}

func FuzzMsgSwapWithinBatch_structured(data []byte) int {
	msg1 := MsgSwapWithinBatch{}

	if len(data) == 0 || bytes.Equal(data, []byte{0}) {
		return fleece.FuzzDiscard
	}

	f := gofuzz.NewFromGoFuzz(data)
	f.NilChance(0)
	f.NumElements(1, 10)
	f.Fuzz(&msg1)

	msg1Data, err := msg1.Marshal()
	if err != nil {
		panic(err)
	}

	msg2 := MsgSwapWithinBatch{}
	if err := msg2.Unmarshal(msg1Data); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(msg1, msg2) {
		// TODO: only run during triage
		//diff, _ := messagediff.PrettyDiff(msg1, msg2)
		//fmt.Println(diff)

		panic(errors.New("deserialized messages didn't match"))
	}

	return fleece.FuzzNormal
}
func FuzzMsgWithdrawWithinBatch_raw(data []byte) int {
	msg1 := MsgWithdrawWithinBatch{}
	if err := msg1.Unmarshal(data); err != nil {
		return fleece.FuzzNormal
	}

	msg1Data, err := msg1.Marshal()
	if err != nil {
		panic(err)
	}

	msg2 := MsgWithdrawWithinBatch{}
	if err := msg2.Unmarshal(msg1Data); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(msg1, msg2) {
		// TODO: only run during triage
		//diff, _ := messagediff.PrettyDiff(msg1, msg2)
		//fmt.Println(diff)

		panic(errors.New("deserialized messages didn't' match"))
	}

	return fleece.FuzzInteresting
}

func FuzzMsgWithdrawWithinBatch_structured(data []byte) int {
	msg1 := MsgWithdrawWithinBatch{}

	if len(data) == 0 || bytes.Equal(data, []byte{0}) {
		return fleece.FuzzDiscard
	}

	f := gofuzz.NewFromGoFuzz(data)
	f.NilChance(0)
	f.NumElements(1, 10)
	f.Fuzz(&msg1)

	msg1Data, err := msg1.Marshal()
	if err != nil {
		panic(err)
	}

	msg2 := MsgWithdrawWithinBatch{}
	if err := msg2.Unmarshal(msg1Data); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(msg1, msg2) {
		panic(errors.New("deserialized messages didn't match"))
	}

	return fleece.FuzzNormal
}
