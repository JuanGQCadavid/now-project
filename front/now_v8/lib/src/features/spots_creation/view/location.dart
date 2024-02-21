import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/login/view/widgets/text_input.dart';

class LocationSeletorViewV2 extends ConsumerWidget {
  const LocationSeletorViewV2({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return const Placeholder();
  }
}

class LocationSelectorView extends StatefulWidget {
  final Future<List<PlaceInfo>> Function(String) onSearch;
  final void Function(PlaceInfo) onChosen;
  late PlaceInfo placeSelected;

  LocationSelectorView({
    super.key,
    required this.onChosen,
    required this.onSearch,
    required this.placeSelected,
  });

  @override
  State<LocationSelectorView> createState() => _LocationSelectorViewState();
}

class _LocationSelectorViewState extends State<LocationSelectorView> {
  late GoogleMapController mapController;

  void onChosen(PlaceInfo place) {
    setState(() {
      widget.placeSelected = place;
    });
    widget.onChosen(place);

    var bounds = MapUtilities.getCameraLatLngBounds(
      [
        Spot.withOutSpotColors(
          principalTag: "",
          secondaryTags: [],
          latLng: LatLng(
            widget.placeSelected.lat,
            widget.placeSelected.lon,
          ),
          spotId: "",
          date: DateTime.now(),
        ),
      ],
    );
    mapController.animateCamera(
      CameraUpdate.newLatLngBounds(
        bounds,
        50,
      ),
    );
  }

  String resume() {
    return widget.placeSelected.name.replaceFirst(" -#AT#- ", "\n");
  }

  @override
  Widget build(BuildContext context) {
    String selectedResume = resume();
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: ClipRRect(
        borderRadius: const BorderRadius.all(
          Radius.circular(15),
        ),
        child: Column(
          children: [
            Stack(
              children: [
                SizedBox(
                  height: 600,
                  width: double.infinity,
                  child: NowMapV2(
                    centerMapOnSpots: true,
                    includeUserLocation: false,
                    camaraPosition: LatLng(
                      widget.placeSelected.lat,
                      widget.placeSelected.lon,
                    ),
                    mapController: Completer(),
                    onMapCreated: (mapController) {
                      this.mapController = mapController;
                    },
                    spots: [
                      Spot.withOutSpotColors(
                        principalTag: "",
                        secondaryTags: [],
                        latLng: LatLng(
                          widget.placeSelected.lat,
                          widget.placeSelected.lon,
                        ),
                        spotId: "",
                        date: DateTime.now(),
                      )
                    ],
                  ),
                ),
                SizedBox(
                  child: SeachLocationSF(
                    onSearch: widget.onSearch,
                    onChosen: onChosen,
                  ),
                  width: double.infinity,
                ),
              ],
            ),
            Container(
              decoration: BoxDecoration(
                color: Theme.of(context).colorScheme.primary,
              ),
              constraints: const BoxConstraints.expand(
                height: 70,
              ),
              child: Center(
                child: Text(
                  selectedResume,
                  style: Theme.of(context)
                      .textTheme
                      .bodyLarge!
                      .copyWith(color: Colors.white),
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}

class SeachLocationSF extends StatefulWidget {
  final Future<List<PlaceInfo>> Function(String) onSearch;
  final void Function(PlaceInfo) onChosen;

  const SeachLocationSF({
    super.key,
    required this.onChosen,
    required this.onSearch,
  });

  @override
  State<SeachLocationSF> createState() => _SeachLocationSFState();
}

class _SeachLocationSFState extends State<SeachLocationSF> {
  List<PlaceInfo> options = [];
  String data = "";

  Future onSearch(String txt) async {
    var locations = await widget.onSearch(data);
    setState(() {
      options = locations;
    });
  }

  void onChosen(PlaceInfo place) {
    widget.onChosen(place);
    print(place.name);
    setState(() {
      options = [];
    });
  }

  void onText(String txt) {
    data = txt;
  }

  @override
  Widget build(BuildContext context) {
    if (options.isEmpty) {
      return searchInputText();
    }

    return ListView.builder(
      itemCount: options.length + 1,
      shrinkWrap: true,
      itemBuilder: (context, index) {
        if (index == 0) {
          return searchInputText();
        }
        return PlaceSearchResult(
          place: options[index - 1],
          onChosen: onChosen,
        );
      },
    );
  }

  Widget searchInputText() {
    return TextField(
      onChanged: onText,
      decoration: InputDecoration(
        filled: true,
        fillColor: Theme.of(context).colorScheme.background,
        border: const OutlineInputBorder(
          borderRadius: BorderRadius.only(
            topLeft: Radius.circular(15),
            topRight: Radius.circular(15),
          ),
        ),
        suffixIcon: IconButton(
          icon: const Icon(Icons.search),
          tooltip: "Search",
          onPressed: () {
            onSearch(data);
          },
        ),
      ),
    );
  }
}

class PlaceSearchResult extends StatelessWidget {
  final PlaceInfo place;
  final void Function(PlaceInfo) onChosen;

  const PlaceSearchResult({
    super.key,
    required this.place,
    required this.onChosen,
  });

  @override
  Widget build(BuildContext context) {
    var adresss = place.name.split("-#AT#-");
    return InkWell(
      onTap: () => onChosen(place),
      child: Container(
        decoration: BoxDecoration(
          color: Theme.of(context).colorScheme.background,
          border: const Border(
            bottom: BorderSide(
              color: Colors.black,
              width: 0.5,
            ),
            left: BorderSide(
              color: Colors.black,
              width: 0.5,
            ),
            right: BorderSide(
              color: Colors.black,
              width: 0.5,
            ),
          ),
        ),
        padding: const EdgeInsets.only(left: 15, bottom: 8, top: 8, right: 15),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  adresss[0],
                  style: Theme.of(context).textTheme.bodyLarge,
                ),
                Text(
                  adresss[1],
                  style: Theme.of(context).textTheme.bodySmall,
                ),
              ],
            ),
            const Icon(Icons.location_searching),
          ],
        ),
      ),
    );
  }
}
