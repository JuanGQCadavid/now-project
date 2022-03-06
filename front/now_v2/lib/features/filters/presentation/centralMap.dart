import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';
import 'package:now_v2/features/filters/application/filter_providers.dart';
import 'package:now_v2/features/filters/presentation/map.dart';

class CentralMap extends ConsumerStatefulWidget {
  const CentralMap({Key? key}) : super(key: key);

  @override
  CentralMapState createState() => CentralMapState();
}

class CentralMapState extends ConsumerState<CentralMap> {
  final String apiKey =
      "sk.eyJ1IjoianVhbmdvbnphbG8iLCJhIjoiY2wwNWo0ZWRyMXlnNzNicGtza3JjdGl5cSJ9.G-lRBbPXFx3zKLxp7ekudg";

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

    return AppMap(
      markers: filterNotifier.markers.values.toList(),
    );
  }
}
