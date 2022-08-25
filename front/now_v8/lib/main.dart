import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/general_view/views/generalView.dart';

void main() {
  Hive.initFlutter();
  runApp(
      ProviderScope(
        child: MyApp(),
      ),
    );
} 

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Flutter Demo',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: GeneralViewFeature(),//NowMapTest(),//GeneralViewFeature() //MyHomePage(title: 'Flutter Demo Home Page'),
        );
  }
}

class NowMapTest extends StatelessWidget {
  const NowMapTest({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: NowMapOld(),
        // child: Container(
        //   width: 100,
        //   color: Colors.blue,
        // ),
      ),
    );
  }
}

class NowMapOld extends StatelessWidget {
  const NowMapOld({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    List<Spot> spot = [];

    spot.add(
      Spot.withOutSpotColors(
        principalTag: "",
        secondaryTags: [""],
        latLng: LatLng(6.241791, -75.595383),
        spotId: "1",
      ),
    );

    spot.add(
      Spot.withOutSpotColors(
        principalTag: "",
        secondaryTags: [""],
        latLng: LatLng(6.2386133, -75.5858772),
        spotId: "2",
      ),
    );

    return Container(
      height: 300,
      // width: 400,
      color: Colors.red,
      child: NowMapV2(
          spots: spot,
          includeUserLocation: true,
          myLocationButtonEnabled: true, centerMapOnSpots: true,),
    );
  }
}
