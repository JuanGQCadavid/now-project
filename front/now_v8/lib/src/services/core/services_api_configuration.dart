class ApiConfig {
  // Service resources
  final String filterServiceResource;
  final String spotServiceResource;
  final String userServiceResource;
  final String profileServiceResource;

  // Endpoint
  final String apiGatewayEndpoint;

  ApiConfig({
    // Service resources
    required this.filterServiceResource,
    required this.spotServiceResource,
    required this.userServiceResource,
    required this.profileServiceResource,

    // Endpoint
    required this.apiGatewayEndpoint,
  });

  ApiConfig.toProd({
    // Service resources
    this.filterServiceResource = "/filter",
    this.userServiceResource = "/user",
    this.spotServiceResource = "/spots/core",
    this.profileServiceResource = "/profile",

    // Endpoint
    this.apiGatewayEndpoint = "https://prod.pululapp.com",
  });

  ApiConfig.toStaging({
    // Service resources
    this.filterServiceResource = "/filter",
    this.userServiceResource = "/user",
    this.spotServiceResource = "/spots/core",
    this.profileServiceResource = "/profile",

    // Endpoint
    this.apiGatewayEndpoint = "https://staging.pululapp.com",
  });

  ApiConfig.toLocal({
    // Service resources
    this.filterServiceResource = ":8000/filter",
    this.userServiceResource = ":8001/user",
    this.spotServiceResource = ":8000/spot",
    this.profileServiceResource = ":8000/profile",

    // Endpoint
    this.apiGatewayEndpoint = "http://10.97.121.66",
  });

  // Service resources
  String getFilterEndpoint() => "${apiGatewayEndpoint}${filterServiceResource}";
  String getSpotCoreEndpoint() => "${apiGatewayEndpoint}${spotServiceResource}";
  String getUserEndpoint() => "${apiGatewayEndpoint}${userServiceResource}";
  String getUserProfileEndpoint() =>
      "$apiGatewayEndpoint$profileServiceResource";

  String getBaseURL() => "${apiGatewayEndpoint}";
}
