import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/login/view/widgets/text_input.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';
import 'package:now_v8/src/features/spots_creation/view/description.dart';
import 'package:now_v8/src/features/spots_creation/view/done_or_cancel.dart';
import 'package:now_v8/src/features/spots_creation/view/location.dart';
import 'package:now_v8/src/features/spots_creation/view/review.dart';
import 'package:now_v8/src/features/spots_creation/view/tags.dart';
import 'package:now_v8/src/features/spots_creation/view_model/providers.dart';
import 'package:now_v8/src/features/spots_creation/view_model/state_notifier.dart';

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

class Body extends ConsumerWidget {
  const Body({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    Widget pageBody;

    SpotCreatorState state = ref.watch(spotsCreatorNotiferProvider);
    SpotCreator notifer = ref.watch(spotsCreatorNotiferProvider.notifier);

    switch (state.onState) {
      case OnState.onDescription:
        pageBody = const SpotGeneralInfo();
        break;
      case OnState.onLocation:
        pageBody = const LocationSelectorView();
        break;
      case OnState.onTags:
        pageBody = const TagsSelectorView();
        break;
      case OnState.onReview:
        pageBody = const ReviewView();
        break;
      case OnState.onDone:
        pageBody = const DoneOrCancelView(
          state: "DONE",
        );
        break;
      case OnState.onCancelle:
        pageBody = const DoneOrCancelView(
          state: "CANCELLED",
        );
        break;
    }

    return Row(
      children: [
        Expanded(
          child: SafeArea(
            child: Container(
              constraints: const BoxConstraints(
                maxHeight: double.infinity,
              ),
              child: Center(
                child: PageNavigator(
                  child: pageBody, //Text("Hi dude how are you? "),
                  next: () {
                    notifer.onNext();
                  },
                  back: () {
                    notifer.onBack();
                  },
                ),
              ),
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

class PageNavigator extends StatelessWidget {
  final void Function()? back;
  final void Function()? next;
  final Widget child;

  const PageNavigator({
    super.key,
    required this.child,
    this.back,
    this.next,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        back != null
            ? NavigationIconButton(
                icon: Icons.arrow_upward,
                onTap: back!,
              )
            : const SizedBox(height: 50),
        child,
        next != null
            ? NavigationIconButton(
                icon: Icons.arrow_downward,
                onTap: next!,
              )
            : const SizedBox(height: 50),
      ],
    );
  }
}

class NavigationIconButton extends StatelessWidget {
  final IconData icon;
  final void Function() onTap;
  const NavigationIconButton({
    super.key,
    required this.icon,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      child: Container(
        constraints: const BoxConstraints(
          minHeight: 50,
          minWidth: double.infinity,
        ),
        child: Icon(icon),
      ),
      onTap: onTap,
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
