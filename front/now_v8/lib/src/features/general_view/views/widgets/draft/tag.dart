import 'package:flutter/material.dart';

class SpotTagV2 extends StatelessWidget {
  final Color bgColor = Colors.green.shade900;
  final Color textColor = Colors.white;

  final String tag = "CoffeLovers";

  SpotTagV2({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(5),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(25),
        color: bgColor,
      ),
      //border: Border.all(color: color)),
      child: Text(
        "#${tag}",
        style: TextStyle(color: textColor),
      ),
    );
  }
}