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
    this.apiGatewayEndpoint = "http://172.20.10.10:8000",
    this.spotServiceResource = "spot",
    this.stage = "prod"
  });

  String getFilterEndpoint() => "${apiGatewayEndpoint}/${filterServiceResource}";

  String getBaseURL() => "${apiGatewayEndpoint}/${stage}";

}
