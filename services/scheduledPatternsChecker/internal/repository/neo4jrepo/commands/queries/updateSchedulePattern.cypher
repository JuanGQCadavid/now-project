MATCH (schedulePattern_%[1]d:SchedulePattern {UUID: $%s})
SET schedulePattern_%[1]d.checkedUpTo = $%s
WITH schedulePattern_%[1]d