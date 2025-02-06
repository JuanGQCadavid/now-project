import 'dart:async';
import 'dart:developer';
import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/granular_view/views/main.dart';
import 'package:now_v8/src/features/spots_creation/main.dart';

class HomeTest extends StatelessWidget {
  HomeTest({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: AllTogether(),
      // floatingActionButton: FooterGeneralView(),
    );
  }
}

class AllTogether extends StatelessWidget {
  AllTogether({super.key});
  final Completer<GoogleMapController> mapController = Completer();

  @override
  Widget build(BuildContext context) {
    return Stack(
      fit: StackFit.expand,
      children: [
        NowMapV2(
          mapController: mapController,
          myLocationButtonEnabled: false,
        ),
        Align(
          child: FooterGeneralView(),
          alignment: Alignment.bottomCenter,
        ),
      ],
    );
  }
}

class FooterGeneralView extends StatelessWidget {
  const FooterGeneralView({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      child: ClipRRect(
          borderRadius: BorderRadius.circular(10),
          child: BackdropFilter(
            filter: ImageFilter.blur(sigmaX: 10, sigmaY: 10),
            child: BottomNavigationBar(
              backgroundColor: Colors.white.withAlpha(175),
              elevation: 0,
              unselectedItemColor: Colors.black,
              currentIndex: 1,
              items: const [
                BottomNavigationBarItem(
                  icon: Icon(Icons.tune_outlined),
                  label: "Filter events",
                ),
                BottomNavigationBarItem(
                    icon: Icon(Icons.add_circle_outline),
                    label: "Create event"),
                BottomNavigationBarItem(
                  icon: Icon(Icons.zoom_in),
                  label: "Look closer",
                ),
              ],
            ),
          )),
    );
  }
}

class RowOption extends StatelessWidget {
  const RowOption({super.key});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceAround,
      children: [
        const Expanded(
          flex: 1,
          child: SizedBox(),
        ),
        Expanded(
            flex: 1,
            child: TextButton.icon(
              onPressed: () {},
              label: const Text('OutlinedButton'),
              icon: const Icon(
                Icons.add_sharp,
              ),
            )),
        Expanded(
          flex: 1,
          child: NowTextCTA(
            ctaText: "Zoom in!",
            onPressed: () {
              log("Zoom in");
            },
          ),
        )
      ],
    );
  }
}
