import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/spots_creation/model/core.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';
import 'package:now_v8/src/features/spots_creation/view_model/state_notifier.dart';
import 'package:now_v8/src/services/core/providers.dart';

final coreProvider = Provider<SpotsCreatorCore>((ref) {
  var provider = ref.read(gpcServicesProvider);
  return SpotsCreatorCore(
    gpcService: provider,
  );
});

final spotsCreatorNotiferProvider =
    StateNotifierProvider.autoDispose<SpotCreator, SpotCreatorState>((ref) {
  var core = ref.read(coreProvider);
  return SpotCreator(
    core: core,
  );
});
