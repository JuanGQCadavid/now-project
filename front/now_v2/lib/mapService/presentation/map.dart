import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';
import 'package:now_v2/env/prod.dart';
import 'package:now_v2/mapService/application/map_notifier.dart';
import 'package:now_v2/mapService/application/map_provider.dart';
import 'package:latlong2/latlong.dart';
import 'package:geolocator/geolocator.dart';


class MapPage extends ConsumerStatefulWidget {
  final List<Marker> markers;
  final double minZoom;
  final double maxZoom;
  final double zoom;
  final bool blockMove;
  final bool boundOnPoints;
  final bool addUserLocation;

  const MapPage({
      Key? key, 
      required this.markers,
      this.minZoom = 11.0,
      this.maxZoom = 17.0,
      this.blockMove = false,
      this.zoom = 13.0,
      this.boundOnPoints = false,
      this.addUserLocation = false
    }): super(key: key);

  @override
  _MapPageState createState() => _MapPageState();

}

class _MapPageState extends ConsumerState<MapPage> {

  @override
  void initState() {
    super.initState();

    final provider = ref.read(mapStateProvider.notifier);
    provider.setMarkers(widget.markers);

    if (widget.addUserLocation) {
      provider.adduserLocation();
    }
  }

  @override
  Widget build(BuildContext context) {
    MapState state = ref.watch(mapStateProvider);

    return FlutterMap(
      mapController: state.mapController,
      options: MapOptions(
        center: LatLng(.0, .0),
        // minZoom: widget.minZoom,
        // maxZoom: widget.maxZoom,
        interactiveFlags: widget.blockMove
            ? InteractiveFlag.none
            : InteractiveFlag.pinchZoom | InteractiveFlag.drag,
        zoom: widget.zoom,
      ),
      layers: [
        TileLayerOptions(
          urlTemplate:
              "https://api.mapbox.com/styles/v1/juangonzalo/cl05djr2i000115m1n529qyh8/tiles/256/{z}/{x}/{y}@2x?access_token=${AppProperties.map_api_key}",
          additionalOptions: {
            "accessToken": AppProperties.map_api_key,
            "id": "mapbox.mapbox-streets-v8"
          },
        ),
        MarkerLayerOptions(
          markers: state.markers,
        ),
      ],
    );
  }
}



class AppMapStateful extends StatefulWidget {
  final List<Marker> markers;
  final double minZoom;
  final double maxZoom;
  final double zoom;
  final bool blockMove;
  final bool boundOnPoints;
  final bool addUserLocation;

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
  late List<Marker> markers;
  bool isAlready = false;

  void mapCreated(MapController mapController) async {
    Position? locationData = await Geolocator.getLastKnownPosition();
    this.mapController = mapController;

    if (widget.boundOnPoints) {
      List<LatLng> _points = [];

      markers.forEach((element) {
        _points.add(element.point);
      });

      mapController.fitBounds(
        LatLngBounds.fromPoints(_points),
      );
    } else {
      mapController.move(
          LatLng(locationData!.latitude, locationData.longitude), widget.zoom);
    }
  }

  @override
  void initState() {
      super.initState();
      //WidgetsBinding.instance!.addPostFrameCallback((_) => ());
  }

  @override
  Widget build(BuildContext context) {

    return FlutterMap(
      options: MapOptions(
        onMapCreated: mapCreated,
        center: LatLng(.0, .0), //LatLng(6.2471017, -75.5874348),
        // minZoom: widget.minZoom,
        // maxZoom: widget.maxZoom,
        interactiveFlags: widget.blockMove
            ? InteractiveFlag.none
            : InteractiveFlag.pinchZoom | InteractiveFlag.drag,
        zoom: widget.zoom,
      ),
      layers: [
        TileLayerOptions(
          urlTemplate:
              "https://api.mapbox.com/styles/v1/juangonzalo/cl05djr2i000115m1n529qyh8/tiles/256/{z}/{x}/{y}@2x?access_token=${AppProperties.map_api_key}",
          additionalOptions: {
            "accessToken": AppProperties.map_api_key,
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
