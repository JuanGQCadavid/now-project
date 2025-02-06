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
        margin: const EdgeInsets.all(15),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          children: [
            Container(
              padding: const EdgeInsets.all(5),
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

class UserLoggedButton extends StatelessWidget {
  final void Function() onTap;
  final String displayName;
  final List<Color> colors;

  const UserLoggedButton({
    super.key,
    required this.onTap,
    required this.displayName,
    required this.colors,
  });

  String genTwoLettersName(String name) {
    if (name.length >= 2) {
      var splits = name.split(" ");

      if (splits.length >= 2) {
        return splits[0].toUpperCase()[0] + splits[1].toUpperCase()[0];
      }

      return splits[0].toUpperCase()[0] + splits[0].toUpperCase()[1];
    }

    if (name.isEmpty) {
      return "";
    }

    return name.toUpperCase();
  }

  @override
  Widget build(BuildContext context) {
    var twoLettersName = genTwoLettersName(displayName);
    Widget icon = SizedBox(
      width: 25,
      height: 25,
      child: Center(
        child: Text(
          twoLettersName,
          style: Theme.of(context).textTheme.labelLarge,
        ),
      ),
    );

    if (twoLettersName.isEmpty) {
      icon = const Icon(Icons.person);
    }

    return InkWell(
      onTap: onTap,
      child: ClipOval(
        child: Container(
          padding: const EdgeInsets.all(2),
          decoration: BoxDecoration(
            gradient: LinearGradient(
              begin: Alignment.topLeft,
              end: Alignment.bottomRight,
              colors: colors,
            ),
          ),
          child: ClipOval(
            child: Container(
              padding: const EdgeInsets.all(6),
              color: Colors.white,
              child: icon,
            ),
          ),
        ),
      ),
    );
  }
}
