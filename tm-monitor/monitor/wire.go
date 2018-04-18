package monitor

import (
	amino "github.com/tendermint/go-amino"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	//"github.com/tendermint/go-crypto"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
	//cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	//cdc.RegisterConcrete(crypto.PubKeyEd25519{},
	//	"tendermint/PubKeyEd25519", nil)
	//cdc.RegisterConcrete(crypto.PubKeySecp256k1{},
	//	"tendermint/PubKeySecp256k1", nil)
}
