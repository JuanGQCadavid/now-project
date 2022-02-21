import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';
import 'package:now/core/services/filterService.dart';
import 'package:now/core/domain/models/spot.dart' as models;

import 'spots_granular_view.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/services.dart' show rootBundle;

class SpotGeneralView extends StatelessWidget {
  const SpotGeneralView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const _PageScafold();
  }
}

class _PageScafold extends StatelessWidget {
  const _PageScafold({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: const _Body(),
      appBar: AppBar(
        title: Center(
          child: Text("Ongoing events!"),
        ),
      ),
      floatingActionButton: FloatingActionButton(onPressed: () {
        Navigator.push(context,
            MaterialPageRoute(builder: (context) => const SpotGranularView()));
      }),
    );
  }
}

class _Body extends StatelessWidget {
  const _Body({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: _MapBody(),
      decoration: const BoxDecoration(
        color: Colors.blueGrey,
      ),
    );
  }
}

class _MapBody extends StatefulWidget {
  const _MapBody({Key? key}) : super(key: key);

  @override
  __MapBodyState createState() => __MapBodyState();
}

class __MapBodyState extends State<_MapBody> {
  late GoogleMapController mapController;
  final Map<String, Marker> _markers = {};
  late String _mapStyle;
  final LatLng _center = const LatLng(0, 0); //LatLng(6.246408, -75.590666);

  late FilterService filterService;

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

  void _onMapCreated(GoogleMapController controller) async {
    mapController = controller;
    mapController.setMapStyle(_mapStyle);

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
      markers: _markers.values.toSet(),
    );
  }
}
