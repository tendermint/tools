package main

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	rpc "github.com/tendermint/tendermint/rpc/lib/server"
	"github.com/tendermint/tmlibs/log"
	monitor "github.com/tendermint/tools/tm-monitor/monitor"
)

const float64EqualityThreshold = 1e-9

func startRPC(listenAddr string, m *monitor.Monitor, logger log.Logger) {
	routes := routes(m)

	mux := http.NewServeMux()
	wm := rpc.NewWebsocketManager(routes, nil)
	mux.HandleFunc("/websocket", wm.WebsocketHandler)
	rpc.RegisterRPCFuncs(mux, routes, cdc, logger)
	if _, err := rpc.StartHTTPServer(listenAddr, mux, logger); err != nil {
		panic(err)
	}
}

func routes(m *monitor.Monitor) map[string]*rpc.RPCFunc {
	return map[string]*rpc.RPCFunc{
		"status":         rpc.NewRPCFunc(RPCStatus(m), ""),
		"status/network": rpc.NewRPCFunc(RPCNetworkStatus(m), ""),
		"status/node":    rpc.NewRPCFunc(RPCNodeStatus(m), "name"),
		"monitor":        rpc.NewRPCFunc(RPCMonitor(m), "endpoint"),
		"unmonitor":      rpc.NewRPCFunc(RPCUnmonitor(m), "endpoint"),

		// "start_meter": rpc.NewRPCFunc(network.StartMeter, "chainID,valID,event"),
		// "stop_meter":  rpc.NewRPCFunc(network.StopMeter, "chainID,valID,event"),
		// "meter":       rpc.NewRPCFunc(GetMeterResult(network), "chainID,valID,event"),
	}
}

// RPCStatus returns common statistics for the network and statistics per node.
func RPCStatus(m *monitor.Monitor) interface{} {
	return func() (NetworkAndNodes, error) {

		status := NewNetworkAndNodes()
		status.SetNetworkStatus(m)
		status.SetNodesStatus(m)

		return *status, nil
	}
}

// RPCNetworkStatus returns common statistics for the network.
func RPCNetworkStatus(m *monitor.Monitor) interface{} {
	return func() (*NetworkStatus, error) {

		status := NewNetworkAndNodes()
		status.SetNetworkStatus(m)

		return status.Network, nil
	}
}

// RPCNodeStatus returns statistics for the given node.
func RPCNodeStatus(m *monitor.Monitor) interface{} {
	return func(name string) (*NodeStatus, error) {

		if i, n := m.NodeByName(name); i != -1 {

			status := &NodeStatus{}
			status.IsValidator = n.IsValidator
			status.Name = n.Name
			status.Online = n.Online
			status.Height = n.Height
			status.BlockLatency = strconv.FormatFloat(n.BlockLatency, 'f', 6, 64)

			return status, nil
		}
		return nil, errors.New("Cannot find node with that name")
	}
}

// RPCMonitor allows to dynamically add a endpoint to under the monitor. Safe
// to call multiple times.
func RPCMonitor(m *monitor.Monitor) interface{} {
	return func(endpoint string) (*NodeStatus, error) {
		i, n := m.NodeByName(endpoint)

		status := &NodeStatus{}
		if i == -1 {
			n = monitor.NewNode(endpoint)
			if err := m.Monitor(n); err != nil {
				return nil, err
			}

			for math.Abs(n.BlockLatency) <= float64EqualityThreshold {
			}

			status.IsValidator = n.IsValidator
			status.Name = n.Name
			status.Online = n.Online
			status.Height = n.Height
			status.BlockLatency = strconv.FormatFloat(n.BlockLatency, 'f', 6, 64)

		}
		return status, nil
	}
}

// RPCUnmonitor removes the given endpoint from under the monitor.
func RPCUnmonitor(m *monitor.Monitor) interface{} {
	return func(endpoint string) (bool, error) {
		if i, n := m.NodeByName(endpoint); i != -1 {
			m.Unmonitor(n)
			return true, nil
		}
		return false, errors.New("Cannot find node with that name")
	}
}

// func (tn *TendermintNetwork) StartMeter(chainID, valID, eventID string) error {
// 	tn.mtx.Lock()
// 	defer tn.mtx.Unlock()
// 	val, err := tn.getChainVal(chainID, valID)
// 	if err != nil {
// 		return err
// 	}
// 	return val.EventMeter().Subscribe(eventID, nil)
// }

// func (tn *TendermintNetwork) StopMeter(chainID, valID, eventID string) error {
// 	tn.mtx.Lock()
// 	defer tn.mtx.Unlock()
// 	val, err := tn.getChainVal(chainID, valID)
// 	if err != nil {
// 		return err
// 	}
// 	return val.EventMeter().Unsubscribe(eventID)
// }

// func (tn *TendermintNetwork) GetMeter(chainID, valID, eventID string) (*eventmeter.EventMetric, error) {
// 	tn.mtx.Lock()
// 	defer tn.mtx.Unlock()
// 	val, err := tn.getChainVal(chainID, valID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return val.EventMeter().GetMetric(eventID)
// }

//--> types

type NetworkAndNodes struct {
	Network *NetworkStatus `json:"network_status"`
	Nodes   []*NodeStatus  `json:"nodes_status"`
}

func NewNetworkAndNodes() *NetworkAndNodes {

	status := &NetworkAndNodes{}

	status.Network = &NetworkStatus{}
	status.Nodes = []*NodeStatus{}

	return status
}

func (n *NetworkAndNodes) SetNetworkStatus(m *monitor.Monitor) {

	n.Network.Height = m.Network.Height

	n.Network.AvgBlockTime = strconv.FormatFloat(m.Network.AvgBlockTime, 'f', 6, 64)
	n.Network.AvgTxThroughput = strconv.FormatFloat(m.Network.AvgTxThroughput, 'f', 6, 64)
	n.Network.AvgBlockLatency = strconv.FormatFloat(m.Network.AvgBlockLatency, 'f', 6, 64)

	n.Network.NumValidators = m.Network.NumValidators
	n.Network.NumNodesMonitored = m.Network.NumNodesMonitored
	n.Network.NumNodesMonitoredOnline = m.Network.NumNodesMonitoredOnline

	n.Network.Health = m.Network.Health

	n.Network.UptimeData.StartTime = m.Network.UptimeData.StartTime
	n.Network.UptimeData.Uptime = strconv.FormatFloat(m.Network.UptimeData.Uptime, 'f', 6, 64)

}

func (n *NetworkAndNodes) SetNodesStatus(m *monitor.Monitor) {

	for i := range m.Nodes {
		node := NodeStatus{
			m.Nodes[i].IsValidator,
			m.Nodes[i].Name,
			m.Nodes[i].Online,
			m.Nodes[i].Height,
			strconv.FormatFloat(m.Nodes[i].BlockLatency, 'f', 6, 64),
		}
		n.Nodes = append(n.Nodes, &node)
	}

}

type UptimeDataStatus struct {
	StartTime time.Time `json:"start_time"`
	Uptime    string    `json:"uptime" wire:"unsafe"`
}

type NetworkStatus struct {
	Height int64 `json:"height"`

	AvgBlockTime    string `json:"avg_block_time"`
	AvgTxThroughput string `json:"avg_tx_throughput"`
	AvgBlockLatency string `json:"avg_block_latency"`

	NumValidators           int `json:"num_validators"`
	NumNodesMonitored       int `json:"num_nodes_monitored"`
	NumNodesMonitoredOnline int `json:"num_nodes_monitored_online"`

	Health monitor.Health `json:"health"`

	UptimeData UptimeDataStatus `json:"uptime_data"`
}

type NodeStatus struct {
	IsValidator bool `json:"is_validator"` // validator or non-validator?
	//pubKey      crypto.PubKey `json:"pub_key"`

	Name         string `json:"name"`
	Online       bool   `json:"online"`
	Height       int64  `json:"height"`
	BlockLatency string `json:"block_latency"` // ms, interval between block commits
}
