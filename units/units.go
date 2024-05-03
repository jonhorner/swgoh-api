package	units

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

type SwgohGgUnitResponse struct {
	Data []Unit `json:"data"`
}

type Unit struct {
	Name               string                 `json:"name"`
	BaseId             string                 `json:"base_id"`
	AbilityClasses     string                 `json:"ability_classes"`
	ActivateShardCount int                    `json:"activate_shard_count"`
	Alignment          int                    `json:"alignment"`
	Categories         []string               `json:"categories"`
	CombatType         int                    `json:"combat_type"`
	CrewBaseIds        string                 `json:"crew_base_ids"`
	Description        string                 `json:"description"`
	GearLevels         map[string]interface{} `json:"gear_levels"`
	Image              string                 `json:"image"`
	IsCapitalShip      bool                   `json:"is_capital_ship"`
	IsGalacticLegend   bool                   `json:"is_galactic_legend"`
	MadeAvailableOn    string                 `json:"made_available_on"`
	OmicronAbilityIds  []string               `json:"omicron_ability_ids"`
	Power              int                    `json:"power"`
	Role               string                 `json:"role"`
	ShipBaseId         string                 `json:"ship_base_id"`
	ShipSlot           string                 `json:"ship_slot"`
	Url                string                 `json:"url"`
	ZetaAbilityIds     []string               `json:"zeta_ability_ids"`
}

func GetUnits(c Credentials) SwgohGgUnitResponse {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, nil)
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var data SwgohGgUnitResponse
	json.Unmarshal(body, &data)
	return data;
}
