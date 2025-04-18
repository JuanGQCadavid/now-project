import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/features/general_view/views/main.dart';

// HERE
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/login/view/main.dart';
import 'dart:math';

import 'package:now_v8/src/features/spots_creation/main.dart';

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      // home: GeneralViewFeature(),
      home: SpotsCreationFeature(),
      // home: DetailsWorkingHome(), //()); //LoginFeature());
    );
  }
}

class DetailsWorkingHome extends StatelessWidget {
  const DetailsWorkingHome({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(child: UserAccountButton()),
    );
  }
}

class UserAccountButton extends StatelessWidget {
  const UserAccountButton({super.key});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        UserLoggedButton(
          onTap: () {},
          displayName: "JG",
          colors: [
            Colors.cyan.shade700,
            Colors.pink.shade600,
          ],
        ),
        const SizedBox(
          width: 10,
        ),
        // Option D
        ClipOval(
          child: Container(
            padding: const EdgeInsets.all(2),
            decoration: const BoxDecoration(
              gradient: LinearGradient(
                  begin: Alignment.topLeft,
                  end: Alignment.bottomRight,
                  colors: [
                    Colors.blue,
                    Colors.purpleAccent,
                  ]),
            ),
            child: ClipOval(
              child: Container(
                padding: const EdgeInsets.all(6),
                color: Colors.white,
                child: const Text(
                  "JG",
                  // style: TextStyle(color: Colors.blueGrey),
                ),
              ),
            ),
          ),
        ),
        const SizedBox(
          width: 10,
        ),
        // Option B
        ClipOval(
          child: Container(
            padding: const EdgeInsets.all(2),
            color: Colors.blue,
            child: ClipOval(
              child: Container(
                padding: const EdgeInsets.all(6),
                color: Colors.white,
                child: const Text(
                  "JG",
                  // style: TextStyle(color: Colors.blueGrey),
                ),
              ),
            ),
          ),
        ),
        // Option A
        IconButton(
          onPressed: () {},
          icon: const Icon(
            Icons.person,
            color: Colors.blue,
          ),
        ),
        // Option C
        IconButton(
          onPressed: () {},
          icon: const Icon(
            Icons.person,
            // color: Colors.blue,
          ),
        )
      ],
    );
  }
}

class TestMapView extends StatefulWidget {
  const TestMapView({super.key});

  @override
  State<TestMapView> createState() => _TestMapViewState();
}

class _TestMapViewState extends State<TestMapView> {
  late Completer<GoogleMapController> mapController = Completer();
  late CameraPosition lastCameraPosition;
  LatLng initialLatLng = const LatLng(6.2045744, 75.5817167);
  List<Spot> spots = const [];

  @override
  Widget build(BuildContext context) {
    return NowMapV2(
      spots: spots,
      centerMapOnSpots: true,
      mapController: mapController,
      camaraPosition: initialLatLng,
      onCameraIdle: onCameraIdle,
      onCameraMove: onCameraMove,
      onCameraMoveStarted: onCameraMoveStarted,
    );
  }

  void onCameraMove(CameraPosition cameraPosition) {
    print(
        "Moving on lat ${cameraPosition.target.latitude} lng ${cameraPosition.target.longitude} zoom ${cameraPosition.zoom}");
    lastCameraPosition = cameraPosition;
  }

  void onCameraIdle() {
    print("Stopped");
    List<Spot> newSpots = [];

    var metersPerPx = (156543.03392 *
            cos(lastCameraPosition.target.latitude * pi / 180) /
            pow(2, lastCameraPosition.zoom)) /
        250; // 250 300;   //1000;  //1000;
    // Center
    newSpots.add(Spot(
      principalTag: "center",
      secondaryTags: [],
      latLng: lastCameraPosition.target,
      spotId: "center",
      spotsColor: SpotsColors(
          color: Colors.green,
          hue:
              BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueGreen)),
      date: DateTime.now(),
    ));

    newSpots.add(Spot(
      principalTag: "up",
      secondaryTags: [],
      latLng: LatLng(lastCameraPosition.target.latitude + metersPerPx,
          lastCameraPosition.target.longitude),
      spotId: "up",
      spotsColor: SpotsColors(
          color: Colors.blue,
          hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueBlue)),
      date: DateTime.now(),
    ));

    newSpots.add(Spot(
      principalTag: "down",
      secondaryTags: [],
      latLng: LatLng(lastCameraPosition.target.latitude - metersPerPx,
          lastCameraPosition.target.longitude),
      spotId: "down",
      spotsColor: SpotsColors(
          color: Colors.blue,
          hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueBlue)),
      date: DateTime.now(),
    ));

    newSpots.add(Spot(
      principalTag: "rigth",
      secondaryTags: [],
      latLng: LatLng(lastCameraPosition.target.latitude,
          lastCameraPosition.target.longitude + metersPerPx),
      spotId: "rigth",
      spotsColor: SpotsColors(
          color: Colors.blue,
          hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueBlue)),
      date: DateTime.now(),
    ));

    newSpots.add(Spot(
      principalTag: "left",
      secondaryTags: [],
      latLng: LatLng(lastCameraPosition.target.latitude,
          lastCameraPosition.target.longitude - metersPerPx),
      spotId: "left",
      spotsColor: SpotsColors(
          color: Colors.blue,
          hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueBlue)),
      date: DateTime.now(),
    ));

    setState(() {
      spots = newSpots;
    });
  }

  void onCameraMoveStarted() {
    print("HERE WE GO");
  }
}
// var  metersPerPx = 156543.03392 * Math.cos(latLng.lat() * Math.PI / 180) / Math.pow(2, zoom)
