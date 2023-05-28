
MATCH (event:Event {UUID: $spotId})<-[at:AT]-(sp:SchedulePattern {UUID: $spId } )-[:HOST_BY]->(host {id: $hostId})
FOREACH (i in CASE WHEN sp.days = $spDays AND sp.endTime = $spEndTime AND sp.StartTime = $spStartTime AND sp.fromDate = $spFromDate AND sp.toDate = $spToDate THEN [sp] ELSE [] END |
    SET sp.checkedUpTo = $spCheckedUpTo
    FOREACH (props IN $props | 

        MERGE (date:Date {StartTime: props.StartTime, Date: props.Date })
        ON CREATE
        SET date.UUID = props.UUID
        SET date.DurationApproximatedInSeconds = props.DurationApproximatedInSeconds
        SET date.Confirmed = props.Confirmed
        SET date.MaximunCapacty = props.MaximunCapacty

        MERGE (host)-[:HOST]->(date)-[atDate:AT {status: props.Status}]->(event)
        ON CREATE
        SET atDate.timestamp = props.Timestamp
    )
)