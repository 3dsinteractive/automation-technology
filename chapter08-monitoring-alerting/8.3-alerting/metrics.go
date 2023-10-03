package main

type Metrics struct {
	ID        string                 `json:"id"`
	NodeName  string                 `json:"node_name"`
	CreatedAt *Timestamp             `json:"created_at"`
	Metrics   map[string]interface{} `json:"metrics"`
}

func NewMetrics(nodeName string, createdAt *Timestamp) *Metrics {
	return &Metrics{
		ID:        NewUUIDGen().NewUUID(),
		NodeName:  nodeName,
		CreatedAt: createdAt,
		Metrics:   map[string]interface{}{},
	}
}

func (m *Metrics) SetMetric(name string, value interface{}) {
	m.Metrics[name] = value
}
