import 'package:flutter/material.dart';

class NowTextCTA extends StatelessWidget {
  final String ctaText;
  final void Function() onPressed;
  final Color textColot ;
  NowTextCTA({
    Key? key,
    required this.ctaText,
    required this.onPressed,
    this.textColot = Colors.black
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: TextButton(
        child: Text(
          ctaText,
          style: TextStyle(color: textColot),
        ),
        onPressed: onPressed,
      ),
    );
  }
}
