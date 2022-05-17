# The service folder

In this folder the idea is to collect all the implementation of the core external services in order to be reusable across all other features.


Each service must be independent enough in order to be encapsulated, also they should provide a provider implementation in order to only has a instance of it across all the app life time, that intance should be created here at the service root.