import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/services/core/providers.dart';

// Statefull with consumer
const LatLng _empty = LatLng(-100000, -1000000);
const ClusterManagerId clusterId =
    ClusterManagerId("general_cluster_manager_id");

class NowMapV2 extends ConsumerStatefulWidget {
  final List<Spot> spots;
  final bool centerMapOnSpots;
  final bool blockMap;
  final double mapZoom;
  final bool myLocationButtonEnabled;
  final bool includeUserLocation;
  final EdgeInsets padding;
  late LatLng? camaraPosition;
  final Completer<GoogleMapController> mapController;
  final Function(CameraPosition)? onCameraMove;
  final Function()? onCameraIdle;
  final Function()? onCameraMoveStarted;
  final Function(GoogleMapController mapController)? onMapCreated;

  // Internally
  final double mapPaddingOnCentered = 50;

  NowMapV2({
    super.key,
    this.spots = const [],
    this.centerMapOnSpots = true,
    this.blockMap = false,
    this.mapZoom = 14.5,
    this.myLocationButtonEnabled = true,
    this.includeUserLocation = true,
    this.camaraPosition,
    this.onCameraIdle,
    this.onCameraMove,
    this.onCameraMoveStarted,
    this.onMapCreated,
    required this.mapController,
    this.padding = EdgeInsets.zero,
  });

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
    if (widget.centerMapOnSpots && widget.spots.isNotEmpty) {
      if (userLocation != null) {
        bounds = MapUtilities.getCameraLatLngBounds(widget.spots,
            userLocation: userLocation);
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

    var callBack =
        widget.onMapCreated ?? (GoogleMapController mapController) {};
    callBack(mapController);
  }

  Set<Marker> generateMarkers(List<Spot> spots) {
    Set<Marker> markers = {};

    for (var spot in spots) {
      markers.add(
        Marker(
          markerId: MarkerId(spot.spotId),
          position: spot.latLng,
          visible: true,
          icon: spot.spotsColor.hue,
          infoWindow: InfoWindow(
            title: "${spot.date}",
          ),
          onTap: () {
            // TODO: Here we could add the transition to something else
            print("${spot.spotId} were tapped");
          },
        ),
      );
    }
    return markers;
  }

  @override
  Widget build(BuildContext context) {
    final locationService = ref.read(locationServiceProvider);
    Set<Marker> markers = generateMarkers(widget.spots);

    if (widget.camaraPosition == null) {
      if (widget.spots.isNotEmpty && widget.spots.length == 1) {
        widget.camaraPosition = widget.spots.first.latLng;
      } else {
        widget.camaraPosition = const LatLng(0, 0);
      }
    }

    if (widget.includeUserLocation) {
      return FutureBuilder(
        future: locationService.getUserCurrentLocation(),
        builder: (context, AsyncSnapshot<LatLng> snapshot) {
          if (snapshot.hasData) {
            widget.camaraPosition = snapshot.data!;

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
              userLocation: snapshot.data!,
              onCameraMove: widget.onCameraMove,
              onCameraIdle: widget.onCameraIdle,
              onCameraMoveStarted: widget.onCameraMoveStarted,
              padding: widget.padding,
            );
          } else if (snapshot.hasError) {
            return const Text("Ops we are having problems to didplay the map");
          } else {
            return const Center(child: CircularProgressIndicator());
          }
        },
      );
    } else {
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
        onCameraMove: widget.onCameraMove,
        onCameraIdle: widget.onCameraIdle,
        onCameraMoveStarted: widget.onCameraMoveStarted,
        padding: widget.padding,
      );
    }
  }
}

class GoogleMapLocal extends StatelessWidget {
  final bool blockMap;
  late double lastZoom;
  late LatLng lastCamare;
  final LatLng userLocation;
  final Set<Marker> markers;
  final bool myLocationButtonEnabled;
  final CameraPosition initialCameraPosition;
  final Function(GoogleMapController, {LatLng? userLocation}) onMapCreated;
  final Function(CameraPosition)? onCameraMove;
  final Function()? onCameraIdle;
  final Function()? onCameraMoveStarted;
  final EdgeInsets padding;

  final MinMaxZoomPreference defaulMinMaxZoom =
      const MinMaxZoomPreference(11.5, 100);

  GoogleMapLocal({
    super.key,
    required this.markers,
    required this.initialCameraPosition,
    required this.myLocationButtonEnabled,
    required this.onMapCreated,
    required this.blockMap,
    this.onCameraMove,
    this.onCameraIdle,
    this.onCameraMoveStarted,
    this.userLocation = _empty,
    this.padding = EdgeInsets.zero,
  });

  @override
  Widget build(BuildContext context) {
    Map<MarkerId, Marker> newMarkers = <MarkerId, Marker>{};

    ClusterManager general = ClusterManager(
      clusterManagerId: clusterId,
      onClusterTap: (argument) {
        print("Cluster tab");
      },
    );

    Map<ClusterManagerId, ClusterManager> clusterManagers =
        <ClusterManagerId, ClusterManager>{};

    clusterManagers[general.clusterManagerId] = general;

    for (var i = 0; i < markers.length; i++) {
      var actualMarker = markers.elementAt(i);

      if (markers.length > 1) {
        newMarkers[actualMarker.markerId] = actualMarker.copyWith(
            clusterManagerIdParam: general.clusterManagerId);
      } else {
        newMarkers[actualMarker.markerId] = actualMarker;
      }
    }

    return GoogleMap(
      markers: Set<Marker>.of(newMarkers.values),
      padding: padding,
      mapType: MapType.normal,
      zoomControlsEnabled: false,
      initialCameraPosition: initialCameraPosition,
      mapToolbarEnabled: false,
      myLocationButtonEnabled: myLocationButtonEnabled,
      myLocationEnabled: true,
      onMapCreated: (controller) {
        onMapCreated(controller,
            userLocation: userLocation == _empty ? null : userLocation);
      },
      scrollGesturesEnabled: !blockMap,
      zoomGesturesEnabled: !blockMap,
      minMaxZoomPreference:
          !blockMap ? defaulMinMaxZoom : MinMaxZoomPreference.unbounded,
      onCameraMove: onCameraMove,
      onCameraIdle: onCameraIdle,
      onCameraMoveStarted: onCameraMoveStarted,
      clusterManagers: Set<ClusterManager>.of(clusterManagers.values),
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
