
MATCH (p:Person {id: $user_id})-[r]->(e:Event{UUID: $event_id })
return p.id as "user_id", p.name as "user_name", type(r) as "relation_type"