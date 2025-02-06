import 'dart:ui';
import 'package:flutter/material.dart';

class FooterGeneralView extends StatelessWidget {
  final String filterMessage;
  final Function() onFilterPressed;

  final String createMessage;
  final Function() onCreatePressed;

  final String lookCloserMessage;
  final Function() onLookCloserPressed;

  const FooterGeneralView({
    super.key,
    required this.filterMessage,
    required this.onFilterPressed,
    required this.createMessage,
    required this.onCreatePressed,
    required this.lookCloserMessage,
    required this.onLookCloserPressed,
  });

  void _onItemTapped(int index) {
    switch (index) {
      case 0:
        onFilterPressed();
        return;
      case 1:
        onCreatePressed();
        return;
      case 2:
        onLookCloserPressed();
        return;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      child: ClipRRect(
          borderRadius: BorderRadius.circular(10),
          child: BackdropFilter(
            filter: ImageFilter.blur(sigmaX: 10, sigmaY: 10),
            child: BottomNavigationBar(
              onTap: _onItemTapped,
              backgroundColor: Colors.white.withAlpha(175),
              elevation: 0,
              unselectedItemColor: Colors.black,
              currentIndex: 1,
              items: [
                BottomNavigationBarItem(
                  icon: const Icon(Icons.tune_outlined),
                  label: filterMessage,
                ),
                BottomNavigationBarItem(
                  icon: const Icon(Icons.add_circle_outline),
                  label: createMessage,
                ),
                BottomNavigationBarItem(
                  icon: const Icon(Icons.zoom_in),
                  label: lookCloserMessage,
                ),
              ],
            ),
          )),
    );
  }
}
