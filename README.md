# basic-lambda
getting started with Lambda and Go in AWS

## High-level deployment steps
- develop code
- build binary
- zip to push
## AWS Lambda: Overview
- An AWS Lambda application is a combination of Lambda functions, event sources, and other resources that work together to perform tasks.
-  You can use AWS CloudFormation and other tools to collect your application's components into a single package that can be deployed and managed as one resource.
-  Applications make your Lambda projects portable and enable you to integrate with additional developer tools, such as AWS CodePipeline, AWS CodeBuild, and the AWS Serverless Application Model command line interface (AWS SAM CLI).
## create simple lambda from console
1. navigate to Lambda in the AWS console
2. choose the **Author from scratch** tabe
3. choose the `Go 1.x` runtime
4. choose `x86_64` arc
5. Permissions: 
  - *By default, Lambda will create an execution role with permissions to upload logs to Amazon CloudWatch Logs. You can customize this default role later when adding triggers*
  - LEAVE THIS SETTING AT DEFAULT FOR THE FIRST LAMBDA, THEN OBSERVE WHAT THE ROLE DOES
6. choose **create function**
7. after the Lambda builds, navigate to the Test tab
8. give the Test a name name and accept the default JSON strings
9. run the test
10. at bottom of the page named after your test, you will see the Test log:
```
START RequestId: 80fa443c-5a3c-4349-89a2-b1e8f475d649 Version: $LATEST
END RequestId: 80fa443c-5a3c-4349-89a2-b1e8f475d649
REPORT RequestId: 80fa443c-5a3c-4349-89a2-b1e8f475d649	Duration: 1.58 ms	Billed Duration: 2 ms	Memory Size: 512 MB	Max Memory Used: 29 MB	Init Duration: 204.52 ms	
```
11. there will also be a link to *view the corresponding CloudWatch log group.*
12. your test will have created a Cloudwatch log group similar to this:
  + `arn:aws:logs:us-east-1:257749931526:log-group:/aws/lambda/simpleGoLambda:*`
  + the log group will have a *log stream*
  + click on the log stream link
13. go back to the lambda page and copy the SRC code that the Lambda created
14. notice the message: The code editor does not support the Go 1.x runtime.
15. examine the IAM configuration that was created automatically to support this Lambda deployment
  - navigate to IAM | roles and search for "lambda"
  - the role name is `simpleGoLambda-role-qtx1glpa`
  - the role ARN is `arn:aws:iam::257749931526:role/service-role/simpleGoLambda-role-qtx1glpa`
  - the role has a **policy** attached to it:
  - policy ARN `arn:aws:iam::257749931526:policy/service-role/AWSLambdaBasicExecutionRole-0d24a91e-d899-4d6f-9ba7-3e29bfc5cd17`
  - examine the JSON policy:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": "arn:aws:logs:us-east-1:257749931526:*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
                "arn:aws:logs:us-east-1:257749931526:log-group:/aws/lambda/simpleGoLambda:*"
            ]
        }
    ]
}
```
  - also note the **trust relationship* of the policy:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Action": "sts:AssumeRole"
        }
    ]
}
```
## github example Lambda using AWS example code
1. `go mod init github.com/robertocamp/basic-lambda`
2. develop code in main.go
3. Preparing a binary to deploy to AWS Lambda requires that it is compiled for Linux and placed into a .zip file:
  + build an executable called "bootstrap" at the root of the project
  + `GOOS=linux GOARCH=amd64 go build -o main main.go`
  + `zip lambda-handler.zip main`
## Links
https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html
https://github.com/aws/aws-lambda-go