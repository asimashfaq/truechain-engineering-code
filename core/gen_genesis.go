// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package core

import (
	"encoding/json"

	"github.com/truechain/truechain-engineering-code/core/fastchain"
	"github.com/truechain/truechain-engineering-code/core/snailchain"
)

//var _ = (*genesisSpecMarshaling)(nil)

func (g Genesis) MarshalJSON() ([]byte, error) {

	var enc Genesis
	enc.Snail = g.Snail
	enc.Fast = g.Fast
	return json.Marshal(&enc)
}

func (g *Genesis) UnmarshalJSON(input []byte) error {

	var snaildec snailchain.Genesis
	var fastdec fastchain.Genesis
	if err := json.Unmarshal(input, &snaildec); err != nil {
		return err
	}
	if err := json.Unmarshal(input, &fastdec); err != nil {
		return err
	}

	if &snaildec != nil {
		g.Snail = &snaildec
	}
	if &fastdec != nil {
		g.Fast = &fastdec
	}
	return nil
}
