import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/simple_state.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
import 'package:now_v8/src/features/spots_creation/model/core.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';

class LocationState extends StateNotifier<SimpleState<PlaceInfo>> {
  final ILocationService locationService;
  final SpotsCreatorCore core;
  late CameraPosition cameraPosition;

  late GoogleMapController controller;
  late void Function(PlaceInfo) onChosenCallBack = (place) {};

  LocationState({required this.locationService, required this.core})
      : super(
          const SimpleState(
            value: PlaceInfo(
              lat: 0,
              lon: 0,
              mapProviderId: "",
              name: "",
            ),
            onState: SimpleOnState.onLoading,
          ),
        ) {
    initState();
  }

  Future initState() async {
    var currentLocation = await locationService.getUserCurrentLocation();
    var response = await core.getAproximatedPlaces(
      currentLocation.latitude,
      currentLocation.longitude,
    );

    response.fold((l) {
      onChosenCallBack(l[0]);
      state = SimpleState(
        value: l[0],
        onState: SimpleOnState.onDone,
      );
    }, (r) => null);
  }

  String resume(PlaceInfo place) {
    return place.name.replaceFirst(" -#AT#- ", "\n");
  }

  void setCallback(void Function(PlaceInfo) onChosenCallBack) {
    this.onChosenCallBack = onChosenCallBack;
  }

  Future onChosen(PlaceInfo placeInfo) async {
    var currentLocation = await locationService.getUserCurrentLocation();
    var bounds = MapUtilities.getCameraLatLngBounds(
      [
        Spot.withOutSpotColors(
          principalTag: "",
          secondaryTags: [],
          latLng: LatLng(
            placeInfo.lat,
            placeInfo.lon,
          ),
          spotId: "",
          date: DateTime.now(),
        ),
      ],
      userLocation: currentLocation,
    );

    // A
    // controller.animateCamera(
    //   CameraUpdate.newLatLngBounds(
    //     bounds,
    //     50,
    //   ),
    // );

    //B

    controller.animateCamera(
      CameraUpdate.newLatLngZoom(
        LatLng(placeInfo.lat, placeInfo.lon),
        18.5,
      ),
    );

    state = state.copyWith(value: placeInfo);
    onChosenCallBack(placeInfo);
  }

  Future<List<PlaceInfo>> onSearch(String locationName) async {
    var called = await core.getOptions(locationName);

    return called.fold((l) {
      return l;
    }, (r) {
      return [];
    });
  }

  Future onMapCreated(GoogleMapController controller) async {
    this.controller = controller;
  }

  Future onCameraMoveStarted() async {
    print("Here we go!");
  }

  Future onCameraIdle() async {
    print("STOP!");
    var response = await core.getAproximatedPlaces(
      cameraPosition.target.latitude,
      cameraPosition.target.longitude,
    );

    response.fold((l) => onChosen(l[0]), (r) => null);
  }

  Future onCameraMove(CameraPosition cameraPosition) async {
    this.cameraPosition = cameraPosition;
  }
}

class SpotCreator extends StateNotifier<SpotCreatorState> {
  late Map<OnState, Function(bool, LongSpot spot)> mapStates;
  final SpotsCreatorCore core;

  SpotCreator({required this.core})
      : super(
          const SpotCreatorState(
            actualStep: 0,
            totalSteps: 4,
            onState: OnState.onLocation,
            onError: "",
            spot: LongSpot(
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
            ),
          ),
        ) {
    mapStates = {
      OnState.onDone: onDone,
      OnState.onDescription: onDescription,
      OnState.onLocation: onLocation,
      OnState.onTags: onTags,
      OnState.onReview: onReview,
      OnState.onCancelle: onCancelle,
    };
  }

  void onNext(LongSpot spot) {
    Function(bool, LongSpot) func = mapStates[super.state.onState]!;
    func(true, spot);
  }

  void onBack() {
    Function(bool, LongSpot) func = mapStates[super.state.onState]!;
    func(false, state.spot);
  }

  void onDescription(bool next, LongSpot spot) {
    print("onDescription");
    print(spot.eventInfo.description);
    print(spot.eventInfo.name);
    if (next) {
      if (spot.eventInfo.description.isEmpty || spot.eventInfo.name.isEmpty) {
        state = state.copyWith(onError: "Title and description are required");
        return;
      }

      var newEvents = state.spot.eventInfo.copyWith(
          name: spot.eventInfo.name, description: spot.eventInfo.description);

      state = state.copyWith(
        onState: OnState.onLocation,
        actualStep: 1,
        spot: spot.copyWith(eventInfo: newEvents),
        onError: "",
      );
    }
  }

  Future onLocation(bool next, LongSpot spot) async {
    if (next) {
      print("-------------");
      print(spot.placeInfo);
      state = state.copyWith(
        onState: OnState.onTags,
        spot: spot.copyWith(placeInfo: spot.placeInfo),
        actualStep: 2,
        onError: "",
      );
    } else {
      state = state.copyWith(
        onState: OnState.onDescription,
        actualStep: 0,
        onError: "",
      );
    }
  }

  void onTags(bool next, LongSpot spot) {
    if (next) {
      state = state.copyWith(
        onState: OnState.onReview,
        actualStep: 3,
        onError: "",
      );
    } else {
      state = state.copyWith(
        onState: OnState.onLocation,
        actualStep: 1,
        onError: "",
      );
    }
  }

  void onReview(bool next, LongSpot spot) {
    if (next) {
      state = state.copyWith(onState: OnState.onDone);
    } else {
      state = state.copyWith(
        onState: OnState.onTags,
        actualStep: 2,
        onError: "",
      );
    }
  }

  void onCancelle(bool next, LongSpot spot) {}
  void onDone(bool next, LongSpot spot) {}
}
