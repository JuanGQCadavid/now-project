import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';
import 'package:now_v2/features/filters/application/filter_providers.dart';

class CentralMap extends ConsumerStatefulWidget {
  const CentralMap({Key? key}) : super(key: key);

  @override
  CentralMapState createState() => CentralMapState();
}

class CentralMapState extends ConsumerState<CentralMap> {
  final String apiKey = "INSERT_KEY";

  @override
  void initState() {
    super.initState();

    final filterNotifier = ref.read(filterNotifierProvier);
    WidgetsBinding.instance?.addPostFrameCallback((_) {
      filterNotifier.fetchSpotsFrom(LatLng(6.2471017, -75.5874348));
    });
  }

  @override
  Widget build(BuildContext context) {
    // También podemos usar "ref" para escuchar a un provider dentro del método build
    final filterNotifier = ref.watch(filterNotifierProvier);
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
          markers: filterNotifier.markers.values.toList(),
        ),
      ],
    );
  }
}
