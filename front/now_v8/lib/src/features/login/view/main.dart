import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/login/model/login_state.dart';
import 'package:now_v8/src/core/widgets/text_input.dart';
import 'package:now_v8/src/core/widgets/code_input.dart';
import 'package:now_v8/src/core/widgets/phone_input.dart';
import 'package:now_v8/src/features/login/view_model/providers.dart';

class LoginFeature extends ConsumerWidget {
  LoginFeature({super.key});
  void onPressed() {}

  String phoneNumber = "";
  void onPhoneNumberChanged(String value) {
    phoneNumber = value;
  }

  String userName = "";
  void onUserNameChanged(String value) {
    userName = value;
  }

  List<String> verificationCodes = ["", "", "", ""];
  void onCodeVerificationChanged(int pos, String value) {
    verificationCodes[pos] = value;
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var state = ref.watch(loginStateNotifierProvider);
    var notifier = ref.read(loginStateNotifierProvider.notifier);

    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (state.onState == OnState.onDone) {
        Navigator.of(context).pop();
      }
    });

    return Scaffold(
      appBar: AppBar(
          // backgroundColor: Theme.of(context).colorScheme.inversePrimary,
          ),
      body: _Body(
        onPhoneChange: onPhoneNumberChanged,
        onCodeVerificationChanged: onCodeVerificationChanged,
        onUserNameChanged: onUserNameChanged,
        state: state,
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          notifier.onNext(phoneNumber, userName, verificationCodes);
        },
        tooltip: "Loging",
        child: const Icon(Icons.login_rounded),
      ),
    );
  }
}

class _Body extends StatelessWidget {
  final void Function(String) onPhoneChange;
  final void Function(String) onUserNameChanged;
  final void Function(int, String) onCodeVerificationChanged;
  final LoginState state;

  const _Body({
    super.key,
    required this.onPhoneChange,
    required this.state,
    required this.onUserNameChanged,
    required this.onCodeVerificationChanged,
  });

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 30),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Visibility(
              child: Container(
                margin: const EdgeInsets.only(bottom: 30),
                child: PhoneNumberV2(
                  onPhoneNomberEdited: onPhoneChange,
                ),
              ),
              visible: state.stateConfig.showPhoneNumber,
            ),
            Visibility(
              child: Container(
                margin: const EdgeInsets.only(bottom: 30),
                child: TextInput(
                  hint: "What is your name?",
                  onTextChanged: onUserNameChanged,
                ),
              ),
              visible: state.stateConfig.showUserName,
            ),
            Visibility(
              child: Container(
                margin: const EdgeInsets.only(bottom: 30),
                child: CodeInputV2(
                  onCodeChange: onCodeVerificationChanged,
                  size: 4,
                ),
              ),
              visible: state.stateConfig.showCodeInput,
            ),
            Visibility(
              child: Text(state.errorMessage),
              visible: state.errorMessage.isNotEmpty,
            )
          ],
        ),
      ),
    );
  }
}
