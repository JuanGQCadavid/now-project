import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter/services.dart' show rootBundle;

import '../../../core/domain/models/spot.dart'; // WHAT !?

Future<Locations> getGoogleOfficies() async {
  const googleLocationsURL = 'https://about.google/static/data/locations.json';

  try {
    final response = await http.get(Uri.parse(googleLocationsURL));

    if (response.statusCode == 200) {
      return Locations.fromJson(json.decode(response.body));
    }
  } catch (e) {
    print(e);
  }

  return Locations.fromJson(
    json.decode(
      await rootBundle.loadString('assets/locations.json'),
    ),
  );
}
