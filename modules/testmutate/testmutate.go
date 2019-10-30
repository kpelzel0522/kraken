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

type tsmut struct {
	f       tpb.Test_Test
	t       tspb.TestScaling_Scaling
	reqs    map[string]reflect.Value
	excs    map[string]reflect.Value
	timeout string
}

type tmut struct {
	// tf      tpb.Test_Test
	// tt      tpb.Test_Test
	// tsf     tspb.TestScaling_Scaling
	// tst     tspb.TestScaling_Scaling
	f       tpb.Test_Test
	t       tpb.Test_Test
	reqs    map[string]reflect.Value
	excs    map[string]reflect.Value
	timeout string
}

type mut struct {
	tf      tpb.Test_Test
	tt      tpb.Test_Test
	tsf     tspb.TestScaling_Scaling
	tst     tspb.TestScaling_Scaling
	reqs    map[string]reflect.Value
	excs    map[string]reflect.Value
	timeout string
}

// var muts = map[string]mut{
// 	"LOW_TEMPtoHIGH_PERF": {
// 		tf:  tpb.Test_UNKNOWN,
// 		tt:  tpb.Test_LOW,
// 		tsf: tspb.TestScaling_NONE,
// 		tst: tspb.TestScaling_HIGH,
// 		reqs: map[string]reflect.Value{
// 			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
// 			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
// 			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_NONE),
// 		},
// 		timeout: "60s",
// 	},
// 	"HIGH_TEMPtoLOW_PERF": {
// 		tf:  tpb.Test_UNKNOWN,
// 		tt:  tpb.Test_HIGH,
// 		tsf: tspb.TestScaling_NONE,
// 		tst: tspb.TestScaling_LOW,
// 		reqs: map[string]reflect.Value{
// 			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
// 			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
// 			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_NONE),
// 		},
// 		timeout: "60s",
// 	},
// }

var tmuts = map[string]tmut{
	"WARNINGtoNORMAL": {
		f: tpb.Test_WARNING,
		t: tpb.Test_NORMAL,
		reqs: map[string]reflect.Value{
			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
			TempStateURL: reflect.ValueOf(tpb.Test_WARNING),
		},
		timeout: "60s",
	},
	"CRITICALtoNORMAL": {
		f: tpb.Test_CRITICAL,
		t: tpb.Test_NORMAL,
		reqs: map[string]reflect.Value{
			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
			TempStateURL: reflect.ValueOf(tpb.Test_CRITICAL),
		},
		timeout: "60s",
	},
	"UNKNOWNtoNORMAL": {
		f: tpb.Test_UNKNOWN,
		t: tpb.Test_NORMAL,
		reqs: map[string]reflect.Value{
			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
			TempStateURL: reflect.ValueOf(tpb.Test_UNKNOWN),
		},
		timeout: "60s",
	},
	"NORMALtoWARNING": {
		f: tpb.Test_NORMAL,
		t: tpb.Test_WARNING,
		reqs: map[string]reflect.Value{
			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_NONE),
		},
		timeout: "60s",
	},
	"NORMALtoCRITICAL": {
		f: tpb.Test_NORMAL,
		t: tpb.Test_CRITICAL,
		reqs: map[string]reflect.Value{
			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_NONE),
		},
		timeout: "60s",
	},
	"NORMALtoUNKNOWN": {
		f: tpb.Test_NORMAL,
		t: tpb.Test_UNKNOWN,
		reqs: map[string]reflect.Value{
			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_NONE),
		},
		timeout: "60s",
	},
}

// var tsmuts = map[string]tsmut{
// 	"NONEtoHIGH": {
// 		f: tspb.TestScaling_NONE,
// 		t: tspb.TestScaling_HIGH,
// 		reqs: map[string]reflect.Value{
// 			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
// 			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
// 			TempStateURL: reflect.ValueOf(tpb.Test_LOW),
// 			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_HIGH),
// 		},
// 		timeout: "60s",
// 	},
// 	"HIGHtoLOW": {
// 		f: tspb.TestScaling_HIGH,
// 		t: tspb.TestScaling_LOW,
// 		reqs: map[string]reflect.Value{
// 			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
// 			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
// 			TempStateURL: reflect.ValueOf(tpb.Test_HIGH),
// 			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_HIGH),
// 		},
// 		timeout: "60s",
// 	},
// 	"LOWtoHIGH": {
// 		f: tspb.TestScaling_LOW,
// 		t: tspb.TestScaling_HIGH,
// 		reqs: map[string]reflect.Value{
// 			"/PhysState": reflect.ValueOf(cpb.Node_POWER_ON),
// 			"/RunState":  reflect.ValueOf(cpb.Node_SYNC),
// 			TempStateURL: reflect.ValueOf(tpb.Test_LOW),
// 			// ScalingStateURL: reflect.ValueOf(tspb.TestScaling_LOW),
// 		},
// 		timeout: "60s",
// 	},
// }

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
		case e := <-t.mchan:
			if e.Type() != lib.Event_STATE_MUTATION {
				t.api.Log(lib.LLERROR, "got unexpected non-mutation event")
				break
			}
			m := e.Data().(*core.MutationEvent)
			go t.handleMutation(m)
			break
		}
	}
}

func (t *TestMutate) handleMutation(m *core.MutationEvent) {
	t.api.Logf(lib.LLDEBUG, "Got mutation: %+v", m)
	switch m.Type {
	case core.MutationEvent_MUTATE:
		switch m.Mutation[1] {
		case "WARNINGtoNORMAL": // starting a new mutation, register the node
			t.api.Logf(lib.LLDEBUG, "Got warning -> normal. sending normal: %+v", m)
			url := lib.NodeURLJoin(m.NodeCfg.ID().String(), TempStateURL)
			ev := core.NewEvent(
				lib.Event_DISCOVERY,
				url,
				&core.DiscoveryEvent{
					Module:  t.Name(),
					URL:     url,
					ValueID: tpb.Test_UNKNOWN.String(),
				},
			)
			t.dchan <- ev
		case "CRITICALtoNORMAL": // starting a new mutation, register the node
			t.api.Logf(lib.LLDEBUG, "Got critical -> normal. sending normal: %+v", m)
			url := lib.NodeURLJoin(m.NodeCfg.ID().String(), TempStateURL)
			ev := core.NewEvent(
				lib.Event_DISCOVERY,
				url,
				&core.DiscoveryEvent{
					Module:  t.Name(),
					URL:     url,
					ValueID: tpb.Test_UNKNOWN.String(),
				},
			)
			t.dchan <- ev
		case "UNKNOWNtoNORMAL": // starting a new mutation, register the node
			t.api.Logf(lib.LLDEBUG, "Got unknown -> normal. sending normal: %+v", m)
			url := lib.NodeURLJoin(m.NodeCfg.ID().String(), TempStateURL)
			ev := core.NewEvent(
				lib.Event_DISCOVERY,
				url,
				&core.DiscoveryEvent{
					Module:  t.Name(),
					URL:     url,
					ValueID: tpb.Test_UNKNOWN.String(),
				},
			)
			t.dchan <- ev
		}
	}
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
	dtestm[tspb.TestScaling_LOW.String()] = reflect.ValueOf(tspb.TestScaling_LOW)
	dtestm[tspb.TestScaling_HIGH.String()] = reflect.ValueOf(tspb.TestScaling_HIGH)

	discovers[ScalingStateURL] = dtestm

	discovers[ModuleStateURL] = map[string]reflect.Value{
		"RUN": reflect.ValueOf(cpb.ServiceInstance_RUN)}

	// for name, m := range muts {
	// 	dur, _ := time.ParseDuration(m.timeout)
	// 	mutations[name] = core.NewStateMutation(
	// 		map[string][2]reflect.Value{
	// 			ScalingStateURL: {
	// 				reflect.ValueOf(m.tsf),
	// 				reflect.ValueOf(m.tst),
	// 			},
	// 			TempStateURL: {
	// 				reflect.ValueOf(m.tf),
	// 				reflect.ValueOf(m.tt),
	// 			},
	// 		},
	// 		m.reqs,
	// 		m.excs,
	// 		lib.StateMutationContext_CHILD,
	// 		dur,
	// 		[3]string{module.Name(), ScalingStateURL, tspb.TestScaling_NONE.String()},
	// 	)
	// }

	// for name, m := range tsmuts {
	// 	dur, _ := time.ParseDuration(m.timeout)
	// 	mutations[name] = core.NewStateMutation(
	// 		map[string][2]reflect.Value{
	// 			ScalingStateURL: {
	// 				reflect.ValueOf(m.f),
	// 				reflect.ValueOf(m.t),
	// 			},
	// 		},
	// 		m.reqs,
	// 		m.excs,
	// 		lib.StateMutationContext_CHILD,
	// 		dur,
	// 		[3]string{module.Name(), ScalingStateURL, tspb.TestScaling_NONE.String()},
	// 	)
	// }

	for name, m := range tmuts {
		dur, _ := time.ParseDuration(m.timeout)
		mutations[name] = core.NewStateMutation(
			map[string][2]reflect.Value{
				TempStateURL: {
					reflect.ValueOf(m.f),
					reflect.ValueOf(m.t),
				},
			},
			m.reqs,
			m.excs,
			lib.StateMutationContext_CHILD,
			dur,
			[3]string{module.Name(), ScalingStateURL, tspb.TestScaling_NONE.String()},
		)
	}

	si := core.NewServiceInstance("testmutate", module.Name(), module.Entry, nil)
	// Register it all
	core.Registry.RegisterModule(module)
	core.Registry.RegisterServiceInstance(module, map[string]lib.ServiceInstance{si.ID(): si})
	core.Registry.RegisterDiscoverable(module, discovers)
	core.Registry.RegisterMutations(module, mutations)
}
