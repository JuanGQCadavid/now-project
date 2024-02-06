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
            child: StatusBar(
              actualStatus: state.actualStep,
              totalStatus: state.totalSteps,
            ),
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
  final int totalStatus;
  final int actualStatus;
  const StatusBar({
    super.key,
    required this.actualStatus,
    required this.totalStatus,
  });

  @override
  Widget build(BuildContext context) {
    List<Widget> childrens = [];

    for (var i = 0; i < totalStatus; i++) {
      if (i == actualStatus) {
        childrens.add(
          Status(
            actualStep: true,
            stepNumber: i,
          ),
        );
      } else {
        childrens.add(
          Status(
            actualStep: false,
            stepNumber: i,
          ),
        );
      }
    }

    return Column(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: childrens,
    );
  }
}

class Status extends StatelessWidget {
  final bool actualStep;
  final int stepNumber;
  const Status({
    super.key,
    required this.actualStep,
    required this.stepNumber,
  });

  @override
  Widget build(BuildContext context) {
    Color color = Colors.blue;

    if (actualStep) {
      color = Colors.redAccent;
    }
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
          decoration: BoxDecoration(
            color: color,
            borderRadius: const BorderRadius.all(
              Radius.circular(50),
            ),
          ),
          child: Center(
            child: Text((stepNumber + 1).toString()),
          ),
        ),
      ),
    );
  }
}
