
MATCH
	(event)<-[at:AT]-(schedulePattern)-[:HOST_BY]->(host)
WHERE  
	NOT (event)-[:IS_DELETED]->(event) AND at.status = "activate" AND date(schedulePattern.fromDate) <= date() AND date(schedulePattern.toDate) >= date() 
RETURN
	event.UUID as event_UUID,
	collect(
		{   
            schedulePattern_id: schedulePattern.UUID,
			schedulePattern_days: schedulePattern.days,
			schedulePattern_fromDate: schedulePattern.fromDate,
			schedulePattern_toDate: schedulePattern.toDate,
			schedulePattern_StartTime: schedulePattern.StartTime,
			schedulePattern_endTime: schedulePattern.endTime,
			schedulePattern_checkedUpTo: schedulePattern.checkedUpTo,
			host_id: host.id
		}
	) as schedulePatterns

