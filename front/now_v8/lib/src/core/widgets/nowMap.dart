import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/services/core/providers.dart';

// Statefull with consumer
const LatLng _empty =  LatLng(-100000, -1000000);
class NowMapV2 extends ConsumerStatefulWidget {
  final List<Spot> spots;
  final bool centerMapOnSpots;
  final bool blockMap;
  final double mapZoom;
  final bool myLocationButtonEnabled;
  final bool includeUserLocation;
  late LatLng? camaraPosition;
  final Completer<GoogleMapController> mapController;

  // Internally
  final double mapPaddingOnCentered = 50;

  NowMapV2(
      {Key? key,
      this.spots = const [],
      this.centerMapOnSpots = true,
      this.blockMap = false,
      this.mapZoom = 14.5,
      this.myLocationButtonEnabled = true,
      this.includeUserLocation = true,
      this.camaraPosition,
      required this.mapController})
      : super(key: key);

  factory NowMapV2.fromFilteredSpots(
    FilteredSpots filteredSpots,
    Completer<GoogleMapController> mapController, {
    bool centerMapOnSpots = true,
    bool blockMap = false,
    double mapZoom = 14.5,
    bool myLocationButtonEnabled = false,
    bool includeUserLocation = true,
    LatLng? camaraPosition,
  }) {
    List<Spot> spots = [];

    for (var spot in filteredSpots.spots) {
      spots.add(Spot(
        principalTag: spot.principalTag,
        secondaryTags: spot.secondaryTags,
        latLng: spot.latLng,
        spotId: spot.spotId,
        date: spot.date,
        spotsColor: filteredSpots.tagsSelected.isEmpty
            ? spot.spotsColor
            : filteredSpots.onFilterColor,
      ));
    }

    return NowMapV2(
      spots: spots,
      centerMapOnSpots: centerMapOnSpots,
      blockMap: blockMap,
      mapZoom: mapZoom,
      myLocationButtonEnabled: myLocationButtonEnabled,
      camaraPosition: camaraPosition,
      includeUserLocation: includeUserLocation,
      mapController: mapController,
    );
  }

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _NowMapV2State();
}

class _NowMapV2State extends ConsumerState<NowMapV2> {
  late CameraPosition initialCameraPosition;
  late GoogleMapController _mapController;

  void onMapCreated(GoogleMapController mapController, {LatLng? userLocation}) {
    widget.mapController.complete(mapController);

    setState(() {
      _mapController = mapController;
    });

    LatLngBounds bounds;
    print("Hello?");
    if (widget.centerMapOnSpots && widget.spots.isNotEmpty) {
      if (userLocation != null) {
        print("Hello?");
        bounds =
            MapUtilities.getCameraLatLngBounds(widget.spots, userLocation: userLocation);
      } else {
        bounds = MapUtilities.getCameraLatLngBounds(widget.spots);
      }

      mapController.animateCamera(
        CameraUpdate.newLatLngBounds(
          bounds,
          widget.mapPaddingOnCentered,
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final locationService = ref.read(locationServiceProvider);
    Set<Marker> markers = Set();

    widget.spots.forEach((spot) {
      markers.add(
        Marker(
            markerId: MarkerId(spot.spotId),
            position: spot.latLng,
            visible: true,
            icon: spot.spotsColor.hue,
            infoWindow: InfoWindow(
              title: "${spot.date}",
            )),
      );
    });

    markers.add(Marker(markerId: MarkerId("mydudeIamHere"), position: LatLng(6.251723158537203, -75.59277109801769), visible: true, icon: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueGreen) ));

    if (widget.includeUserLocation) {
      return FutureBuilder(
        future: locationService.getUserCurrentLocation(),
        builder: (context, AsyncSnapshot<LatLng> snapshot) {
          if (snapshot.hasData) {
            print("snapshot.hasData");
            print(snapshot.data!);
            if (widget.camaraPosition == null) {
              if (widget.includeUserLocation) {
                widget.camaraPosition = snapshot.data!;
              } else if (widget.spots.isNotEmpty && widget.spots.length == 1) {
                widget.camaraPosition = widget.spots.first.latLng;
              } else {
                widget.camaraPosition = const LatLng(0, 0);
              }
            }

            initialCameraPosition = CameraPosition(
              target: widget.camaraPosition!,
              zoom: widget.mapZoom,
            );

            return  GoogleMapLocal(
              blockMap: widget.blockMap,
              markers: markers,
              initialCameraPosition: initialCameraPosition,
              myLocationButtonEnabled: widget.myLocationButtonEnabled,
              onMapCreated: onMapCreated,
              userLocation: snapshot.data!,
            );
          } else if (snapshot.hasError) {
            print("snapshot.hasError");
            return const Text("Ops we are having problems to didplay the map");
          } else {
            print("snapshot loading");
            return const Center(child: CircularProgressIndicator());
          }
        },
      );
    } else {
      if (widget.camaraPosition == null) {
        if (widget.spots.isNotEmpty && widget.spots.length == 1) {
          widget.camaraPosition = widget.spots.first.latLng;
        } else {
          widget.camaraPosition = const LatLng(0, 0);
        }
      }

      initialCameraPosition = CameraPosition(
        target: widget.camaraPosition!,
        zoom: widget.mapZoom,
      );

      return GoogleMapLocal(
        blockMap: widget.blockMap,
        markers: markers,
        initialCameraPosition: initialCameraPosition,
        myLocationButtonEnabled: widget.myLocationButtonEnabled,
        onMapCreated: onMapCreated,
      );
    }
  }
}

class GoogleMapLocal extends StatelessWidget {
  final Set<Marker> markers;
  final CameraPosition initialCameraPosition;
  final bool myLocationButtonEnabled;
  final Function(GoogleMapController, {LatLng? userLocation}) onMapCreated;
  final bool blockMap;
  final LatLng userLocation;
  final MinMaxZoomPreference defaulMinMaxZoom = const MinMaxZoomPreference(11.5, 100);

  late LatLng lastCamare;
  late double lastZoom;

  GoogleMapLocal({
    super.key, 
    required this.markers, 
    required this.initialCameraPosition, 
    required this.myLocationButtonEnabled, 
    required this.onMapCreated, 
    required this.blockMap,
    this.userLocation = _empty
    }
  );

  @override
  Widget build(BuildContext context) {
    return GoogleMap(
        markers: markers,
        mapType: MapType.normal,
        zoomControlsEnabled: false,
        initialCameraPosition: initialCameraPosition,
        mapToolbarEnabled: false,
        myLocationButtonEnabled: myLocationButtonEnabled,
        myLocationEnabled: true,
        onMapCreated: (controller) {
          onMapCreated(controller, userLocation: userLocation == _empty ? null : userLocation );
        },
        scrollGesturesEnabled: !blockMap,
        zoomGesturesEnabled: !blockMap,
        minMaxZoomPreference: !blockMap ? defaulMinMaxZoom : MinMaxZoomPreference.unbounded,
        onCameraMove: ((position) {
          print("User camera ${position.target.latitude} ${position.target.longitude}");
          lastCamare = position.target;
          lastZoom = position.zoom;
        }),
        onCameraIdle: () {
          print("LAST POSITION  ${lastCamare.latitude} ${lastCamare.longitude} ZOOM - ${lastZoom}");
        },
        onCameraMoveStarted: () {
          print("HERE WE GOOO!");
        },
      );
  }
}

class MapUtilities {
  static LatLngBounds getCameraLatLngBounds(List<Spot> spots,
      {LatLng userLocation = const LatLng(0, 0)}) {
    Spot spot = spots.first;
    double down, up, left, rigth;
    down = up = spot.latLng.latitude;
    left = rigth = spot.latLng.longitude;

    List<LatLng> spotsToCheck = List.from(spots.map((spot) => spot.latLng));
    if (userLocation.latitude != 0 && userLocation.longitude != 0) {
      spotsToCheck.add(userLocation);
    }

    for (final spotLatLng in spotsToCheck) {
      if (spotLatLng.latitude > up) up = spotLatLng.latitude;
      if (spotLatLng.latitude < down) down = spotLatLng.latitude;

      if (spotLatLng.latitude < left) left = spotLatLng.longitude;
      if (spotLatLng.latitude > rigth) rigth = spotLatLng.longitude;
    }

    return LatLngBounds(
        southwest: LatLng(down, left), northeast: LatLng(up, rigth));
  }
}
