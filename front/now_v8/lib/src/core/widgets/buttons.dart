import 'package:flutter/material.dart';

class NowTextCTA extends StatelessWidget {
  final String ctaText;
  final void Function() onPressed;
  final Color textColot;
  NowTextCTA(
      {Key? key,
      required this.ctaText,
      required this.onPressed,
      this.textColot = Colors.black})
      : super(key: key);

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

class IconTextButtom extends StatelessWidget {
  final String message;
  final Color mainColor;
  final IconData icon;
  final Color iconColor;
  final bool addUnderline;
  final Function() onTap;

  IconTextButtom(
      {Key? key,
      required this.message,
      required this.mainColor,
      required this.icon,
      required this.iconColor,
      required this.onTap,
      this.addUnderline = false})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onTap,
      child: Container(
        margin: EdgeInsets.all(15),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            Container(
              padding: EdgeInsets.all(5),
              decoration: BoxDecoration(
                color: mainColor,
                borderRadius: BorderRadius.circular(50),
              ),
              child: Icon(
                icon,
                color: iconColor,
              ),
            ),
            const SizedBox(
              width: 10,
            ),
            Container(
              decoration: BoxDecoration(
                border: addUnderline
                    ? Border(
                        bottom: BorderSide(color: mainColor, width: 2.0),
                      )
                    : const Border(),
              ),
              child: Text(
                message,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
