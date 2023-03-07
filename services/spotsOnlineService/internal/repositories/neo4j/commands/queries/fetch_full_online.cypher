
MATCH
	(owner:Person)-[:OWNS]->(event:Event {UUID : $spot_uuid})-[:ON]->(place:Place)
WHERE NOT 
	(event)-[:IS_DELETED]->(event)
OPTIONAL MATCH 
	(tags:Topic)-[tagged:TAGGED]->(event)
OPTIONAL MATCH 
	(event)<-[:BELONGS_TO]-(date:Date)<-[:IS_HOSTING]-(event_host:Person)
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
	owner.id as host_id,
	owner.name as host_name,
	collect(tags.tag) as tag_tags,
	collect(tagged.isPrincipal) as tag_principals,
	collect(
		{
			date_duration_in_seconds: date.DurationApproximatedInSeconds,
			date_start_time: date.StartTime,
			date_date: date.Date,
			date_confirmed: date.Confirmed,
			date_maximun_capacity: date.MaximunCapacty,
			hosted_by: {
				host_id: event_host.id,
				host_name: event_host.name
			}
		}
	) as date_online
