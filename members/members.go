package members

import (
	"net/http"
	"log"
	"io"
	"encoding/json"
)

type Credentials struct {
	Username string
	Password string
	Url string
}

type Unit struct {
	BaseID    string `json:"base_id"`
	Name      string `json:"name"`
	RelicTier int      `json:"relic_tier"`
}

type MemberData struct {
	AllyCode  int    `json:"ally_code"`
	Name      string `json:"name"`
}

type Roster struct {
	Units []Unit `json:"units"`
	Data MemberData `json:"data"`
}

type GuildMembers struct {
	Members []struct {
		GalacticPower       int    `json:"galactic_power"`
		GuildJoinTime       string `json:"guild_join_time"`
		LifetimeSeasonScore int    `json:"lifetime_season_score"`
		MemberLevel         int    `json:"member_level"`
		AllyCode            int    `json:"ally_code"`
		PlayerLevel         int    `json:"player_level"`
		PlayerName          string `json:"player_name"`
		LeagueID            string `json:"league_id"`
		LeagueName          string `json:"league_name"`
		LeagueFrameImage    string `json:"league_frame_image"`
		PortraitImage       string `json:"portrait_image"`
		Title               string `json:"title"`
		SquadPower          int    `json:"squad_power"`
	} `json:"members"`
}

type GuildMemberData struct {
	GalacticPower       int    `json:"galactic_power"`
	GuildJoinTime       string `json:"guild_join_time"`
	LifetimeSeasonScore int    `json:"lifetime_season_score"`
	MemberLevel         int    `json:"member_level"`
	AllyCode            int    `json:"ally_code"`
	PlayerLevel         int    `json:"player_level"`
	PlayerName          string `json:"player_name"`
	LeagueID            string `json:"league_id"`
	LeagueName          string `json:"league_name"`
	LeagueFrameImage    string `json:"league_frame_image"`
	PortraitImage       string `json:"portrait_image"`
	Title               string `json:"title"`
	SquadPower          int    `json:"squad_power"`
}

func GetRoster(c Credentials) Roster {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, nil)
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var data Roster
	json.Unmarshal(body, &data)
	return data;
}
