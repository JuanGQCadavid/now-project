import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/features/granular_view/model/granular_spot.dart';

class GanularHeader extends StatelessWidget {
  final double headerSize = 300;
  final double mapSize = 250;
  final double spotsHeaderSize = 80;
  final Color appColor;
  final GranularSpotWindow spotWindow = const GranularSpotWindow(
      actualOne: "Actual One",
      nextOne: "Next one",
      previousOne: "Previous one");

  const GanularHeader({Key? key, required this.appColor}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: headerSize,
      child: Stack(
        children: [
          Container(
            height: mapSize,
            color: appColor,
          ),
          Align(
            alignment: Alignment.bottomLeft,
            child: Container(
              height: spotsHeaderSize,
              color: Theme.of(context).scaffoldBackgroundColor,
              child: Row(
                children: [
                  Expanded(
                    flex: 1,
                    child: TextHeaderOption(
                      onPressed: () {},
                      text: spotWindow.previousOne,
                      isPrincipal: false,
                    ),
                  ),
                  Expanded(
                    flex: 2,
                    child: TextHeaderOption(
                      onPressed: () {},
                      text: spotWindow.actualOne,
                      isPrincipal: true,
                    ),
                  ),
                  Expanded(
                    flex: 1,
                    child: TextHeaderOption(
                      onPressed: () {},
                      text: spotWindow.nextOne,
                      isPrincipal: false,
                    ),
                  )
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}

class TextHeaderOption extends StatelessWidget {
  final bool isPrincipal;
  final String text;
  final void Function() onPressed;

  const TextHeaderOption({
    Key? key,
    required this.text,
    required this.isPrincipal,
    required this.onPressed,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        margin: const EdgeInsets.symmetric(horizontal: 5),
        child: TextButton(
          onPressed: onPressed,
          child: Text(
            text,
            maxLines: 3,
            textAlign: TextAlign.center,
            overflow: TextOverflow.ellipsis,
            style: isPrincipal
                ? Theme.of(context).textTheme.bodyLarge
                : Theme.of(context).textTheme.bodyMedium,
          ),
        ),
      ),
    );
  }
}
