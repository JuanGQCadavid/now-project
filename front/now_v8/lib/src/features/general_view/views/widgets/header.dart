import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/profile.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/services/core/providers.dart';

class GeneralViewHeader extends ConsumerWidget {
  final void Function() onRequestToLogin;
  final void Function() onRequestToGoToProfile;
  final void Function() onRequestToGoToMenu;

  const GeneralViewHeader({
    super.key,
    required this.onRequestToGoToMenu,
    required this.onRequestToGoToProfile,
    required this.onRequestToLogin,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var userProfile = ref.watch(userProfileStateProvider);

    return userProfile.fold(
      (l) => DefaultHeader(
        userHeader: UserLogged(
          userProfile: l,
          onMenuTap: onRequestToGoToMenu,
          onUserTap: onRequestToGoToProfile,
        ),
      ),
      (r) => DefaultHeader(
        userHeader: NotLoggedHeader(
          onMenuTap: onRequestToGoToMenu,
          onUserTap: onRequestToLogin,
        ),
      ),
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
  final UserProfile userProfile;

  const UserLogged({
    super.key,
    required this.userProfile,
    required this.onMenuTap,
    required this.onUserTap,
  });

  @override
  Widget build(BuildContext context) {
    String greetingMessage = "Welcome back,";

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 10.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(greetingMessage, style: Theme.of(context).textTheme.titleLarge,),
              Text(userProfile.userName, style: Theme.of(context).textTheme.bodySmall,)
            ],
          ),
          
          UserLoggedButton(onTap: onUserTap, displayName: userProfile.userName)
        ],
      ),
    );
  }
}

class NotLoggedHeader extends StatelessWidget {
  final void Function() onUserTap;
  final void Function() onMenuTap;
  final String header = "Welcome to Pululapp";
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
        Text(header, style: Theme.of(context).textTheme.titleLarge,),
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
