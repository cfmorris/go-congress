package model

type Member struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		Congress   string `json:"congress"`
		Chamber    string `json:"chamber"`
		NumResults int    `json:"num_results"`
		Offset     int    `json:"offset"`
		Members    []struct {
			ID                   string      `json:"id"`          //p
			ShortTitle           string      `json:"short_title"` //p
			APIURI               string      `json:"api_uri"`     //p
			FirstName            string      `json:"first_name"`  //p
			MiddleName           interface{} `json:"middle_name"`
			LastName             string      `json:"last_name"` //p
			Suffix               interface{} `json:"suffix"`
			DateOfBirth          string      `json:"date_of_birth"`    //p
			Gender               string      `json:"gender"`           //p
			Phone                string      `json:"phone"`            //p
			OcdID                string      `json:"ocd_id"`           //p
			TwitterAccount       string      `json:"twitter_account"`  //p
			FacebookAccount      string      `json:"facebook_account"` //p
			YoutubeAccount       interface{} `json:"youtube_account"`
			GovtrackID           string      `json:"govtrack_id"`             //p
			CspanID              string      `json:"cspan_id"`                //p
			VotesmartID          string      `json:"votesmart_id"`            //p
			IcpsrID              string      `json:"icpsr_id"`                //p
			CrpID                string      `json:"crp_id"`                  //p
			GoogleEntityID       string      `json:"google_entity_id"`        //p
			FecCandidateID       string      `json:"fec_candidate_id"`        //p
			URL                  string      `json:"url"`                     //p
			RssURL               string      `json:"rss_url"`                 //p
			ContactForm          interface{} `json:"contact_form"`            //p
			InOffice             bool        `json:"in_office"`               //p
			CookPvi              string      `json:"cook_pvi"`                //p
			NextElection         string      `json:"next_election,omitempty"` //p
			LastUpdated          string      `json:"last_updated"`            //p
			IdealPoint           interface{} `json:"ideal_point"`
			Seniority            string      `json:"seniority"`       //r
			TotalVotes           int         `json:"total_votes"`     //r
			MissedVotes          int         `json:"missed_votes"`    //r
			TotalPresent         int         `json:"total_present"`   //r
			DwNominate           float64     `json:"dw_nominate"`     //r
			Office               string      `json:"office"`          //r
			Party                string      `json:"party"`           //r
			LeadershipRole       string      `json:"leadership_role"` //r
			Fax                  interface{} `json:"fax"`
			State                string      `json:"state"`    //r
			Title                string      `json:"title"`    //r
			District             string      `json:"district"` //r
			AtLarge              bool        `json:"at_large"` //r
			Geoid                string      `json:"geoid"`    //r
			MissedVotesPct       float64     `json:"missed_votes_pct,omitempty"`
			VotesWithPartyPct    float64     `json:"votes_with_party_pct,omitempty"`
			VotesAgainstPartyPct float64     `json:"votes_against_party_pct,omitempty"`
		} `json:"members"`
	} `json:"results"`
}

type Comittee struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		Congress   string `json:"congress"`
		Chamber    string `json:"chamber"`
		NumResults int    `json:"num_results"`
		Committees []struct {
			ID              string      `json:"id"`
			Name            string      `json:"name"`
			Chamber         string      `json:"chamber"`
			URL             string      `json:"url"`
			APIURI          string      `json:"api_uri"`
			Chair           interface{} `json:"chair"`
			ChairID         interface{} `json:"chair_id"`
			ChairParty      interface{} `json:"chair_party"`
			ChairState      interface{} `json:"chair_state"`
			ChairURI        interface{} `json:"chair_uri"`
			RankingMemberID interface{} `json:"ranking_member_id"`
			Subcommittees   []struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				APIURI string `json:"api_uri"`
			} `json:"subcommittees"`
		} `json:"committees"`
	} `json:"results"`
}
