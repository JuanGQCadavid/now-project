import 'dart:async';

import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/login/view/widgets/text_input.dart';

class LocationSelectorView extends StatelessWidget {
  late Completer<GoogleMapController> mapController = Completer();

  LocationSelectorView({super.key});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 15.0),
      child: Stack(
        children: [
          Container(
            height: 600,
            width: double.infinity,
            child: NowMapV2(
              mapController: mapController,
            ),
          ),
          TextField()
        ],
      ),
      // child: Column(
      //   children: [Text("LocationSelectorView")],
      // ),
    );
  }
}
