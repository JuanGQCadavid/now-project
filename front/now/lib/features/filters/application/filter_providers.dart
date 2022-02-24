import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now/core/domain/ports/i_now_filter_service.dart';
import 'package:now/core/services/filterService.dart';
import 'package:now/features/filters/application/filter_notifier.dart';

final filterServiceProvider = Provider<INowFIlterService>((ref) {
  return FilterService();
});

final filterNotifierProvier =
    ChangeNotifierProvider<FilterChangeNotifier>((ref) {
  final INowFIlterService service = ref.read(filterServiceProvider);

  return FilterChangeNotifier(filterService: service);
});
