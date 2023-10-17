import 'package:flutter/material.dart';

class PhoneNumberV2 extends StatefulWidget {
  final void Function(String) onPhoneNomberEdited;
  const PhoneNumberV2({super.key, required this.onPhoneNomberEdited});

  @override
  State<PhoneNumberV2> createState() => _PhoneNumberV2State();
}

class CountryCodes {
  final String code;
  final String countryName;
  final String countryID;

  CountryCodes(
      {required this.code, required this.countryID, required this.countryName});
}

class _PhoneNumberV2State extends State<PhoneNumberV2> {
  TextEditingController countriesController = TextEditingController();

  String phoneNumber = "";
  String countryCode = "";
  String countryCodeName = "";
  bool showCountriesSearch = false;

  List<CountryCodes> countryCodes = [
    CountryCodes(code: "+57", countryID: "COL", countryName: "colombia"),
    CountryCodes(code: "+1", countryID: "TRK", countryName: "turkia"),
    CountryCodes(code: "+2", countryID: "CNL", countryName: "venezuela"),
    CountryCodes(code: "+3", countryID: "GER", countryName: "alemania"),
    CountryCodes(code: "+4", countryID: "CHI", countryName: "china"),
    CountryCodes(code: "+5", countryID: "PEU", countryName: "peru"),
    CountryCodes(code: "+6", countryID: "BOL", countryName: "bolivia"),
    CountryCodes(code: "+7", countryID: "CHL", countryName: "chile"),
    CountryCodes(code: "+8", countryID: "ARG", countryName: "argentina"),
    CountryCodes(code: "+9", countryID: "ESP", countryName: "espana"),
  ];

  late CountryCodes countrySelected = CountryCodes(
    countryID: "",
    code: "",
    countryName: "",
  );

  void onCountryPressed() {
    setState(() {
      showCountriesSearch = !showCountriesSearch;
      countryCodeName = "";
      countriesController.text = "";
    });
  }

  void onCountryNameChanged(value) {
    setState(() {
      countryCodeName = value;
    });
  }

  void onPhoneNumberChanged(String value) {
    phoneNumber = value;

    widget.onPhoneNomberEdited("${countrySelected.code}${value}");

    if (showCountriesSearch) {
      setState(() {
        showCountriesSearch = false;
      });
    }
  }

  void onCountryNameSelected(CountryCodes country) {
    setState(() {
      countrySelected = country;
      countryCodeName = "";
      countriesController.text = "";
      showCountriesSearch = false;
    });
  }

  List<CountryCodes> getCountriesToDisplay(List<CountryCodes> countryCodes) {
    List<CountryCodes> countryCodesToShow = [];

    if (countryCodeName.isNotEmpty && showCountriesSearch) {
      for (var country in countryCodes) {
        if (country.countryName
                .toLowerCase()
                .contains(countryCodeName.toLowerCase()) ||
            country.countryID
                .toLowerCase()
                .contains(countryCodeName.toLowerCase()) ||
            country.code
                .toLowerCase()
                .contains(countryCodeName.toLowerCase())) {
          countryCodesToShow.add(country);
        }
      }
    } else {
      countryCodesToShow = countryCodes;
    }

    return countryCodesToShow;
  }

  @override
  Widget build(BuildContext context) {
    if (countrySelected.code.isEmpty) {
      countrySelected = countryCodes.first;
    }
    List<CountryCodes> countryCodesToShow = getCountriesToDisplay(countryCodes);

    return Column(
      children: [
        TextFormField(
          decoration: InputDecoration(
            hintText: "Phone number",
            prefixIconConstraints: const BoxConstraints.tightFor(height: 60),
            prefixIcon: Container(
              margin: const EdgeInsets.only(
                right: 10,
              ),
              padding: const EdgeInsets.only(
                left: 5,
              ),
              decoration: BoxDecoration(
                  color: Theme.of(context)
                      .primaryColor
                      .withAlpha(30), // Colors.grey.shade200,
                  borderRadius: const BorderRadius.only(
                    topLeft: Radius.circular(15),
                    bottomLeft: Radius.circular(15),
                  )),
              child: IconButton(
                  onPressed: onCountryPressed,
                  icon: Row(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      Text(
                        "( ${countrySelected.code} )",
                        style: Theme.of(context).textTheme.bodyLarge,
                      ),
                      const SizedBox(
                        width: 5,
                      ),
                      RotatedBox(
                        quarterTurns: showCountriesSearch ? 1 : 3,
                        child: const Icon(
                          Icons.arrow_back_ios_sharp,
                          size: 20,
                        ),
                      ),
                    ],
                  )),
            ),
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(15.0),
            ),
          ),
          keyboardType: TextInputType.number,
          autocorrect: false,
          onTapOutside: (event) {
            WidgetsBinding.instance.focusManager.primaryFocus?.unfocus();
          },
          onChanged: onPhoneNumberChanged,
        ),

        //// This is the search
        Visibility(
          maintainAnimation: true,
          maintainState: true,
          visible: showCountriesSearch,
          child: AnimatedOpacity(
            duration: const Duration(milliseconds: 500),
            curve: Curves.easeInOut,
            opacity: showCountriesSearch ? 1 : 0,
            child: Container(
              margin: const EdgeInsets.only(top: 10),
              padding: const EdgeInsets.symmetric(horizontal: 10),
              decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(10),
                  border: Border.all(
                    color: Colors.black.withAlpha(150),
                  )
                  // color: Colors.blue,
                  ),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  ////////////
                  //
                  // Country input name
                  //
                  ////////////
                  TextFormField(
                    controller: countriesController,
                    decoration: InputDecoration(
                      hintText: "Search for countries",
                      prefixIcon: Container(
                        margin: const EdgeInsets.only(
                          right: 10,
                          left: 10,
                        ),
                        child: RotatedBox(
                          quarterTurns: 4,
                          child: Icon(
                            Icons.search,
                            // size: 30,
                            color: Theme.of(context).primaryColor,
                          ),
                        ),
                      ),
                    ),
                    keyboardType: TextInputType.text,
                    autocorrect: false,
                    onTapOutside: (event) {
                      WidgetsBinding.instance.focusManager.primaryFocus
                          ?.unfocus();
                    },
                    onChanged: onCountryNameChanged,
                  ),

                  ////////////
                  //
                  // list  options
                  //
                  ////////////
                  ConstrainedBox(
                    constraints: const BoxConstraints(
                      minHeight: 50,
                      maxHeight: 250,
                    ),
                    child: ListView.builder(
                      shrinkWrap: true,
                      scrollDirection: Axis.vertical,
                      itemCount: countryCodesToShow.length, // 5,
                      itemBuilder: (buildContext, index) {
                        return CountryCodeInfo(
                          countryCode: countryCodesToShow[index],
                          countryCodeSelected: countrySelected,
                          onPressed: () {
                            onCountryNameSelected(countryCodesToShow[index]);
                          },
                        );
                      },
                    ),
                  ),
                ],
              ),
            ),
          ),
        )
      ],
    );
  }
}

class CountryCodeInfo extends StatelessWidget {
  final CountryCodes countryCode;
  final void Function() onPressed;
  final CountryCodes countryCodeSelected;

  const CountryCodeInfo({
    super.key,
    required this.countryCode,
    required this.onPressed,
    required this.countryCodeSelected,
  });
  String generateCountryDisplay(CountryCodes countryCode) {
    return "${countryCode.countryID}  ${countryCode.countryName.replaceRange(0, 1, countryCode.countryName[0].toUpperCase())} ( ${countryCode.code} )";
  }

  @override
  Widget build(BuildContext context) {
    return TextButton(
      onPressed: onPressed,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(generateCountryDisplay(countryCode)),
          Visibility(
            visible: countryCodeSelected.code == countryCode.code,
            child: const Icon(Icons.check),
          )
        ],
      ),
    );
  }
}
