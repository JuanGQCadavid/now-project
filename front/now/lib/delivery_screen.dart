import 'dart:math';

import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';

class DeliveryScreen extends StatefulWidget {
  final LatLng fromPoint = const LatLng(-38.956176, -67.920666);
  final LatLng toPoint = const LatLng(-38.953724, -67.923921);

  const DeliveryScreen({Key? key}) : super(key: key);

  @override
  _DeliveryScreenState createState() => _DeliveryScreenState();
}

class _DeliveryScreenState extends State<DeliveryScreen> {
  late GoogleMapController _controller;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Now project!"),
      ),
      body: _map(),
      floatingActionButton: FloatingActionButton(
        onPressed: _centerView,
        child: const Icon(Icons.zoom_out_map),
      ),
    );
  }

  Widget _map() {
    return GoogleMap(
      initialCameraPosition: CameraPosition(target: widget.fromPoint, zoom: 16),
      markers: _createMarkers(),
      onMapCreated: _onMapCreated,
      mapToolbarEnabled: false,
      myLocationButtonEnabled: false,
      zoomControlsEnabled: false,
    );
  }

  Set<Marker> _createMarkers() {
    var markers = Set<Marker>();

    markers.add(
      Marker(
          markerId: const MarkerId("fromMarker"),
          position: widget.fromPoint,
          infoWindow: const InfoWindow(title: "You are here.")),
    );

    markers.add(
      Marker(
          markerId: const MarkerId("toMarker"),
          position: widget.toPoint,
          infoWindow: const InfoWindow(title: "You will be here.")),
    );

    return markers;
  }

  _onMapCreated(GoogleMapController controller) {
    _controller = controller;

    _centerView();
  }

  void _centerView() async {
    await _controller.getVisibleRegion();

    var left = min(widget.fromPoint.latitude, widget.toPoint.latitude);
    var right = max(widget.fromPoint.latitude, widget.toPoint.latitude);

    var top = max(widget.fromPoint.longitude, widget.toPoint.longitude);
    var down = min(widget.fromPoint.longitude, widget.toPoint.longitude);

    var bounds = LatLngBounds(
        southwest: LatLng(left, down), northeast: LatLng(right, top));
    var cameraUpdate = CameraUpdate.newLatLngBounds(bounds, 50);
    _controller.animateCamera(cameraUpdate);
  }
}
