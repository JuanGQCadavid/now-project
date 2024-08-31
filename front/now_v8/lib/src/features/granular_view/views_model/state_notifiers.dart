import 'dart:async';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/location_service.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/granular_view/model/ganular_model.dart';
import 'package:now_v8/src/features/granular_view/model/granular_spot.dart';
import 'package:now_v8/src/features/granular_view/views_model/providers.dart';

class DetailedSpotsState extends StateNotifier<List<LongSpot>> {
  final GranularModel granularModel;

  DetailedSpotsState({
    required this.granularModel,
  }) : super([]);

  Future<List<LongSpot>> refreshSpots() async {
    List<LongSpot> spots = await granularModel.getSpots();

    for (LongSpot spotState in state) {
      bool alreadySaved = false;
      for (LongSpot spotReturned in spots) {
        if (spotState.eventInfo.id == spotReturned.eventInfo.id) {
          alreadySaved = true;
        }
      }
      if (!alreadySaved) {
        spots.add(spotState);
      }
    }
    state = spots;

    return state;
  }
}

// TODO: We are not thinking in the case that there is only 1, 2 or not events.
class OnSpotState extends StateNotifier<GranularSpot> {
  late List<LongSpot> actualList;
  final GranularModel granularModel;
  final ILocationService locationService;
  final StateNotifierProviderRef ref;

  OnSpotState({
    this.actualList = const [],
    required this.granularModel,
    required this.ref,
    required this.locationService,
  }) : super(GranularSpot.empty()) {
    if (actualList.isNotEmpty) {
      state = granularModel.generateNewModel(0, actualList);
    }
  }

  void previousOne(Completer<GoogleMapController> mapController) {
    moveWindow(mapController, backward: true);
  }

  void nextOne(Completer<GoogleMapController> mapController) async {
    moveWindow(mapController, forward: true);
  }

  void moveWindow(
    Completer<GoogleMapController> mapController, {
    bool forward = false,
    bool backward = false,
  }) {
    int spotIndex =
        granularModel.findSpotIndex(state.spot.eventInfo.id, actualList);

    int newIndex = 0;

    if (forward) {
      if ((spotIndex + 1) > (actualList.length - 1)) {
        newIndex = spotIndex;
      } else {
        newIndex = spotIndex + 1;
      }
    } else if (backward) {
      if ((spotIndex - 1) >= 0) {
        newIndex = spotIndex - 1;
      } else {
        newIndex = actualList.length - 1;
      }
    }

    GranularSpot newState =
        granularModel.generateNewModel(newIndex, actualList);

    animateMapOnSpotChange(newState.spot, mapController, locationService);
    state = newState;
  }

  void refresh(Completer<GoogleMapController> mapController) async {
    int newIndex = 0;

    List<LongSpot> newSpots =
        await ref.read(detailedSpotProvider.notifier).refreshSpots();

    animateMapOnSpotChange(newSpots[newIndex], mapController, locationService);
  }

  void animateMapOnSpotChange(
    LongSpot spot,
    Completer<GoogleMapController> mapController,
    ILocationService locationService,
  ) async {
    GoogleMapController mapControllerFuture = await mapController.future;
    LatLng location = await locationService.getUserCurrentLocation();
    LatLngBounds bounds = MapUtilities.getCameraLatLngBounds([
      Spot.fromLongSpot(spot),
    ], userLocation: location);
    mapControllerFuture.animateCamera(
      CameraUpdate.newLatLngBounds(
        bounds,
        50,
      ),
    );
  }
}
