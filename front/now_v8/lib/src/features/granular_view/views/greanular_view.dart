import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/header.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/host_updates.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/place_label.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/tags_list.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/text_formats.dart';
import 'dart:math';

import 'package:now_v8/src/features/granular_view/views_model/providers.dart';

class GranularView extends StatelessWidget {
  GranularView({Key? key}) : super(key: key);
  final Color appColor = Colors.cyan.shade400;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: _Body(
          appColor: appColor,
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
      ),
    );
  }
}

class _Body extends ConsumerWidget {
  final Color appColor;
  const _Body({Key? key, required this.appColor}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final onSpot = ref.watch(onSpotProvider);

    if (onSpot.window.isEmpty()) {
      return Container(
        child: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            new CircularProgressIndicator(),
            new Text("Loading"),
          ],
        ),
      );

    } else {
      return Container(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            GanularHeader(
              spotWindow: onSpot.window,
              appColor: appColor,
            ),
            Center(
              child: PlaceLabel(
                placeName: onSpot.spot.placeInfo.name,
                appColor: appColor,
              ),
            ),
            const SizedBox(
              height: 15,
            ),
            CreatorLabel(
              onTap: () {},
              highlightedText: onSpot.spot.hostInfo.name,
              appColor: appColor,
            ),
            const SizedBox(
              height: 15,
            ),
            ReadMoreBox(
              textBody: onSpot.spot.eventInfo.description,
            ),
            const SizedBox(
              height: 15,
            ),
            TagsList(
                primaryTag: onSpot.spot.topicInfo.principalTag,
                secondaryTags: onSpot.spot.topicInfo.secondaryTags,
                appColor: appColor),
            const SizedBox(
              height: 15,
            ),
            Container(
              margin: EdgeInsets.symmetric(horizontal: 15),
              child: const Divider(),
            ),
            const SizedBox(
              height: 15,
            ),
            const HostUpdates()
          ],
        ),
      );
    }
  }
}
