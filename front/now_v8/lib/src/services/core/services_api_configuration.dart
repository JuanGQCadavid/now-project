class ApiConfig {
  final String filterServiceResource;
  final String spotServiceResource;
  final String userServiceResource;
  final String apiGatewayEndpoint;
  final String stage;

  ApiConfig(
      {required this.filterServiceResource,
      required this.spotServiceResource,
      required this.userServiceResource,
      required this.apiGatewayEndpoint,
      required this.stage});

  ApiConfig.toProd({
    this.filterServiceResource = "filter",
    this.apiGatewayEndpoint = "http://192.168.20.39:8000",
    this.userServiceResource = "user",
    this.spotServiceResource = "spot",
    this.stage = "prod",
  });

  String getFilterEndpoint() =>
      "${apiGatewayEndpoint}/${filterServiceResource}";

  String getUserEndpoint() => "${apiGatewayEndpoint}/${userServiceResource}";

  String getBaseURL() => "${apiGatewayEndpoint}/${stage}";
}
