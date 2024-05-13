package db

import (
	"log"
	"time"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type StoreMemberData struct {
	Table string
	KeyName string
	KeyValue string
	ValueName string
	Value string
}

type GuildItem struct {
	SwgohGuild string `json:"swgohGuild,omitempty" dynamodbav:",omitempty"`
	MemberData string `json:"memberData,omitempty" dynamodbav:",omitempty"`
	UpdateTime string `json:"updateTime,omitempty" dynamodbav:",omitempty"`
}

func StoreGuildMembers(data StoreMemberData) {
	storeItem(data)
}

func storeItem(data StoreMemberData) {
	sesssion := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sesssion)
	now := time.Now()      // current local time
	ts := now.Unix()

	item := GuildItem{
		SwgohGuild:  data.KeyValue,
		MemberData: data.Value,
		UpdateTime: string(ts),
	}

	av, err := dynamodbattribute.MarshalMap(item)
	// log.Println(av)
	if err != nil {
		log.Println("Error marshalling map:")
		log.Println(err.Error())
		os.Exit(1)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(data.Table),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		log.Println("Got error calling PutItem:")
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println("Successfully added item to table")
}
