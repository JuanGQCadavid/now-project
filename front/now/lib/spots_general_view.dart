import 'package:flutter/material.dart';
import 'spots_granular_view.dart';

class SpotGeneralView extends StatelessWidget {
  const SpotGeneralView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const _PageScafold(),
    );
  }
}

class _PageScafold extends StatelessWidget {
  const _PageScafold({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: const _Body(),
      appBar: AppBar(
        title: Center(
          child: Text("Ongoing events!"),
        ),
      ),
      floatingActionButton: FloatingActionButton(onPressed: () {
        Navigator.push(context,
            MaterialPageRoute(builder: (context) => const SpotGranularView()));
      }),
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
