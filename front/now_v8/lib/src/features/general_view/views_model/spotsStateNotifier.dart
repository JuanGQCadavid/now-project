
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/spot.dart';

/**
 * We are going to have singles state notifiers
 * where we could only set the state, there is not need 
 * to call listeners, this will do the job for us.
 * 
 * The state is inmutable so we need to create new states
 * and assiigned it to the state.
 */

class SpotsNotifer extends StateNotifier<List<Spot>> {
  final int spotService;
  SpotsNotifer({required this.spotService}):super([]);

}