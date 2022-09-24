import { Stack, StackProps } from "aws-cdk-lib";
import { Construct } from "constructs";
import {
  aws_events as events,
  aws_events_targets as eventsTargets,
  aws_lambda as lambda,
  aws_iam as iam,
} from "aws-cdk-lib";
import { join } from "path";

export class InfraMoneySavers extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    let path = join(__dirname, "assets", "main.zip");
    const lambdaPrincipal = new iam.ServicePrincipal("lambda.amazonaws.com");

    const rdsMoneySaverLambdaRole = new iam.Role(
      this,
      "rdsMoneySaverLambdaRole",
      {
        assumedBy: lambdaPrincipal,
        managedPolicies: [
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaBasicExecutionRole"
          ),
          iam.ManagedPolicy.fromAwsManagedPolicyName(
            "service-role/AWSLambdaRole"
          ),
        ],
        roleName: "rdsMoneySaverLambdaRole",
        description: "Role for calling rds api and basic lambda roles.",
      }
    );

    rdsMoneySaverLambdaRole.addToPolicy(
      new iam.PolicyStatement({
        actions: [
          "rds:StartDBCluster",
          "rds:StopDBCluster",
          "rds:StartDBInstanceAutomatedBackupsReplication",
          "rds:StopDBInstance",
          "rds:StartDBInstance",
        ],
        resources: ["*"],
        effect: iam.Effect.ALLOW,
        sid: "AllowRDSinstaceStartAndStop",
      })
    );

    const rdsMoneySaverLambda = new lambda.Function(
      this,
      "rdsMoneySaverLambda",
      {
        runtime: lambda.Runtime.GO_1_X,
        code: lambda.Code.fromAsset(path),
        handler: "main",
        functionName: "RdsMoneySaver",
        role: rdsMoneySaverLambdaRole,
      }
    );

    const stopInstaceBody = {
      action: "STOP",
      instances: ["locations-db"],
    };
    const rdsMoneySaverStopLambdatarget = new eventsTargets.LambdaFunction(
      rdsMoneySaverLambda,
      {
        event: events.RuleTargetInput.fromObject(stopInstaceBody),
      }
    );

    const RdsMoneySaverStopRDSCron = new events.Rule(this, "RdsMoneySaverStopRDSCron", {
      enabled: true,
      schedule: events.Schedule.cron({
        minute: "0",
        hour: "14",
        day: "*",
        month: "*",
        year: "*",
      }),
      targets: [rdsMoneySaverStopLambdatarget],
      ruleName: "RdsMoneySaverStopRDSCron",
    });

    const startInstaceBody = {
      action: "START",
      instances: ["locations-db"],
    };

    const rdsMoneySaverStartLambdatarget = new eventsTargets.LambdaFunction(
      rdsMoneySaverLambda,
      {
        event: events.RuleTargetInput.fromObject(startInstaceBody),
      }
    );

    const RdsMoneySaverStartRDSCron = new events.Rule(
      this,
      "RdsMoneySaverStartRDSCron",
      {
        enabled: true,
        schedule: events.Schedule.cron({
          minute: "55",
          hour: "10",
          day: "*",
          month: "*",
          year: "*",
        }),
        targets: [rdsMoneySaverStartLambdatarget],
        ruleName: "RdsMoneySaverStartRDSCron",
      }
    );
  }
}
