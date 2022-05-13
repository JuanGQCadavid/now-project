


import 'package:flutter_map/plugin_api.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';
import 'package:geolocator/geolocator.dart';
import 'package:now_v2/features/filters/presentation/marker.dart';

class MapStateNotifier extends StateNotifier<MapState> {
  MapStateNotifier() : super( MapState([], false));

  void adduserLocation(){

    Geolocator.getLastKnownPosition().then((value) {
      state.containUserLocation = true;
      state.markers.add(
        Marker(
          point: LatLng(
            value?.latitude ?? 0.0 , 
            value?.longitude ?? 0.0
            ), 
            builder: (context)=> const HumanMarkerContent(),height: 10, width: 10,));
      });

  }

  void setMarkers(List<Marker> markers) {
    state.markers = markers;
  }

}

class MapState {
  List<Marker> markers;
  bool containUserLocation;
  late MapController mapController;

  MapState(this.markers, this.containUserLocation) {
    this.mapController = MapController();
  }
}