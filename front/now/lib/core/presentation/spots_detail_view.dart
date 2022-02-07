import 'package:flutter/material.dart';

class SpotDetaillView extends StatelessWidget {
  const SpotDetaillView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const _PageScafold();
  }
}

class _PageScafold extends StatelessWidget {
  const _PageScafold({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: const _Body(),
      appBar: AppBar(
        /*
        actions: [
          IconButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              icon: const Icon(Icons.arrow_back))
        ],
        */
        title: const Center(
          child: Text("Detail view"),
        ),
      ),
      floatingActionButton: FloatingActionButton(onPressed: () {}),
    );
  }
}

class _Body extends StatelessWidget {
  const _Body({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        color: Colors.blueGrey,
      ),
    );
  }
}
