import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:location/location.dart';
import 'package:latlong2/latlong.dart';
import 'package:flutter/services.dart' show rootBundle;

class Body extends StatelessWidget {
  const Body({Key? key}) : super(key: key);
  final String apiKey = "INSERT_HERE";

  @override
  Widget build(BuildContext context) {
    return FlutterMap(
      options: MapOptions(
        center: LatLng(51.5, -0.09),
        //zoom: 19.0,
        //maxZoom: 18,
        minZoom: 11.0,
        maxZoom: 17.0,
        interactiveFlags: InteractiveFlag.pinchZoom | InteractiveFlag.drag,
        zoom: 13.0,
      ),
      layers: [
        TileLayerOptions(
          urlTemplate:
              "https://api.mapbox.com/styles/v1/juangonzalo/cl05djr2i000115m1n529qyh8/tiles/256/{z}/{x}/{y}@2x?access_token=${apiKey}",
          additionalOptions: {
            "accessToken": apiKey,
            "id": "mapbox.mapbox-streets-v8"
          },
        ),
        MarkerLayerOptions(
          markers: [
            Marker(
              width: 80.0,
              height: 80.0,
              point: LatLng(51.5, -0.09),
              builder: (ctx) => Container(
                  color: Colors.red,
                  child: TextButton(
                      child: Text("dsfsdf"),
                      onPressed: () {
                        print("DUDEEE");
                      })),
            ),
          ],
        ),
      ],
    );
  }
}


// class MapBody extends ConsumerStatefulWidget {
//   const MapBody({Key? key}) : super(key: key);

//   @override
//   _MapBodyState createState() => _MapBodyState();
// }

// class _MapBodyState extends ConsumerState<MapBody> {
//   late GoogleMapController mapController;
//   late String _mapStyle;
//   final LatLng _center = const LatLng(0, 0);

//   final Location _location = Location();

//   @override
//   void initState() {
//     super.initState();

//     WidgetsBinding.instance?.addPostFrameCallback((_) {
//       final provider = ref.read(filterNotifierProvier);
//       _location.getLocation().then(
//             (value) => provider.fetchSpotsFrom(
//               LatLng(
//                 value.latitude ?? 0.0,
//                 value.longitude ?? 0.0,
//               ),
//             ),
//           );
//     });

//     rootBundle
//         .loadString('assets/maps/mapStyle.json', cache: true)
//         .then((string) {
//       _mapStyle = string;
//     });
//   }

//   void findMe() async {
//     final locationData = await _location.getLocation();
//     mapController.animateCamera(CameraUpdate.newCameraPosition(CameraPosition(
//         target: LatLng(
//           locationData.latitude ?? 0.0,
//           locationData.longitude ?? 0.0,
//         ),
//         zoom: 15)));
//   }

//   void _onMapCreated(GoogleMapController controller) async {
//     mapController = controller;
//     mapController.setMapStyle(_mapStyle);
//     findMe();
//   }

//   @override
//   Widget build(BuildContext context) {
//     final provider = ref.watch(filterNotifierProvier);
//     return GoogleMap(
//       onMapCreated: _onMapCreated,
//       initialCameraPosition: CameraPosition(
//         target: _center,
//         zoom: 2, //20.0,
//       ),
//       mapToolbarEnabled: false,
//       myLocationButtonEnabled: false,
//       liteModeEnabled: false,
//       zoomControlsEnabled: false,
//       minMaxZoomPreference: const MinMaxZoomPreference(15, 30),
//       markers: provider.markers.values.toSet(),
//     );
//   }
// }
