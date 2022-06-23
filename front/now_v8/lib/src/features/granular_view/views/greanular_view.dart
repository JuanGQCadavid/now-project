import 'package:flutter/material.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/header.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/place_label.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/tags_list.dart';
import 'package:now_v8/src/features/granular_view/views/widgets/text_formats.dart';

class GranularView extends StatelessWidget {
  const GranularView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SafeArea(
        child: Scaffold(
      body: _Body(),
    ));
  }
}

class _Body extends StatelessWidget {
  const _Body({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const GanularHeader(),
        const Center(
          child: PlaceLabel(
            placeName: "La tia o ladraba",
          ),
        ),
        const SizedBox(
          height: 15,
        ),
        CreatorLabel(
          onTap: () {},
          highlightedText: "Juan Gonzalo",
        ),
        const SizedBox(
          height: 15,
        ),
        const ReadMoreBox(
          textBody: "adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf adsfdasfdasfadsf dsafadsfadsfadsfads fadsfadsfadsfdas fadsf adsfadsfadsfadsf dasfdsafa",
        ),
        const SizedBox(
          height: 15,
        ),
        const TagsList(
          primaryTag: "PrimaryTag",
          secondaryTags: [
            "second1", "second2"
          ],
        )
      ],
    ));
  }
}