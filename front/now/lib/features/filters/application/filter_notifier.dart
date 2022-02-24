import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:location/location.dart';
import 'package:now/core/domain/models/spot.dart';
import 'package:now/core/domain/ports/i_now_filter_service.dart';
import 'package:now/core/presentation/utils/marksIcons.dart';
import 'package:now/features/filters/application/filter_state.dart';

class FilterChangeNotifier extends ChangeNotifier {
  late Map<String, Spot> _spotData;
  late Map<String, Marker> _markersData;
  late FilterState _filterState;
  final Location _location = Location();

  final INowFIlterService filterService;

  FilterChangeNotifier({required this.filterService}) {
    _spotData = {};
    _markersData = {};
    _filterState = const FilterState.initial();
  }

  set _spots(Map<String, Spot> locations) {
    _spotData = locations;
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
        var markerGen = MarkerGenerator(100);
        var icon = await markerGen.createBitmapDescriptorFromIconData(
            Icons.ac_unit, Colors.blue, Colors.white, Colors.yellow);

        Map<String, Spot> newSpotData = {};
        Map<String, Marker> newMarkersData = {};
        for (final spot in response.places) {
          final marker = Marker(
            icon: icon,
            markerId: MarkerId(spot.eventInfo.id),
            position: LatLng(spot.placeInfo.lat, spot.placeInfo.lon),
            infoWindow: InfoWindow(
                title: spot.eventInfo.name, snippet: spot.eventInfo.emoji),
          );
          newSpotData[spot.eventInfo.id] = spot;
          newMarkersData[spot.eventInfo.id] = marker;
        }

        _markersData = newMarkersData;
        _spotData = newSpotData;
      },
    );

    notifyListeners();
  }
}
