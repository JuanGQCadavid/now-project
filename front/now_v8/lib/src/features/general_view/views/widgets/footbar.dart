import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';
import 'package:now_v8/src/features/granular_view/views/greanular_view.dart';
import 'package:now_v8/src/features/granular_view/views_model/providers.dart';
import 'package:now_v8/src/features/spots_creation/main.dart';

class BottomBar extends ConsumerWidget {
  const BottomBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Container(
      margin: EdgeInsets.all(15),
      //color: Colors.blue,
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
              final detailedSpot = ref.read(detailedSpotProvider.notifier);
              detailedSpot.refreshSpots();

              Navigator.of(context).push(
                  MaterialPageRoute(builder: (context) => GranularView()));
            },
          )
        ],
      ),
    );
  }
}
