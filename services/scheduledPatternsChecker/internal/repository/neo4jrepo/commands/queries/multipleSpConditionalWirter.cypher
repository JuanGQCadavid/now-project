
UNWIND $schedulePatterns AS spProp
MATCH (event:Event {UUID: $spotId})<-[at:AT]-(sp:SchedulePattern {UUID: spProp.id } )-[:HOST_BY]->(host {id: spProp.hostId})
FOREACH (i in CASE WHEN sp.days = spProp.days AND sp.endTime = spProp.endTime AND sp.StartTime = spProp.startTime AND sp.fromDate = spProp.fromDate AND sp.toDate = spProp.toDate THEN [sp] ELSE [] END |
    SET sp.checkedUpTo = spProp.checkedUpTo
    FOREACH (prop IN spProp.datesProps | 

        MERGE (host)-[:HOST]->(date:Date {StartTime: prop.StartTime, Date: prop.Date })-[atDate:AT {status: prop.Status}]->(event)
        ON CREATE
        SET date.UUID = prop.UUID
        SET date.DurationApproximatedInSeconds = prop.DurationApproximatedInSeconds
        SET date.Confirmed = prop.Confirmed
        SET date.MaximunCapacty = prop.MaximunCapacty
        SET atDate.timestamp = prop.Timestamp

        MERGE (sp)<-[:CREATED_FROM]-(date)
    )
)
