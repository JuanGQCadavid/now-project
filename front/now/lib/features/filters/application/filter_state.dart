import 'package:freezed_annotation/freezed_annotation.dart';

part 'filter_state.freezed.dart';

@freezed
class FilterState with _$FilterState {
  const FilterState._();
  const factory FilterState.initial() = Initial;
  const factory FilterState.loading() = Loading;
  const factory FilterState.ready() = Ready;
  const factory FilterState.onError() = Error;
}
