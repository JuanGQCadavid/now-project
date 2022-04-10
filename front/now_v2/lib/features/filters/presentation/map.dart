import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';

class AppMap extends StatelessWidget {
  final List<Marker> markers;
  const AppMap({
    Key? key,
    required this.markers,
  }) : super(key: key);

  final String apiKey = "sk.eyJ1IjoianVhbmdvbnphbG8iLCJhIjoiY2wwNWo0ZWRyMXlnNzNicGtza3JjdGl5cSJ9.G-lRBbPXFx3zKLxp7ekudg";

  @override
  Widget build(BuildContext context) {
    return FlutterMap(
      options: MapOptions(
        center: LatLng(6.2471017, -75.5874348),
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
          markers: markers,
        ),
      ],
    );
  }
}
