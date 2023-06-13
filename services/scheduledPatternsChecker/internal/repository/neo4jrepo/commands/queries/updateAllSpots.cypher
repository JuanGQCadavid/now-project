UNWIND $schedulePatterns AS sp
MATCH (host:Person {id: sp.HostId})<-[host_relation:HOST_BY]-(schedulePattern:SchedulePattern {UUID: sp.Id})-[:AT {status: "activate"}]->(event:Event {UUID: sp.SpotId } )
SET schedulePattern.checkedUpTo = sp.CheckedUpTo
FOREACH (dateProp IN sp.Dates | 
    MERGE (host)-[:HOST]->(date:Date {StartTime: dateProp.StartTime, Date: dateProp.Date })-[at:AT {status: dateProp.Status}]->(event)
    ON CREATE
        SET date.UUID = dateProp.Id
        SET date.DurationApproximatedInSeconds = dateProp.DurationApproximatedInSeconds
        SET date.Confirmed = dateProp.Confirmed
        SET date.MaximunCapacty = dateProp.MaximunCapacty
        SET at.timestamp = dateProp.Timestamp
)