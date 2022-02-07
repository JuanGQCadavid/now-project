import 'package:flutter/material.dart';
import 'package:now/filters/application/fetch_data.dart';
import 'spots_granular_view.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';

class SpotGeneralView extends StatelessWidget {
  const SpotGeneralView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const _PageScafold();
  }
}

class _PageScafold extends StatelessWidget {
  const _PageScafold({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: const _Body(),
      appBar: AppBar(
        title: Center(
          child: Text("Ongoing events!"),
        ),
      ),
      floatingActionButton: FloatingActionButton(onPressed: () {
        Navigator.push(context,
            MaterialPageRoute(builder: (context) => const SpotGranularView()));
      }),
    );
  }
}

class _Body extends StatelessWidget {
  const _Body({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: _MapBody(),
      decoration: const BoxDecoration(
        color: Colors.blueGrey,
      ),
    );
  }
}

class _MapBody extends StatefulWidget {
  const _MapBody({Key? key}) : super(key: key);

  @override
  __MapBodyState createState() => __MapBodyState();
}

class __MapBodyState extends State<_MapBody> {
  late GoogleMapController mapController;
  final Map<String, Marker> _markers = {};

  final LatLng _center = const LatLng(0, 0); //LatLng(6.246408, -75.590666);

  void _onMapCreated(GoogleMapController controller) async {
    mapController = controller;

    final googleOficies = await getGoogleOfficies();
    final BitmapDescriptor pinLocation = await BitmapDescriptor.fromAssetImage(
        ImageConfiguration(devicePixelRatio: 2.5), 'assets/custo_marker.png');

    setState(() {
      _markers.clear();
      for (final office in googleOficies.offices) {
        final marker = Marker(
          icon: pinLocation,
          markerId: MarkerId(office.name),
          position: LatLng(office.lat, office.lng),
          infoWindow: InfoWindow(title: office.name, snippet: office.address),
        );
        _markers[office.name] = marker;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return GoogleMap(
      onMapCreated: _onMapCreated,
      initialCameraPosition: CameraPosition(
        target: _center,
        zoom: 2, //20.0,
      ),
      markers: _markers.values.toSet(),
    );
  }
}
