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

package testmutate

import (
	"os"
	"reflect"
	"time"

	"github.com/hpc/kraken/core"
	cpb "github.com/hpc/kraken/core/proto"
	tpb "github.com/hpc/kraken/extensions/Test/proto"
	tspb "github.com/hpc/kraken/extensions/TestScaling/proto"
	"github.com/hpc/kraken/lib"
)

const (
	ModuleStateURL  = "/Services/testmutate/State"
	ScalingStateURL = "type.googleapis.com/proto.TestScaling/State"
	TempStateURL    = "type.googleapis.com/proto.Test/State"
)

var _ lib.Module = (*TestMutate)(nil)
var _ lib.ModuleWithDiscovery = (*TestMutate)(nil)
var _ lib.ModuleWithMutations = (*TestMutate)(nil)
var _ lib.ModuleSelfService = (*TestMutate)(nil)

// PXE provides PXE-boot capabilities
type TestMutate struct {
	api   lib.APIClient
	dchan chan<- lib.Event
	mchan <-chan lib.Event
}

type tmut struct {
	f       tspb.TestScaling_Scaling
	t       tspb.TestScaling_Scaling
	reqs    map[string]reflect.Value
	timeout string
}

var muts = map[string]tmut{
	"HIGHtoLOW": {
		f: tspb.TestScaling_HIGH,
		t: tspb.TestScaling_LOW,
		reqs: map[string]reflect.Value{
			"/PhysState":    reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":     reflect.ValueOf(cpb.Node_SYNC),
			TempStateURL:    reflect.ValueOf(tpb.Test_HIGH),
			ScalingStateURL: reflect.ValueOf(tspb.TestScaling_HIGH),
		},
		timeout: "60s",
	},
	"LOWtoHIGH": {
		f: tspb.TestScaling_LOW,
		t: tspb.TestScaling_HIGH,
		reqs: map[string]reflect.Value{
			"/PhysState":    reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":     reflect.ValueOf(cpb.Node_SYNC),
			TempStateURL:    reflect.ValueOf(tpb.Test_LOW),
			ScalingStateURL: reflect.ValueOf(tspb.TestScaling_LOW),
		},
		timeout: "60s",
	},
}

var excs = map[string]reflect.Value{}

// Name returns the FQDN of the module
func (*TestMutate) Name() string { return "github.com/hpc/kraken/modules/testmutate" }

// SetDiscoveryChan sets the current discovery channel
// this is generally done by the API
func (t *TestMutate) SetDiscoveryChan(c chan<- lib.Event) { t.dchan = c }

// SetMutationChan sets the current mutation channel
// this is generally done by the API
func (t *TestMutate) SetMutationChan(c <-chan lib.Event) { t.mchan = c }

// Entry is the module's executable entrypoint
func (t *TestMutate) Entry() {
	url := lib.NodeURLJoin(t.api.Self().String(), ModuleStateURL)
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

	// main loop
	for {
		select {
		case m := <-t.mchan: // mutation request
			go t.handleMutation(m)
			break
		}
	}
}

func (t *TestMutate) handleMutation(m lib.Event) {
	t.api.Logf(lib.LLDEBUG, "Got muationed: %+v", m)
}

// Init is used to intialize an executable module prior to entrypoint
func (t *TestMutate) Init(api lib.APIClient) {
	t.api = api
}

// Stop should perform a graceful exit
func (t *TestMutate) Stop() {
	os.Exit(0)
}

func init() {
	module := &TestMutate{}
	mutations := make(map[string]lib.StateMutation)
	discovers := make(map[string]map[string]reflect.Value)
	dtestm := make(map[string]reflect.Value)

	dtestm[tspb.TestScaling_NONE.String()] = reflect.ValueOf(tspb.TestScaling_NONE)
	dtestm[tspb.TestScaling_WAIT.String()] = reflect.ValueOf(tspb.TestScaling_WAIT)
	dtestm[tspb.TestScaling_LOW.String()] = reflect.ValueOf(tspb.TestScaling_LOW)
	dtestm[tspb.TestScaling_HIGH.String()] = reflect.ValueOf(tspb.TestScaling_HIGH)

	discovers[ScalingStateURL] = dtestm

	discovers[ModuleStateURL] = map[string]reflect.Value{
		"RUN": reflect.ValueOf(cpb.ServiceInstance_RUN)}

	for m := range muts {
		dur, _ := time.ParseDuration(muts[m].timeout)
		mutations[m] = core.NewStateMutation(
			map[string][2]reflect.Value{
				ScalingStateURL: {
					reflect.ValueOf(muts[m].f),
					reflect.ValueOf(muts[m].t),
				},
			},
			muts[m].reqs,
			excs,
			lib.StateMutationContext_CHILD,
			dur,
			[3]string{module.Name(), "/PhysState", "PHYS_HANG"},
		)
	}

	si := core.NewServiceInstance("testmutate", module.Name(), module.Entry, nil)
	// Register it all
	core.Registry.RegisterModule(module)
	core.Registry.RegisterServiceInstance(module, map[string]lib.ServiceInstance{si.ID(): si})
	core.Registry.RegisterDiscoverable(module, discovers)
	core.Registry.RegisterMutations(module, mutations)
}
