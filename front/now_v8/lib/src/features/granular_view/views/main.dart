import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/header.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/host_updates.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/loaders.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/place_label.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/tags_list.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/text_formats.dart';
import 'dart:math';
import 'dart:async';
import 'package:now_v8/src/features/granular_view/views_model/providers.dart';

class GranularView extends StatelessWidget {
  GranularView({Key? key}) : super(key: key);
  final Color appColor = Colors.cyan.shade400;
  late Completer<GoogleMapController> mapController = Completer();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: _Body(
          appColor: appColor,
          mapController: mapController,
        ),
      ),
      floatingActionButton: Wrap(
        direction: Axis.vertical,
        children: [
          Container(
            margin: const EdgeInsets.only(top: 15),
            child: FloatingActionButton(
              heroTag: "message",
              backgroundColor: appColor,
              onPressed: () {},
              child: const Icon(
                Icons.message,
              ),
            ),
          ),
          Container(
            margin: const EdgeInsets.only(top: 15),
            child: FloatingActionButton(
              heroTag: "go_back",
              backgroundColor: appColor,
              onPressed: () {},
              child: Transform.rotate(
                angle: -pi,
                child: const Icon(
                  Icons.arrow_back,
                ),
              ),
            ),
          )
        ],
      ),
    );
  }
}

class _Body extends ConsumerWidget {
  final Color appColor;
  final Completer<GoogleMapController> mapController;
  final spacer = const SizedBox(
    height: 15,
  );

  const _Body({Key? key, required this.appColor, required this.mapController})
      : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final onSpot = ref.watch(onSpotProvider);

    if (onSpot.window.isEmpty()) {
      // TODO - Should we leave this here ?
      final detailedSpot = ref.read(detailedSpotProvider.notifier);
      detailedSpot.refreshSpots();
      return const FindingSpotsLoadingScreen();
    }

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        GanularHeader(
          onSpot: onSpot,
          appColor: appColor,
          mapController: mapController,
        ),
        Center(
          child: PlaceLabel(
            placeName: onSpot.spot.placeInfo.name,
            appColor: appColor,
          ),
        ),
        spacer,
        CreatorLabel(
          onTap: () {},
          highlightedText: onSpot.spot.hostInfo.name,
          appColor: appColor,
        ),
        spacer,
        ReadMoreBox(
          textBody: onSpot.spot.eventInfo.description,
        ),
        spacer,
        Wrap(
          direction: Axis.horizontal,
          children: [
            IconTextButtom(
              message: "Instagram",
              mainColor: Colors.grey.shade700,
              icon: Icons.link,
              iconColor: Colors.white,
              onTap: () {},
            ),
            IconTextButtom(
              message: "WhatsApp",
              mainColor: Colors.green.shade500,
              icon: Icons.phone,
              iconColor: Colors.white,
              onTap: () {},
            ),
            IconTextButtom(
              message: "Call me",
              mainColor: Colors.blue.shade500,
              icon: Icons.phone,
              iconColor: Colors.white,
              onTap: () {},
            ),
          ],
        ),
        spacer,
        TagsList(
            primaryTag: onSpot.spot.topicInfo.principalTopic,
            secondaryTags: onSpot.spot.topicInfo.secondaryTopics,
            appColor: appColor),
        spacer,
        Container(
          margin: const EdgeInsets.symmetric(
            horizontal: 15,
          ),
          child: const Divider(),
        ),
        spacer,
        const HostUpdates()
      ],
    );
  }
}
