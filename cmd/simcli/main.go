package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"aethersim/sim"
)

type Stats struct {
	Battles     int
	Wins        int
	DamageDealt float64
	DamageTaken float64
	HealingDone float64
	DoTDealt    float64
}

func main() {
	completed := 0
	start := time.Now()
	totalRounds := 0
	totalTurns := 0

	testCalc := 30 * float64(27+5) / math.Pow(float64(17+13), 0.9) * (float64((6*2)/4+5) / 30)
	fmt.Printf("Test Calculation: %10.3f\n", testCalc)

	// 1) Gather your full roster of character‐keys:
	roster := sim.AllCharacterKeys() // returns []string, e.g. ["chocolate_chip", "flitterfyre", …]

	// 2) Build all unique 3‐member teams (combinations without repetition)
	teams := generateUniqueTeams(roster, TEAM_SIZE)

	matchupCount := len(teams) * (len(teams) + 1) / 2
	totalBattles := matchupCount * TRIALS
	fmt.Printf(
		"Simulating %d unique team matchups × %d trials = %d total battles…\n\n",
		matchupCount, TRIALS, totalBattles,
	)

	// 3) Prepare stats map: key → { battles, wins }
	stats := make(map[string]*Stats, len(roster))
	for _, id := range roster {
		stats[id] = &Stats{}
	}

	// 4) Simulate each distinct pairing i ≤ j
	for i := 0; i < len(teams); i++ {
		aKeys := teams[i]

		for j := i; j < len(teams); j++ {
			bKeys := teams[j]

			// simulate TRIALS battles for this matchup
			for t := 0; t < TRIALS; t++ {
				// rebuild brand‐new Character structs each time
				aTeam := sim.MakeTeam(aKeys, LEVEL, true)
				bTeam := sim.MakeTeam(bKeys, LEVEL, false)

				engine := sim.NewEngine(aTeam, bTeam)

				for !engine.GameOver {
					engine.Step(sim.UtilityDecision)

					// record per-target delta stats for this round
					for _, imp := range engine.LastImpacts {
						rec := stats[imp.ActorID]
						if imp.Delta > 0 {
							rec.DamageDealt += imp.Delta
							stats[imp.TargetID].DamageTaken += imp.Delta
							if imp.IsDebuff {
								rec.DoTDealt += imp.Delta
							}
						} else if imp.Delta < 0 {
							// negative delta means healing
							rec.HealingDone += -imp.Delta
							// healed HP is not “damage taken”
						}
					}
				}

				// tally results
				alliesWin := engine.PlayerWon
				// for each char in A-team
				for _, c := range aTeam {
					stats[c.ID].Battles++
					if alliesWin {
						stats[c.ID].Wins++
					}
				}
				// for each char in B-team
				for _, c := range bTeam {
					stats[c.ID].Battles++
					if !alliesWin {
						stats[c.ID].Wins++
					}
				}

				totalRounds += engine.TotalRounds
				totalTurns += engine.TotalTurns

				completed++
				if completed%BATCH_SIZE == 0 {
					drawProgress(completed, totalBattles, start)
				}
			}
		}
	}

	drawProgress(completed, totalBattles, start)

	// 5) Print out aggregate win‐rates plus per‐game averages
	fmt.Print("\n\n")
	fmt.Println("Character              | Win%     | Tot Dmg▶  | Tot Dmg◀  | Tot Heal  | Tot DoT ")
	fmt.Println("-----------------------+----------+-----------+-----------+-----------+----------")
	for _, k := range roster {
		s := stats[k]
		if s.Battles == 0 {
			fmt.Printf("%-22s |    N/A   |    N/A    |    N/A    |   N/A    |   N/A\n", k)
			continue
		}

		winPct := 100 * float64(s.Wins) / float64(s.Battles)
		avgDealt := s.DamageDealt / float64(s.Battles)
		avgTaken := s.DamageTaken / float64(s.Battles)
		avgHeal := s.HealingDone / float64(s.Battles)
		avgDoT := s.DoTDealt / float64(s.Battles)

		fmt.Printf("%-22s | %6.2f%%  | %9.1f | %9.1f | %9.1f | %9.1f\n",
			k,
			winPct,
			avgDealt,
			avgTaken,
			avgHeal,
			avgDoT,
		)
	}
	avgRoundsPerGame := float64(totalRounds) / float64(totalBattles)
	avgTurnsPerGame := float64(totalTurns) / float64(totalBattles)
	fmt.Print("\n")
	fmt.Printf("Average Rounds per Game: %12.5f\n", avgRoundsPerGame)
	fmt.Printf("Average Turns per Game: %12.5f\n", avgTurnsPerGame)
}

// generateUniqueTeams returns all k‐sized combinations of keys, without repetition.
func generateUniqueTeams(keys []string, k int) [][]string {
	var res [][]string
	var helper func(start int, combo []string)
	helper = func(start int, combo []string) {
		if len(combo) == k {
			// copy
			c := make([]string, k)
			copy(c, combo)
			res = append(res, c)
			return
		}
		for i := start; i < len(keys); i++ {
			helper(i+1, append(combo, keys[i]))
		}
	}
	helper(0, []string{})
	return res
}

// drawProgress overwrites the current line with a bar, pct, counts, rate, elapsed & ETA.
func drawProgress(completed, total int, start time.Time) {
	now := time.Now()
	elapsed := now.Sub(start)
	secs := elapsed.Seconds()

	pct := float64(completed) / float64(total)
	rate := float64(completed) / secs // sims per second
	remaining := total - completed
	eta := time.Duration(float64(time.Second) * float64(remaining) / rate)

	// build a 30-char bar
	width := 30
	filled := int(pct * float64(width))
	if filled > width {
		filled = width
	}
	bar := strings.Repeat("█", filled) + strings.Repeat(" ", width-filled)

	fmt.Printf(
		"\r[%s] %6.2f%%  %6d/%6d (%5.1f sim/s) elapsed %s ETA %s",
		bar,
		pct*100,
		completed,
		total,
		rate,
		formatDuration(elapsed),
		formatDuration(eta),
	)
	if completed == total {
		fmt.Print("\n")
	}
}

// formatDuration prints h m s compactly
func formatDuration(d time.Duration) string {
	secTotal := int(d.Seconds())
	h := secTotal / 3600
	m := (secTotal % 3600) / 60
	s := secTotal % 60
	return fmt.Sprintf("%dh%02dm%02ds", h, m, s)
}
