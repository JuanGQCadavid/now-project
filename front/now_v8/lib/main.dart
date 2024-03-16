import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:now_v8/src/features/spots_creation/main.dart';

void main() async {
  await Hive.initFlutter();

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
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: SpotsCreationFeature(),
    );
  }
}

class TagView extends StatelessWidget {
  const TagView({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: SafeArea(
        child: Body(),
      ),
    );
  }
}

class Body extends ConsumerWidget {
  const Body({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return PageNavigator(
      child:childViewTry(),  
      next: () {
        print("OnNetx");
      },
      back: () {
        print("onBac");
      },
    );
  }

  Widget columnTry(){
    return const Column(
        children: [
          Text(
            "Do you want to add some tags?",
          ),
           SizedBox(height: 25),
           TextField(
            autofocus: true,
          )
        ],
      );
  }

  Widget childViewTry(){
    return SingleChildScrollView(
      child: columnTry(),
    );
  }

}

class PageNavigator extends StatelessWidget {
  final void Function()? back;
  final void Function()? next;
  final Widget child;
  final IconData upIcon;
  final IconData downIcon;

  const PageNavigator({
    super.key,
    required this.child,
    this.back,
    this.next,
    this.downIcon = Icons.arrow_downward,
    this.upIcon = Icons.arrow_upward,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.spaceBetween, // Updated line
      children: [
        back != null
            ? NavigationIconButton(
                icon: upIcon,
                onTap: back!,
              )
            : const SizedBox(height: 50),
        const SizedBox(
          height: 15,
        ),
        child,
        // child,
        const SizedBox(
          height: 15,
        ),
        next != null
            ? NavigationIconButton(
                icon: downIcon,
                onTap: next!,
              )
            : const SizedBox(height: 50),
      ],
    );
  }
}

class NavigationIconButton extends StatelessWidget {
  final IconData icon;
  final void Function() onTap;
  const NavigationIconButton({
    super.key,
    required this.icon,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      child: Container(
        constraints: const BoxConstraints(
          minHeight: 50,
          minWidth: double.infinity,
        ),
        child: Icon(icon),
      ),
      onTap: onTap,
    );
  }
}
