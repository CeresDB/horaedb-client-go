// Copyright 2022 CeresDB Project Authors. Licensed under Apache-2.0.

package test

import (
	"testing"
	"time"

	"github.com/CeresDB/ceresdb-client-go/ceresdb"
	"github.com/stretchr/testify/require"
)

func TestRouteGc(t *testing.T) {
	t.Skip("ignore local test")

	client, err := ceresdb.NewClient(clusterEndpoint, ceresdb.Direct,
		ceresdb.EnableLoggerDebug(true),
		ceresdb.WithRouteMaxCacheSize(3),
	)
	require.NoError(t, err, "init ceresdb client failed")

	timestamp := currentMS()

	testBaseWrite(t, client, "ceresdb_route_test1", timestamp, 1)
	time.Sleep(time.Second)
	testBaseWrite(t, client, "ceresdb_route_test2", timestamp, 2)
	time.Sleep(time.Second)
	testBaseWrite(t, client, "ceresdb_route_test3", timestamp, 3)
	testBaseQuery(t, client, "ceresdb_route_test2", timestamp, 2)
	time.Sleep(time.Second)
	testBaseWrite(t, client, "ceresdb_route_test4", timestamp, 4)
	time.Sleep(time.Second)
	testBaseWrite(t, client, "ceresdb_route_test5", timestamp, 5)

	// Under single-threaded test conditions,
	// the expected result is that the routes of ceresdb_route_test1 and table ceresdb_route_test3 are cleaned up

	time.Sleep(12 * time.Second)
}

func TestRouteProxy(t *testing.T) {
	t.Skip("ignore local test")

	client, err := ceresdb.NewClient(clusterEndpoint, ceresdb.Proxy,
		ceresdb.EnableLoggerDebug(true),
		ceresdb.WithRouteMaxCacheSize(3),
	)
	require.NoError(t, err, "init ceresdb client failed")

	timestamp := currentMS()
	testBaseWrite(t, client, "ceresdb_route_test1", timestamp, 1)
}
