package main

import (
	"log"
	"os"
	"encoding/json"
	"time"
	"swgoh-api/units"
	"swgoh-api/guild"
	"swgoh-api/db"

	"github.com/davecgh/go-spew/spew"
)

// SWGOH GG API
const API_BASE_URL = "https://swgoh.gg/api/"
const API_PLAYER_URL = API_BASE_URL + "player/"
const API_UNIT_URL = API_BASE_URL + "units/"

// const API_SHIPS_URL = self::API_BASE_URL . '/api/ships/';
// const API_CHARACTERS_URL = self::API_BASE_URL . '/api/characters/';

func main() {
	updateGuild, getRequiredUnits := false, true

	// 1) Get list of required units
	//
	// 2) Get list of active users and their rosters - http://api.swgoh.gg/guild-profile/guild_code
	//		- store ally code, unit key, relic level, gear level, GP (for getting best person to farm a unit)])
	// 3) Get count of characters required - go we have that many characters?
	//

	if getRequiredUnits {
		requiredUnits, err := db.GetRequiredUnits("Mustafar")
		if err != nil {
			log.Println("Got error calling GetRequiredUnits:")
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Println("requiredUnits:");
		//spew.Dump(requiredUnits.Value.Tb3Platoon1)

		// count characters
		for i, v := range requiredUnits.Value {

		}
	}

	if updateGuild {
		c := guild.Credentials{
			Username: os.Getenv("USERNAME"),
			Password: os.Getenv("PASSWORD"),
			Url: "http://api.swgoh.gg/guild-profile/6Q1Rhhi0T26BnkByV1NmxQ",
		}

		data := guild.GetMembers(c)
		content, err := json.Marshal(data.Data.Members)
		if err != nil {
			log.Fatal("Error getting data: ", err)
		}
		dataStore := db.StoreMemberData{
			Table: "SwgohGuildData",
			KeyName: "SwgohGuild",
			KeyValue: "6Q1Rhhi0T26BnkByV1NmxQ",
			ValueName: "MemberData",
			Value: string(content),
		}

		db.StoreGuildMembers(dataStore)

		getUnits := false
		if getUnits == true {
			c := units.Credentials{
				Username: os.Getenv("USERNAME"),
				Password: os.Getenv("PASSWORD"),
				Url: API_UNIT_URL,
			}
			data := units.GetUnits(c)
			log.Println(data.Data[0])
		}
	}
}
