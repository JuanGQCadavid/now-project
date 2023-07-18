MATCH (:Event {UUID: $spot_uuid})<-[at:AT]-(date:Date)-[:CREATED_FROM]->(sp:SchedulePattern)-[:HOST_BY]->(host:Person)
WHERE date(date.Date) >= date()
return collect(
    {
        dateMaximunCapacity: date.MaximunCapacty,
        dateDurationApproximatedInSeconds: date.DurationApproximatedInSeconds,
        dateStartTime: date.StartTime,
        dateConfirmed: date.Confirmed,
        dateUUID: date.UUID,
        dateDate: date.Date,
        schedulePatternId: sp.UUID,
        hostedBy: {
            hostId: host.id,
            hostName: host.name
        },
        dateState: {
            status: at.status,
            timestamp: at.timestamp
        }
    }
) as dates