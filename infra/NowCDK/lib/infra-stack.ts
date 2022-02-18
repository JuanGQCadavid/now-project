import { Stack, StackProps } from 'aws-cdk-lib';

import 
{ 
  aws_apigateway as apigateway,
  aws_lambda as lambda,
  aws_sns as sns,
  aws_sns_subscriptions as subscriptions,
  aws_sqs as sqs,
  aws_lambda_event_sources as lambdaEvent,
  aws_iam as iam
} from 'aws-cdk-lib';

import { Construct } from 'constructs';
import { join } from 'path';
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class InfraStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    function addMethodToApiGateway(lambdaToAdd: lambda.Function, apigatewayToHandle: apigateway.RestApi, methodName: string ): apigateway.Resource{
      const method = apigatewayToHandle.root.addResource( methodName);
      const proxyMethod = method.addResource(`{${methodName}+}`);

      const lambdaIntegration = new apigateway.LambdaIntegration(lambdaToAdd, {
        proxy: true
      });
      method.addMethod('ANY', lambdaIntegration)
      
      proxyMethod.addMethod('GET', lambdaIntegration);
      proxyMethod.addMethod('POST', lambdaIntegration);
      proxyMethod.addMethod('PUT', lambdaIntegration);
      proxyMethod.addMethod('DELETE', lambdaIntegration);

      return method;
    }

    // MAIN API GATEWAY
    const mainApiGateway = new apigateway.RestApi(this, 'mainApiGateway', {

    })

    mainApiGateway.root.addMethod("ANY");

    // SNS
    // TODO -> extra config ?
    const spotActivityTopic = new sns.Topic(this, 'spotActivityTopic' , {
      displayName: "spotActivityTopic",
      topicName:"spotActivityTopic"
    });

    // SQS

    // TODO -> extra config ?
    const updateLocationDataSQS = new sqs.Queue(this, "updateLocationDataSQS", {
      queueName:"updateLocationDataSQS",
    })

    spotActivityTopic.addSubscription( new subscriptions.SqsSubscription(updateLocationDataSQS))

    // LAMBDAS

    let path = join(__dirname, 'assets', 'main.zip');
    console.log(path)

    // -> spotLambda

    const spotLambdaRole = new iam.Role(this, 'spotLambdaRole', {
      assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com'),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName("service-role/AWSLambdaBasicExecutionRole"),
        iam.ManagedPolicy.fromAwsManagedPolicyName("service-role/AWSLambdaRole")
      ],
      roleName: "spotLambdaRole",  
      description: 'Spot Lambda Role',
    });

    spotLambdaRole.addToPolicy(new iam.PolicyStatement({
      actions:[
        "sns:*"
      ],
      effect: iam.Effect.ALLOW,
      resources: [
        spotActivityTopic.topicArn
      ]
    }))

    const spotLambda = new lambda.Function(this, 'spotLambda', {
      runtime : lambda.Runtime.GO_1_X,
      handler: 'main',
      role: spotLambdaRole, 
      code: lambda.Code.fromAsset(path),
      functionName: "SpotService",
      environment: { // TODO -> How can we manage this env variables ?
        neo4jUser: "neo4jUser",
        neo4jPassword: "neo4jPassword",
        neo4jUri: "neo4jUri",
        snsArn: spotActivityTopic.topicArn
      }
    })



    const spotMethod = addMethodToApiGateway(spotLambda, mainApiGateway, "spot")
    
    // -> filterLambda
    const filterLambda = new lambda.Function(this, 'filterLambda', {
      runtime : lambda.Runtime.GO_1_X,
      handler: 'main',
      code: lambda.Code.fromAsset(path),
      functionName: "FilterService",
      environment: {
        dbUser: "dbUser",
        dbPassword: "dbPassword",
        dbName: "dbName",
        dbUrl: "dbUrl",
        spotServiceURL: `https://${mainApiGateway.restApiId}.execute-api.${this.region}.amazonaws.com/prod${spotMethod.path}`
      }
    })

    addMethodToApiGateway(filterLambda, mainApiGateway, "filter")

    // -> locationDataUpdater
    const locationDataUpdater = new lambda.Function(this, 'locationDataUpdater', {
      runtime : lambda.Runtime.GO_1_X,
      handler: 'main',
      code: lambda.Code.fromAsset(path),
      functionName: "LocationDataUpdater",
      environment:{
        dbUser: "dbUser",
        dbPassword: "dbPassword",
        dbName: "dbName",
        dbUrl: "dbUrl"
      }
    })

    const locationDataUpdaterEvent = new lambdaEvent.SqsEventSource(updateLocationDataSQS,{
      enabled: true,
      batchSize: 10
    })

    locationDataUpdater.addEventSource(locationDataUpdaterEvent)



  }
}
