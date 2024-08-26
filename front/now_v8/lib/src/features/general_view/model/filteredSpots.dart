import 'package:flutter/foundation.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';

part 'filteredSpots.freezed.dart';

class FilteredSpots {
  late Set<String> tagsSelected;
  late Set<String> tagsOff;
  late List<Spot> spots;
  late SpotsColors onFilterColor;

  FilteredSpots({
    required this.spots,
    required this.tagsOff,
    required this.tagsSelected,
    required this.onFilterColor,
  });

  FilteredSpots.empty() {
    tagsOff = {};
    spots = [];
    tagsSelected = {};
    onFilterColor = const SpotsColors.empty();
  }
}

@freezed
class MapStatus with _$MapStatus {
  factory MapStatus.movingOnMap() = MovingOnMap;
  factory MapStatus.movingIdle() = MovingIdle;
  factory MapStatus.movingStarted() = MovingStarted;
}

final MapState emptyMapState = MapState(
    lastPositionKnowed: const LatLng(0, 0),
    zoom: 0,
    status: MapStatus.movingIdle());

@freezed
class MapState with _$MapState {
  const factory MapState({
    required LatLng lastPositionKnowed,
    required double zoom,
    required MapStatus status,
  }) = _MapState;
}
