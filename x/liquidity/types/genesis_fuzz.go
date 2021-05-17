// +build gofuzz

package types

import (
	"bytes"
	"errors"
	"reflect"

	gofuzz "github.com/google/gofuzz"
	fleece "github.com/leastauthority/fleece/fuzzing"
)

func FuzzGenesisState_raw(data []byte) int {
	gen1 := GenesisState{}

	if len(data) == 0 {
		return fleece.FuzzDiscard
	}

	if err := gen1.Unmarshal(data); err != nil {
		return fleece.FuzzNormal
	}

	gen1Data, err := gen1.Marshal()
	if err != nil {
		panic(err)
	}

	gen2 := GenesisState{}
	if err := gen2.Unmarshal(gen1Data); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(gen1, gen2) {
		// TODO: only run during triage
		//diff, _ := messagediff.PrettyDiff(gen1, gen2)
		//fmt.Println(diff)

		panic(errors.New("deserialized states didn't' match"))
	}

	return fleece.FuzzInteresting
}

func FuzzGenesisState_structured(data []byte) int {
	gen1 := GenesisState{}

	if len(data) == 0 || bytes.Equal(data, []byte{0}) {
		return fleece.FuzzDiscard
	}

	f := gofuzz.NewFromGoFuzz(data)
	f.NilChance(0)
	f.NumElements(1, 10)
	f.Fuzz(&gen1)

	gen1Data, err := gen1.Marshal()
	if err != nil {
		panic(err)
	}

	gen2 := MsgSwapWithinBatch{}
	if err := gen2.Unmarshal(gen1Data); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(gen1, gen2) {
		// TODO: only run during triage
		//diff, _ := messagediff.PrettyDiff(gen1, gen2)
		//fmt.Println(diff)

		panic(errors.New("deserialized states didn't match"))
	}

	return fleece.FuzzNormal
}
