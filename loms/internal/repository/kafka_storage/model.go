package kafka_storage

import "database/sql"

const (
	StatusNew    string = "new"
	StatusSent   string = "sent"
	StatusFailed string = "failed"
)

type Outbox struct {
	ID         string
	CreatedAt  sql.NullTime
	Event      string
	EntityType string
	EntityID   string
	Status     string
	Data       string
}
