MATCH 
(host:Person {id: "33ddab28-006b-4790-bf42-1832f90dc8d4"})-[:OWNS]->(event:Event {UUID: "0bf5182b-33b1-4ad1-b5e6-992d1c034609" })
MERGE 
(schedulePattern:SchedulePattern {UUID: "UUID_123"})
ON CREATE
	SET schedulePattern.days = 42
	SET schedulePattern.fromDate = "2007-03-01"
	SET schedulePattern.toDate = "2007-07-01"
	SET schedulePattern.StartTime = "13:00:00"
	SET schedulePattern.endTime = "16:00:00"
MERGE 
(host)<-[:HOST_BY]-(schedulePattern)-[:AT {status: "activate", timestamp: 12220321 }]->(event)
RETURN host, event, schedulePattern

MATCH
	(owner:Person {id: "33ddab28-006b-4790-bf42-1832f90dc8d4"})-[:OWNS]->(event:Event {UUID : "0bf5182b-33b1-4ad1-b5e6-992d1c034609"})
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

MATCH 
	(event:Event {UUID:"0bf5182b-33b1-4ad1-b5e6-992d1c034609"})<-[at:AT]-(schedulePattern:SchedulePattern {UUID: "UUID_123"})
SET at.status = "freezed"
SET at.timestamp = 123455