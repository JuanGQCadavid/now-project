import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/granular_view/model/granular_spot.dart';
import 'package:now_v8/src/features/granular_view/views_model/providers.dart';

class GanularHeader extends ConsumerWidget {
  final double headerSize = 300;
  final double mapSize = 250;
  final double spotsHeaderSize = 80;
  final Color appColor;
  final GranularSpot onSpot;
  GanularHeader({Key? key, required this.appColor, required this.onSpot})
      : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Container(
      height: headerSize,
      child: Stack(
        children: [
          Container(
            height: mapSize,
            color: appColor,
            child: NowMapV2(
              blockMap: true,
              myLocationButtonEnabled: false,
              centerMapOnSpots: true,
              includeUserLocation: true,
              spots: [
                Spot.fromLongSpot(onSpot.spot),
              ],
            ),
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
                    child: onSpot.window.previousOne.isEmpty
                        ? Container()
                        : TextHeaderOption(
                            onPressed: () {
                              final onSpot = ref.read(onSpotProvider.notifier);
                              onSpot.previousOne();
                            },
                            text: onSpot.window.previousOne,
                            isPrincipal: false,
                          ),
                  ),
                  Expanded(
                    flex: 2,
                    child: TextHeaderOption(
                      onPressed: () {},
                      text: onSpot.window.actualOne,
                      isPrincipal: true,
                    ),
                  ),
                  Expanded(
                    flex: 1,
                    child: onSpot.window.nextOne.isEmpty
                        ? IconButton(
                            onPressed: () {
                              final onSpot = ref.read(onSpotProvider.notifier);
                              onSpot.nextOne();
                            },
                            icon: Icon(Icons.refresh))
                        : TextHeaderOption(
                            onPressed: () {
                              final onSpot = ref.read(onSpotProvider.notifier);
                              onSpot.nextOne();
                            },
                            text: onSpot.window.nextOne,
                            isPrincipal: false,
                          ),
                  )
                ],
              ),
            ),
          ),
          Align(
              alignment: Alignment.topLeft,
              child: Container(
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(50),
                  color: Colors.white,
                ),
                child: IconButton(
                  icon: Icon(
                    Icons.arrow_back_ios_new,
                    size: 15,
                  ),
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                ),
              ))
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
          onPressed: isPrincipal ? null : onPressed,
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
