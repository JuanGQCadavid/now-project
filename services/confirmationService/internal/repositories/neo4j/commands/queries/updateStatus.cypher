MATCH (date:Date {UUID: $dateId})
SET date.Confirmed = $confirmed