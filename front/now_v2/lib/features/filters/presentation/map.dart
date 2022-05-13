import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:geolocator/geolocator.dart';

class AppMapStateful extends StatefulWidget {
  final List<Marker> markers;
  final double minZoom;
  final double maxZoom;
  final double zoom;
  final bool blockMove;
  final bool boundOnPoints;
  final bool addUserLocation;

  final String apiKey =
      "sk.eyJ1IjoianVhbmdvbnphbG8iLCJhIjoiY2wwNWo0ZWRyMXlnNzNicGtza3JjdGl5cSJ9.G-lRBbPXFx3zKLxp7ekudg";

  const AppMapStateful(
      {Key? key,
      required this.markers,
      this.minZoom = 11.0,
      this.maxZoom = 17.0,
      this.blockMove = false,
      this.zoom = 13.0,
      this.boundOnPoints = false,
      this.addUserLocation = false})
      : super(key: key);

  @override
  State<AppMapStateful> createState() => _AppMapStatefulState();
}

class _AppMapStatefulState extends State<AppMapStateful> {
  late MapController mapController;

  void mapCreated(MapController mapController) async {
    Position? locationData = await Geolocator.getLastKnownPosition();
    this.mapController = mapController;

    print(locationData);

    if (widget.boundOnPoints) {
      List<LatLng> _points = [];

      widget.markers.forEach((element) {
        _points.add(element.point);
      });

      _points.add(LatLng(locationData!.latitude, locationData.longitude));

      mapController.fitBounds(
        LatLngBounds.fromPoints(_points),
      );
    } else {
      mapController.move(
          LatLng(locationData!.latitude, locationData.longitude), widget.zoom);
    }
  }

  @override
  Widget build(BuildContext context) {

    if (widget.addUserLocation) {
      
    }


    return FlutterMap(
      options: MapOptions(
        onMapCreated: mapCreated,
        center: LatLng(.0, .0), //LatLng(6.2471017, -75.5874348),
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
