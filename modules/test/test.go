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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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

// PXE provides PXE-boot capabilities
type Test struct {
	api   lib.APIClient
	cfg   *pb.TestConfig
	dchan chan<- lib.Event

	pollTicker *time.Ticker
	router     *mux.Router
	srv        *http.Server
}

type reqTest struct {
	ID    string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
}

// Name returns the FQDN of the module
func (*Test) Name() string { return "github.com/hpc/kraken/modules/test" }

// NewConfig returns a fully initialized default config
func (*Test) NewConfig() proto.Message {
	r := &pb.TestConfig{
		IpUrl: "type.googleapis.com/proto.IPv4OverEthernet/Ifaces/0/Ip/Ip",
		Addr:  "127.0.0.1",
		Port:  3143,
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
	// dur, _ := time.ParseDuration("10s")
	// t.pollTicker = time.NewTicker(dur)

	// // main loop
	// for {
	// 	select {
	// 	case <-t.pollTicker.C:
	// 		go t.discoverAll()
	// 		break
	// 	}
	// }

	t.setupRouter()
	for {
		t.startServer()
	}

}

func (t *Test) setupRouter() {
	t.router = mux.NewRouter()
	t.router.HandleFunc("/set", t.setTest).Methods("POST")
}

func (t *Test) startServer() {
	t.srv = &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"PUT", "GET", "POST", "DELETE"}),
		)(t.router),
		Addr:         fmt.Sprintf("%s:%d", t.cfg.Addr, t.cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	t.api.Logf(lib.LLINFO, "restapi is listening on: %s\n", t.srv.Addr)
	if e := t.srv.ListenAndServe(); e != nil {
		if e != http.ErrServerClosed {
			t.api.Logf(lib.LLNOTICE, "http stopped: %v\n", e)
		}
	}
	t.api.Log(lib.LLNOTICE, "restapi listener stopped")
}

func (t *Test) srvStop() {
	t.api.Log(lib.LLDEBUG, "restapi is shutting down listener")
	t.srv.Shutdown(context.Background())
}

func (t *Test) setTest(w http.ResponseWriter, req *http.Request) {
	t.api.Logf(lib.LLDEBUG, "Setting test thing")
	defer req.Body.Close()
	body, e := ioutil.ReadAll(req.Body)
	if e != nil {
		t.api.Logf(lib.LLERROR, "http GET failed to read body: %v", e)
		return
	}
	rt := reqTest{}
	e = json.Unmarshal(body, &rt)
	t.api.Logf(lib.LLDEBUG, "got this from request: %+v", rt)

	w.WriteHeader(http.StatusOK)
	t.fakeDiscover(rt.ID, rt.State)
}

func (t *Test) fakeDiscover(id string, state string) {

	vid := testpb.Test_Test_value[state]
	n, e := t.api.QueryRead(id)
	if e != nil {
		t.api.Logf(lib.LLERROR, "error getting node from id")
	}

	url := lib.NodeURLJoin(n.ID().String(), "type.googleapis.com/proto.Test/State")
	v := core.NewEvent(
		lib.Event_DISCOVERY,
		url,
		&core.DiscoveryEvent{
			Module:  t.Name(),
			URL:     url,
			ValueID: testpb.Test_Test_name[vid],
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

	dtest[testpb.Test_NORMAL.String()] = reflect.ValueOf(testpb.Test_NORMAL)
	dtest[testpb.Test_WARNING.String()] = reflect.ValueOf(testpb.Test_WARNING)
	dtest[testpb.Test_CRITICAL.String()] = reflect.ValueOf(testpb.Test_CRITICAL)
	dtest[testpb.Test_UNKNOWN.String()] = reflect.ValueOf(testpb.Test_UNKNOWN)

	discovers["type.googleapis.com/proto.Test/State"] = dtest

	discovers[SrvStateURL] = map[string]reflect.Value{
		"RUN": reflect.ValueOf(cpb.ServiceInstance_RUN)}
	si := core.NewServiceInstance("test", module.Name(), module.Entry, nil)
	// Register it all
	core.Registry.RegisterModule(module)
	core.Registry.RegisterServiceInstance(module, map[string]lib.ServiceInstance{si.ID(): si})
	core.Registry.RegisterDiscoverable(module, discovers)
}
