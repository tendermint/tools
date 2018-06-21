package main

import (
	amino "github.com/tendermint/go-amino"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
	cdc.RegisterConcrete(&NetworkAndNodes{}, "monitor/NetworkAndNodes", nil)
	cdc.RegisterConcrete(&NetworkStatus{}, "monitor/NetworkStatus", nil)
	cdc.RegisterConcrete(&NodeStatus{}, "monitor/NodeStatus", nil)
}
