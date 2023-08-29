
// import 'dart:developer';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/model/generalViewModel.dart';

/**
 * We are going to have singles state notifiers
 * where we could only set the state, there is not need 
 * to call listeners, this will do the job for us.
 * 
 * The state is inmutable so we need to create new states
 * and assiigned it to the state.
 */

class SpotsNotifer extends StateNotifier<List<Spot>> {
  final GeneralViewModel generalViewModel;

  SpotsNotifer({required this.generalViewModel}):super([]) {
    refreshSpots();
  }

  void refreshSpots() async{
    List<Spot> spots = await generalViewModel.getSpots();
    state = spots;
  }
}

class TagsNotifier extends StateNotifier<Set<String>> {

  TagsNotifier(): super({});

  void tagSelected(String tag){
    Set<String> newState = new Set.from(state);
    if (state.contains(tag)) {
      newState.remove(tag);
    }else {
      newState.add(tag);
    } 

    state = newState;
  }

  void cleanTags() {
    state = {};
  }
}