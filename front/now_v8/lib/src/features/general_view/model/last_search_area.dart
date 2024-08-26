import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';

part 'last_search_area.freezed.dart';

@freezed
class LastSearchArea with _$LastSearchArea {
  const factory LastSearchArea({
    required MapState mapState,
    required bool jump,
  }) = _MapState;
}
