import { Stack, StackProps } from 'aws-cdk-lib';
import { aws_apigateway as apigateway, aws_lambda as lambda } from 'aws-cdk-lib';
import { Construct } from 'constructs';
// import * as sqs from 'aws-cdk-lib/aws-sqs';

export class InfraStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    function addMethodToApiGateway(lambdaToAdd: lambda.Function, apigatewayToHandle: apigateway.RestApi, methodName: string ): apigateway.Resource{
      const spotMethod = apigatewayToHandle.root.addResource(methodName);

      const lambdaIntegration = new apigateway.LambdaIntegration(lambdaToAdd, {
        proxy: true
      });

      spotMethod.addMethod('GET', lambdaIntegration);
      spotMethod.addMethod('POST', lambdaIntegration);
      spotMethod.addMethod('PUT', lambdaIntegration);
      spotMethod.addMethod('DELETE', lambdaIntegration);

      return spotMethod;
    }

    const mainApiGateway = new apigateway.RestApi(this, 'mainApiGateway', {

    })

    mainApiGateway.root.addMethod("ANY");

    const spotLambda = new lambda.Function(this, 'spotLambda', {
      runtime : lambda.Runtime.GO_1_X,
      handler: 'main',
      code: lambda.Code.fromInline('')
    })

    addMethodToApiGateway(spotLambda, mainApiGateway, "spot")
  }
}
