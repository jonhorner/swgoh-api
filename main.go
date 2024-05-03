package main

import (
	"log"
	"os"
	"swgoh-api/units"
	"github.com/joho/godotenv"
)

// SWGOH GG API
const API_BASE_URL = "https://swgoh.gg/api/"
const API_PLAYER_URL = API_BASE_URL + "player/"
const API_UNIT_URL = API_BASE_URL + "units/"

// const API_SHIPS_URL = self::API_BASE_URL . '/api/ships/';
// const API_CHARACTERS_URL = self::API_BASE_URL . '/api/characters/';
// const API_GUILD_URL = self::API_BASE_URL . '/api/guild-profile/{guild_id}';


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c := units.Credentials{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Url: API_UNIT_URL,
	}

	//allycode := os.Getenv("ALLYCODE")
	data := units.GetUnits(c)
	log.Println(data.Data[0])
}
