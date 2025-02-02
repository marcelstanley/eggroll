package honeypot

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gligneul/eggroll"
	"github.com/holiman/uint256"
)

// Owner of the honeypot that can withdraw all funds.
var Owner common.Address

func init() {
	Owner = common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
}

type Withdraw struct {
	Value *uint256.Int
}

type Honeypot struct {
	Balance *uint256.Int
}

func Codecs() []eggroll.Codec {
	return []eggroll.Codec{
		eggroll.NewJSONCodec[Withdraw](),
		eggroll.NewJSONCodec[Honeypot](),
	}
}
