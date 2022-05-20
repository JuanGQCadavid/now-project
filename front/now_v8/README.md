# now_v8

## How to manage the state ?

The idea is to have a combination of StateNotifierProvider with StateNotifier, where we are going to simplify our lifes just managin one state per notifier, thus we are going to be more granular and avoiding creating big view-models monsters!

![States creation](https://drive.google.com/uc?export=view&id=10kBk4MN9ye6ztBe9eUyaCPNvPVj2rKtW)


At the end of the day we are going to have different states connected together.

![States joining together](https://drive.google.com/uc?export=view&id=1h6hTIE6xJ-LIx_fvrN6GKLt-IvjpWpwi)

## Visual studio code setup

Extensions:
* Flutter + dart
* Riverpod -> https://marketplace.visualstudio.com/items?itemName=robert-brunhage.flutter-riverpod-snippets

## M1 setup

If you find problems while running it due to cocopod you should run this command:

```
arch -x86_64 sudo gem install ffi
flutter build ios --no-codesign
```

Taken from https://medium.com/p-society/cocoapods-on-apple-silicon-m1-computers-86e05aa10d3e

