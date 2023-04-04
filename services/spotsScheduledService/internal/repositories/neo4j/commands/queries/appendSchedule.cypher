MATCH 
(host:Person {id: $host_id})-[:OWNS]->(event:Event {UUID: $event_uuid })
MERGE 
(schedulePattern:SchedulePattern {UUID: $schedulePattern_uuid})
ON CREATE
	SET schedulePattern.days = $schedulePattern_days
	SET schedulePattern.fromDate = $schedulePattern_fromDate
	SET schedulePattern.toDate = $schedulePattern_toDate
	SET schedulePattern.startTime = $schedulePattern_startTime
	SET schedulePattern.endTime = $schedulePattern_endTime
MERGE 
(host)<-[:HOST_BY]-(schedulePattern)-[:AT {status: $status, timestamp: $timestamp }]->(event)