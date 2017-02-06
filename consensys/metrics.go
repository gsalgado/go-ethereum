package consensys

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"

	client "github.com/influxdata/influxdb/client/v2"
)

type MetricsCollector struct {
	eventmux *event.TypeMux
	client   client.Client
	nodeID   string
}

func NewMetricsCollector(eventmux *event.TypeMux, url, nodeID string) (*MetricsCollector, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{Addr: url})
	if err != nil {
		return nil, err
	}
	return &MetricsCollector{nodeID: nodeID, client: c, eventmux: eventmux}, nil
}

func (c *MetricsCollector) Start() {
	glog.V(logger.Info).Infoln("Starting consensys metrics collector")
	sub := c.eventmux.Subscribe(
		core.ChainHeadEvent{},
		// core.TxPreEvent{},
		// core.NewMinedBlockEvent{}
	)

	go func() {
		defer sub.Unsubscribe()
		for event := range sub.Chan() {
			c.process(event)
		}
	}()
}

func (c *MetricsCollector) process(event *event.Event) {
	switch e := event.Data.(type) {
	case core.ChainHeadEvent:
		fmt.Println(e.Block.Header().Number)
	}

	// Node-specific metrics
	//   delta between time.Now() and block.Header().Time?

	// Common metrics
	//   block rate

}
