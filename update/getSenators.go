package update

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cfmorris/go-congress/config"
	"github.com/cfmorris/go-congress/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetSenators() {
	config, _ := config.ReadConfig("./config.yaml")

	fmt.Println("updating Senators....")
	for i := 80; i <= 117; i++ {
		time.Sleep(3 * time.Second)
		house := "Senate" // "Senate" or "House"
		// baseUrl := "https://api.propublica.org/congress/v1/"
		url := "https://api.propublica.org/congress/v1/" + strconv.Itoa(int(i)) + "/" + house + "/members.json"

		congressClient := http.Client{
			Timeout: time.Second * 5, // Timeout after 5 seconds
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("X-API-Key", config.Keys.PpApi)

		res, getErr := congressClient.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		people1 := model.Member{}
		jsonErr := json.Unmarshal(body, &people1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		fmt.Println(people1.Results)
		for _, k := range people1.Results {
			for b, l := range k.Members {
				fmt.Println(b, l.LastName)
				fmt.Println("")

				uri := "bolt://" + config.Database.Host + ":7687"
				auth := neo4j.BasicAuth(config.Database.Username, config.Keys.Database, "")
				// auth := neo4j.BasicAuth(config.Database.Username, config.Keys.AuraDatabase, "")

				driver, err := neo4j.NewDriver(uri, auth)
				if err != nil {
					panic(err)
				}

				defer driver.Close()

				session := driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
				defer session.Close()

				_, err = session.WriteTransaction(
					func(tx neo4j.Transaction) (interface{}, error) {
						// To learn more about the Cypher syntax, see https://neo4j.com/docs/cypher-manual/current/
						// The Reference Card is also a good resource for keywords https://neo4j.com/docs/cypher-refcard/current/
						createRelationshipBetweenPeopleQuery := `
						MERGE (house:Congress:Senate {number: $congressNumber })
						MERGE (member:Person:CongressMembers:Senators { id: $id })
						SET
						member.firstName = $firstName,
						member.lastName = $lastName,
						member.dob = $dob,
						member.gender = $gender,
						member.govtrackId = $govtrackId,
						member.cspanId = $cspanId,
						member.votesmartId = $votesmartId,
						member.icpsrId = $icpsrId,
						member.crpId = $crpId,
						member.googleEntityId = $googleEntityId,
						member.fecCandidateId = $fecCandidateId,
						member.rssUrl = $rssUrl,
						member.cookPvi = $cookPvi,
						member.lastUpdated = $lastUpdated,
						member.url = $url,
						member.ocdId = $ocdId,
						member.apiUri = $apiUri
						
						MERGE (member)-[r:MEMBER_OF]->(house)
						ON CREATE SET
						r.house = $house,
						r.shortTitle = $shortTitle,
						r.title = $title,
						r.atLarge = $atLarge,
						r.office = $office,
						r.twitter = $twitter,
						r.facebook = $facebook,
						r.party = $party,
						r.leadershipRole = $leadershipRole,
						r.senority = $senority,
						r.nextElection = $nextElection,
						r.totalVotes = $nextElection,
						r.missedVotes = $missedVotes,
						r.totalPresent = $totalPresent,
						r.phone = $phone,
						r.state = $state,
						r.geoId = $geoId,
						r.district = $district,
						r.missedVotesPct = $missedVotesPct,
						r.partyVotesPct = $partyVotesPct,
						r.nonPartyVotesPct = $nonPartyVotesPct
						RETURN house, member`
						result, err := tx.Run(createRelationshipBetweenPeopleQuery, map[string]interface{}{
							"house":            house,
							"congressNumber":   i,
							"id":               string(l.ID),
							"title":            string(l.Title),
							"shortTitle":       string(l.ShortTitle),
							"apiUri":           string(l.APIURI),
							"firstName":        string(l.FirstName),
							"lastName":         string(l.LastName),
							"dob":              string(l.DateOfBirth),
							"gender":           string(l.Gender),
							"party":            string(l.Party),
							"leadershipRole":   string(l.LeadershipRole),
							"twitter":          string(l.TwitterAccount),
							"facebook":         string(l.FacebookAccount),
							"govtrackId":       string(l.GovtrackID),
							"cspanId":          string(l.CspanID),
							"votesmartId":      string(l.VotesmartID),
							"icpsrId":          string(l.IcpsrID),
							"crpId":            string(l.CrpID),
							"googleEntityId":   string(l.GoogleEntityID),
							"fecCandidateId":   string(l.FecCandidateID),
							"url":              string(l.URL),
							"rssUrl":           string(l.RssURL),
							"cookPvi":          string(l.CookPvi),
							"senority":         string(l.Seniority),
							"nextElection":     string(l.NextElection),
							"totalVotes":       l.TotalVotes,
							"missedVotes":      l.MissedVotes,
							"totalPresent":     l.TotalPresent,
							"lastUpdated":      string(l.LastUpdated),
							"ocdId":            string(l.OcdID),
							"office":           string(l.Office),
							"phone":            string(l.Phone),
							"state":            string(l.State),
							"district":         string(l.District),
							"geoId":            string(l.Geoid),
							"atLarge":          l.AtLarge,
							"missedVotesPct":   l.MissedVotesPct,
							"partyVotesPct":    l.VotesWithPartyPct,
							"nonPartyVotesPct": l.VotesAgainstPartyPct,

							// "middleName": string(l.MiddleName),
							// "contactForm": string(l.ContactForm),
							// "suffix": string(l.Suffix),
							// "youtube": string(l.YoutubeAccount),
							// "inOffice": string(l.InOffice),
							// "fax": string(l.Fax),
							// "dwNominate": string(l.DwNominate),
							// "idealPoint": string(l.IdealPoint),
						})
						if err != nil {

							fmt.Println(err)
							return nil, err
						}

						return result.Collect()
					})
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
