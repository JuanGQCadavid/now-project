import 'package:flutter/material.dart';
import 'package:now/features/filters/presentation/centralMap.dart';
import 'spots_granular_view.dart';

class SpotGeneralView extends StatelessWidget {
  const SpotGeneralView({Key? key}) : super(key: key);

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
      floatingActionButton: FloatingActionButton(
        elevation: 6,
        backgroundColor: Colors.white,
        child: const Icon(
          Icons.arrow_forward_ios,
          color: Colors.black,
        ),
        onPressed: () {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) => const SpotGranularView()));
        },
      ),
    );
  }
}

class _Body extends StatelessWidget {
  const _Body({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Stack(
        children: [
          Container(
            child: const MapBody(),
            decoration: const BoxDecoration(
              color: Colors.blueGrey,
            ),
          ),
          Material(
            borderRadius: const BorderRadius.only(
                bottomLeft: Radius.circular(25),
                bottomRight: Radius.circular(25)),
            elevation: 6,
            child: Container(
              margin: const EdgeInsets.fromLTRB(0, 0, 0, 10),
              padding: const EdgeInsets.fromLTRB(10, 10, 10, 10),
              child: Row(
                children: [
                  const IconButton(
                    onPressed: null,
                    icon: Icon(Icons.menu),
                  ),
                  Flexible(
                    child: Container(
                      margin: const EdgeInsets.symmetric(horizontal: 10),
                      child: const TextField(
                        decoration: InputDecoration(
                            enabled: false,
                            suffixIcon: Icon(Icons.search),
                            border: UnderlineInputBorder(),
                            hintText: "Que hacemos?"),
                      ),
                    ),
                  ),
                  const IconButton(
                    onPressed: null,
                    icon: Icon(Icons.filter_list),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class _Header extends StatelessWidget {
  const _Header({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container();
  }
}
