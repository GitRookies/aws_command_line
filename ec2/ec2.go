package ec2

import (
	"fmt"
	"os"

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

	// // On/Off
	// svc := ec2.New(sess)
	// action := os.Args[2]
	// action = strings.ToLower(action)

	// fmt.Println("********************************8")
	// fmt.Println(svc)
	// fmt.Println(action)
	// fmt.Println("###############################33")
	// if action == "on" {
	// 	input := &ec2.MonitorInstancesInput{
	// 		InstanceIds: []*string{
	// 			aws.String(conf.InstanceId),
	// 		},
	// 		DryRun: aws.Bool(true),
	// 	}
	// 	result, err := svc.MonitorInstances(input)
	// 	awsErr, ok := err.(awserr.Error)

	// 	if ok && awsErr.Code() == "DryRunOperation" {
	// 		input.DryRun = aws.Bool(false)
	// 		result, err = svc.MonitorInstances(input)
	// 		if err != nil {
	// 			fmt.Println("Error", err)
	// 		} else {
	// 			fmt.Println("Success", result.InstanceMonitorings)
	// 		}
	// 	} else {
	// 		fmt.Println("Error", err)
	// 	}

	// } else if action == "off" {
	// 	// Turn the instance off
	// } else if action == "restart" {
	// 	// Restart the server
	// } else {
	// 	fmt.Println("No action performed")
	// }

	// Create new EC2 client
	svc := ec2.New(sess)

	// Turn monitoring on
	if os.Args[2] == "START" {
		// We set DryRun to true to check to see if the instance exists and we have the
		// necessary permissions to monitor the instance.
		input := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(conf.InstanceId),
			},
			DryRun: aws.Bool(true),
		}
		fmt.Println("Printing the input", input)
		result, err := svc.StartInstances(input)
		awsErr, ok := err.(awserr.Error)

		// If the error code is `DryRunOperation` it means we have the necessary
		// permissions to Start this instance
		if ok && awsErr.Code() == "DryRunOperation" {
			// Let's now set dry run to be false. This will allow us to start the instances
			input.DryRun = aws.Bool(false)
			result, err = svc.StartInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Success", result.StartingInstances)
			}
		} else { // This could be due to a lack of permissions
			fmt.Println("Error", err)
		}
	} else if os.Args[2] == "STOP" { // Turn instances off
		input := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String(conf.InstanceId),
			},
			DryRun: aws.Bool(true),
		}
		result, err := svc.StopInstances(input)
		awsErr, ok := err.(awserr.Error)
		if ok && awsErr.Code() == "DryRunOperation" {
			input.DryRun = aws.Bool(false)
			result, err = svc.StopInstances(input)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("Success", result.StoppingInstances)
			}
		} else {
			fmt.Println("Error", err)
		}
	}
}
