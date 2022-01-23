import { Stack, StackProps } from 'aws-cdk-lib';
import { aws_apigateway as apigateway, aws_lambda as lambda } from 'aws-cdk-lib';
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

    const mainApiGateway = new apigateway.RestApi(this, 'mainApiGateway', {

    })

    mainApiGateway.root.addMethod("ANY");

    let path = join(__dirname, 'assets', 'main.zip');
    console.log(path)
    const spotLambda = new lambda.Function(this, 'spotLambda', {
      runtime : lambda.Runtime.GO_1_X,
      handler: 'main',
      code: lambda.Code.fromAsset(path)
    })

    addMethodToApiGateway(spotLambda, mainApiGateway, "spot")
  }
}
