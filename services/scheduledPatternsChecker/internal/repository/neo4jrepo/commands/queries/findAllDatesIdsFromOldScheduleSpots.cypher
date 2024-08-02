MATCH (d:Date {Confirmed: true})-[at:AT {status:"SCHEDULED"}]->(event:Event)
where d.Date =~ "[0-9]*-[0-9]*-[0-9]*" and datetime(d.Date+ "T"+ d.StartTime) + duration({seconds: d.DurationApproximatedInSeconds}) < datetime()
RETURN d.UUID, d.Date, datetime(d.Date+ "T"+ d.StartTime) + duration({seconds: d.DurationApproximatedInSeconds}), datetime()