MATCH (event:Event {UUID: $event_uuid } )<-[at:AT]-(schedulePattern:SchedulePattern {UUID: $schedulePattern_uuid})
SET schedulePattern.checkedUpTo = $checkedUpTo