
UNWIND $SchedulePatterns AS spProp
MATCH (:Event)<-[:AT {status: $DateStatus}]-(dates:Date WHERE date(dates.Date) >= date())-[:CREATED_FROM]->(:SchedulePattern {UUID: spProp })
WITH dates, COLLECT(dates.UUID) AS datesId
DETACH DELETE dates
RETURN datesId