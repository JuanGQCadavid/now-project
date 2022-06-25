import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/state_response.dart';
import 'package:now_v8/src/features/granular_view/model/granular_spot.dart';

class GranularModel {
  final IFilterService filterService;
  final ILocationService locationService;

  GranularModel({required this.filterService, required this.locationService});

  GranularSpot generateNewModel(int newIndex, List<LongSpot> actualList) {
    SpotWindow spotWindow = generateSpotWindow(newIndex, actualList);

    GranularSpot newModel = GranularSpot(
      spot: actualList[newIndex],
      window: spotWindow,
    );

    return newModel;
  }

  Future<List<LongSpot>> getSpots() async {
    LatLng userLocation = locationService.getUserCurrentLocation();

    StateResponse<List<LongSpot>, String> filterResponse =
        await filterService.getByProximityWithState(
            cpLat: userLocation.latitude, cpLng: userLocation.longitude);

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
          previousOne: spotsList[spotsList.length - 1].eventInfo.name,
          actualOne: spotsList[middleSpotIndex].eventInfo.name,
          nextOne: spotsList[middleSpotIndex + 1].eventInfo.name,
        );
      } else if (middleSpotIndex == (spotsList.length - 1)) {
        return SpotWindow(
          previousOne: spotsList[middleSpotIndex - 1].eventInfo.name,
          actualOne: spotsList[middleSpotIndex].eventInfo.name,
          nextOne: spotsList[0].eventInfo.name,
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
