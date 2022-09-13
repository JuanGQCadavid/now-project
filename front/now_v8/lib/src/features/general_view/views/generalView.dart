import 'dart:async';

import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/features/general_view/views/widgets/footbar.dart';
import 'package:now_v8/src/features/general_view/views/widgets/header.dart';
import 'package:now_v8/src/features/general_view/views/widgets/nowMap.dart';

class GeneralViewFeature extends StatelessWidget {
  GeneralViewFeature({Key? key}) : super(key: key);
  final Completer<GoogleMapController> mapController = Completer();

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: GeneralViewBody(
          mapController: mapController,
        ),
        bottomNavigationBar: BottomBar(),
      ),
    );
  }
}

class GeneralViewBody extends StatelessWidget {
  final Completer<GoogleMapController> mapController;
  const GeneralViewBody({Key? key, required this.mapController}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          child: GeneralViewHeader(),
          margin: const EdgeInsets.all(10),
        ),
        Expanded(
          child: ClipRRect(
            borderRadius: BorderRadius.only(
              topLeft: Radius.circular(25),
              topRight: Radius.circular(25),
              bottomLeft: Radius.circular(25),
            ),
            child: MapSample(
              mapController: mapController 
            ),
          ),
        )
      ],
    );
  }
}
