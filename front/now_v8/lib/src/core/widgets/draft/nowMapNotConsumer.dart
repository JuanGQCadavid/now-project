import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';

class NowMap extends StatefulWidget {
  final List<Spot> spots;
  final bool centerMapOnSpots;
  final bool blockMap;
  final double mapZoom;
  final bool myLocationButtonEnabled;

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
      this.myLocationButtonEnabled = true,
      this.camaraPosition,
      this.userLocation})
      : super(key: key) {
    if (camaraPosition == null) {
      if (spots.isNotEmpty && spots.length == 1) {
        camaraPosition = spots.first.latLng;
      } else {
        if (userLocation != null) {
          camaraPosition = userLocation;
        } else {
          camaraPosition = const LatLng(0, 0);
        }
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
        date: spot.date,
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
  late GoogleMapController googleMapController;

  void onMapCreated(GoogleMapController mapController) {
    googleMapController = mapController;

    if (widget.centerMapOnSpots && widget.spots.length > 1) {
      LatLngBounds bounds;
      if (widget.userLocation != null) {
        bounds = getCameraLatLngBounds(widget.spots,
            userLocation: widget.userLocation!);
      } else {
        bounds = getCameraLatLngBounds(widget.spots);
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
      myLocationButtonEnabled: widget.myLocationButtonEnabled,
      myLocationEnabled: true,
      onMapCreated: onMapCreated,
    );
  }

  LatLngBounds getCameraLatLngBounds(List<Spot> spots,
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
