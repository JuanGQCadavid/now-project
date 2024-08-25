import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/features/general_view/model/general_view_model.dart';
import 'package:now_v8/src/features/general_view/views/widgets/footbar.dart';
import 'package:now_v8/src/features/general_view/views/widgets/header.dart';
import 'package:now_v8/src/features/general_view/views/widgets/map.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';
import 'package:now_v8/src/features/login/view/main.dart';
import 'package:now_v8/src/features/profile/view/main.dart';

class GeneralViewFeature extends ConsumerWidget {
  GeneralViewFeature({Key? key}) : super(key: key);
  final Completer<GoogleMapController> mapController = Completer();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Scaffold(
      body: SafeArea(
        child: GeneralViewBody(
          mapController: mapController,
        ),
      ),
      bottomNavigationBar: const BottomBar(),
    );
  }
}

class GeneralViewBody extends StatelessWidget {
  final Completer<GoogleMapController> mapController;
  const GeneralViewBody({Key? key, required this.mapController})
      : super(key: key);

  void openMenu() {}

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          margin: const EdgeInsets.all(10),
          child: GeneralViewHeader(
            onRequestToGoToMenu: openMenu,
            onRequestToGoToProfile: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => const ProfileFeature(),
                ),
              );
            },
            onRequestToLogin: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => LoginFeature(),
                ),
              );
            },
          ),
        ),
        Expanded(
          child: ClipRRect(
            borderRadius: const BorderRadius.only(
              topLeft: Radius.circular(25),
              topRight: Radius.circular(25),
              bottomLeft: Radius.circular(25),
            ),
            child: Stack(
              children: [
                GeneralViewMap(mapController: mapController),
              ],
            ),
          ),
        )
      ],
    );
  }
}
