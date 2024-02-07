import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';
import 'package:now_v8/src/core/models/long_spot.dart';

part 'spot_creator_state.freezed.dart';

enum OnState {
  onDescription,
  onLocation,
  onTags,
  onReview,
  onDone,
  onCancelle,
}

@freezed
class SpotCreatorState with _$SpotCreatorState {
  const factory SpotCreatorState({
    required OnState onState,
    required int totalSteps,
    required int actualStep,
    required LongSpot spot,
    required String onError,
  }) = _SpotCreatorState;
}
