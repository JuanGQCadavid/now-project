package domain

const (
	DB_INSTANCE_STARTED   EventId = "RDS-EVENT-0088"
	DB_INSTANCE_STOPED    EventId = "RDS-EVENT-0087"
	DB_INSTANCE_RESTARTED EventId = "RDS-EVENT-0006"
	//DB_INSTANCE_RESTARTED EventId = "RDS-EVENT-0006"
)

type EventId string
