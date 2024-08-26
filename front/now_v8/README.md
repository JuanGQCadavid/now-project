# now_v8

## Useful commands

This project use some tools that generate code as json_serelizable, in order to generate the classes we could run:

* flutter pub run build_runner build 
* flutter pub run build_runner watch

In order to update packages to their last one allowed we should use 

* flutter pub upgrade

## How to manage the state ?

The idea is to have a combination of StateNotifierProvider with StateNotifier, where we are going to simplify our lifes just managin one state per notifier, thus we are going to be more granular and avoiding creating big view-models monsters!

![States creation](https://drive.google.com/uc?export=view&id=10kBk4MN9ye6ztBe9eUyaCPNvPVj2rKtW)

At the end of the day we are going to have different states connected together for a single feature.

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


## To ignore files

https://stackoverflow.com/questions/1274057/how-do-i-make-git-forget-about-a-file-that-was-tracked-but-is-now-in-gitignore

Update, a better option
Since this answer was posted, a new option has been created and that should be preferred. You should use --skip-worktree which is for modified tracked files that the user don't want to commit anymore and keep --assume-unchanged for performance to prevent git to check status of big tracked files. See https://stackoverflow.com/a/13631525/717372 for more details...

git update-index --skip-worktree <file>
To cancel

git update-index --no-skip-worktree <file>

