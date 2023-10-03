// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"encoding/json"
	"fmt"
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
	case "metrics-agent":
		startMetricsAgent(ms, cfg)
	}

	ms.Start()
}

func startMetricsAgent(ms *Microservice, cfg IConfig) {
	ms.Schedule(time.Minute, func(ctx IContext) error {
		// Query last 5 metrics
		query := `{
			"size": 5,
			"sort": [
			  {
				"created_at": {
				  "order": "desc"
				}
			  }
			]
		  }`
		idx := ctx.Indexer(cfg.IndexerServers())
		items, _, err := idx.Query("metrics", query)
		if err != nil {
			ctx.Log(fmt.Sprintf("error=%s", err.Error()))
			return err
		}
		if len(items) < 5 {
			return nil
		}

		metrics := make([]*Metrics, len(items))
		for i, item := range items {
			var metric *Metrics
			json.Unmarshal([]byte(item), &metric)
			metrics[i] = metric
		}

		// {
		// 	"id": "2WFTiOiHvuEW9EO9xEFw7XqtRRF",
		// 	"node_name": "ecs-a81d",
		// 	"created_at": "2023-10-03 10:04:14",
		// 	"metrics": {
		// 	  "cpu_used_percent": 9,
		// 	  "mem_used_percent": 81
		// 	}
		// }
		// if last 5 metrics has cpu_used_percent > 90, then alert
		exceedCPUCounter := 0
		for _, metric := range metrics {
			if metric == nil {
				continue
			}

			cpuPercent := InterfaceToInt64(metric.Metrics["cpu_used_percent"])
			if cpuPercent > 90 {
				exceedCPUCounter++
			}
		}
		if exceedCPUCounter >= 5 {
			ctx.Log("ALERT CPU EXCEED")
		}

		// if last 5 metrics has mem_used_percent > 90, then alert
		exceedMemCounter := 0
		for _, metric := range metrics {
			if metric == nil {
				continue
			}

			memPercent := InterfaceToInt64(metric.Metrics["mem_used_percent"])
			if memPercent > 90 {
				exceedMemCounter++
			}
		}
		if exceedMemCounter >= 5 {
			ctx.Log("ALERT MEMORY EXCEED")
		}
		return nil
	})
}

func startNodeDaemon(ms *Microservice, cfg IConfig) {
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
