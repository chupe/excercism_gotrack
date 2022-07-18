package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Outcome string

const (
	win  Outcome = "win"
	draw         = "draw"
	loss         = "loss"
)

func ParseOutcome(s string) (Outcome, error) {
	switch s {
	case "win":
		return win, nil
	case "draw":
		return draw, nil
	case "loss":
		return loss, nil
	default:
		return "", errors.New("unable to parse " + s)
	}
}

type TableInput struct {
	team1   string
	team2   string
	outcome Outcome
}
type TableRow struct {
	Team   string
	Wins   int
	Draws  int
	Losses int
}

func (tr *TableRow) Points() int {
	return tr.Wins*3 + tr.Draws
}
func (tr *TableRow) MatchesPlayed() int {
	return tr.Wins + tr.Draws + tr.Losses
}
func (tr *TableRow) IncrementOutcomeForT1(o Outcome) {
	switch o {
	case win:
		tr.Wins++
	case draw:
		tr.Draws++
	case loss:
		tr.Losses++
	}
}
func (tr *TableRow) IncrementOutcomeForT2(o Outcome) {
	switch o {
	case win:
		tr.Losses++
	case draw:
		tr.Draws++
	case loss:
		tr.Wins++
	}
}
func (t Table) Add(ti *TableInput) {
	if tr, ok := t[ti.team1]; ok {
		tr.IncrementOutcomeForT1(ti.outcome)
	} else {
		ntr := &TableRow{
			Team: ti.team1,
		}
		ntr.IncrementOutcomeForT1(ti.outcome)
		t[ti.team1] = ntr
	}
	if tr, ok := t[ti.team2]; ok {
		tr.IncrementOutcomeForT2(ti.outcome)
	} else {
		ntr := &TableRow{
			Team: ti.team2,
		}
		ntr.IncrementOutcomeForT2(ti.outcome)
		t[ti.team2] = ntr
	}
}

type Table map[string]*TableRow

func (t *Table) GetHeader() []byte {
	header := fmt.Sprintf("Team                           | MP |  W |  D |  L |  P\n")
	return []byte(header)
}
func (t Table) ToBytes() []byte {
	result := t.GetHeader()
	sorted := make([]TableRow, len(t))
	i := 0
	for _, row := range t {
		sorted[i] = *row
		i++
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Points() == sorted[j].Points() {
			return sorted[i].Team < sorted[j].Team
		}
		return sorted[i].Points() > sorted[j].Points()
	})
	for i, row := range sorted {
		team := sorted[i].Team
		line := fmt.Sprintf("%-31s| % d | % d | % d | % d | % d\n", team, row.MatchesPlayed(), row.Wins, row.Draws, row.Losses, row.Points())
		result = append(result, []byte(line)...)
	}
	return result
}
func Tally(reader io.Reader, writer io.Writer) error {
	table := make(Table)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ";")
		if tokens[0] == "" ||
			strings.HasPrefix(tokens[0], "#") {
			continue
		}
		if len(tokens) != 3 {
			return errors.New("token len != 3")
		}
		outcome, err := ParseOutcome(tokens[2])
		if err != nil {
			return err
		}
		ti1 := &TableInput{
			team1:   tokens[0],
			team2:   tokens[1],
			outcome: outcome,
		}
		table.Add(ti1)
	}
	_, err := writer.Write(table.ToBytes())
	if err != nil {
		return err
	}
	return nil
}
