import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/material.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'dart:async';

import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';

class MapSample extends ConsumerWidget {
  const MapSample({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final spotsState = ref.watch(spotsStateProvider);

    return Stack(
      children: [
        NowMap(
          spots: spotsState,
        ),
        Align(
          alignment: Alignment.bottomLeft,
          child: MapTags(
            spots: spotsState,
          ),
        )
      ],
    );
  }
}

class MapTags extends StatelessWidget {
  final List<Spot> spots;
  const MapTags({
    required this.spots,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    List<Widget> rowTags = [];

    spots.forEach((spot) {
      rowTags.add(Container(
        margin: EdgeInsets.only(right: 15),
        child: SpotTag(
          color: spot.spotsColor.color,
          tag: spot.principalTag,
          onPressed: () {},
        ),
      ));
    });

    return Container(
      margin: EdgeInsets.only(left: 15, bottom: 15),
      child: SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: Row(
          children: rowTags,
        ),
      ),
    );
  }
}

class NowMap extends StatefulWidget {
  final List<Spot> spots;
  NowMap({
    Key? key,
    required this.spots,
  }) : super(key: key);

  @override
  State<NowMap> createState() => _NowMapState();
}

class _NowMapState extends State<NowMap> {
  Completer<GoogleMapController> _controller = Completer();
  // final Completer<GoogleMapController> _controller;

  static final CameraPosition _kGooglePlex = CameraPosition(
    target: LatLng(6.251723, -75.592771),
    zoom: 14.4746,
  );

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
            infoWindow: InfoWindow(title: spot.principalTag)),
      );
    });

    return GoogleMap(
      markers: markers,
      mapType: MapType.normal,
      initialCameraPosition: _kGooglePlex,
      mapToolbarEnabled: false,
      myLocationButtonEnabled: false,
      padding: EdgeInsets.only(bottom: 65, left: 15),
      // onMapCreated: (GoogleMapController controller) {
      //   _controller.complete(controller);
      // },
    );
  }
}
