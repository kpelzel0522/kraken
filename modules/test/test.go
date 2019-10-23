/* pxe.go: provides generic PXE/iPXE-boot capabilities
 *           this manages both DHCP and TFTP/HTTP services.
 *			 If <file> doesn't exist, but <file>.tpl does, tftp will fill it as as template.
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2018, Triad National Security, LLC
 * See LICENSE file for details.
 */

//go:generate protoc -I ../../core/proto/include -I proto --go_out=plugins=grpc:proto proto/pxe.proto

package test

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/proto"
	"github.com/hpc/kraken/core"
	cpb "github.com/hpc/kraken/core/proto"
	testpb "github.com/hpc/kraken/extensions/Test/proto"
	"github.com/hpc/kraken/lib"
	pb "github.com/hpc/kraken/modules/test/proto"
)

const (
	TestStateURL = "type.googleapis.com/proto.Test/State"
	SrvStateURL  = "/Services/test/State"
)

var _ lib.Module = (*Test)(nil)
var _ lib.ModuleWithConfig = (*Test)(nil)
var _ lib.ModuleWithDiscovery = (*Test)(nil)
var _ lib.ModuleSelfService = (*Test)(nil)

// modify these if you want different requires for mutations
var reqs = map[string]reflect.Value{
	"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
}

// modify this if you want excludes
var excs = map[string]reflect.Value{}

/* we use channels and a node manager rather than locking
   to make our node store safe.  This is a simpple query
   language for that service */

type nodeQueryBy string

const (
	queryByIP  nodeQueryBy = "IP"
	queryByMAC nodeQueryBy = "MAC"
)

// PXE provides PXE-boot capabilities
type Test struct {
	api   lib.APIClient
	cfg   *pb.TestConfig
	dchan chan<- lib.Event

	pollTicker *time.Ticker
}

// Name returns the FQDN of the module
func (*Test) Name() string { return "github.com/hpc/kraken/modules/test" }

// NewConfig returns a fully initialized default config
func (*Test) NewConfig() proto.Message {
	r := &pb.TestConfig{
		IpUrl: "type.googleapis.com/proto.IPv4OverEthernet/Ifaces/0/Ip/Ip",
		Servers: map[string]*pb.TestServer{
			"testServer": {
				Name: "testServer",
				Ip:   "localhost",
				Port: 8269,
			},
		},
	}
	return r
}

// UpdateConfig updates the running config
func (t *Test) UpdateConfig(cfg proto.Message) (e error) {
	if tcfg, ok := cfg.(*pb.TestConfig); ok {
		t.cfg = tcfg
		return
	}
	return fmt.Errorf("invalid config type")
}

// ConfigURL gives the any resolver URL for the config
func (*Test) ConfigURL() string {
	cfg := &pb.TestConfig{}
	any, _ := ptypes.MarshalAny(cfg)
	return any.GetTypeUrl()
}

// SetDiscoveryChan sets the current discovery channel
// this is generally done by the API
func (t *Test) SetDiscoveryChan(c chan<- lib.Event) { t.dchan = c }

// Entry is the module's executable entrypoint
func (t *Test) Entry() {
	url := lib.NodeURLJoin(t.api.Self().String(), SrvStateURL)
	ev := core.NewEvent(
		lib.Event_DISCOVERY,
		url,
		&core.DiscoveryEvent{
			Module:  t.Name(),
			URL:     url,
			ValueID: "RUN",
		},
	)
	t.dchan <- ev

	// setup a ticker for polling discovery
	dur, _ := time.ParseDuration("10s")
	t.pollTicker = time.NewTicker(dur)

	// main loop
	for {
		select {
		case <-t.pollTicker.C:
			go t.discoverAll()
			break
		}
	}
}

// discoverAll is used to do polling discovery of power state
// Note: this is probably not extremely efficient for large systems
func (t *Test) discoverAll() {
	t.api.Log(lib.LLDEBUG, "polling for node state")
	ns, e := t.api.QueryReadAll()
	if e != nil {
		t.api.Logf(lib.LLERROR, "polling node query failed: %v", e)
		return
	}
	ipmap := make(map[string]lib.NodeID)

	// get ip addresses for nodes
	for _, n := range ns {
		v, e := n.GetValue(t.cfg.GetIpUrl())
		if e != nil {
			t.api.Logf(lib.LLERROR, "problem getting ip address of nodes")
		}
		ip := v.String()
		ipmap[ip] = n.ID()
	}

	t.api.Logf(lib.LLDEBUG, "got ip addresses: %v", ipmap)
	for _, n := range ns {
		t.fakeDiscover(n)
	}
}

func (t *Test) fakeDiscover(node lib.Node) {

	var vid testpb.Test_Test
	vid = testpb.Test_HIGH

	url := lib.NodeURLJoin(node.ID().String(), "type.googleapis.com/proto.Test/State")
	v := core.NewEvent(
		lib.Event_DISCOVERY,
		url,
		&core.DiscoveryEvent{
			Module:  t.Name(),
			URL:     url,
			ValueID: vid.String(),
		},
	)
	t.dchan <- v
}

// Init is used to intialize an executable module prior to entrypoint
func (t *Test) Init(api lib.APIClient) {
	t.api = api
	t.cfg = t.NewConfig().(*pb.TestConfig)
}

// Stop should perform a graceful exit
func (t *Test) Stop() {
	os.Exit(0)
}

func init() {
	module := &Test{}
	discovers := make(map[string]map[string]reflect.Value)
	dtest := make(map[string]reflect.Value)

	dtest[testpb.Test_NONE.String()] = reflect.ValueOf(testpb.Test_NONE)
	dtest[testpb.Test_LOW.String()] = reflect.ValueOf(testpb.Test_LOW)
	dtest[testpb.Test_MED.String()] = reflect.ValueOf(testpb.Test_MED)
	dtest[testpb.Test_HIGH.String()] = reflect.ValueOf(testpb.Test_HIGH)

	discovers["type.googleapis.com/proto.Test/State"] = dtest

	discovers[SrvStateURL] = map[string]reflect.Value{
		"RUN": reflect.ValueOf(cpb.ServiceInstance_RUN)}
	si := core.NewServiceInstance("test", module.Name(), module.Entry, nil)
	// Register it all
	core.Registry.RegisterModule(module)
	core.Registry.RegisterServiceInstance(module, map[string]lib.ServiceInstance{si.ID(): si})
	core.Registry.RegisterDiscoverable(module, discovers)
}
