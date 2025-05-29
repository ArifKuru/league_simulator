// services/predictor.go
// / This service defines the core prediction logic used to estimate league outcomes based on simulations.
// / I wanted to use only Strength of each team and not their morale, in order to make prediction realistic
// / because we cannot know the morale of a team in real life , on the other hand strength can be approximately estimated
// / Both `DefaultPredictor` and `MonteCarloPredictor` simulate the remainder of the football season multiple times
// / The main difference lies in how they simulate the match outcomes:
// / - `DefaultPredictor` uses a simple random goal generation based on team strength (Strength/10),
// /   resulting in more deterministic and evenly distributed scores.
// / - `MonteCarloPredictor` introduces randomness through a multiplier (0.5 to 1.5) applied to each team’s strength,
// /   creating a more dynamic and realistic simulation by simulating performance variability per match.
// / This makes `MonteCarloPredictor` more reflective of real-life uncertainty and fluctuations in team performance.
package services

import (
	"league-simulator/config"
	"league-simulator/models"
	"math/rand"
	"sort"
	"time"
)

type Predictor interface {
	Predict(simCount int) ([]PredictionResult, error)
}

type PredictionResult struct {
	Team   string `json:"team"`
	Chance int    `json:"chance"`
}

// Default Predictor

type DefaultPredictor struct{}

func (p DefaultPredictor) Predict(simCount int) ([]PredictionResult, error) {
	db := config.DB

	var teams []models.Team
	if err := db.Find(&teams).Error; err != nil {
		return nil, err
	}

	var state models.SeasonState
	if err := db.First(&state).Error; err != nil {
		state.CurrentWeek = 1
	}

	var matches []models.Match
	if err := db.Find(&matches).Error; err != nil {
		return nil, err
	}

	counts := map[string]int{}
	rand.Seed(time.Now().UnixNano())

	for s := 0; s < simCount; s++ {
		standings := map[uint]*models.Standing{}
		for _, team := range teams {
			standings[team.ID] = &models.Standing{
				TeamID:   team.ID,
				TeamName: team.Name,
			}
		}

		for _, m := range matches {
			applyMatchToStanding(m.HomeTeamID, m.AwayTeamID, m.HomeScore, m.AwayScore, standings)
		}

		for week := state.CurrentWeek; week <= 38; week++ {
			fixtures := generateWeeklyFixtures(week, teams)
			for _, pair := range fixtures {
				t1 := pair[0]
				t2 := pair[1]

				goals1 := rand.Intn(t1.Strength/10 + 1)
				goals2 := rand.Intn(t2.Strength/10 + 1)

				applyMatchToStanding(t1.ID, t2.ID, goals1, goals2, standings)
			}
		}

		var result []models.Standing
		for _, s := range standings {
			s.GoalDifference = s.GoalsFor - s.GoalsAgainst
			result = append(result, *s)
		}

		sort.Slice(result, func(i, j int) bool {
			if result[i].Points == result[j].Points {
				return result[i].GoalDifference > result[j].GoalDifference
			}
			return result[i].Points > result[j].Points
		})

		counts[result[0].TeamName]++
	}

	var output []PredictionResult
	for _, team := range teams {
		count := counts[team.Name]
		percent := int(float64(count) / float64(simCount) * 100)
		output = append(output, PredictionResult{
			Team:   team.Name,
			Chance: percent,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Chance > output[j].Chance
	})

	return output, nil
}

// Monte Carlo Predictor

type MonteCarloPredictor struct{}

func (m MonteCarloPredictor) Predict(simCount int) ([]PredictionResult, error) {
	db := config.DB

	var teams []models.Team
	if err := db.Find(&teams).Error; err != nil {
		return nil, err
	}

	var state models.SeasonState
	if err := db.First(&state).Error; err != nil {
		state.CurrentWeek = 1
	}

	var matches []models.Match
	if err := db.Find(&matches).Error; err != nil {
		return nil, err
	}

	counts := map[string]int{}
	rand.Seed(time.Now().UnixNano())

	for s := 0; s < simCount; s++ {
		standings := map[uint]*models.Standing{}
		for _, team := range teams {
			standings[team.ID] = &models.Standing{
				TeamID:   team.ID,
				TeamName: team.Name,
			}
		}

		for _, m := range matches {
			applyMatchToStanding(m.HomeTeamID, m.AwayTeamID, m.HomeScore, m.AwayScore, standings)
		}

		for week := state.CurrentWeek; week <= 38; week++ {
			fixtures := generateWeeklyFixtures(week, teams)
			for _, pair := range fixtures {
				t1 := pair[0]
				t2 := pair[1]

				eff1 := float64(t1.Strength) * (rand.Float64() + 0.5)
				eff2 := float64(t2.Strength) * (rand.Float64() + 0.5)

				goals1 := int(eff1 / 20)
				goals2 := int(eff2 / 20)

				applyMatchToStanding(t1.ID, t2.ID, goals1, goals2, standings)
			}
		}

		var result []models.Standing
		for _, s := range standings {
			s.GoalDifference = s.GoalsFor - s.GoalsAgainst
			result = append(result, *s)
		}

		sort.Slice(result, func(i, j int) bool {
			if result[i].Points == result[j].Points {
				return result[i].GoalDifference > result[j].GoalDifference
			}
			return result[i].Points > result[j].Points
		})

		counts[result[0].TeamName]++
	}

	var output []PredictionResult
	for _, team := range teams {
		count := counts[team.Name]
		percent := int(float64(count) / float64(simCount) * 100)
		output = append(output, PredictionResult{
			Team:   team.Name,
			Chance: percent,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Chance > output[j].Chance
	})

	return output, nil
}
func applyMatchToStanding(homeID, awayID uint, homeGoals, awayGoals int, standings map[uint]*models.Standing) {
	home := standings[homeID]
	away := standings[awayID]

	home.Played++
	away.Played++

	home.GoalsFor += homeGoals
	home.GoalsAgainst += awayGoals

	away.GoalsFor += awayGoals
	away.GoalsAgainst += homeGoals

	if homeGoals > awayGoals {
		home.Won++
		home.Points += 3
		away.Lost++
	} else if homeGoals < awayGoals {
		away.Won++
		away.Points += 3
		home.Lost++
	} else {
		home.Drawn++
		away.Drawn++
		home.Points++
		away.Points++
	}
}
