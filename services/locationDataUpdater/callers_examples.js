
Attributes: 
Operation String onlineStart

Body:
{
    "spotId":"f11495e4-3b25-4aa8-9b7d-33583b7242fe",
    "userId":"33ddab28-006b-4790-bf42-1832f90dc8d4",
    "aditionalpayload":{
       "spotInfo":{
          "SpotId":"f11495e4-3b25-4aa8-9b7d-33583b7242fe"
       },
       "dateInfo":[
          {
             "dateId":"c41be642-2b9c-4978-90ee-29db8e896487",
             "durationApproximated":300,
             "startTime":"2023-08-01 13:59:02.307892389 +0000 UTC m=+0.892990774",
             "date":"2023-08-01 13:59:02.307892389 +0000 UTC m=+0.892990774",
             "maximunCapacty":58,
             "hostInfo":{
                "host_name":""
             },
             "state":{
                "confirmed":true
             }
          }
       ],
       "placeInfo":{
          "name":"Place Laureles",
          "lat":6.245887,
          "lon":-75.589868,
          "mapProviderId":"Place_Laureles"
       }
    }
 }

 Attributes:
Operation String dateUnconfirmed

Body:
{
   "dateId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
   "spotId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
   "userId":"33ddab28-006b-4790-bf42-1832f90dc8d4",
   "aditionalpayload":{
      "startTime":"15:00:00",
      "confirmed":false,
      "id":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
      "spotId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
      "place":{
         "name":"Place Laureles",
         "lat":6.245887,
         "lon":-75.589868,
         "mapProviderId":"Place_Laureles"
      },
      "dateStamp":"2023-08-06",
      "host":{
         "hostId":"33ddab28-006b-4790-bf42-1832f90dc8d4",
         "hostName":"karol G"
      },
      "status":"SCHEDULED"
   }
}


Attributes: 
Operation String onlineResume

Body:
{
   "spotId":"f11495e4-3b25-4aa8-9b7d-33583b7242fe",
   "userId":"33ddab28-006b-4790-bf42-1832f90dc8d4"
}

Attributes:
Operation String onlineStop

Body:
{
   "spotId":"f11495e4-3b25-4aa8-9b7d-33583b7242fe",
   "userId":"33ddab28-006b-4790-bf42-1832f90dc8d4"
}

Attributes:
Operation String onlineFinalize

Body:
{
   "spotId":"f11495e4-3b25-4aa8-9b7d-33583b7242fe",
   "userId":"33ddab28-006b-4790-bf42-1832f90dc8d4"
}

Attributes:
Operation String dateConfirmed

Body:
{
   "dateId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
   "spotId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
   "userId":"33ddab28-006b-4790-bf42-1832f90dc8d4",
   "aditionalpayload":{
      "startTime":"15:00:00",
      "confirmed":true,
      "id":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
      "spotId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
      "place":{
         "name":"Place Laureles",
         "lat":6.245887,
         "lon":-75.589868,
         "mapProviderId":"Place_Laureles"
      },
      "dateStamp":"2023-08-06",
      "host":{
         "hostId":"33ddab28-006b-4790-bf42-1832f90dc8d4",
         "hostName":"karol G"
      },
      "status":"SCHEDULED"
   }
}

Attributes:
Operation String dateUnconfirmed

Body:
{
   "dateId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
   "spotId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
   "userId":"33ddab28-006b-4790-bf42-1832f90dc8d4",
   "aditionalpayload":{
      "startTime":"15:00:00",
      "confirmed":false,
      "id":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
      "spotId":"23b70dcd-14a9-4607-9740-7f6bf889e5de",
      "place":{
         "name":"Place Laureles",
         "lat":6.245887,
         "lon":-75.589868,
         "mapProviderId":"Place_Laureles"
      },
      "dateStamp":"2023-08-06",
      "host":{
         "hostId":"33ddab28-006b-4790-bf42-1832f90dc8d4",
         "hostName":"karol G"
      },
      "status":"SCHEDULED"
   }
}



---

{
   MessageId:74e41748-a83c-4076-8a83-4a56866888b6 
   Body:{"dateId":"23b70dcd-14a9-4607-9740-7f6bf889e5de","spotId":"f11495e4-3b25-4aa8-9b7d-33583b7242fe","userId":"33ddab28-006b-4790-bf42-1832f90dc8d4","aditionalpayload":{"startTime":"15:00:00","confirmed":true,"id":"23b70dcd-14a9-4607-9740-7f6bf889e5de","spotId":"f11495e4-3b25-4aa8-9b7d-33583b7242fe","place":{"name":"Place Laureles","lat":6.245887,"lon":-75.589868,"mapProviderId":"Place_Laureles"},"dateStamp":"2023-08-06","host":{"hostId":"33ddab28-006b-4790-bf42-1832f90dc8d4","hostName":"karol G"},"status":"SCHEDULED"}} 
   Attributes:map[
      ApproximateFirstReceiveTimestamp:1691281143321 
      ApproximateReceiveCount:13 
      SenderId:AIDAJQR6QDGQ7PATMSYEY 
      SentTimestamp:1691281143321
   ] 
   MessageAttributes:map[
      Operation:{
         StringValue:0xc00001e860 
         BinaryValue:[] 
         StringListValues:[] 
         BinaryListValues:[] 
         DataType:String
      }
   ] 
   EventSourceARN:arn:aws:sqs:us-east-2:732596568988:updateLocationDataSQS 
   EventSource:aws:sqs 
   AWSRegion:us-east-2
}
