import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:now_v2/core/domain/models/spot.dart';
import '../../features/filters/presentation/map.dart';
import 'spots_detail_view.dart';
import 'package:latlong2/latlong.dart';

class SpotGranularView extends StatelessWidget {
  final Spot centerSpot;

  const SpotGranularView({
    Key? key,
    required this.centerSpot,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _Page(centerSpot: centerSpot),
      floatingActionButton: FloatingActionButton(onPressed: () {
        Navigator.push(context,
            MaterialPageRoute(builder: (context) => const SpotDetaillView()));
      }),
    );
  }
}

class _Page extends StatefulWidget {
  final Spot centerSpot;
  const _Page({
    Key? key,
    required this.centerSpot,
  }) : super(key: key);

  @override
  __PageState createState() => __PageState();
}

class __PageState extends State<_Page> {
  final iconSize = 30.0;

  @override
  Widget build(BuildContext context) {
    final description =
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur interdum volutpat mi, vitae hendrerit nisi gravida sit amet. In sodales tellus erat, quis iaculis nulla commodo a. Sed id sem nibh. Integer ullamcorper sollicitudin nunc, non accumsan massa dapibus consectetur. Integer aliquet laoreet consequat. Fusce malesuada ligula nibh, sed convallis erat suscipit id. Nullam consectetur eros commodo iaculis tempus. Vestibulum dignissim nibh ante, id dictum ex mollis non. Morbi nisi velit, molestie eu iaculis at, aliquam sit amet diam. Aenean posuere urna magna, id ullamcorper nunc tristique sed. Donec ultrices ipsum eget turpis rhoncus, vitae pharetra odio molestie.";

    final creator = "Juan Gonzalo";

    return SafeArea(
      child: Column(
        children: [
          // HEADER
          Container(
            margin: EdgeInsets.all(10),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                // TITLE and buttom
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    IconButton(
                      onPressed: () {
                        Navigator.of(context).pop();
                      },
                      icon: Icon(
                        Icons.arrow_back,
                        size: iconSize,
                      ),
                    ),
                    Text(
                      widget.centerSpot.eventInfo.name,
                    ),
                    Icon(
                      Icons.menu,
                      size: iconSize,
                    ),
                  ],
                ),

                // User info!
                Center(
                  child: Text("Creado por $creator"),
                ),

                const SizedBox(
                  height: 10,
                ),

                // Description place

                Container(
                  margin: EdgeInsets.only(
                    left: iconSize,
                    right: iconSize,
                  ),
                  child: Text(
                    "Descripcion: $description",
                    overflow: TextOverflow.ellipsis,
                    maxLines: 4,
                  ),
                )
              ],
            ),
          ),

          // Map
          Container(
            child: Expanded(
              child: AppMap(
                markers: List.empty(),
              ),
            ),
          ),

          // Foother
          Container(),
        ],
      ),
    );
  }
}
