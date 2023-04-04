MATCH 
	(event:Event {UUID:$spot_uuid})<-[at:AT]-(schedulePattern:SchedulePattern {UUID: $schedulePattern_uuid})
SET at.status = $status
SET at.timestamp = $timestamp