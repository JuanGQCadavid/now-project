MATCH 
(host:Person {id: $host_id})-[:OWNS]->(event:Event {UUID: $event_uuid })
MERGE 
(date:Date {UUID: $date_uuid})
ON CREATE
	SET date.DurationApproximatedInSeconds = $date_approximated_seconds
	SET date.StartTime = $date_start_time
	SET date.Date = $date_date
	SET date.Confirmed = $date_confirmed
	SET date.MaximunCapacty = $date_maximun_capacity
MERGE 
(host)-[:HOST]->(date)-[:AT {status: $status, timestamp: $timestamp }]->(event)
