package services

/// Main logic behind simulating match
/// from fixture pattern decide match to played
/// teams score will be determined randomly (rand interval based on their effective_power,stronger teams have more chance to have higher scores)
/// Write it down to DB
import (
	"fmt"
	"league-simulator/config"
	"league-simulator/models"
	"math/rand"
	"time"
)

func SimulateCurrentWeek() (string, error) {
	db := config.DB

	// Check Season DATA
	var state models.SeasonState
	if err := db.First(&state).Error; err != nil {
		state.CurrentWeek = 1
		if err := db.Create(&state).Error; err != nil {
			return "", err
		}
	}

	if state.CurrentWeek > 38 {
		return "League finished. Please reset.", nil
	}

	var teams []models.Team
	if err := db.Find(&teams).Error; err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())

	matchups := generateWeeklyFixtures(state.CurrentWeek, teams)

	for _, m := range matchups {
		home := m[0]
		away := m[1]

		homeGoals := rand.Intn(int(home.EffectivePower()/10) + 1)
		awayGoals := rand.Intn(int(away.EffectivePower()/10) + 1)

		match := models.Match{
			HomeTeamID: home.ID,
			AwayTeamID: away.ID,
			HomeScore:  homeGoals,
			AwayScore:  awayGoals,
			Week:       state.CurrentWeek,
		}
		if err := db.Create(&match).Error; err != nil {
			return "", err
		}

		updateMorale(&home, &away, homeGoals, awayGoals)
		db.Save(&home)
		db.Save(&away)
	}

	// Procced week
	state.CurrentWeek++
	if err := db.Save(&state).Error; err != nil {
		return "", err
	}

	return fmt.Sprintf("Week %d simulated.", state.CurrentWeek-1), nil
}

func generateWeeklyFixtures(week int, teams []models.Team) [][2]models.Team {
	var matchups [][2]models.Team

	if len(teams) != 4 || week > 38 {
		return matchups
	}

	// Fixture pattern to make sure all team have played 38 matches at end of the season
	pairsByWeek := map[int][][2]int{
		1: {{0, 1}, {2, 3}},
		2: {{0, 2}, {1, 3}},
		3: {{0, 3}, {1, 2}},
		4: {{1, 0}, {3, 2}},
		5: {{2, 0}, {3, 1}},
		6: {{3, 0}, {2, 1}},
	}

	// For repeat above structure
	weekKey := ((week - 1) % 6) + 1

	for _, pair := range pairsByWeek[weekKey] {
		matchups = append(matchups, [2]models.Team{
			teams[pair[0]],
			teams[pair[1]],
		})
	}

	return matchups
}

func updateMorale(home *models.Team, away *models.Team, homeGoals, awayGoals int) {
	if homeGoals > awayGoals {
		home.Morale += rand.Intn(6) + 5 // +5 to +10
		away.Morale -= rand.Intn(6) + 5 // -5 to -10
	} else if homeGoals < awayGoals {
		home.Morale -= rand.Intn(6) + 5
		away.Morale += rand.Intn(6) + 5
	} else {
		home.Morale += rand.Intn(5) - 2 // -2 to +2
		away.Morale += rand.Intn(5) - 2
	}

	// Minimum morale is 25 because if it can be 0 some teams might lose their
	// effective_power at all
	home.Morale = clamp(home.Morale, 25, 100)
	away.Morale = clamp(away.Morale, 25, 100)
}

func clamp(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

func SimulateAllRemainingWeeks() (string, error) {
	db := config.DB

	var state models.SeasonState
	if err := db.First(&state).Error; err != nil {
		state.CurrentWeek = 1
		if err := db.Create(&state).Error; err != nil {
			return "", err
		}
	}

	if state.CurrentWeek > 38 {
		return "League finished. Please reset.", nil
	}

	var teams []models.Team
	if err := db.Find(&teams).Error; err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())

	for week := state.CurrentWeek; week <= 38; week++ {
		matchups := generateWeeklyFixtures(week, teams)

		for _, m := range matchups {
			home := m[0]
			away := m[1]

			homeGoals := rand.Intn(int(home.EffectivePower()/10) + 1)
			awayGoals := rand.Intn(int(away.EffectivePower()/10) + 1)

			match := models.Match{
				HomeTeamID: home.ID,
				AwayTeamID: away.ID,
				HomeScore:  homeGoals,
				AwayScore:  awayGoals,
				Week:       week,
			}
			if err := db.Create(&match).Error; err != nil {
				return "", err
			}

			updateMorale(&home, &away, homeGoals, awayGoals)
			db.Save(&home)
			db.Save(&away)
		}
	}

	state.CurrentWeek = 39
	if err := db.Save(&state).Error; err != nil {
		return "", err
	}

	return "Season fully simulated from current week to week 38.", nil
}
