import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:location/location.dart';
import 'package:latlong2/latlong.dart';

import '../../../core/domain/models/spot.dart';
import '../../../core/domain/ports/i_now_filter_service.dart';
import 'filter_state.dart';

class FilterChangeNotifier extends ChangeNotifier {
  late Map<String, Spot> _spotData;
  late Map<String, Marker> _markersData;
  late FilterState _filterState;
  final INowFIlterService filterService;

  FilterChangeNotifier({required this.filterService}) {
    _spotData = {};
    _markersData = {};
    _filterState = const FilterState.initial();
  }

  set _spots(Map<String, Spot> locations) {
    _spotData = locations;
    notifyListeners();
  }

  Map<String, Spot> get spotData {
    return _spotData;
  }

  FilterState get filterState {
    return _filterState;
  }

  set __filterState(FilterState newState) {
    _filterState = newState;
  }

  set _markers(Map<String, Marker> newMarkers) {
    _markersData = newMarkers;
    notifyListeners();
  }

  Map<String, Marker> get markers {
    return _markersData;
  }

  void fetchSpotsFrom(LatLng latLng) async {
    final backendResponse = await filterService.fetchByProximity(
        lat: latLng.latitude, lon: latLng.longitude, radious: 1);

    backendResponse.fold(
      (l) => __filterState = const FilterState.onError(),
      (response) async {
        print(response);
      },
    );
  }
}
