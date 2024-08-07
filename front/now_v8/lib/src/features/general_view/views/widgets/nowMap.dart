import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/material.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'dart:async';

import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';
import 'package:now_v8/src/services/core/providers.dart';

class MapSample extends ConsumerWidget {
  final Completer<GoogleMapController> mapController;
  const MapSample({Key? key, required this.mapController}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final filteredSpots = ref.watch(filteredSpotsProvider);
    final mapInteraction = ref.read(mapInteractionProvider.notifier);
    const location = LatLng(6.251723, -75.592771);

    // if (widget.includeUserLocation != null || widget.centerMapOnSpots != true) {
    //   locationData = await location.getLocation();
    // }

    return Stack(
      children: [
        NowMapV2(
          spots: filteredSpots.spots,
          centerMapOnSpots: true,
          mapController: mapController, 
          camaraPosition: location,
          onCameraIdle: mapInteraction.onCameraIdle,
          onCameraMove: mapInteraction.onCameraMove,
          onCameraMoveStarted: mapInteraction.onCameraMoveStarted,
        ),
        Align(
          alignment: Alignment.bottomLeft,
          child: MapTags(
            filteredSpots: filteredSpots,
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
    Key? key,
  }) : super(key: key);

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

      filteredSpots.spots.forEach((spot) {
        Color tagColor = Colors.black;

        if (!tags.containsKey(spot.principalTag)) {
          tags[spot.principalTag] = spot.spotsColor.color;
        }

        spot.secondaryTags.forEach((secondaryTag) {
          if (!tags.containsKey(secondaryTag)) {
            tags[secondaryTag] = spot.spotsColor.color;
          }
        });
      });

      tags.forEach((tag, color) {
        rowTags.add(generateTag(color, tag, ref));
      });
    }

    return Container(
      margin: EdgeInsets.only(left: 15, bottom: 15),
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
      margin: EdgeInsets.only(right: 15),
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