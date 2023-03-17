MATCH 
	(event:Event {UUID:$spot_uuid})<-[at:AT]-(date:Date {UUID: $date_uuid})
SET at.status = $status
SET at.timestamp = $timestamp