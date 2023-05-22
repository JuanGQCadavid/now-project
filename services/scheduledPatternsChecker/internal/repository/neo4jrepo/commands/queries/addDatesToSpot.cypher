MATCH (host_%[1]d:Person {id: $host_id})-[host_relation_%[1]d:OWNS]->(event_%[1]d:Event {UUID: $event_uuid } )
MERGE (date_%[1]d:Date {UUID: $date_uuid})
ON CREATE
	SET date_%[1]d.DurationApproximatedInSeconds = $date_approximated_seconds
	SET date_%[1]d.StartTime = $date_start_time
	SET date_%[1]d.Date = $date_date
	SET date_%[1]d.Confirmed = $date_confirmed
	SET date_%[1]d.MaximunCapacty = $date_maximun_capacity
MERGE (host_%[1]d)-[:HOST]->(date_%[1]d)-[:AT {status: $status, timestamp: $timestamp }]->(event_%[1]d)
