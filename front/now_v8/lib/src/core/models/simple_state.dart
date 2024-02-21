import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';

part 'simple_state.freezed.dart';

enum SimpleOnState { onLoading, onDone, onError }

@freezed
class SimpleState<T> with _$SimpleState<T> {
  const factory SimpleState({
    required T value,
    required SimpleOnState onState,
  }) = _SimpleState<T>;
}
