

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v2/mapService/application/map_notifier.dart';

final mapStateProvider = StateNotifierProvider<MapStateNotifier, MapState>((ref){
  return MapStateNotifier();
});