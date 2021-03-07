package cluster

import (
	"github.com/hashicorp/memberlist"
)

/*
Cluster AutoDiscovery via GOSSIP protocol

*/

type WhatnotCluster struct {
	members *memberlist.Memberlist
}

func (c *WhatnotCluster) BeginDiscovery() (err error) {
	cnf := memberlist.DefaultLANConfig()
	c.members, err = memberlist.Create(cnf)
	if err != nil {
		return err
	}

	return nil
}

func NewClusterDefinition() {

}
