import 'package:flutter/material.dart';
import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';

class TagsList extends StatelessWidget {
  final EdgeInsets generalMaring;
  final String primaryTag;
  final List<String> secondaryTags;
  final String emptyMessage;
  final Color appColor;

  const TagsList({
    Key? key,
    this.primaryTag = "",
    this.secondaryTags = const [],
    this.emptyMessage = "No tags",
    required this.appColor,
    this.generalMaring = const EdgeInsets.symmetric(horizontal: 15),
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    if (primaryTag.isEmpty && secondaryTags.isEmpty) {
      return Container(
        child: Center(
          child: Text(
            emptyMessage,
            style: Theme.of(context)
                .textTheme
                .bodyMedium!
                .copyWith(fontStyle: FontStyle.italic),
          ),
        ),
      );
    }

    List<Widget> children = <Widget>[];

    if (primaryTag.isNotEmpty) {
      children.add(Container(
        margin: EdgeInsets.only(right: 15),
        child: SpotTag(
          tag: primaryTag,
          color: appColor,
          onPressed: () {},
        ),
      ));
    }

    if (secondaryTags.isNotEmpty) {
      secondaryTags.forEach(
        (tag) {
          children.add(
            Container(
              margin: EdgeInsets.only(right: 15),
              child: SpotTag(
                tag: tag,
                color: appColor,
                onPressed: () {},
              ),
            ),
          );
        },
      );
    }

    return Container(
      margin: generalMaring,
      child: SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: Row(
          children: children,
        ),
      ),
    );
  }
}
