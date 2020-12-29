package conf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	AccessTokenEC2  string
	AccessSecretEC2 string
)

type AWS struct {
	EC2 AwsCred `yaml:"Cred"`
}

type AwsCred struct {
	AccessKey string `yaml:"AccessKey"`
	Secret    string `yaml:"Secret"`
}

func init() {
	file, err := os.Open("/home/dave/Conf/aws.yaml")
	if err != nil {
		log.Println("Couldn't find credentials file!")
	}

	accessBits, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error occoured in reading the file")
	}

	awsCred := AWS{}

	if err = yaml.Unmarshal([]byte(accessBits), &awsCred); err != nil {
		fmt.Println("Can't unmarshal a credentials file!")
	}

	AccessTokenEC2 = awsCred.EC2.AccessKey
	AccessSecretEC2 = awsCred.EC2.Secret
}
