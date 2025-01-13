MERGE (event:Event {UUID: $event_uuid })
ON CREATE
    SET event.description = $event_desc
    SET event.maximunCapacty = $event_max_capacity
    SET event.name = $event_name
    SET event.emoji = $event_emoji
MERGE (place:Place {mapProviderId: $place_provider_id})
ON CREATE
    SET place.lat = toFloat($place_lat)
    SET place.lon = toFloat($place_lon)
    SET place.name = $place_name
MERGE (host:Person {phoneNumber:$host_phone_number})
ON CREATE 
    SET host.name = $host_name
    SET host.id = $host_id
MERGE (host)-[:OWNS]->(event)-[:ON]->(place)