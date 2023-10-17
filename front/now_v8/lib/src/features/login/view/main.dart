import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/login/model/login_state.dart';
import 'package:now_v8/src/features/login/view/widgets/text_input.dart';
import 'package:now_v8/src/features/login/view/widgets/code_input.dart';
import 'package:now_v8/src/features/login/view/widgets/phone_input.dart';
import 'package:now_v8/src/features/login/view_model/providers.dart';
import 'package:now_v8/src/features/login/view_model/state_notifier.dart';

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

  String codeVerification = "";
  void onCodeVerificationChanged(String value) {
    codeVerification = value;
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var state = ref.watch(loginStateNotifierProvider);
    var notifier = ref.read(loginStateNotifierProvider.notifier);

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('Flutter Demo Home Page'),
      ),
      body: _Body(
        onPhoneChange: onPhoneNumberChanged,
        onCodeVerificationChanged: onCodeVerificationChanged,
        onUserNameChanged: onUserNameChanged,
        state: state,
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          notifier.onNext(phoneNumber, userName, codeVerification);
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
  final void Function(String) onCodeVerificationChanged;
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
              visible: state.onState == OnState.onLogin ||
                  state.onState == OnState.onInit ||
                  state.onState == OnState.onSingUp,
            ),
            Visibility(
              child: Container(
                margin: const EdgeInsets.only(bottom: 30),
                child: TextInput(
                  hint: "What is your name?",
                  onTextChanged: onUserNameChanged,
                ),
              ),
              visible: state.onState == OnState.onSingUp,
            ),
            Visibility(
              child: Container(
                margin: const EdgeInsets.only(bottom: 30),
                child: CodeInputV2(
                  size: 5,
                ),
              ),
              visible: state.onState == OnState.onLogin ||
                  state.onState == OnState.onSingUpPhoneValidation,
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
