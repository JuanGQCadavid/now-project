import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'core/presentation/main_view.dart';

void main() {
  runApp(
    const ProviderScope(
      child: MainView(),
    ),
  );
}
