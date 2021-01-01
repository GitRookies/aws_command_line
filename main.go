package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davetweetlive/aws_command_line/ec2"
)

func main() {
	choices := os.Args[1:]
	fmt.Println(choices[0])

	if len(choices) == 0 {
		fmt.Println(strings.Repeat("#", 40))
		fmt.Println("Make sure you peovide the service name as the first parameter!")
		fmt.Println("Eg. 'aws ec2 start'")
		fmt.Println(strings.Repeat("#", 40))
	}
	if choices[0] == "ec2" {
		ec2.EC2Workflow()
	}
}
