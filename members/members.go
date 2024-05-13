package members


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
