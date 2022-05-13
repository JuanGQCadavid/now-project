import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:now_v2/features/filters/presentation/marker.dart';

import '../../../core/domain/models/spot.dart';
import '../../../core/domain/ports/i_now_filter_service.dart';
import 'filter_state.dart';

class FilterChangeNotifier extends ChangeNotifier {
  late Map<String, Spot> _spotData;
  late Map<String, Marker> markersData;
  late FilterState _filterState;
  final INowFIlterService filterService;
  late MapController mapController;

  FilterChangeNotifier({required this.filterService}) {
    _spotData = {};
    markersData = {};
    _filterState = const FilterState.initial();
  }

  set _spots(Map<String, Spot> locations) {
    _spotData = locations;
    //notifyListeners();
  }

  Map<String, Spot> get spotData {
    return _spotData;
  }

  FilterState get filterState {
    return _filterState;
  }

  set filterState(FilterState newState) {
    _filterState = newState;
  }

  set markers(Map<String, Marker> newMarkers) {
    markersData = newMarkers;
  }

  Map<String, Marker> get markers {
    return markersData;
  }

  void fetchSpotsFrom(LatLng latLng) async {
    final backendResponse = await filterService.fetchByProximity(
        lat: latLng.latitude, lon: latLng.longitude, radious: 1);

    backendResponse.fold(
      (l) => filterState = const FilterState.onError(),
      (response) async {
        Map<String, Spot> newSpotData = {};
        Map<String, Marker> newMarkersData = {};
        for (final spot in response.places) {
          final marker = Marker(
            width: 50.0,
            height: 50.0,
            point: LatLng(spot.placeInfo.lat, spot.placeInfo.lon),
            builder: (ctx) => SpotMarker(
              spotData: spot,
            ),
          );
          newSpotData[spot.eventInfo.id] = spot;
          newMarkersData[spot.eventInfo.id] = marker;
        }

        markersData = newMarkersData;
        _spotData = newSpotData;
      },
    );
    notifyListeners();
  }
}
