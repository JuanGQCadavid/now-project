import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/general_view/views/generalView.dart';
import 'package:now_v8/src/features/granular_view/views/greanular_view.dart';

void main() => runApp(
      ProviderScope(
        child: MyApp(),
      ),
    );

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
        title: 'Flutter Demo',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: GeneralViewFeature() //MyHomePage(title: 'Flutter Demo Home Page'),
        );
  }
}