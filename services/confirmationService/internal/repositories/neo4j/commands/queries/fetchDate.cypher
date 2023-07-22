MATCH (place:Place)<-[:ON]-(event:Event)<-[at:AT]-(date:Date {UUID: $dateId})<-[:HOST]-(host:Person)
WHERE date(date.Date) >= date()
return
    date.StartTime as dateStartTime,
    date.Confirmed as dateConfirmed,
    date.UUID as dateUUID,
    date.Date as dateDate,
    host.id as hostId,
    host.name as hostName,
    at.status as status,
    event.UUID as eventUUID,
    place.lat as placeLat ,
    place.lon as placeLon ,
    place.mapProviderId as placeMapProviderId ,
    place.name as placeName