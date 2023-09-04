import 'package:flutter/material.dart';
import 'package:now_v8/src/core/widgets/locationInfo.dart';

class GeneralViewHeader extends StatelessWidget {
  final greetingMessage = "Welcome back, \n Juan Gonzalo!";

  const GeneralViewHeader({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      mainAxisSize: MainAxisSize.max,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            IconButton(onPressed: () {}, icon: const Icon(Icons.menu)),
            Text(greetingMessage),
            IconButton(onPressed: () {}, icon: const Icon(Icons.person))
          ],
        ),
        const SizedBox(height: 10,),
        const MapDescriptor(),
      ],
    );
  }
}


class MapDescriptor extends StatelessWidget {
  const MapDescriptor({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: const Color(0xF7f7f7f7).withOpacity(0.5),
      child: Padding(
        padding: const EdgeInsets.all(10.0),
        child: Row(
          // mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,

          children: [
            DescriptorLocationType(Colors.cyan.shade700, "Events now"),
            // const SizedBox(
            //   //height: 10,
            //   width: 15,
            // ),
            DescriptorLocationType(Colors.pink.shade600, "Events comming")
          ],
        ),
      ),
    );
  }
}


class DescriptorLocationType extends StatelessWidget {
  final Color color;
  final String type;
  const DescriptorLocationType(this.color, this.type, {super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Row(
          children: [
            Container(
              height: 20,
              width: 20,
              decoration: BoxDecoration(
                color: color,
                borderRadius: BorderRadius.circular(50)
              ),
            ),
            const SizedBox(width: 10,),
            Text(type)
          ],
        ),
    );
  }
}
