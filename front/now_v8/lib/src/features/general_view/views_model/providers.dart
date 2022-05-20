import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/views_model/spotsStateNotifier.dart';
import 'package:now_v8/src/services/providers.dart';

final spotsStateProvider = StateNotifierProvider<SpotsNotifer, List<Spot>>(
  ((ref) {
    final spotService = ref.read(spotServiceProvider);

    return SpotsNotifer(
      spotService: spotService,
    );
  }),
);
