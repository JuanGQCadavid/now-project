import 'package:dartz/dartz.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/features/general_view/views_model/providers.dart';

class GeneralViewHeader extends ConsumerWidget {
  final void Function() onRequestToLogin;
  final void Function() onRequestToGoToProfile;
  final void Function() onRequestToGoToMenu;

  const GeneralViewHeader({
    Key? key,
    required this.onRequestToGoToMenu,
    required this.onRequestToGoToProfile,
    required this.onRequestToLogin,
  }) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var userDetails = ref.read(generalViewModelProvider).getUserInfo();

    return FutureBuilder<Either<UserDetails, None>>(
      future: userDetails,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return snapshot.data!.fold(
            (l) => DefaultHeader(
                userHeader: UserLogged(
              userDetails: l,
              onMenuTap: onRequestToGoToMenu,
              onUserTap: onRequestToGoToProfile,
            )),
            (r) => DefaultHeader(
              userHeader: NotLoggedHeader(
                onMenuTap: onRequestToGoToMenu,
                onUserTap: onRequestToLogin,
              ),
            ),
          );
        } else {
          return DefaultHeader(
            userHeader: NotLoggedHeader(
              onMenuTap: onRequestToGoToMenu,
              onUserTap: onRequestToLogin,
            ),
          );
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
  final void Function() onUserTap;
  final void Function() onMenuTap;
  final UserDetails userDetails;

  const UserLogged({
    super.key,
    required this.userDetails,
    required this.onMenuTap,
    required this.onUserTap,
  });

  @override
  Widget build(BuildContext context) {
    String greetingMessage = "Welcome back, \n ${userDetails.userName}!";

    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        IconButton(onPressed: onMenuTap, icon: const Icon(Icons.menu)),
        Text(greetingMessage),
        UserLoggedButton(onTap: onUserTap, displayName: userDetails.userName)
      ],
    );
  }
}

class NotLoggedHeader extends StatelessWidget {
  final void Function() onUserTap;
  final void Function() onMenuTap;
  const NotLoggedHeader({
    super.key,
    required this.onMenuTap,
    required this.onUserTap,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        IconButton(onPressed: onMenuTap, icon: const Icon(Icons.menu)),
        const Text("Welcome to Pululapp"),
        IconButton(onPressed: onUserTap, icon: const Icon(Icons.person))
      ],
    );
  }
}

class MapDescriptor extends StatelessWidget {
  const MapDescriptor({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          DescriptorLocationType(Colors.cyan.shade700, "Events now"),
          DescriptorLocationType(Colors.pink.shade600, "Events comming")
        ],
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
    return Row(
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
    );
  }
}
