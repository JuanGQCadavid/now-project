import 'package:dartz/dartz.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/core/widgets/locationInfo.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';
import 'package:now_v8/src/services/core/providers.dart';

class GeneralViewHeader extends ConsumerWidget {
  const GeneralViewHeader({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var userDetails = ref.read(generalViewModelProvider).getUserInfo();

    // Widget optionsHeader = userDetails.fold(
    //   (l) => UserLogged(userDetails: l),
    //   (r) => const NotLoggedHeader(),
    // );

    return FutureBuilder<Either<UserDetails, None>>(
      future: userDetails,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return snapshot.data!.fold(
            (l) => DefaultHeader(userHeader: UserLogged(userDetails: l)),
            (r) => const DefaultHeader(userHeader: NotLoggedHeader()),
          );
        } else {
          return const DefaultHeader(userHeader: NotLoggedHeader());
        }
      },
    );
  }
}

class DefaultHeader extends StatelessWidget {
  final Widget userHeader;
  const DefaultHeader({super.key, required this.userHeader});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      mainAxisSize: MainAxisSize.max,
      children: [
        userHeader,
        const SizedBox(
          height: 10,
        ),
        const MapDescriptor(),
      ],
    );
  }
}

class UserLogged extends StatelessWidget {
  final UserDetails userDetails;

  const UserLogged({super.key, required this.userDetails});

  @override
  Widget build(BuildContext context) {
    String greetingMessage = "Welcome back, \n ${userDetails.userName}!";
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        IconButton(onPressed: () {}, icon: const Icon(Icons.menu)),
        Text(greetingMessage),
        IconButton(onPressed: () {}, icon: const Icon(Icons.person))
      ],
    );
  }
}

class NotLoggedHeader extends StatelessWidget {
  const NotLoggedHeader({super.key});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        IconButton(onPressed: () {}, icon: const Icon(Icons.menu)),
        const Text("Welcome to Pululapp"),
        IconButton(onPressed: () {}, icon: const Icon(Icons.person))
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
                color: color, borderRadius: BorderRadius.circular(50)),
          ),
          const SizedBox(
            width: 10,
          ),
          Text(type)
        ],
      ),
    );
  }
}
