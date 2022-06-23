import 'package:flutter/material.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/header.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/host_updates.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/place_label.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/tags_list.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/text_formats.dart';
import 'dart:math';

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

class _Body extends StatelessWidget {
  final Color appColor;
  const _Body({Key? key, required this.appColor}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          GanularHeader(
            appColor: appColor,
          ),
          Center(
            child: PlaceLabel(
              placeName: "La tia o ladraba",
              appColor: appColor,
            ),
          ),
          const SizedBox(
            height: 15,
          ),
          CreatorLabel(
            onTap: () {},
            highlightedText: "Juan Gonzalo",
            appColor: appColor,
          ),
          const SizedBox(
            height: 15,
          ),
          const ReadMoreBox(
            textBody:
                "adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf dsafadsfadsfadsfads fadsfadsfadsfdas fadsf adsfadsfadsfadsf dasfdsafa",
          ),
          const SizedBox(
            height: 15,
          ),
          TagsList(
              primaryTag: "PrimaryTag",
              secondaryTags: ["second1", "second2"],
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


