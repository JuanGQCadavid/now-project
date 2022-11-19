import { Stack, StackProps, Tags } from "aws-cdk-lib";

import {
  aws_apigateway as apigateway,
  aws_lambda as lambda,
  aws_sns as sns,
  aws_sns_subscriptions as subscriptions,
  aws_sqs as sqs,
  aws_lambda_event_sources as lambdaEvent,
  aws_iam as iam,
  aws_dynamodb as dynamodb,
  aws_ssm as ssm,
  aws_events as events,
  aws_events_targets as eventsTargets,
} from "aws-cdk-lib";
import { RuleTargetInput } from "aws-cdk-lib/aws-events";

import { Construct } from "constructs";
import { join } from "path";
// import * as sqs from 'aws-cdk-lib/aws-sqs';

// Tags const
const SPOT_FAMILY = "Spot";
const FILTER_FAMILY = "Filter";
const SERVICE_TYPE = "Service";
const FUNCTION_TYPE = "Function";
const CONFIRMATION_FAMILY = "Confirmation";

export class InfraStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    function addMethodToApiGateway(
      lambdaToAdd: lambda.Function,
      apigatewayToHandle: apigateway.RestApi,
      methodName: string
    ): apigateway.Resource {
      const method = apigatewayToHandle.root.addResource(methodName);
      const proxyMethod = method.addResource(`{${methodName}+}`);

      const lambdaIntegration = new apigateway.LambdaIntegration(lambdaToAdd, {
        proxy: true,
      });
      method.addMethod("ANY", lambdaIntegration);

      proxyMethod.addMethod("GET", lambdaIntegration);
      proxyMethod.addMethod("POST", lambdaIntegration);
      proxyMethod.addMethod("PUT", lambdaIntegration);
      proxyMethod.addMethod("DELETE", lambdaIntegration);

      return method;
    }

    function addResourceToMethod(
      lambdaToAdd: lambda.Function,
      baseResource: apigateway.Resource,
      methodName: string
    ): apigateway.Resource {
      var methodResource = baseResource.addResource(methodName);
      var methoodProxy = methodResource.addResource(`{${methodName}+}`);

      const lambdaIntegration = new apigateway.LambdaIntegration(lambdaToAdd, {
        proxy: true,
      });
      methodResource.addMethod("ANY", lambdaIntegration);

      methoodProxy.addMethod("GET", lambdaIntegration);
      methoodProxy.addMethod("POST", lambdaIntegration);
      methoodProxy.addMethod("PUT", lambdaIntegration);
      methoodProxy.addMethod("DELETE", lambdaIntegration);

      return methodResource;
    }

    // ssm parameters

    // neo4jUser: "neo4jUser",
    // neo4jPassword: "neo4jPassword",
    // neo4jUri: "neo4jUri",
    // dbUser: "dbUser",
    // dbPassword: "dbPassword",
    // dbName: "dbName",
    // dbUrl: "dbUrl",

    const neo4jUserParameter = new ssm.StringParameter(
      this,
      "neo4jUserParameter",
      {
        parameterName: "neo4jUser",
        stringValue: "neo4jUser",
      }
    );

    const neo4jPasswordParameter = new ssm.StringParameter(
      this,
      "neo4jPasswordParameter",
      {
        parameterName: "neo4jPassword",
        stringValue: "neo4jPassword",
      }
    );

    const neo4jUriParameter = new ssm.StringParameter(
      this,
      "neo4jUriParameter",
      {
        parameterName: "neo4jUri",
        stringValue: "neo4jUri",
      }
    );

    const dbUserParameter = new ssm.StringParameter(this, "dbUserParameter", {
      parameterName: "dbUser",
      stringValue: "dbUser",
    });

    const dbPasswordParameter = new ssm.StringParameter(
      this,
      "dbPasswordParameter",
      {
        parameterName: "dbPassword",
        stringValue: "dbPassword",
      }
    );

    const dbUrlParameter = new ssm.StringParameter(this, "dbUrlParameter", {
      parameterName: "dbUrl",
      stringValue: "dbUrl",
    });

    const dbNameParameter = new ssm.StringParameter(this, "dbNameParameter", {
      parameterName: "dbName",
      stringValue: "dbName",
    });

    // MAIN API GATEWAY
    const mainApiGateway = new apigateway.RestApi(this, "mainApiGateway", {});

    mainApiGateway.root.addMethod("ANY");

    // SNS
    // TODO -> extra config ?
    const spotActivityTopic = new sns.Topic(this, "spotActivityTopic", {
      displayName: "spotActivityTopic",
      topicName: "spotActivityTopic",
    });

    // SQS

    // TODO -> extra config ?
    const updateLocationDataSQS = new sqs.Queue(this, "updateLocationDataSQS", {
      queueName: "updateLocationDataSQS",
    });

    spotActivityTopic.addSubscription(
      new subscriptions.SqsSubscription(updateLocationDataSQS)
    );

    const schedulePatternSQS = new sqs.Queue(this, "schedulePatternSQS", {
      queueName: "schedulePatternSQS",
    });
    spotActivityTopic.addSubscription(
      new subscriptions.SqsSubscription(schedulePatternSQS)
    );

    // LAMBDAS

    let path = join(__dirname, "assets", "main.zip");
    console.log(path);

    // SPOTS FAMILY

    const rootMethod = mainApiGateway.root.addResource("spots");

    // Spot Core lambda
    const spotsCoreLambdaRole = new iam.Role(this, "spotsCoreLambdaRole", {
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaBasicExecutionRole"
        ),
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaRole"
        ),
      ],
      roleName: "spotsCoreLambdaRole",
      description: "Spots Shchedule Lambda Role",
    });

    spotsCoreLambdaRole.addToPolicy(
      new iam.PolicyStatement({
        actions: ["sns:*"],
        effect: iam.Effect.ALLOW,
        resources: [spotActivityTopic.topicArn],
      })
    );

    const spotsCoreLambda = new lambda.Function(this, "spotsCoreLambda", {
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
      role: spotsCoreLambdaRole,
      code: lambda.Code.fromAsset(path),
      functionName: "SpotsCoreService",
      environment: {
        // TODO -> How can we manage this env variables ?
        neo4jUser: neo4jUserParameter.parameterName,
        neo4jPassword: neo4jPasswordParameter.parameterName,
        neo4jUri: neo4jUriParameter.parameterName,
        snsArn: spotActivityTopic.topicArn,
      },
    });
    // const spotsScheduledMethod = addMethodToApiGateway(spotsScheduledLambda, mainApiGateway, "spots/scheduled")
    const spotsCoredMethod = addResourceToMethod(
      spotsCoreLambda,
      rootMethod,
      "core"
    );
    Tags.of(spotsCoreLambda).add("Type", SERVICE_TYPE);
    Tags.of(spotsCoreLambda).add("Family", SPOT_FAMILY);

    // Spots Online

    const spotsOnlineLambdaRole = new iam.Role(this, "spotsOnlineLambdaRole", {
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaBasicExecutionRole"
        ),
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaRole"
        ),
      ],
      roleName: "spotsOnlineLambdaRole",
      description: "Spot Online Lambda Role",
    });

    const spotsOnlineLambda = new lambda.Function(this, "spotsOnlineLambda", {
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
      role: spotsOnlineLambdaRole,
      code: lambda.Code.fromAsset(path),
      functionName: "SpotsOnlineService",
      environment: {
        // TODO -> How can we manage this env variables ?
        neo4jUser: neo4jUserParameter.parameterName,
        neo4jPassword: neo4jPasswordParameter.parameterName,
        neo4jUri: neo4jUriParameter.parameterName,
        snsArn: spotActivityTopic.topicArn,
        spotsCoreServiceURL: `https://${mainApiGateway.restApiId}.execute-api.${this.region}.amazonaws.com/prod${spotsCoredMethod.path}`,
      },
    });
    Tags.of(spotsOnlineLambda).add("Type", SERVICE_TYPE);
    Tags.of(spotsOnlineLambda).add("Family", SPOT_FAMILY);
    // const spotsOnlineLambdaMethod = addMethodToApiGateway(spotsOnlineLambda, mainApiGateway, "spots/online")
    const spotsOnlineLambdaMethod = addResourceToMethod(
      spotsOnlineLambda,
      rootMethod,
      "online"
    );

    // Spot schedule lambda
    const spotsScheduledLambdaRole = new iam.Role(
      this,
      "spotsScheduledLambdaRole",
      {
        assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
        managedPolicies: [
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaBasicExecutionRole"
          ),
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaRole"
          ),
        ],
        roleName: "spotsScheduledLambdaRole",
        description: "Spots Shchedule Lambda Role",
      }
    );

    const spotsScheduledLambda = new lambda.Function(
      this,
      "spotsScheduledLambda",
      {
        runtime: lambda.Runtime.GO_1_X,
        handler: "main",
        role: spotsScheduledLambdaRole,
        code: lambda.Code.fromAsset(path),
        functionName: "SpotsScheduledService",
        environment: {
          // TODO -> How can we manage this env variables ?
          neo4jUser: neo4jUserParameter.parameterName,
          neo4jPassword: neo4jPasswordParameter.parameterName,
          neo4jUri: neo4jUriParameter.parameterName,
          snsArn: spotActivityTopic.topicArn,
          spotsCoreServiceURL: `https://${mainApiGateway.restApiId}.execute-api.${this.region}.amazonaws.com/prod${spotsCoredMethod.path}`,
        },
      }
    );
    // const spotsScheduledMethod = addMethodToApiGateway(spotsScheduledLambda, mainApiGateway, "spots/scheduled")
    const spotsScheduledMethod = addResourceToMethod(
      spotsScheduledLambda,
      rootMethod,
      "scheduled"
    );
    Tags.of(spotsScheduledLambda).add("Type", SERVICE_TYPE);
    Tags.of(spotsScheduledLambda).add("Family", SPOT_FAMILY);
    // Handler

    const spotHandlerLambdaRole = new iam.Role(this, "spotHandlerLambdaRole", {
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaBasicExecutionRole"
        ),
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaRole"
        ),
      ],
      roleName: "spotHandlerLambdaRole",
      description: "Spot handler lambda role",
    });

    const spotHandlerLambda = new lambda.Function(this, "spotHandlerLambda", {
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
      role: spotHandlerLambdaRole,
      code: lambda.Code.fromAsset(path),
      functionName: "SpotsHandlerService",
      environment: {
        // TODO -> How can we manage this env variables ?
        // mainApiGateway: mainApiGateway.url,
        spotsOnlineServiceURL: `https://${mainApiGateway.restApiId}.execute-api.${this.region}.amazonaws.com/prod${spotsOnlineLambdaMethod.path}`,
        spotsScheduledServiceURL: `https://${mainApiGateway.restApiId}.execute-api.${this.region}.amazonaws.com/prod${spotsScheduledMethod.path}`,
      },
    });
    const spotsHandlerMethod = addResourceToMethod(
      spotHandlerLambda,
      rootMethod,
      "handler"
    );
    Tags.of(spotHandlerLambda).add("Type", SERVICE_TYPE);
    Tags.of(spotHandlerLambda).add("Family", SPOT_FAMILY);

    // Confirmation Spot Service
    const confirmationSpotLambdaRole = new iam.Role(
      this,
      "confirmationSpotLambdaRole",
      {
        assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
        managedPolicies: [
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaBasicExecutionRole"
          ),
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaRole"
          ),
        ],
        roleName: "confirmationSpotLambdaRole",
        description: "Spots Shchedule Lambda Role",
      }
    );

    const confirmationSpotLambda = new lambda.Function(
      this,
      "confirmationSpotLambda",
      {
        runtime: lambda.Runtime.GO_1_X,
        handler: "main",
        role: confirmationSpotLambdaRole,
        code: lambda.Code.fromAsset(path),
        functionName: "ConfirmationSpotService",
        environment: {
          // TODO -> How can we manage this env variables ?
          neo4jUser: neo4jUserParameter.parameterName,
          neo4jPassword: neo4jPasswordParameter.parameterName,
          neo4jUri: neo4jUriParameter.parameterName,
          snsArn: spotActivityTopic.topicArn,
        },
      }
    );

    Tags.of(confirmationSpotLambda).add("Type", SERVICE_TYPE);
    Tags.of(confirmationSpotLambda).add("Family", CONFIRMATION_FAMILY);
    addMethodToApiGateway(
      confirmationSpotLambda,
      mainApiGateway,
      "confirmation"
    );

    // Filter

    const filterSessionsDynamoTable = new dynamodb.Table(
      this,
      "filterSessionsDynamoTable",
      {
        partitionKey: {
          name: "SessionId",
          type: dynamodb.AttributeType.STRING,
        },
        sortKey: {
          name: "State",
          type: dynamodb.AttributeType.STRING,
        },
        billingMode: dynamodb.BillingMode.PROVISIONED,
        writeCapacity: 20,
        readCapacity: 20,
        tableName: "FilterSessions",
        timeToLiveAttribute: "TTL",
        stream: dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
      }
    );

    const filterLambdaRole = new iam.Role(this, "filterLambdaRole", {
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      description: "Role for filter service",
    });

    // Lambda Permisions
    filterLambdaRole.addManagedPolicy(
      iam.ManagedPolicy.fromAwsManagedPolicyName(
        "service-role/AWSLambdaBasicExecutionRole"
      )
    );
    // filterLambdaRole.addManagedPolicy(iam.ManagedPolicy.fromAwsManagedPolicyName("service-role/AWSLambdaVPCAccessExecutionRole"))
    filterSessionsDynamoTable.grantReadWriteData(filterLambdaRole);

    // -> filterLambda
    const filterLambda = new lambda.Function(this, "filterLambda", {
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
      code: lambda.Code.fromAsset(path),
      functionName: "FilterService",
      role: filterLambdaRole,
      environment: {
        dbUser: dbUserParameter.parameterName,
        dbPassword: dbPasswordParameter.parameterName,
        dbName: dbNameParameter.parameterName,
        dbUrl: dbUrlParameter.parameterName,
        sessionTableName: filterSessionsDynamoTable.tableName,
        spotServiceURL: `https://${mainApiGateway.restApiId}.execute-api.${this.region}.amazonaws.com/prod${spotsHandlerMethod.path}`,
      },
    });

    addMethodToApiGateway(filterLambda, mainApiGateway, "filter");
    Tags.of(filterLambda).add("Type", SERVICE_TYPE);
    Tags.of(filterLambda).add("Family", FILTER_FAMILY);

    // -> locationDataUpdater
    const locationDataUpdater = new lambda.Function(
      this,
      "locationDataUpdater",
      {
        runtime: lambda.Runtime.GO_1_X,
        handler: "main",
        code: lambda.Code.fromAsset(path),
        functionName: "LocationDataUpdater",
        environment: {
          dbUser: dbUserParameter.parameterName,
          dbPassword: dbPasswordParameter.parameterName,
          dbName: dbNameParameter.parameterName,
          dbUrl: dbUrlParameter.parameterName,
        },
      }
    );
    Tags.of(locationDataUpdater).add("Type", FUNCTION_TYPE);
    Tags.of(locationDataUpdater).add("Family", FILTER_FAMILY);

    const locationDataUpdaterEvent = new lambdaEvent.SqsEventSource(
      updateLocationDataSQS,
      {
        enabled: true,
        batchSize: 10,
      }
    );
    locationDataUpdater.addEventSource(locationDataUpdaterEvent);

    const userTrackFilterSearchSession = new lambda.Function(
      this,
      "userTrackFilterSearchSession",
      {
        runtime: lambda.Runtime.GO_1_X,
        handler: "main",
        code: lambda.Code.fromAsset(path),
        functionName: "UserTrackFilterSearchSession",
      }
    );

    Tags.of(userTrackFilterSearchSession).add("Type", FUNCTION_TYPE);
    Tags.of(userTrackFilterSearchSession).add("Family", FILTER_FAMILY);

    const userTrackFilterSearchSessionEvent = new lambdaEvent.DynamoEventSource(
      filterSessionsDynamoTable,
      {
        startingPosition: lambda.StartingPosition.LATEST,
        batchSize: 10,
        enabled: true,
        retryAttempts: 2,
        reportBatchItemFailures: true, // Check later on code its implications.
      }
    );
    userTrackFilterSearchSession.addEventSource(
      userTrackFilterSearchSessionEvent
    );

    const scheduledPatternsCheckerRole = new iam.Role(
      this,
      "scheduledPatternsCheckerRole",
      {
        assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
        managedPolicies: [
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaBasicExecutionRole"
          ),
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaRole"
          ),
        ],
        roleName: "scheduledPatternsCheckerRole",
      }
    );

    const scheduledPatternsChecker = new lambda.Function(
      this,
      "scheduledPatternsChecker",
      {
        runtime: lambda.Runtime.GO_1_X,
        handler: "main",
        role: scheduledPatternsCheckerRole,
        code: lambda.Code.fromAsset(path),
        functionName: "ScheduledPatternsChecker",
        environment: {
          // TODO -> How can we manage this env variables ?
          neo4jUser: neo4jUserParameter.parameterName,
          neo4jPassword: neo4jPasswordParameter.parameterName,
          neo4jUri: neo4jUriParameter.parameterName,
        },
      }
    );

    Tags.of(scheduledPatternsChecker).add("Type", FUNCTION_TYPE);
    Tags.of(scheduledPatternsChecker).add("Family", SPOT_FAMILY);

    const scheduledPatternsCheckerEvent = new lambdaEvent.SqsEventSource(
      schedulePatternSQS,
      {
        enabled: true,
        batchSize: 10,
      }
    );
    scheduledPatternsChecker.addEventSource(scheduledPatternsCheckerEvent);

    const checkerObject = {
      action: "CreateScheduledSpots",
    };

    const scheduledPatternsCheckerLambdatarget = new eventsTargets.SqsQueue(
      schedulePatternSQS,
      {
        message: RuleTargetInput.fromObject(checkerObject),
      }
    );

    const scheduledPatternsCheckerCron = new events.Rule(
      this,
      "scheduledPatternsCheckerCron",
      {
        enabled: true,
        schedule: events.Schedule.cron({
          minute: "0",
          hour: "0",
          day: "*",
          month: "*",
          year: "*",
        }),
        targets: [scheduledPatternsCheckerLambdatarget],
        ruleName: "ScheduledPatternsCheckerCron",
      }
    );
  }
}
