import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/features/general_view/views/generalView.dart';
import 'package:now_v8/src/features/granular_view/views/greanular_view.dart';

void main() => runApp(
      ProviderScope(
        child: MyApp(),
      ),
    );

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Flutter Demo',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: NowMapTest() //MyHomePage(title: 'Flutter Demo Home Page'),
        );
  }
}

class NowMapTest extends StatelessWidget {
  const NowMapTest({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(child: NowMapOld()),
    );
  }
}

class NowMapOld extends StatelessWidget {
  const NowMapOld({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    List<Spot> spot = [];

    spot.add(
      Spot.withOutSpotColors(
        principalTag: "",
        secondaryTags: [""],
        latLng: LatLng(6.241791, -75.595383),
        spotId: "1",
      ),
    );

    spot.add(
      Spot.withOutSpotColors(
        principalTag: "",
        secondaryTags: [""],
        latLng: LatLng(6.2386133, -75.5858772),
        spotId: "2",
      ),
    );

    return Container(
      height: 300,
      // width: 400,
      color: Colors.red,
      child: NowMap(spots: spot),
    );
  }
}

class NowMap extends StatefulWidget {
  final List<Spot> spots;
  final bool centerMapOnSpots;
  final bool blockMap;
  final double mapZoom;

  // Optional nulls
  final LatLng? userLocation;
  late LatLng? camaraPosition;

  // Internally
  late CameraPosition initialCameraPosition;
  final double mapPaddingOnCentered = 50;

  NowMap(
      {Key? key,
      this.spots = const [],
      this.centerMapOnSpots = true,
      this.blockMap = false,
      this.mapZoom = 14.5,
      this.camaraPosition,
      this.userLocation})
      : super(key: key) {
    if (camaraPosition == null) {
      if (spots.isNotEmpty && spots.length == 1) {
        camaraPosition = spots.first.latLng;
      } else {
        camaraPosition = const LatLng(0, 0);
      }
    }
    initialCameraPosition = CameraPosition(
      target: camaraPosition!,
      zoom: mapZoom,
    );
  }

  factory NowMap.fromFilteredSpots(FilteredSpots filteredSpots) {
    List<Spot> spots = [];

    filteredSpots.spots.forEach((spot) {
      spots.add(Spot(
        principalTag: spot.principalTag,
        secondaryTags: spot.secondaryTags,
        latLng: spot.latLng,
        spotId: spot.spotId,
        spotsColor: filteredSpots.tagsSelected.isEmpty
            ? spot.spotsColor
            : filteredSpots.onFilterColor,
      ));
    });

    return NowMap(spots: spots);
  }

  @override
  State<NowMap> createState() => _NowMapState();
}

class _NowMapState extends State<NowMap> {
  // Completer<GoogleMapController> _controller = Completer();
  // final Completer<GoogleMapController> _controller;
  late GoogleMapController googleMapController;
  void onMapCreated(GoogleMapController mapController) {
    googleMapController = mapController;

    if (widget.centerMapOnSpots && widget.spots.length > 1) {
      mapController.animateCamera(
        CameraUpdate.newLatLngBounds(
          getCameraLatLngBounds(widget.spots),
          widget.mapPaddingOnCentered,
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    Set<Marker> markers = Set();

    widget.spots.forEach((spot) {
      markers.add(
        Marker(
            markerId: MarkerId(spot.spotId),
            position: spot.latLng,
            visible: true,
            icon: spot.spotsColor.hue,
            infoWindow: InfoWindow(
              title: spot.principalTag,
            )),
      );
    });

    return GoogleMap(
      markers: markers,
      mapType: MapType.normal,
      zoomControlsEnabled: false,
      initialCameraPosition: widget.initialCameraPosition,
      mapToolbarEnabled: false,
      myLocationButtonEnabled: false,
      myLocationEnabled: true,
      onMapCreated: onMapCreated,
      // padding: const EdgeInsets.only(bottom: 65, left: 15),
      // cameraTargetBounds: ,
    );
  }

  LatLngBounds getCameraLatLngBounds(List<Spot> spots) {
    Spot spot = spots.first;
    double down, up, left, rigth;
    down = up = spot.latLng.latitude;
    left = rigth = spot.latLng.longitude;

    for (final spot in spots) {
      LatLng spotLatLng = spot.latLng;

      if (spotLatLng.latitude > up) up = spotLatLng.latitude;
      if (spotLatLng.latitude < down) down = spotLatLng.latitude;

      if (spotLatLng.latitude < left) left = spotLatLng.longitude;
      if (spotLatLng.latitude > rigth) rigth = spotLatLng.longitude;
    }

    return LatLngBounds(
        southwest: LatLng(down, left), northeast: LatLng(up, rigth));
  }
}
