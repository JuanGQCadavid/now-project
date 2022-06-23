import 'dart:developer';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:flutter/material.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'dart:async';

import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';

class MapSample extends ConsumerWidget {
  const MapSample({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final filteredSpots = ref.watch(filteredSpotsProvider);

    return Stack(
      children: [
        NowMap(
          filteredSpots: filteredSpots,
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
      filteredSpots.tagsSelected.forEach((tag) {
        rowTags.add(
          generateTag(filteredSpots.onFilterColor.color, tag, ref),
        );
      });

      // Then Adding the ones that are not selected as gray

      filteredSpots.tagsOff.forEach((tag) {
        rowTags.add(
          generateTag(Colors.black38, tag, ref),
        );
      });
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
      margin: EdgeInsets.only(right: 15),
      child: IconButton(
        onPressed: () => onClearButtom(ref),
        icon: const Icon(Icons.delete_outline),
      ),
    );
  }
}

class NowMap extends StatefulWidget {
  final FilteredSpots filteredSpots;
  NowMap({
    Key? key,
    required this.filteredSpots,
  }) : super(key: key);

  @override
  State<NowMap> createState() => _NowMapState();
}

class _NowMapState extends State<NowMap> {
  Completer<GoogleMapController> _controller = Completer();
  // final Completer<GoogleMapController> _controller;

  static final CameraPosition _kGooglePlex = CameraPosition(
    target: LatLng(6.251723, -75.592771),
    zoom: 14.4746,
  );

  @override
  Widget build(BuildContext context) {
    Set<Marker> markers = Set();

    widget.filteredSpots.spots.forEach((spot) {
      markers.add(
        Marker(
            markerId: MarkerId(spot.spotId),
            position: spot.latLng,
            visible: true,
            icon: widget.filteredSpots.tagsSelected.isEmpty
                ? spot.spotsColor.hue
                : widget.filteredSpots.onFilterColor.hue,
            infoWindow: InfoWindow(title: spot.principalTag)),
      );
    });

    return GoogleMap(
      markers: markers,
      mapType: MapType.normal,
      initialCameraPosition: _kGooglePlex,
      mapToolbarEnabled: false,
      myLocationButtonEnabled: false,
      padding: EdgeInsets.only(bottom: 65, left: 15),
      // onMapCreated: (GoogleMapController controller) {
      //   _controller.complete(controller);
      // },
    );
  }
}