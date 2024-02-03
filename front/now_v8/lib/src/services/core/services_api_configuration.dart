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
    this.filterServiceResource = ":8000/filter",
    this.apiGatewayEndpoint = "http://10.97.121.66",
    this.userServiceResource = ":8001/user",
    this.spotServiceResource = ":8000/spot",
    this.stage = "prod",
  });

  String getFilterEndpoint() => "${apiGatewayEndpoint}${filterServiceResource}";

  String getUserEndpoint() => "${apiGatewayEndpoint}${userServiceResource}";

  String getBaseURL() => "${apiGatewayEndpoint}/${stage}";
}
