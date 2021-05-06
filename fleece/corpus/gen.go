package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"path/filepath"

	gofuzz "github.com/google/gofuzz"
	"github.com/spf13/cobra"

	"github.com/tendermint/liquidity/x/liquidity/types"
)

var (
	cmdRoot = &cobra.Command{
		//Use: "gen <fuzz func name>",
		//Args: cobra.ExactArgs(1),
		RunE: root,
	}

	// TODO: convert to flags / args
	count = 100
	outputDir = "/home/bwhite/Projects/liquidity/fleece/corpus/test"
)

func main() {
	if err := cmdRoot.Execute(); err != nil {
		log.Fatal(err)
	}
}

func root(cmd *cobra.Command, args []string) error {
	for i := 0; i < count; i++ {
		msgData := generate()
		if err := write(msgData); err != nil {
		}
	}
	return nil
}

func generate() []byte {
	msg := types.MsgSwapWithinBatch{}
	f := gofuzz.New()
	f.NilChance(0)
	f.NumElements(1, 100)
	f.Fuzz(&msg)

	msgData, err := msg.Marshal()
	if err != nil {
		return generate()
	}
	return msgData
}

func write(data []byte) error {
	hash := sha256.Sum256(data)
	hexHash := hex.EncodeToString(hash[:])
	outputPath := filepath.Join(outputDir, hexHash + ".fleece")
	err := ioutil.WriteFile(outputPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
