#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { InfraStack } from '../lib/infra-stack';
import { InfraMoneySavers } from '../lib/infra-helpers';

const app = new cdk.App();
const defaultRegion = "us-east-2"

new InfraStack(app, 'InfraStack', {
  env: { account: process.env.CDK_DEFAULT_ACCOUNT, region:  defaultRegion},
});

new InfraMoneySavers(app, 'InfraMoneySavers', {
  env: { account: process.env.CDK_DEFAULT_ACCOUNT, region:  defaultRegion},
})