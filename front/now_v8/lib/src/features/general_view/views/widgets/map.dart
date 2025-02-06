import 'dart:developer';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/material.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/features/general_view/views/widgets/footbar.dart';
import 'dart:async';

import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';
import 'package:now_v8/src/features/granular_view/views/main.dart';
import 'package:now_v8/src/features/spots_creation/main.dart';

class GeneralViewMap extends ConsumerWidget {
  final String filterMessage = "Filter events";
  final String createMessage = "Create event";
  final String lookCloserMessage = "Look closer";

  final Completer<GoogleMapController> mapController;
  const GeneralViewMap({super.key, required this.mapController});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final filteredSpots = ref.watch(filteredSpotsProvider);
    final mapInteraction = ref.read(mapInteractionProvider.notifier);

    return Stack(
      children: [
        NowMapV2(
          spots: filteredSpots.spots,
          centerMapOnSpots: true,
          mapController: mapController,
          onCameraIdle: mapInteraction.onCameraIdle,
          onCameraMove: mapInteraction.onCameraMove,
          onCameraMoveStarted: mapInteraction.onCameraMoveStarted,
          padding: const EdgeInsets.only(bottom: 100),
        ),
        Align(
          alignment: Alignment.bottomLeft,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              MapTags(
                filteredSpots: filteredSpots,
              ),
              FooterGeneralView(
                filterMessage: filterMessage,
                createMessage: createMessage,
                lookCloserMessage: lookCloserMessage,
                onFilterPressed: () {
                  log(filterMessage);
                },
                onCreatePressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) => const SpotsCreationFeature(),
                    ),
                  );
                },
                onLookCloserPressed: () {
                  Navigator.of(context).push(
                    MaterialPageRoute(
                      builder: (context) => GranularView(),
                    ),
                  );
                },
              ),
            ],
          ),
        )
      ],
    );
  }
}

class MapTags extends ConsumerWidget {
  final FilteredSpots filteredSpots;

  const MapTags({
    required this.filteredSpots,
    super.key,
  });

  void onTagClick(WidgetRef ref, String tag) {
    final tagsNotifier = ref.read(tagsSelectedProvider.notifier);
    tagsNotifier.tagSelected(tag);
  }

  void onClearButtom(WidgetRef ref) {
    final tagsNotifier = ref.read(tagsSelectedProvider.notifier);
    tagsNotifier.cleanTags();
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    List<Widget> rowTags = [];

    if (filteredSpots.tagsSelected.isNotEmpty) {
      // Clear buttom
      rowTags.add(clearButtom(ref));

      // First Adding the tags that are selected.
      for (var tag in filteredSpots.tagsSelected) {
        rowTags.add(
          generateTag(filteredSpots.onFilterColor.color, tag, ref),
        );
      }

      // Then Adding the ones that are not selected as gray

      for (var tag in filteredSpots.tagsOff) {
        rowTags.add(
          generateTag(Colors.black38, tag, ref),
        );
      }
    } else {
      Map<String, Color> tags = {};

      for (var spot in filteredSpots.spots) {
        if (!tags.containsKey(spot.principalTag)) {
          tags[spot.principalTag] = spot.spotsColor.color;
        }
        for (var secondaryTag in spot.secondaryTags) {
          if (!tags.containsKey(secondaryTag)) {
            tags[secondaryTag] = spot.spotsColor.color;
          }
        }
      }

      tags.forEach((tag, color) {
        rowTags.add(generateTag(color, tag, ref));
      });
    }

    return Container(
      margin: const EdgeInsets.only(left: 15, bottom: 15),
      child: SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: Row(
          children: rowTags,
        ),
      ),
    );
  }

  Widget generateTag(Color color, String tag, WidgetRef ref) {
    return Container(
      margin: const EdgeInsets.only(right: 15),
      child: SpotTag(
        color: color,
        tag: tag,
        onPressed: () => onTagClick(ref, tag),
      ),
    );
  }

  Widget clearButtom(WidgetRef ref) {
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(50),
        color: Colors.white,
      ),
      margin: const EdgeInsets.only(right: 15),
      child: IconButton(
        onPressed: () => onClearButtom(ref),
        icon: const Icon(Icons.delete_outline),
      ),
    );
  }
}
