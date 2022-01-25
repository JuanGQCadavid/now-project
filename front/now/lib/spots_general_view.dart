import 'package:flutter/material.dart';
import 'spots_granular_view.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';

class SpotGeneralView extends StatelessWidget {
  const SpotGeneralView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const _PageScafold(),
    );
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

  final LatLng _center = const LatLng(6.246408, -75.590666);

  void _onMapCreated(GoogleMapController controller) {
    mapController = controller;
  }

  @override
  Widget build(BuildContext context) {
    return GoogleMap(
        onMapCreated: _onMapCreated,
        initialCameraPosition: CameraPosition(
          target: _center,
          zoom: 20.0,
        ));
  }
}
