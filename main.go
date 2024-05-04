package main

import (
	"log"
	"os"
	"swgoh-api/units"
	"swgoh-api/guild"

	"github.com/joho/godotenv"
)

// SWGOH GG API
const API_BASE_URL = "https://swgoh.gg/api/"
const API_PLAYER_URL = API_BASE_URL + "player/"
const API_UNIT_URL = API_BASE_URL + "units/"

// const API_SHIPS_URL = self::API_BASE_URL . '/api/ships/';
// const API_CHARACTERS_URL = self::API_BASE_URL . '/api/characters/';

func main() {
	// 1) Get list of required units
	//
	// 2) Get list of active users and their rosters - http://api.swgoh.gg/guild-profile/guild_code
	//
	// 3) Get count of characters required - go we have that many characters?
	//
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := guild.Credentials{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Url: "http://api.swgoh.gg/guild-profile/6Q1Rhhi0T26BnkByV1NmxQ",
	}

	//allycode := os.Getenv("ALLYCODE")
	data := guild.GetMembers(c)
	getUnits := false
	if getUnits == true {
		c := units.Credentials{
			Username: os.Getenv("USERNAME"),
			Password: os.Getenv("PASSWORD"),
			Url: API_UNIT_URL,
		}

		//allycode := os.Getenv("ALLYCODE")
		data := units.GetUnits(c)
		log.Println(data.Data[0])
	}
}
