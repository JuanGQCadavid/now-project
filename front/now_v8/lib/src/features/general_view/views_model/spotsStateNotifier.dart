import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/features/general_view/model/general_view_model.dart';

/**
 * We are going to have singles state notifiers
 * where we could only set the state, there is not need 
 * to call listeners, this will do the job for us.
 * 
 * The state is inmutable so we need to create new states
 * and assiigned it to the state.
 */

class SpotsNotifer extends StateNotifier<Map<String, Spot>> {
  final GeneralViewModel generalViewModel;
  final double threshold = 11;
  LatLng? lastPositionKnown;
  double? lastZoomKnown;

  SpotsNotifer({required this.generalViewModel}) : super({}) {
    refreshSpots();
  }

  void refreshSpots({LatLng? latLng, double? zoom}) async {
    List<Spot> spots = await generalViewModel.getSpots(
      centralPosition: latLng,
      zoom: zoom,
    );
    Map<String, Spot> newState = Map<String, Spot>.from(state);

    for (var newSpot in spots) {
      newState[newSpot.spotId] = newSpot;
    }
    state = newState;
  }
}

class TagsNotifier extends StateNotifier<Set<String>> {
  TagsNotifier() : super({});

  void tagSelected(String tag) {
    Set<String> newState = Set.from(state);
    if (state.contains(tag)) {
      newState.remove(tag);
    } else {
      newState.add(tag);
    }

    state = newState;
  }

  void cleanTags() {
    state = {};
  }
}

class MapInteractions extends StateNotifier<MapState> {
  MapInteractions() : super(emptyMapState);

  void onCameraMove(CameraPosition position) {
    var newState = state.copyWith(
        lastPositionKnowed: position.target,
        zoom: position.zoom,
        status: MapStatus.movingOnMap());
    state = newState;
  }

  void onCameraIdle() {
    print(
        "LAST POSITION  ${state.lastPositionKnowed.latitude} ${state.lastPositionKnowed.longitude} ZOOM - ${state.zoom}");
    state = state.copyWith(status: MapStatus.movingIdle());
  }

  void onCameraMoveStarted() {
    state = state.copyWith(status: MapStatus.movingStarted());
  }
}
