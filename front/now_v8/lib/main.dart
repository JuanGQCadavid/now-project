import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:now_v8/src/features/general_view/views/generalView.dart';

void main() {
  Hive.initFlutter();
  runApp(
    ProviderScope(
      child: MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Flutter Demo',
        theme: ThemeData(
          primarySwatch: Colors.grey,
        ),
        home:
            Container(
              color: Colors.grey.shade100,//Theme.of(context).primaryColorLight,
              child: GeneralViewFeature(),
            )
        );
  }
}
