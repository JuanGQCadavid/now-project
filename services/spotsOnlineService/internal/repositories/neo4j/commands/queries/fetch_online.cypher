
MATCH
	(owner:Person)-[:OWNS]->(event:Event {UUID : $spot_uuid})-[:ON]->(place:Place)
WHERE NOT 
	(event)-[:IS_DELETED]->(event)
OPTIONAL MATCH 
	(event)<-[:BELONGS_TO]-(date:Date)<-[:IS_HOSTING]-(event_host:Person)
RETURN
	event.UUID as event_UUID,
	place.name as place_name,
	place.lon as place_lon,
	place.mapProviderId as place_provider_id,
	place.lat as place_lat,
	owner.id as host_id,
	collect(
		{
			date_uuid: date.UUID,
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
	) as date_onlieonline
