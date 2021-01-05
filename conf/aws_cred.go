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
	Region          string
	InstanceId      string
)

type AWS struct {
	AwsCred `yaml:"Cred"`
	General `yaml:"General"`
}

type AwsCred struct {
	AccessKey string `yaml:"AccessKey"`
	Secret    string `yaml:"Secret"`
}

type General struct {
	Region     string `yaml:"Region"`
	InstanceID string `yaml:"InstanceID"`
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

	AccessTokenEC2 = awsCred.AccessKey
	AccessSecretEC2 = awsCred.Secret
	Region = awsCred.Region
	InstanceId = awsCred.InstanceID
}
