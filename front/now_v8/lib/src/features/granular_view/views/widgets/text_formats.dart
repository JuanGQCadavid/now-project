import 'package:flutter/material.dart';

class CreatorLabel extends StatelessWidget {
  final EdgeInsets generalMaring;
  final void Function() onTap;
  final String unhighlightedText;
  final String highlightedText;
  final Color appColor;
  const CreatorLabel({
    Key? key,
    this.generalMaring = const EdgeInsets.symmetric(horizontal: 15),
    this.unhighlightedText = "Hosted by",
    required this.highlightedText,
    required this.onTap,
    required this.appColor

  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: generalMaring,
      child: InkWell(
        onTap: onTap,
        child: Row(
          children: [
            Text(unhighlightedText),
            const SizedBox(
              width: 10,
            ),
            Container(
              padding: const EdgeInsets.all(5),
              decoration: BoxDecoration(
                  color: appColor,
                  borderRadius: BorderRadius.circular(10)),
              child: Text(
                highlightedText,
                style: Theme.of(context)
                    .textTheme
                    .bodyMedium!
                    .copyWith(color: Colors.white),
              ),
            )
          ],
        ),
      ),
    );
  }
}

class ReadMoreBox extends StatelessWidget {
  final EdgeInsets generalMaring;
  final String textBody;
  const ReadMoreBox({
    Key? key,
    required this.textBody,
    this.generalMaring = const EdgeInsets.symmetric(horizontal: 15),
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: generalMaring,
      child: Column(children: [
        Text(
          textBody
        ),
      ]),
    );
  }
}