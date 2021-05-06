// +build gofuzz

package types

import (
	"bytes"
	"errors"
	gofuzz "github.com/google/gofuzz"
	fleece "github.com/leastauthority/fleece/fuzzing"
	"reflect"
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
		panic(errors.New("deserialized messages didn't' match"))
	}

	msg2Data, err := msg1.Marshal()
	if err != nil {
		panic(err)
	}

	if !bytes.Equal(msg1Data, msg2Data) {
		panic(errors.New("serialized messages didn't match"))
	}

	return fleece.FuzzInteresting
}

func FuzzMsgSwapWithinBatch_structured(data []byte) int {
	msg1 := MsgSwapWithinBatch{}

	if len(data) == 0 {
		return fleece.FuzzDiscard
	}

	f := gofuzz.NewFromGoFuzz(data)
	f.NilChance(0)
	f.NumElements(1, 100)
	f.Fuzz(&msg1)

	msg1Data, err := msg1.Marshal()
	if err != nil {
		panic(err)
	}

	msg2 := MsgSwapWithinBatch{}
	if err := msg2.Unmarshal(msg1Data); err != nil {
		panic(err)
	}

	msg2Data, err := msg2.Marshal()
	if err != nil {
		panic(err)
	}

	if bytes.Equal(msg1Data, msg2Data) {
		panic(errors.New("serialized messages don't match"))
	}

	return fleece.FuzzNormal
}