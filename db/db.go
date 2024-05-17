package db

import (
	"log"
	"os"
	"fmt"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const OperationRequirementsKey = "OperationRequirements"

type StoreMemberData struct {
	Table string
	KeyName string
	KeyValue string
	ValueName string
	Value string
}

type StoreMemberDataV2 struct {
	Table string
	PK string
	SK string
	Key string
	Value interface{}
}

type GuildItem struct {
	SwgohGuild string `json:"swgohGuild,omitempty" dynamodbav:",omitempty"`
	MemberData string `json:"memberData,omitempty" dynamodbav:",omitempty"`
	UpdateTime string `json:"updateTime,omitempty" dynamodbav:",omitempty"`
}

type GuildItemv2 struct {
	PK string `json:"SK,omitempty" dynamodbav:",omitempty"`
	SK string `json:"PK,omitempty" dynamodbav:",omitempty"`
	Key string `json:"key,omitempty" dynamodbav:",omitempty"`
	Value interface{} `json:"value,omitempty" dynamodbav:",omitempty"`
}

// Struct for DB item which stores required characters
// for Territory Battle Operations
type TbDataItem struct {
	PK string `json:"SK,omitempty" dynamodbav:",omitempty"`
	SK string `json:"PK,omitempty" dynamodbav:",omitempty"`
	Key string `json:"key,omitempty" dynamodbav:",omitempty"`
	Value TbPhaseOperations `json:"value,omitempty" dynamodbav:",omitempty"`
}

type TbPhaseOperations struct {
	Tb3Platoon1 TbOperation `json:"tb3-platoon-1,omitempty"`
	Tb3Platoon2 TbOperation `json:"tb3-platoon-2,omitempty"`
	Tb3Platoon3 TbOperation `json:"tb3-platoon-3,omitempty"`
	Tb3Platoon4 TbOperation `json:"tb3-platoon-4,omitempty"`
	Tb3Platoon5 TbOperation `json:"tb3-platoon-5,omitempty"`
	Tb3Platoon6 TbOperation `json:"tb3-platoon-6,omitempty"`
}

type TbOperation struct {
	OperationRequirements OperationRequirements `json:"OperationRequirements"`
}

type OperationRequirements struct {
	UnitBaseID1  string `json:"unitBaseId1"`
	UnitBaseID2  string `json:"unitBaseId2"`
	UnitBaseID3  string `json:"unitBaseId3"`
	UnitBaseID4  string `json:"unitBaseId4"`
	UnitBaseID5  string `json:"unitBaseId5"`
	UnitBaseID6  string `json:"unitBaseId6"`
	UnitBaseID7  string `json:"unitBaseId7"`
	UnitBaseID8  string `json:"unitBaseId8"`
	UnitBaseID9  string `json:"unitBaseId9"`
	UnitBaseID10 string `json:"unitBaseId10"`
	UnitBaseID11 string `json:"unitBaseId11"`
	UnitBaseID12 string `json:"unitBaseId12"`
	UnitBaseID13 string `json:"unitBaseId13"`
	UnitBaseID14 string `json:"unitBaseId14"`
	UnitBaseID15 string `json:"unitBaseId15"`
}

func StoreGuildMembers(data StoreMemberDataV2) {
	storeItem(data)
}

func GetGuildMembers() (GuildItemv2, error) {
	sesssion := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamodb.New(sesssion)

	tableName := "SwgohGuildData"
	key := "1"
	guildId := os.Getenv("GUILD_ID")
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
		    "PK": {S: aws.String(guildId)},
		    "SK": {S: aws.String(key)},
		  },
	})

	if err != nil {
		log.Fatalf("Got error calling GetGuildMembers: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find guild data"
		return GuildItemv2{}, errors.New(msg)
	}

	item := GuildItemv2{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return item, nil
}

func storeItem(data StoreMemberDataV2) {
	sesssion := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sesssion)
	item := GuildItemv2{
		PK:  data.PK,
		SK: data.SK,
		Key: data.Key,
		Value: data.Value,
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

// Get list of required units for the specified planet
func GetRequiredUnits(Planet string) (TbDataItem, error)  {
	if Planet != "Mustafar" {
		log.Println("Invalid planet parameter")
	}
	sesssion := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamodb.New(sesssion)

	tableName := "SwgohGuildData"
	key := "OperationRequirements" + Planet
	guildId := os.Getenv("GUILD_ID")
	log.Println(key);
	log.Println(guildId);
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
		    "PK": {S: aws.String(guildId)},
		    "SK": {S: aws.String(key)},
		  },
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find item"
		return TbDataItem{}, errors.New(msg)
	}

	item := TbDataItem{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return item, nil
}
