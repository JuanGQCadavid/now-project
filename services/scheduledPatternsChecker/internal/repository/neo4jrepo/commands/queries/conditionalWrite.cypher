
MATCH (event:Event {UUID: $spotId})<-[at:AT]-(sp:SchedulePattern {UUID: $spId } )-[:HOST_BY]->(host {id: $hostId})
FOREACH (i in CASE WHEN sp.days = $spDays AND sp.endTime = $spEndTime AND sp.StartTime = $spStartTime AND sp.fromDate = $spFromDate AND sp.toDate = $spToDate THEN [sp] ELSE [] END |
    SET sp.checkedUpTo = $spCheckedUpTo
    FOREACH (prop IN $props | 

        MERGE (date:Date {StartTime: prop.StartTime, Date: prop.Date })
        ON CREATE
        SET date.UUID = prop.UUID
        SET date.DurationApproximatedInSeconds = prop.DurationApproximatedInSeconds
        SET date.Confirmed = prop.Confirmed
        SET date.MaximunCapacty = prop.MaximunCapacty

        MERGE (host)-[:HOST]->(date)-[atDate:AT {status: prop.Status}]->(event)
        ON CREATE
        SET atDate.timestamp = prop.Timestamp
    )
)