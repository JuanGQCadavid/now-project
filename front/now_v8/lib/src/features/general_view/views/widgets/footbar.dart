import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';
import 'package:now_v8/src/features/granular_view/views/main.dart';
import 'package:now_v8/src/features/granular_view/views_model/providers.dart';
import 'package:now_v8/src/features/spots_creation/main.dart';

class BottomBar extends ConsumerWidget {
  const BottomBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Container(
      margin: const EdgeInsets.all(15),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          NowTextCTA(
            ctaText: "Filter",
            onPressed: () {},
          ),
          NowTextCTA(
            ctaText: "Create",
            onPressed: () {
              // HEREEEE
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => const SpotsCreationFeature(),
                ),
              );
            },
          ),
          NowTextCTA(
            ctaText: "Zoom in!",
            onPressed: () {
              Navigator.of(context).push(
                  MaterialPageRoute(builder: (context) => GranularView()));
            },
          )
        ],
      ),
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
