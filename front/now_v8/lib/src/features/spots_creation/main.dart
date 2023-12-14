import 'package:flutter/material.dart';

class SpotsCreationFeature extends StatelessWidget {
  const SpotsCreationFeature({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Body(),
      ),
    );
  }
}

class Body extends StatelessWidget {
  const Body({super.key});

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Expanded(
          child: Container(
            constraints: const BoxConstraints(
              maxHeight: double.infinity,
            ),
            child: Center(
              child: Text("Hi dude how are you? "),
            ),
          ),
        ),
        Container(
          constraints: const BoxConstraints(
            maxHeight: double.infinity,
            minWidth: 50,
            maxWidth: 50,
          ),
          decoration: BoxDecoration(color: Colors.amber.shade100),
          child: Center(
            child: StatusBar(),
          ),
        ),
      ],
    );
  }
}

class StatusBar extends StatelessWidget {
  const StatusBar({super.key});

  @override
  Widget build(BuildContext context) {
    return const Column(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: [
        Status(),
        Status(),
        Status(),
      ],
    );
  }
}

class Status extends StatelessWidget {
  const Status({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 35,
      width: 35,
      decoration: const BoxDecoration(
        color: Colors.green,
        borderRadius: BorderRadius.all(
          Radius.circular(50),
        ),
      ),
      child: Center(
        child: Container(
          height: 30,
          width: 30,
          decoration: const BoxDecoration(
            color: Colors.blue,
            borderRadius: BorderRadius.all(
              Radius.circular(50),
            ),
          ),
          child: const Center(
            child: Text("1"),
          ),
        ),
      ),
    );
  }
}
