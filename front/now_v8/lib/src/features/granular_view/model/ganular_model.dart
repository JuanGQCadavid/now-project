import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/state_response.dart';
import 'package:now_v8/src/features/granular_view/model/granular_spot.dart';

class GranularModel {
  final IFilterService filterService;
  final ILocationService locationService;
  final IKeyValueStorage sessionDatabase;

  final String searchSessionKey = "X-Now-Search-Session-Id";

  GranularModel({
    required this.filterService,
    required this.locationService,
    required this.sessionDatabase,
  });

  GranularSpot generateNewModel(int newIndex, List<LongSpot> actualList) {
    SpotWindow spotWindow = generateSpotWindow(newIndex, actualList);

    GranularSpot newModel = GranularSpot(
      spot: actualList[newIndex],
      window: spotWindow,
    );

    return newModel;
  }

  Future<List<LongSpot>> getSpots() async {
    LatLng userLocation = await locationService.getUserCurrentLocation();
    await sessionDatabase.doInit();

    print("Before calling Search Session key");
    String token = sessionDatabase.getValue(searchSessionKey);
    print("Token -> " + token);

    StateResponse<List<LongSpot>, String> filterResponse =
        await filterService.getLongSpotByProximityWithState(
            cpLat: userLocation.latitude, cpLng: userLocation.longitude, token: token);

    if(filterResponse.token.isEmpty || filterResponse.token != token) {
      print("Differente tokens");
      print("Before calling Search Session key");
      print("filterResponse.token -> " + filterResponse.token);
      print("Token -> " + token);
      sessionDatabase.save(filterResponse.token, searchSessionKey);

      print("Before calling Search Session key again");
      String token2 = sessionDatabase.getValue(searchSessionKey);
      print("token2 -> " + token2);
    } else {
      print("Same token as the one we use to call the service");
      print("filterResponse.token -> " + filterResponse.token);
      print("Token -> " + token);
    }

    return filterResponse.response;
  }

  // TODO: We could simplify a lot this process by usling a Linked list.
  SpotWindow generateSpotWindow(int middleSpotIndex, List<LongSpot> spotsList) {
    if (spotsList.isEmpty) {
      return SpotWindow(
        previousOne: "",
        actualOne: "",
        nextOne: "",
      );
    } else if (spotsList.length == 1) {
      return SpotWindow(
        previousOne: "",
        actualOne: spotsList[middleSpotIndex].eventInfo.name,
        nextOne: "",
      );
    } else if (spotsList.length == 2) {
      if (middleSpotIndex == (spotsList.length - 1)) {
        return SpotWindow(
          previousOne: spotsList[middleSpotIndex - 1].eventInfo.name,
          actualOne: spotsList[middleSpotIndex].eventInfo.name,
          nextOne: "",
        );
      } else {
        return SpotWindow(
          previousOne: "",
          actualOne: spotsList[middleSpotIndex].eventInfo.name,
          nextOne: spotsList[middleSpotIndex + 1].eventInfo.name,
        );
      }
    } else {
      if (middleSpotIndex == 0) {
        return SpotWindow(
          previousOne: "",//spotsList[spotsList.length - 1].eventInfo.name,
          actualOne: spotsList[middleSpotIndex].eventInfo.name,
          nextOne: spotsList[middleSpotIndex + 1].eventInfo.name,
        );
      } else if (middleSpotIndex == (spotsList.length - 1)) {
        return SpotWindow(
          previousOne: spotsList[middleSpotIndex - 1].eventInfo.name,
          actualOne: spotsList[middleSpotIndex].eventInfo.name,
          nextOne: "",
        );
      } else {
        return SpotWindow(
          previousOne: spotsList[middleSpotIndex - 1].eventInfo.name,
          actualOne: spotsList[middleSpotIndex].eventInfo.name,
          nextOne: spotsList[middleSpotIndex + 1].eventInfo.name,
        );
      }
    }
  }

  int findSpotIndex(String spotId, List<LongSpot> spotsList) {
    for (int index = 0; index < spotsList.length; index++) {
      LongSpot actualOne = spotsList[index];

      if (actualOne.eventInfo.id == spotId) {
        return index;
      }
    }
    return -1;
  }
}
