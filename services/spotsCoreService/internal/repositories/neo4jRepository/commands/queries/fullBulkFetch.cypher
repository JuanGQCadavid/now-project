MATCH
	(place:Place)<-[:ON]-(event:Event)<-[at:AT]-(date:Date)<-[:HOST]-(host:Person)
WHERE
	date.UUID IN  $datesIds
	AND NOT (event)-[:IS_DELETED]->(event)
OPTIONAL MATCH (tags:Topic)-[tagged:TAGGED]->(event)
RETURN
	event.description as event_desc,
	event.name as event_name,
	event.maximunCapacty as event_max_capacity,
	event.UUID as event_UUID,
	event.emoji as event_emoji,
	place.name as place_name,
	place.lon as place_lon,
	place.mapProviderId as place_provider_id,
	place.lat as place_lat,
	host.id as host_id,
	host.name as host_name,
	collect(tags.tag) as tag_tags,
	collect(tagged.isPrincipal) as tag_principals,
	date.Confirmed as date_confirmed,
	date.Date as date_date,
	date.DurationApproximatedInSeconds as date_durationApproximatedInSeconds ,
	date.StartTime as date_startTime,
	date.UUID as date_UUID