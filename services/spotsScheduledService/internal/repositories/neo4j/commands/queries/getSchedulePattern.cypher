
MATCH
	(owner:Person {id: $host_id})-[:OWNS]->(event:Event {UUID : $spot_uuid})
WHERE NOT 
	(event)-[:IS_DELETED]->(event)
OPTIONAL MATCH 
    (host)<-[:HOST_BY]-(schedulePattern)-[at:AT]->(event)
RETURN
	event.UUID as event_UUID,
	owner.id as owner_id,
	collect(
		{   
            schedulePattern_id: schedulePattern.UUID,
			schedulePattern_days: schedulePattern.days,
			schedulePattern_fromDate: schedulePattern.fromDate,
			schedulePattern_toDate: schedulePattern.toDate,
			schedulePattern_StartTime: schedulePattern.StartTime,
			schedulePattern_endTime: schedulePattern.endTime,
            state: {
                status: at.status,
                since: at.timestamp
            },
			hosted_by: {
				host_id: host.id,
				host_name: host.name
			}
		}
	) as schedulePatterns