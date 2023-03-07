MATCH 
(host:Person {phoneNumber: $host_phone_number})-[:OWNS]->(event:Event {UUID: $event_uuid })
MERGE 
(date:Date {UUID: $date_uuid})
ON CREATE
	SET date.DurationApproximatedInSeconds = toFloat($date_approximated_seconds)
	SET date.StartTime = $date_start_time
	SET date.Date = $date_date
	SET date.Confirmed = $date_confirmed
	SET date.MaximunCapacty = $date_maximun_capacity
MERGE 
(host)-[:IS_HOSTING]->(date)-[:BELONGS_TO]->(event)
