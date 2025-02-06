import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/features/general_view/views/widgets/header.dart';
import 'package:now_v8/src/features/general_view/views/widgets/map.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';
import 'package:now_v8/src/features/granular_view/views/main.dart';
import 'package:now_v8/src/features/login/view/main.dart';
import 'package:now_v8/src/features/profile/view/main.dart';

class GeneralViewFeature extends ConsumerWidget {
  GeneralViewFeature({super.key});
  final Completer<GoogleMapController> mapController = Completer();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    if (ref.watch(
      lastPositionKnownStateProvider.select((state) => state.jump),
    )) {
      SchedulerBinding.instance.addPostFrameCallback((_) {
        Navigator.push(
          context,
          MaterialPageRoute(
            builder: (context) => GranularView(),
          ),
        );
      });
    }
    return Scaffold(
      body: SafeArea(
        bottom: false,
        child: GeneralViewBody(
          mapController: mapController,
        ),
      ),
    );
  }
}

class GeneralViewBody extends StatelessWidget {
  final Completer<GoogleMapController> mapController;
  const GeneralViewBody({super.key, required this.mapController});

  void openMenu() {}

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          margin: const EdgeInsets.symmetric(horizontal: 10),
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
            child: GeneralViewMap(mapController: mapController),
          ),
        )
      ],
    );
  }
}
