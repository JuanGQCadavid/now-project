MATCH
	(place:Place)<-[:ON]-(event:Event)<-[at:AT]-(date:Date)<-[:HOST]-(host:Person)
WHERE
	date.UUID IN  $datesIds
	AND NOT (event)-[:IS_DELETED]->(event)
OPTIONAL MATCH (tags:Topic)-[tagged:TAGGED]->(event)
RETURN
	event.name as event_name,
	event.UUID as event_UUID,
	event.emoji as event_emoji,
	place.lon as place_lon,
	place.mapProviderId as place_provider_id,
	place.lat as place_lat,
	collect(tags.tag) as tag_tags,
	collect(tagged.isPrincipal) as tag_principals,
	date.Confirmed as date_confirmed,
	date.Date as date_date,
	date.DurationApproximatedInSeconds as date_durationApproximatedInSeconds ,
	date.StartTime as date_startTime,
	date.UUID as date_UUID