import 'dart:async';

import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/login/view/widgets/text_input.dart';

class LocationSelectorView extends StatelessWidget {
  late Completer<GoogleMapController> mapController = Completer();

  LocationSelectorView({super.key});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: ClipRRect(
        borderRadius: const BorderRadius.all(
          Radius.circular(15),
        ),
        child: Stack(
          children: [
            SizedBox(
              height: 600,
              width: double.infinity,
              child: NowMapV2(
                mapController: mapController,
              ),
            ),
            SizedBox(
              child: SeachLocationSF(),
              height: 600,
              width: double.infinity,
            )
          ],
        ),
      ),
      // child: Column(
      //   children: [Text("LocationSelectorView")],
      // ),
    );
  }
}

class SeachLocationSF extends StatefulWidget {
  const SeachLocationSF({super.key});

  @override
  State<SeachLocationSF> createState() => _SeachLocationSFState();
}

class _SeachLocationSFState extends State<SeachLocationSF> {
  List<String> options = [];

  void onSearch(String txt) {
    setState(() {
      options = [txt, txt, txt];
    });
  }

  @override
  Widget build(BuildContext context) {
    if (options.isEmpty) {
      return SearchInputText(
        onSend: onSearch,
      );
    }

    return ListView.builder(
      itemCount: options.length,
      itemBuilder: (context, index) {
        if (index == 0) {
          return SearchInputText(
            onSend: onSearch,
          );
        }
        return PlaceSearchResult();
        //return Text(options[index]);
      },
    );
  }
}

class SearchInputText extends StatelessWidget {
  String data = "";
  final void Function(String) onSend;

  SearchInputText({
    super.key,
    required this.onSend,
  });

  void onText(String txt) {
    data = txt;
  }

  @override
  Widget build(BuildContext context) {
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
            onSend(data);
          },
        ),
      ),
    );
  }
}

// {
//     "name": "places/ChIJZW4d6ZCbP44RkIlGZJpenuE",
//     "id": "ChIJZW4d6ZCbP44RkIlGZJpenuE",
//     "location": {
//         "latitude": 4.6640806,
//         "longitude": -74.0559318
//     },
//     "shortFormattedAddress": "Cra. 12a #78-40"
// },

class PlaceSearchResult extends StatelessWidget {
  const PlaceSearchResult({super.key});

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () {
        print("HIIII");
      },
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
        padding: const EdgeInsets.only(left: 15, bottom: 8, top: 8),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  "A",
                  style: Theme.of(context).textTheme.bodyLarge,
                ),
                Text(
                  "BBBBBBB",
                  style: Theme.of(context).textTheme.bodySmall,
                ),
              ],
            ),
            IconButton(onPressed: () {}, icon: Icon(Icons.location_searching)),
          ],
        ),
      ),
    );
  }
}
