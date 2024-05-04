package	guild

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


type GuildData struct {
	Data Data `json:"data"`
	Message any `json:"message"`
}

type Data struct {
		GuildID          string `json:"guild_id"`
		Name             string `json:"name"`
		ExternalMessage  string `json:"external_message"`
		BannerColorID    string `json:"banner_color_id"`
		BannerLogoID     string `json:"banner_logo_id"`
		EnrollmentStatus int    `json:"enrollment_status"`
		GalacticPower    int    `json:"galactic_power"`
		GuildType        string `json:"guild_type"`
		LevelRequirement int    `json:"level_requirement"`
		MemberCount      int    `json:"member_count"`
		Members          []Members `json:"members"`
		AvgGalacticPower  int     `json:"avg_galactic_power"`
		AvgArenaRank      float64 `json:"avg_arena_rank"`
		AvgFleetArenaRank float64 `json:"avg_fleet_arena_rank"`
		AvgSkillRating    int     `json:"avg_skill_rating"`
		LastSync          string  `json:"last_sync"`
}

type Members struct {
		// GalacticPower       int    `json:"galactic_power"`
		// GuildJoinTime       string `json:"guild_join_time"`
		// LifetimeSeasonScore int    `json:"lifetime_season_score"`
		MemberLevel         int    `json:"member_level"`
		AllyCode            int    `json:"ally_code"`
		// PlayerLevel         int    `json:"player_level"`
		PlayerName          string `json:"player_name"`
		// LeagueID            string `json:"league_id"`
		// LeagueName          string `json:"league_name"`
		// LeagueFrameImage    string `json:"league_frame_image"`
		// PortraitImage       string `json:"portrait_image"`
		// Title               string `json:"title"`
		// SquadPower          int    `json:"squad_power"`
}

func GetMembers(c Credentials) GuildData {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, nil)
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var data GuildData
	json.Unmarshal(body, &data)
	return data;
}
