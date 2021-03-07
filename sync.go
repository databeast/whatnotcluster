package cluster

/*
 gRPC-synchronized synchronization of namespaces across cluster instances

 connections are maintained via gRPC stream mode
*/

import (
	"context"
	"github.com/databeast/whatnot"
	"github.com/databeast/whatnotcluster/transports"
)

type NamespaceSyncService struct {

}

func (n NamespaceSyncService) Sync(server SyncService_SyncServer) error {
	panic("implement me")
}

func NewSyncServiceForManager(ctx context.Context, mgr whatnot.NameSpaceManager) {
	var SyncService SyncServiceServer
	SyncService = &NamespaceSyncService{}
}

// Receive Sync Updates from Cluster Members

// Transmit Changes to Cluster Members

// Resolve State Conflicts with the Quorum Leader

// Populate new Cluster Member with entire state of Namespace

// Lock Entire Namespace for Sync Resolution



