import 'package:flutter/material.dart';

import 'package:now/core/services/filterService.dart';
import 'package:location/location.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/services.dart' show rootBundle;

class MapBody extends StatefulWidget {
  const MapBody({Key? key}) : super(key: key);

  @override
  _MapBodyState createState() => _MapBodyState();
}

class _MapBodyState extends State<MapBody> {
  late GoogleMapController mapController;
  final Map<String, Marker> _markers = {};
  late String _mapStyle;
  final LatLng _center = const LatLng(0, 0);

  late FilterService filterService;
  final Location _location = Location();

  @override
  void initState() {
    super.initState();

    filterService = FilterService();

    rootBundle
        .loadString('assets/maps/mapStyle.json', cache: true)
        .then((string) {
      _mapStyle = string;
    });
  }

  void findMe() async {
    final locationData = await _location.getLocation();
    mapController.animateCamera(CameraUpdate.newCameraPosition(CameraPosition(
        target: LatLng(
          locationData.latitude ?? 0.0,
          locationData.longitude ?? 0.0,
        ),
        zoom: 15)));
  }

  void _onMapCreated(GoogleMapController controller) async {
    mapController = controller;
    mapController.setMapStyle(_mapStyle);

    final locationData = await _location.getLocation();

    findMe();

    final backendResponse =
        await filterService.fetchByProximity(lat: 6.246944, lon: -75.586930);

    backendResponse.fold((l) => null, (response) async {
      final BitmapDescriptor pinLocation =
          await BitmapDescriptor.fromAssetImage(
              ImageConfiguration(devicePixelRatio: 2.5),
              'assets/custo_marker.png');

      setState(() {
        _markers.clear();
        for (final spot in response.places) {
          final marker = Marker(
            icon: pinLocation,
            markerId: MarkerId(spot.eventInfo.id),
            position: LatLng(spot.placeInfo.lat, spot.placeInfo.lon),
            infoWindow: InfoWindow(
                title: spot.eventInfo.name, snippet: spot.eventInfo.emoji),
          );
          _markers[spot.eventInfo.id] = marker;
        }
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return GoogleMap(
      onMapCreated: _onMapCreated,
      initialCameraPosition: CameraPosition(
        target: _center,
        zoom: 2, //20.0,
      ),
      mapToolbarEnabled: false,
      myLocationButtonEnabled: false,
      liteModeEnabled: false,
      zoomControlsEnabled: false,
      minMaxZoomPreference: const MinMaxZoomPreference(15, 30),
      markers: _markers.values.toSet(),
    );
  }
}
