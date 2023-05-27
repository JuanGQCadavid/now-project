

MATCH (sp:SchedulePattern {UUID: "8e9b748c-d5af-42d1-8ef0-70e124b246cf" })
FOREACH (i in CASE WHEN sp.days = 104 AND sp.endTime = "16:00:00" AND sp.StartTime = "13:00:00" AND sp.fromDate = "2022-05-10" AND sp.toDate = "2022-07-01" THEN [sp] ELSE [] END |
  
  
  
  
  
  
  
  
  
  
  SET i.theOtherNew = 5678
)
RETURN sp


MATCH (schedulePattern_0:SchedulePattern {UUID: $schedule_pattern_id_0})
SET schedulePattern_0.checkedUpTo = $checkedUpTo_0
WITH schedulePattern_0

MATCH (host_0:Person {id: $host_id_0})-[host_relation_0:OWNS]->(event_0:Event {UUID: $event_uuid_0 } )

MERGE (date_0:Date {StartTime: $date_start_time_0, Date: $date_date_0 })
ON CREATE
SET date_0.UUID = $date_uuid_0
SET date_0.DurationApproximatedInSeconds = $date_approximated_seconds_0
SET date_0.Confirmed = $date_confirmed_0
SET date_0.MaximunCapacty = $date_maximun_capacity_0


MERGE (host_0)-[:HOST]->(date_0)-[at_0:AT {status: $status}]->(event_0)
ON CREATE
SET at_0.timestamp = $timestamp 

MERGE (date_1:Date {StartTime: $date_start_time_1, Date: $date_date_1 })
ON CREATE
SET date_1.UUID = $date_uuid_1
SET date_1.DurationApproximatedInSeconds = $date_approximated_seconds_1
SET date_1.Confirmed = $date_confirmed_1
SET date_1.MaximunCapacty = $date_maximun_capacity_1


MERGE (host_0)-[:HOST]->(date_1)-[at_1:AT {status: $status}]->(event_0)
ON CREATE
SET at_1.timestamp = $timestamp 

MERGE (date_2:Date {StartTime: $date_start_time_2, Date: $date_date_2 })
ON CREATE
SET date_2.UUID = $date_uuid_2
SET date_2.DurationApproximatedInSeconds = $date_approximated_seconds_2
SET date_2.Confirmed = $date_confirmed_2
SET date_2.MaximunCapacty = $date_maximun_capacity_2


MERGE (host_0)-[:HOST]->(date_2)-[at_2:AT {status: $status}]->(event_0)
ON CREATE
SET at_2.timestamp = $timestamp