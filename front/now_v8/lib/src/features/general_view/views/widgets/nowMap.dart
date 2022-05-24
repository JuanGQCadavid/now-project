import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/material.dart';
import 'dart:async';

import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';

class MapSample extends StatefulWidget {
  @override
  State<MapSample> createState() => MapSampleState();
}

class MapSampleState extends State<MapSample> {
  Completer<GoogleMapController> _controller = Completer();

  static final CameraPosition _kGooglePlex = CameraPosition(
    target: LatLng(6.251723, -75.592771),
    zoom: 14.4746,
  );

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        NowMap(kGooglePlex: _kGooglePlex, controller: _controller),
        Align(
          alignment: Alignment.bottomLeft,
          child: MapTags(),
        )
      ],
    );
  }
}

class MapTags extends StatelessWidget {
  const MapTags({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.only(left: 15, bottom: 15),
      child: SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: Row(
          children: [
            Container(
              margin: EdgeInsets.only(right: 15),
              child: SpotTag(
                color: Colors.green.shade900,
                tag: "CoffeLovers",
                onPressed: () {},
              ),
            ),
            Container(
              margin: EdgeInsets.only(right: 15),
              child: SpotTag(
                color: Colors.blue.shade900,
                tag: "ReadingClub",
                onPressed: () {},
              ),
            ),
            Container(
              margin: EdgeInsets.only(right: 15),
              child: SpotTag(
                color: Colors.red.shade900,
                tag: "StreePainting",
                onPressed: () {},
              ),
            ),
            Container(
              margin: EdgeInsets.only(right: 15),
              child: SpotTag(
                color: Colors.orange.shade900,
                tag: "Dance",
                onPressed: () {},
              ),
            ),
            Container(
              margin: EdgeInsets.only(right: 15),
              child: SpotTag(
                color: Colors.pink.shade900,
                tag: "Lettering",
                onPressed: () {},
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class NowMap extends StatelessWidget {
  const NowMap({
    Key? key,
    required CameraPosition kGooglePlex,
    required Completer<GoogleMapController> controller,
  })  : _kGooglePlex = kGooglePlex,
        _controller = controller,
        super(key: key);

  final CameraPosition _kGooglePlex;
  final Completer<GoogleMapController> _controller;

  @override
  Widget build(BuildContext context) {
    Set<Marker> markers = Set();

    markers.add( Marker(
      markerId: MarkerId("juan"),
      position: LatLng(6.251723, -75.592771),
      visible: true,
      icon: BitmapDescriptor.defaultMarkerWithHue(270)
    ));

    return GoogleMap(
      markers: markers,
      mapType: MapType.normal,
      initialCameraPosition: _kGooglePlex,
      mapToolbarEnabled: false,
      myLocationButtonEnabled: false,
      padding: EdgeInsets.only(bottom: 65, left: 15),
      onMapCreated: (GoogleMapController controller) {
        _controller.complete(controller);
      },
    );
  }
}
