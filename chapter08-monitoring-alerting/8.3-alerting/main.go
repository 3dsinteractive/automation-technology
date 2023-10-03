// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"os"
	"time"
)

func main() {
	cfg := NewConfig()

	ms := NewMicroservice()
	ms.RegisterLivenessProbeEndpoint("/healthz")

	serviceID := cfg.ServiceID()

	switch serviceID {
	case "node-daemon":
		startNodeDaemon(ms, cfg)
	}

	ms.Start()
}

func startNodeDaemon(ms *Microservice, cfg IConfig) {
	ms.Schedule(time.Minute, func(ctx IContext) error {
		// Query last 10 metrics

		return nil
	})

	ms.Schedule(10*time.Second, func(ctx IContext) error {

		osStat, err := osStat()
		if err != nil {
			ctx.Log(err.Error())
			return err
		}

		now := NewTimestampT(ctx.Now())
		nodeName := os.Getenv("MY_NODE_NAME")

		metric := NewMetrics(nodeName, now)
		metric.SetMetric("cpu_used_percent", osStat.CPUUserPercent)
		metric.SetMetric("mem_used_percent", osStat.MemoryUsedPercent)

		idx := ctx.Indexer(cfg.IndexerServers())

		_, err = idx.Index("metrics", metric.ID, metric)
		if err != nil {
			ctx.Log(err.Error())
			return err
		}

		return nil
	})
}
