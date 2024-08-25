import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';

class MapsClusterDemoTwo extends StatefulWidget {
  const MapsClusterDemoTwo({super.key});

  @override
  State<MapsClusterDemoTwo> createState() => _MapsClusterDemoTwoState();
}

class _MapsClusterDemoTwoState extends State<MapsClusterDemoTwo> {
  GoogleMapController? controller;
  // This is my state
  Map<ClusterManagerId, ClusterManager> clusterManagers =
      <ClusterManagerId, ClusterManager>{};

  Map<MarkerId, Marker> markers = <MarkerId, Marker>{};

  void _onMapCreated(GoogleMapController controllerParam) {
    ClusterManager general = ClusterManager(
      clusterManagerId: ClusterManagerId("1"),
      onClusterTap: (argument) {
        print("Cluster tab");
      },
    );
    setState(() {
      controller = controllerParam;
    });

    clusterManagers[general.clusterManagerId] = general;
    for (var i = 0; i < 1000; i++) {
      var id = MarkerId("${i}");
      markers[id] = Marker(
          clusterManagerId: general.clusterManagerId,
          markerId: id,
          position: LatLng(-33.852 + (i * 0.001), 151.25),
          infoWindow: InfoWindow(title: id.value));
    }
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    return GoogleMap(
      onMapCreated: _onMapCreated,
      initialCameraPosition: const CameraPosition(
        target: LatLng(-33.852, 151.25),
        zoom: 11.0,
      ),
      markers: Set<Marker>.of(markers.values),
      clusterManagers: Set<ClusterManager>.of(clusterManagers.values),
    );
  }
}
