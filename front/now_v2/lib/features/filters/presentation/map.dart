import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:location/location.dart';

class AppMapStateful extends StatefulWidget {
  final List<Marker> markers;
  final double minZoom;
  final double maxZoom;
  final double zoom;
  final bool blockMove;
  final bool boundOnPoints;

  final String apiKey =
      "sk.eyJ1IjoianVhbmdvbnphbG8iLCJhIjoiY2wwNWo0ZWRyMXlnNzNicGtza3JjdGl5cSJ9.G-lRBbPXFx3zKLxp7ekudg";
  const AppMapStateful({
    Key? key,
    required this.markers,
    this.minZoom = 11.0,
    this.maxZoom = 17.0,
    this.blockMove = false,
    this.zoom = 13.0,
    this.boundOnPoints = false,
  }) : super(key: key);

  @override
  State<AppMapStateful> createState() => _AppMapStatefulState();
}

class _AppMapStatefulState extends State<AppMapStateful> {
  late MapController mapController;
  late Location location = Location();

  void mapCreated(MapController mapController) async {
    var locationData = await location.getLocation();
    setState(() {
      this.mapController = mapController;

      if (widget.boundOnPoints) {
        List<LatLng> _points = [];

        widget.markers.forEach((element) {
          _points.add(element.point);
        });

        _points.add(LatLng(locationData.latitude!, locationData.longitude!));

        mapController.fitBounds(
          LatLngBounds.fromPoints(_points),
        );
      } else {
        mapController.move(
            LatLng(locationData.latitude!, locationData.longitude!),
            widget.zoom);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return FlutterMap(
      options: MapOptions(
        onMapCreated: mapCreated,
        center: LatLng(.0, .0),
        minZoom: widget.minZoom,
        maxZoom: widget.maxZoom,
        interactiveFlags: widget.blockMove
            ? InteractiveFlag.none
            : InteractiveFlag.pinchZoom | InteractiveFlag.drag,
        zoom: widget.zoom,
      ),
      layers: [
        TileLayerOptions(
          urlTemplate:
              "https://api.mapbox.com/styles/v1/juangonzalo/cl05djr2i000115m1n529qyh8/tiles/256/{z}/{x}/{y}@2x?access_token=${widget.apiKey}",
          additionalOptions: {
            "accessToken": widget.apiKey,
            "id": "mapbox.mapbox-streets-v8"
          },
        ),
        MarkerLayerOptions(
          markers: widget.markers,
        ),
      ],
    );
  }
}

class AppMap extends StatelessWidget {
  final List<Marker> markers;
  final double minZoom;
  final double maxZoom;
  final double zoom;
  final bool blockMove;

  const AppMap({
    Key? key,
    required this.markers,
    this.minZoom = 11.0,
    this.maxZoom = 17.0,
    this.blockMove = false,
    this.zoom = 13.0,
  }) : super(key: key);

  final String apiKey =
      "sk.eyJ1IjoianVhbmdvbnphbG8iLCJhIjoiY2wwNWo0ZWRyMXlnNzNicGtza3JjdGl5cSJ9.G-lRBbPXFx3zKLxp7ekudg";

  Widget buildMap(double centerLat, double centerLon) {
    return FlutterMap(
      options: MapOptions(
        center: LatLng(centerLat, centerLon),
        minZoom: minZoom,
        maxZoom: maxZoom,
        interactiveFlags: blockMove
            ? InteractiveFlag.none
            : InteractiveFlag.pinchZoom | InteractiveFlag.drag,
        zoom: zoom,
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

  @override
  Widget build(BuildContext context) {
    final location = Location();

    return FutureBuilder(
      future: location.getLocation(),
      builder: (ctx, snap) {
        if (snap.connectionState == ConnectionState.done) {
          print("MOR");

          LocationData locationData = snap.data as LocationData;

          print(locationData.latitude ?? 0.0);
          print(locationData.longitude ?? 0.0);

          return buildMap(
            locationData.latitude ?? 0.0,
            locationData.longitude ?? 0.0,
          );
        } else {
          return buildMap(.0, .0);
        }
      },
    );
  }
}
