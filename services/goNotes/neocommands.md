# Find a node and its relationships using its uuid

MATCH (e:Event {UUID:"cd4c17fb-84de-4437-a98a-abbb74eee29c" })<-[]->(n)
RETURN e,n