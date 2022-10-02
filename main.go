// Copyright 2016 Cong Ding
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"

	"github.com/ccding/go-stun/stun"
)

// 스턴 서버 목록 - https://ourcodeworld.com/articles/read/1536/list-of-free-functional-public-stun-servers-2021
// stun.actionvoip.com:3478
// stun.cablenet-as.net:3478
// stun.dcalling.de:3478

func main() {
	var serverAddr = flag.String("s", "stun.rockenstein.de:3478", "STUN server address")
	var b = flag.Bool("b", false, "NAT behavior test mode")
	var v = flag.Bool("v", true, "verbose mode")
	var vv = flag.Bool("vv", false, "double verbose mode (includes -v)")
	var vvv = flag.Bool("vvv", false, "triple verbose mode (includes -v and -vv)")
	flag.Parse()

	// Creates a STUN client. NewClientWithConnection can also be used if you want to handle the UDP listener by yourself.
	client := stun.NewClient()
	// The default addr (stun.DefaultServerAddr) will be used unless we call SetServerAddr.
	client.SetServerAddr(*serverAddr)
	// Non verbose mode will be used by default unless we call SetVerbose(true) or SetVVerbose(true).
	client.SetVerbose(*v || *vv || *vvv)
	client.SetVVerbose(*vv || *vvv)

	if *b {
		behaviorTest(client)
		return
	}

	// Discover the NAT and return the result.
	nat, host, err := client.Discover()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("NAT Type:", nat)
	if host != nil {
		fmt.Println("External IP Family:", host.Family())
		fmt.Println("External IP:", host.IP())
		fmt.Println("External Port:", host.Port())
	}
}

func behaviorTest(c *stun.Client) {
	natBehavior, err := c.BehaviorTest()
	if err != nil {
		fmt.Println(err)
	}

	if natBehavior != nil {
		fmt.Println("  Mapping Behavior:", natBehavior.MappingType)
		fmt.Println("Filtering Behavior:", natBehavior.FilteringType)
		fmt.Println("   Normal NAT Type:", natBehavior.NormalType())
	}
}
