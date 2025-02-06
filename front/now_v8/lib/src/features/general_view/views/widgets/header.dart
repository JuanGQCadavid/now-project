import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/profile.dart';
import 'package:now_v8/src/core/widgets/buttons.dart';
import 'package:now_v8/src/services/core/providers.dart';

class GeneralViewHeader extends ConsumerWidget {
  final void Function() onRequestToLogin;
  final void Function() onRequestToGoToProfile;
  final void Function() onRequestToGoToMenu;

  final Color nowColor = Colors.cyan.shade700;
  final Color commingColor = Colors.pink.shade600;

  final String nowTitle = "Events now";
  final String commingTitle = "Events comming";

  final String welcomeMessage = "Welcome back,";
  final String headerMessage = "Welcome to Pululapp";

  late List<Color> loggingColors = [
    nowColor,
    commingColor,
  ];

  final List<Color> noLoggingColors = [
    Colors.grey.shade400,
    Colors.black,
  ];

  GeneralViewHeader({
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
        nowColor: nowColor,
        nowTitle: nowTitle,
        commingColor: commingColor,
        commingTitle: commingTitle,
        userHeader: UserLogged(
          greetingMessage: welcomeMessage,
          userProfile: l,
          onMenuTap: onRequestToGoToMenu,
          onUserTap: onRequestToGoToProfile,
          colors: loggingColors,
        ),
      ),
      (r) => DefaultHeader(
        nowColor: nowColor,
        nowTitle: nowTitle,
        commingColor: commingColor,
        commingTitle: commingTitle,
        userHeader: NotLoggedHeader(
          headerMessage: headerMessage,
          onMenuTap: onRequestToGoToMenu,
          onUserTap: onRequestToLogin,
          colors: noLoggingColors,
        ),
      ),
    );
  }
}

class DefaultHeader extends StatelessWidget {
  final Widget userHeader;
  final Color nowColor;
  final String nowTitle;

  final Color commingColor;
  final String commingTitle;
  const DefaultHeader(
      {super.key,
      required this.userHeader,
      required this.nowColor,
      required this.nowTitle,
      required this.commingColor,
      required this.commingTitle});

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
        MapDescriptor(
          commingColor: commingColor,
          commingTitle: commingTitle,
          nowColor: nowColor,
          nowTitle: nowTitle,
        ),
      ],
    );
  }
}

class UserLogged extends StatelessWidget {
  final void Function() onUserTap;
  final void Function() onMenuTap;
  final List<Color> colors;
  final UserProfile userProfile;
  final String greetingMessage;

  const UserLogged({
    super.key,
    required this.userProfile,
    required this.onMenuTap,
    required this.onUserTap,
    required this.colors,
    required this.greetingMessage,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 10.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                greetingMessage,
                style: Theme.of(context).textTheme.titleLarge,
              ),
              Text(
                userProfile.userName,
                style: Theme.of(context).textTheme.bodySmall,
              )
            ],
          ),
          UserLoggedButton(
            onTap: onUserTap,
            displayName: userProfile.userName,
            colors: colors,
          )
        ],
      ),
    );
  }
}

class NotLoggedHeader extends StatelessWidget {
  final void Function() onUserTap;
  final void Function() onMenuTap;
  final String headerMessage;
  final List<Color> colors;
  const NotLoggedHeader({
    super.key,
    required this.onMenuTap,
    required this.onUserTap,
    required this.colors,
    required this.headerMessage,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 10.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(
            headerMessage,
            style: Theme.of(context).textTheme.titleLarge,
          ),
          UserLoggedButton(
            onTap: onUserTap,
            displayName: "",
            colors: colors,
          ),
        ],
      ),
    );
  }
}

class MapDescriptor extends StatelessWidget {
  final Color nowColor;
  final String nowTitle;

  final Color commingColor;
  final String commingTitle;
  const MapDescriptor(
      {super.key,
      required this.nowColor,
      required this.nowTitle,
      required this.commingColor,
      required this.commingTitle});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(10.0),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          DescriptorLocationType(nowColor, nowTitle),
          DescriptorLocationType(commingColor, commingTitle)
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
