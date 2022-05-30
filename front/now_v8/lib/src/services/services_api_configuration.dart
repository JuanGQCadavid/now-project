class ApiConfig {
  final String filterServiceResource;
  final String spotServiceResource;
  final String apiGatewayEndpoint;
  final String stage;

  ApiConfig({
    required this.filterServiceResource,
    required this.spotServiceResource,
    required this.apiGatewayEndpoint,
    required this.stage
  });

  ApiConfig.toProd({
    this.filterServiceResource = "filter",
    this.apiGatewayEndpoint = "https://4co5utcub8.execute-api.us-east-2.amazonaws.com",
    this.spotServiceResource = "spot",
    this.stage = "prod"
  });

  String getFilterEndpoint() => "${apiGatewayEndpoint}/${stage}/${filterServiceResource}";

  String getBaseURL() => "${apiGatewayEndpoint}/${stage}";

}
