import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/general_view/views/generalView.dart';
import 'package:flutter_svg/flutter_svg.dart';

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
        home:
            GeneralViewFeature(),//NowMapTest(),//GeneralViewFeature() //MyHomePage(title: 'Flutter Demo Home Page'),
        );
  }
}

class NowMapTest extends StatelessWidget {
  const NowMapTest({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: IconTextButtom(
          icon: Icons.phone,
          iconColor: Colors.white,
          mainColor: Colors.green.shade500,
          message: "Say hi",
          onTap: () {
            
          },
        ),
        // child: Container(
        //   width: 100,
        //   color: Colors.blue,
        // ),
      ),
    );
  }
}
