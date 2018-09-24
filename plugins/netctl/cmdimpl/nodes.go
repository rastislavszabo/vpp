// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//

package cmdimpl

import (
	"encoding/json"
	"fmt"
	nodeinfomodel "github.com/contiv/vpp/plugins/contiv/model/node"

	"github.com/contiv/vpp/plugins/crd/cache/telemetrymodel"
	"github.com/contiv/vpp/plugins/netctl/http"
	"github.com/coreos/etcd/clientv3"
	"github.com/ligato/cn-infra/db/keyval/etcd"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/logging/logrus"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

const timeLayout = "Mon Jan _2 15:04:05 2006"

//PrintNodes will print out all of the cmdimpl in a network in a table format.
func PrintNodes() {
	cfg := &etcd.ClientConfig{
		Config: &clientv3.Config{
			Endpoints: []string{"127.0.0.1:32379"},
		},
		OpTimeout: 1 * time.Second,
	}
	logger := logrus.DefaultLogger()
	logger.SetLevel(logging.FatalLevel)
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	// w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	// Create connection to etcd.
	db, err := etcd.NewEtcdConnectionWithBytes(*cfg, logger)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	itr, err := db.ListValues("/vnf-agent/contiv-ksr/allocatedIDs/")
	if err != nil {
		fmt.Printf("Error getting values")
		return
	}
	fmt.Fprintf(w, "ID\tNODE-NAME\tVPP-IP\tHOST-IP\tSTART-TIME\tSTATE\tBUILD-VERSION\tBUILD-DATE\n")
	for {
		kv, stop := itr.GetNext()
		if stop {
			break
		}

		// Get nodeinfo
		buf := kv.GetValue()
		nodeInfo := &nodeinfomodel.NodeInfo{}
		err = json.Unmarshal(buf, nodeInfo)

		// Get liveness data which contains image version / build time
		bytes := http.GetNodeInfo(nodeInfo.ManagementIpAddress, "liveness")
		var liveness telemetrymodel.NodeLiveness
		err = json.Unmarshal(bytes, &liveness)

		if err != nil {
			fmt.Println(err)
			liveness.BuildDate = "Not Available"
		}

		buildDate := liveness.BuildDate
		bd, err1 := time.Parse("2006-01-02T15:04+00:00", buildDate)
		if err1 == nil {
			buildDate = bd.Format(timeLayout)
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%d\t%s\t%s\n",
			nodeInfo.Id,
			nodeInfo.Name,
			strings.Split(nodeInfo.IpAddress, "/")[0],
			nodeInfo.ManagementIpAddress,
			time.Unix(int64(liveness.StartTime), 0).Format(timeLayout),
			liveness.State,
			liveness.BuildVersion,
			buildDate)
	}
	w.Flush()
	db.Close()
}
