import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/widgets/text_input.dart';
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
    return Scaffold(
      // resizeToAvoidBottomInset: false,
      body: SafeArea(
        child: Body(),
      ),
    );
  }
}

class Body extends ConsumerWidget {
  late LongSpot spot = const LongSpot(
    dateInfo: DateInfo(
      dateTime: "",
      id: "",
      startTime: "",
      durationApproximatedInSeconds: 0,
    ),
    eventInfo: EventInfo(
      name: "",
      id: "",
      description: "",
      maximunCapacty: 0,
      emoji: ":p",
    ),
    hostInfo: HostInfo(
      name: "",
    ),
    placeInfo: PlaceInfo(
      name: "",
      lat: 0.0,
      lon: 0.0,
      mapProviderId: "",
    ),
    topicInfo: TopicsInfo(
      principalTag: "",
      secondaryTags: [],
    ),
  );

  Body({super.key});

  void onTitleChange(String txt) {
    spot = spot.copyWith(eventInfo: spot.eventInfo.copyWith(name: txt));
  }

  void onDescriptionChange(String txt) {
    spot = spot.copyWith(eventInfo: spot.eventInfo.copyWith(description: txt));
  }

  void onTagsChange(List<String> tags) {
    spot =
        spot.copyWith(topicInfo: spot.topicInfo.copyWith(secondaryTags: tags));
  }

  void onLocationSeleted(PlaceInfo place) {
    spot = spot.copyWith(placeInfo: place);
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    Widget pageBody;
    SpotCreatorState state = ref.watch(spotsCreatorNotiferProvider);
    SpotCreator notifer = ref.watch(spotsCreatorNotiferProvider.notifier);

    switch (state.onState) {
      case OnState.onDescription:
        TextEditingController? titleController;
        TextEditingController? descriptionController;
        String? err;

        if (state.spot.eventInfo.name.isNotEmpty) {
          titleController = TextEditingController(
            text: state.spot.eventInfo.name,
          );
        }

        if (state.spot.eventInfo.description.isNotEmpty) {
          descriptionController = TextEditingController(
            text: state.spot.eventInfo.description,
          );
        }

        if (state.onError.isNotEmpty) {
          err = state.onError;
        }

        pageBody = SpotGeneralInfo(
          onDescriptionChange: onDescriptionChange,
          onTitleChanged: onTitleChange,
          titleController: titleController,
          descriptionController: descriptionController,
          errMessage: err,
        );
        break;
      case OnState.onLocation:
        pageBody = LocationSeletorViewV2(
          onChosen: onLocationSeleted,
        );
        break;
      case OnState.onTags:
        pageBody = TagsSelectorView(
          tagsSelected: onTagsChange,
        );
        break;
      case OnState.onReview:
        pageBody = ReviewView(
          spot: state.spot,
        );
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

    return Container(
      constraints: const BoxConstraints(
        maxHeight: double.infinity,
      ),
      child: Center(
        child: PageNavigator(
          child: pageBody,
          next: () {
            notifer.onNext(spot);
          },
          back: () {
            notifer.onBack();
          },
          pageNumber: state.actualStep,
          pageTotal: state.totalSteps,
        ),
      ),
    );
  }
}

class PageNavigator extends StatelessWidget {
  final void Function()? back;
  final void Function()? next;
  final Widget child;
  final IconData upIcon;
  final IconData downIcon;
  final int pageNumber;
  final int pageTotal;

  const PageNavigator({
    super.key,
    required this.child,
    this.back,
    this.next,
    this.downIcon = Icons.arrow_downward,
    this.upIcon = Icons.arrow_upward,
    required this.pageTotal,
    required this.pageNumber,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        back != null
            ? NavigationIconButton(
                icon: upIcon,
                onTap: back!,
              )
            : const SizedBox(height: 50),
        StatusBar(
          totalStatus: pageTotal,
          actualStatus: pageNumber,
        ),
        const SizedBox(
          height: 15,
        ),
        child,
        const SizedBox(
          height: 15,
        ),
        next != null
            ? NavigationIconButton(
                icon: (pageTotal - 1) == pageNumber ? Icons.check : downIcon,
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
      childrens.add(
        Status(
          actualStep: i == actualStatus,
        ),
      );
    }

    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: childrens,
    );
  }
}

class Status extends StatelessWidget {
  final bool actualStep;
  const Status({
    super.key,
    required this.actualStep,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
        height: 10,
        width: 10,
        margin: const EdgeInsets.symmetric(horizontal: 10),
        decoration: BoxDecoration(
          color: actualStep
              ? Theme.of(context).primaryColor
              : Theme.of(context).primaryColor.withAlpha(100), // Colors.grey,
          borderRadius: const BorderRadius.all(
            Radius.circular(50),
          ),
        ));
  }
}
