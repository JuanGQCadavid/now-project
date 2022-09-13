import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:location/location.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/services/core/providers.dart';


// Statefull with consumer
class NowMapV2 extends ConsumerStatefulWidget {
  final List<Spot> spots;
  final bool centerMapOnSpots;
  final bool blockMap;
  final double mapZoom;
  final bool myLocationButtonEnabled;
  final bool includeUserLocation;
  late LatLng? camaraPosition;

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
      this.camaraPosition})
      : super(key: key){
        print("on NowMap");
      }

  factory NowMapV2.fromFilteredSpots(
    FilteredSpots filteredSpots, {
    bool centerMapOnSpots = true,
    bool blockMap = false,
    double mapZoom = 14.5,
    bool myLocationButtonEnabled = false,
    bool includeUserLocation = true,
    LatLng? camaraPosition,
  }) {
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

    return NowMapV2(
      spots: spots,
      centerMapOnSpots: centerMapOnSpots,
      blockMap: blockMap,
      mapZoom: mapZoom,
      myLocationButtonEnabled: myLocationButtonEnabled,
      camaraPosition: camaraPosition,
      includeUserLocation: includeUserLocation,
    );
  }

  @override
  ConsumerState<ConsumerStatefulWidget> createState() => _NowMapV2State();
}

class _NowMapV2State extends ConsumerState<NowMapV2> {
  late GoogleMapController googleMapController;
  late CameraPosition initialCameraPosition;

  void onMapCreated(GoogleMapController mapController, {LatLng? userLocation}) {
    googleMapController = mapController;
    LatLngBounds bounds;
    print("Hello?");
    if (widget.centerMapOnSpots && widget.spots.isNotEmpty) {
      if (userLocation != null) {
        print("Hello?");
        bounds =
            getCameraLatLngBounds(widget.spots, userLocation: userLocation);
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
              title: spot.principalTag,
            )),
      );
    });

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
          return GoogleMap(
            markers: markers,
            mapType: MapType.normal,
            zoomControlsEnabled: false,
            initialCameraPosition: initialCameraPosition,
            mapToolbarEnabled: false,
            myLocationButtonEnabled: widget.myLocationButtonEnabled,
            myLocationEnabled: true,
            onMapCreated: (controller) {
              onMapCreated(controller, userLocation: snapshot.data);
            },
            scrollGesturesEnabled: !widget.blockMap,
            zoomGesturesEnabled: !widget.blockMap,
          );

          } else if(snapshot.hasError) {
            print("snapshot.hasError");
            return Container(child: Text("Ops we are having problems to didplay the map"),);

          } else {
            print("snapshot loading");
            return Center(child: const CircularProgressIndicator());
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

      return GoogleMap(
        markers: markers,
        mapType: MapType.normal,
        zoomControlsEnabled: false,
        initialCameraPosition: initialCameraPosition,
        mapToolbarEnabled: false,
        myLocationButtonEnabled: widget.myLocationButtonEnabled,
        myLocationEnabled: true,
        onMapCreated: (controller) {
          onMapCreated(controller);
        },
        scrollGesturesEnabled: !widget.blockMap,
        zoomGesturesEnabled: !widget.blockMap,
      );
    }
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

