import { Duration, Stack, StackProps } from "aws-cdk-lib";
import { Construct } from "constructs";

import {
  aws_events as events,
  aws_events_targets as events_targets,
  aws_lambda as lambda,
  aws_iam as iam,
  aws_sns as sns,
  aws_sqs as sqs,
  aws_sns_subscriptions as sns_subscriptions,
  aws_lambda_event_sources as lambda_event_sources,
} from "aws-cdk-lib";

import { join } from "path";

export class InfraChangeNotifier extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    let path = join(__dirname, "assets", "main.zip");
    const lambdaPrincipal = new iam.ServicePrincipal("lambda.amazonaws.com");

    const infraChangesSNS = new sns.Topic(this, "infraChangesSNS", {
      displayName: "infraChangesSNS",
      topicName: "infraChangesSNS.fifo",
      fifo: true,
      contentBasedDeduplication: true
    });

    const emailNotifierSQS = new sqs.Queue(this, "emailNotifierSQS", {
      queueName: "emailNotifierSQS.fifo",
      contentBasedDeduplication: true,
      fifo: true,
      visibilityTimeout: Duration.seconds(30),
    });

    const emailSenderLambdaRole = new iam.Role(this, "emailSenderLambdaRole", {
      assumedBy: lambdaPrincipal,
      roleName: "emailSenderLambdaRole",
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaBasicExecutionRole"
        ),
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaRole"
        ),
      ],
      description: "Role for calling ses and basic lambda roles.",
    });

    emailSenderLambdaRole.addToPolicy(
      new iam.PolicyStatement({
        actions: [
          "ses:ListTemplates",
          "ses:SendEmail",
          "ses:SendTemplatedEmail",
          "ses:GetSendQuota",
          "ses:SendRawEmail",
          "ses:GetTemplate",
          "ses:ListIdentities",
        ],
        effect: iam.Effect.ALLOW,
        resources: ["*"],
      })
    );

    const emailSenderLambda = new lambda.Function(this, "emailSenderLambda", {
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromAsset(path),
      handler: "main",
      functionName: "EmailSender",
      role: emailSenderLambdaRole,
      environment: {
        from: "now.inc.project@gmail.com",
        to: "jquirozcadavid@gmail.com",
        TWILIO_ACCOUNT_SID: "_MISSING_",
        TWILIO_AUTH_TOKEN: "_MISSING_",
      },
    });
    
    emailSenderLambda.addEventSource(new lambda_event_sources.SqsEventSource(emailNotifierSQS, {
      batchSize: 10,
      enabled: true,
      reportBatchItemFailures: true // We need to be careful about this.
    } ))

    infraChangesSNS.addSubscription(
      new sns_subscriptions.SqsSubscription(emailNotifierSQS, {
        rawMessageDelivery: true,
      })
    );

    const awsServiceChangesNotifierLambdaRole = new iam.Role(
      this,
      "awsServiceChangesNotifierLambdaRole",
      {
        assumedBy: lambdaPrincipal,
        roleName: "awsServiceChangesNotifierLambdaRole",
        managedPolicies: [
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaBasicExecutionRole"
          ),
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaRole"
          ),
        ],
        description: "Role for calling sns and basic lambda roles.",
      }
    );

    awsServiceChangesNotifierLambdaRole.addToPolicy(
      new iam.PolicyStatement({
        actions: [
          "sns:Publish",
          "sns:GetTopicAttributes",
          "sns:ListTopics",
          "sns:Unsubscribe",
          "sns:Subscribe",
          "sns:AddPermission",
        ],
        effect: iam.Effect.ALLOW,
        resources: ["*"],
      })
    );

    const awsServiceChangesNotifierLambda = new lambda.Function(
      this,
      "awsServiceChangesNotifierLambda",
      {
        runtime: lambda.Runtime.GO_1_X,
        code: lambda.Code.fromAsset(path),
        handler: "main",
        functionName: "AWSServiceChangesNotifier",
        role: awsServiceChangesNotifierLambdaRole,
        environment: {
          infraChangesTopicArn: infraChangesSNS.topicArn,
          infraChangesTopicName: infraChangesSNS.topicName,
        },
      }
    );

    const rdsEventsRuleTarget = new events_targets.LambdaFunction(
      awsServiceChangesNotifierLambda
    );

    const rdsEventsRule = new events.Rule(this, "rdsEventsRule", {
      enabled: true,
      description:
        "This will map all instances changes regarding stop / start instance",
      ruleName: "rdsEventsRule",
      eventPattern: {
        source: ["aws.rds"],
        detailType: ["RDS DB Instance Event"],
      },
      targets: [rdsEventsRuleTarget],
    });
  }
}
