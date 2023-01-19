# Spots core service

## WTF?

It handles Spots at the core level, it is used only for CRUD regarding spots that later are going to be modified by other services as Oline or schedule spots

## Methods

### Create a spot
``` javascript
HTTP  1.1
POST /spots/core/

{
    "eventInfo": {
        "name": "Gatubella",
        "description": "My friend",
        "maximunCapacty": 58,
        "emoji": ":v"
    },
    // This should be remove it.
    "hostInfo": {
        "id": "",
        "name": "Juan 1",
        "phoneNumber": "+573137590102"
    },
    // A service to validate this is needed.
    "placeInfo": {
        "name": "Place Laureles",
        "lat": 6.245887,
        "lon": -75.589868,
        "mapProviderId": "Place_Laureles"
    },
    "topicInfo": {
        "principalTopic": "TestingIt",
        "secondaryTopics" : ["OtherTest", "MaybeSecondary"]
    }
}
```

### Get spot info
``` javascript
HTTP  1.1
GET /spots/core/<spot_id>/


```

### Update spot info
``` javascript
HTTP  1.1
PATCH /spots/core/<spot_id>/

{
    "eventInfo": {
        "name": "Gatubella",
        "description": "My friend",
        "maximunCapacty": 58,
        "emoji": ":v"
    },
    "placeInfo": {
        "name": "Place Laureles",
        "lat": 6.245887,
        "lon": -75.589868,
        "mapProviderId": "Place_Laureles"
    },
    "topicInfo": {
        "principalTopic": "TestingIt",
        "secondaryTopics" : ["OtherTest", "MaybeSecondary"]
    }
}

```

### Delete spot info
``` javascript
HTTP  1.1
DELETE /spots/core/<spot_id>/
```