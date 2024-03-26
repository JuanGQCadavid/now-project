class Token {
  final String header;
  final String value;

  Token({
    required this.header,
    required this.value,
  });

  Map<String, dynamic> toJson() => {header: value};
}
