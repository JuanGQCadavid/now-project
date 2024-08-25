import 'package:flutter/material.dart';

class SpotTag extends StatelessWidget {
  final Color color;
  final Color bgColor;
  final String tag;
  final void Function() onPressed;

  SpotTag(
      {Key? key,
      required this.tag,
      this.bgColor = Colors.white,
      required this.color,
      required this.onPressed})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(5),
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(10),
          border: Border.all(color: color),
          color: bgColor),
      child: InkWell(
        onTap: onPressed,
        child: Text(
          "#${tag}",
          style: TextStyle(color: color),
        ),
      ),
    );
  }
}
