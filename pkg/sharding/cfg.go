package sharding

import (
	"time"

	"github.com/infinivision/prophet"
	"seata.io/server/pkg/core"
	"seata.io/server/pkg/meta"
	"seata.io/server/pkg/transport"
)

// Cfg raftstore configuration
type Cfg struct {
	Addr                string
	ShardingAddr        string
	DataPath            string
	Labels              []prophet.Pair
	ProphetName         string
	ProphetAddr         string
	ProphetOptions      []prophet.Option
	FragHBInterval      time.Duration
	StoreHBInterval     time.Duration
	MaxPeerDownDuration time.Duration
	RMLease             time.Duration
	CoreOptions         []core.Option
	InitFragments       int
	Concurrency         int
	OverloadPercentage  uint64
	OverloadPeriod      uint64
	TransSendCB         func(meta.ResourceManager, meta.Notify) error
	TransWorkerCount    int
	PRWorkerCount       int

	// just for test
	storage       storage
	tc            core.TransactionCoordinator
	shardingTrans Transport
	seataTrans    transport.Transport
	storeID       uint64
}

// Adjust adjust
func (c *Cfg) Adjust() {
	if c.FragHBInterval == 0 {
		c.FragHBInterval = time.Second * 10
	}

	if c.StoreHBInterval == 0 {
		c.StoreHBInterval = time.Second * 30
	}
}
