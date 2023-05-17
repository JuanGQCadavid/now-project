MATCH (event:Event {UUID: $%s } )<-[at:AT]-(schedulePattern:SchedulePattern {UUID: $%s})
SET schedulePattern.checkedUpTo = $%d
