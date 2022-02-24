import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:location/location.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/services.dart' show rootBundle;
import 'package:now/features/filters/application/filter_providers.dart';
import 'package:now/features/filters/application/filter_state.dart';

class MapBody extends ConsumerStatefulWidget {
  const MapBody({Key? key}) : super(key: key);

  @override
  _MapBodyState createState() => _MapBodyState();
}

class _MapBodyState extends ConsumerState<MapBody> {
  late GoogleMapController mapController;
  final Map<String, Marker> _markers = {};
  late String _mapStyle;
  final LatLng _center = const LatLng(0, 0);

  final Location _location = Location();

  @override
  void initState() {
    super.initState();

    WidgetsBinding.instance?.addPostFrameCallback((_) {
      final provider = ref.read(filterNotifierProvier);
      _location.getLocation().then(
            (value) => provider.fetchSpotsFrom(
              LatLng(
                value.latitude ?? 0.0,
                value.longitude ?? 0.0,
              ),
            ),
          );
    });

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
    findMe();
  }

  @override
  Widget build(BuildContext context) {
    final provider = ref.read(filterNotifierProvier);
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
      markers: provider.markers.values.toSet(),
    );
  }
}
