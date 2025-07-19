MATCH (p:Person)-[r]->(d:Date {UUID: $date_id})-[:AT]-(e:Event {UUID: $event_id})
return p.id as user_id, p.name as user_name, type(r) as access_type