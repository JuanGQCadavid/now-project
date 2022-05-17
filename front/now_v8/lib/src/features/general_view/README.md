# The General view feature folder

Each feature is going to be created as if they were microservices, thus
we are going to fallow a Hexagonal architecture.

* model -> This will contain the core/hexagon/bussiness logic of the features, this one will do all the logic and will be connected to the services, in order to only have one of this class instanciated we are going to use a proivider in the view model in oder to have a single instance used across the feature ( this could also have streams, so we are going to use differente kinds of providers!)

* views -> Here we are going to have all about User interface / flutter widgets for the feature.

* Views_model -> here we are going to have only the state manager that will connect the view with the model.