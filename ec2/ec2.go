package ec2

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/davetweetlive/aws_command_line/conf"
)

func EC2Workflow() {
	// Establishing session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Region),
		Credentials: credentials.NewStaticCredentials(conf.AccessTokenEC2, conf.AccessSecretEC2, ""),
	},
	)

	if err != nil {
		// TODO: Handle error

	}

	// On/Off
	svc := ec2.New(sess)
	action := os.Args[2]
	action = strings.ToLower(action)
	if action == "on" {
		input := &ec2.MonitorInstancesInput{
			InstanceIds: []*string{
				aws.String(conf.InstanceId),
			},
			DryRun: aws.Bool(true),
		}
		result, err := svc.MonitorInstances(input)
		awsErr, ok := err.(awserr.Error)

		if ok && awsErr.Code() == "DryRunOperation" {
			input.DryRun = aws.Bool(false)
			result, err = svc.MonitorInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Success", result.InstanceMonitorings)
			}
		} else {
			fmt.Println("Error", err)
		}

	} else if action == "off" {
		// Turn the instance off
	} else if action == "restart" {
		// Restart the server
	} else {
		fmt.Println("No action performed")
	}
}
