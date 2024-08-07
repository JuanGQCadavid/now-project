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
    this.filterServiceResource = "/filter",
    this.apiGatewayEndpoint = "https://prod.pululapp.com",
    this.userServiceResource = "/user",
    this.spotServiceResource = "/spots/core",
    this.stage = "prod",
  });

  ApiConfig.toStaging({
    this.filterServiceResource = "/filter",
    this.apiGatewayEndpoint = "https://staging.pululapp.com",
    this.userServiceResource = "/user",
    this.spotServiceResource = "/spots/core",
    this.stage = "prod",
  });

  ApiConfig.toLocal({
    this.filterServiceResource = ":8000/filter",
    this.apiGatewayEndpoint = "http://10.97.121.66",
    this.userServiceResource = ":8001/user",
    this.spotServiceResource = ":8000/spot",
    this.stage = "prod",
  });

  String getFilterEndpoint() => "${apiGatewayEndpoint}${filterServiceResource}";

  String getSpotCoreEndpoint() => "${apiGatewayEndpoint}${spotServiceResource}";

  String getUserEndpoint() => "${apiGatewayEndpoint}${userServiceResource}";

  String getBaseURL() => "${apiGatewayEndpoint}/${stage}";
}
