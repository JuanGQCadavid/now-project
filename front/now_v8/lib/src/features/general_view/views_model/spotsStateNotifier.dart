// import 'dart:developer';

import 'package:dartz/dartz.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
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

class SpotsNotifer extends StateNotifier<List<Spot>> {
  final GeneralViewModel generalViewModel;
  final double threshold = 11;
  LatLng? lastPositionKnown;
  double? lastZoomKnown;

  SpotsNotifer({required this.generalViewModel}) : super([]) {
    refreshSpots();
  }

  void refreshSpots({LatLng? latLng, double? zoom}) async {
    List<Spot> spots =
        await generalViewModel.getSpots(centralPosition: latLng, zoom: zoom);

    // TODO: This should be a set
    print(
        "Before addAll state len -> ${state.length}  Spots len -> ${spots.length} total -> ${state.length + spots.length} ");
    spots.addAll(state);
    state = spots;
    print("After addAll state len -> ${state.length}");
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
    print(
        "User camera ${position.target.latitude} ${position.target.longitude}");

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
    print("HERE WE GOOO!");
    state = state.copyWith(status: MapStatus.movingStarted());
  }
}
