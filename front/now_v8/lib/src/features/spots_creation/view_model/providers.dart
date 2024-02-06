import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';
import 'package:now_v8/src/features/spots_creation/view_model/state_notifier.dart';

final spotsCreatorNotiferProvider =
    StateNotifierProvider.autoDispose<SpotCreator, SpotCreatorState>((ref) {
  return SpotCreator();
});
