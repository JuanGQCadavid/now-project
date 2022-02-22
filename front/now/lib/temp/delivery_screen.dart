// import 'dart:math';

// import 'package:flutter/material.dart';
// import 'package:google_maps_flutter/google_maps_flutter.dart';
// import 'package:now/temp/directions_provider.dart';
// import 'package:provider/provider.dart';

// class DeliveryScreen extends StatefulWidget {
//   final LatLng fromPoint = const LatLng(6.244725272333634, -75.5889493512081);
//   final LatLng toPoint = const LatLng(6.2464643393698855, -75.59065451559682);

//   const DeliveryScreen({Key? key}) : super(key: key);

//   @override
//   _DeliveryScreenState createState() => _DeliveryScreenState();
// }

// class _DeliveryScreenState extends State<DeliveryScreen> {
//   late GoogleMapController _controller;

//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       appBar: AppBar(
//         title: const Text("Now project!"),
//       ),
//       body: _map(),
//       floatingActionButton: FloatingActionButton(
//         onPressed: _centerView,
//         child: const Icon(Icons.zoom_out_map),
//       ),
//     );
//   }

//   Widget _map() {
//     return Consumer<DirectionsProvider>(
//       builder:
//           (BuildContext context, DirectionsProvider provider, Widget? child) {
//         return Stack(
//           children: [
//             _googleMap(provider),
//             Center(
//                 child: Container(
//               child: Text(
//                 provider.distanceTime,
//                 style: const TextStyle(fontSize: 25),
//               ),
//               color: Colors.amber,
//             )),
//           ],
//         );
//       },
//     );
//   }

//   GoogleMap _googleMap(DirectionsProvider provider) {
//     return GoogleMap(
//       initialCameraPosition: CameraPosition(target: widget.fromPoint, zoom: 16),
//       markers: _createMarkers(),
//       onMapCreated: _onMapCreated,
//       mapToolbarEnabled: false,
//       myLocationButtonEnabled: true,
//       zoomControlsEnabled: false,
//       polylines: provider.currentRoute,
//       myLocationEnabled: true,
//     );
//   }

//   Set<Marker> _createMarkers() {
//     var markers = Set<Marker>();

//     markers.add(
//       Marker(
//           markerId: const MarkerId("fromMarker"),
//           position: widget.fromPoint,
//           infoWindow: const InfoWindow(title: "You are here.")),
//     );

//     markers.add(
//       Marker(
//           markerId: const MarkerId("toMarker"),
//           position: widget.toPoint,
//           infoWindow: const InfoWindow(title: "You will be here.")),
//     );

//     return markers;
//   }

//   _onMapCreated(GoogleMapController controller) {
//     _controller = controller;

//     _centerView();
//   }

//   void _centerView() async {
//     await _controller.getVisibleRegion();
//     var api = Provider.of<DirectionsProvider>(context, listen: false);
//     await api.findDirections(widget.fromPoint, widget.toPoint);

//     var left = min(widget.fromPoint.latitude, widget.toPoint.latitude);
//     var right = max(widget.fromPoint.latitude, widget.toPoint.latitude);
//     var top = max(widget.fromPoint.longitude, widget.toPoint.longitude);
//     var down = min(widget.fromPoint.longitude, widget.toPoint.longitude);

//     api.currentRoute.first.points.forEach((point) {
//       left = min(left, point.latitude);
//       right = max(right, point.latitude);
//       top = max(top, point.longitude);
//       down = min(down, point.longitude);
//     });

//     var bounds = LatLngBounds(
//         southwest: LatLng(left, down), northeast: LatLng(right, top));
//     var cameraUpdate = CameraUpdate.newLatLngBounds(bounds, 50);
//     _controller.animateCamera(cameraUpdate);
//   }
// }
