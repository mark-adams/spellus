package main

type entry struct {
	ID              string
	CorrectSpelling string
}

var dictionary = map[string]entry{
	"colour":      entry{"001", "color"},
	"flavour":     entry{"002", "flavor"},
	"humour":      entry{"003", "humor"},
	"labour":      entry{"004", "labor"},
	"neighbour":   entry{"005", "neighbor"},
	"apologise":   entry{"006", "apologize"},
	"organise":    entry{"007", "organize"},
	"recognise":   entry{"008", "recognize"},
	"analyse":     entry{"009", "analyze"},
	"breathalyse": entry{"010", "breathalyze"},
	"paralyse":    entry{"011", "paralyze"},
	"travelled":   entry{"012", "traveled"},
	"travelling":  entry{"013", "traveling"},
	"traveller":   entry{"014", "traveler"},
	"leukaemia":   entry{"015", "leukemia"},
	"manoeuvre":   entry{"016", "maneuver"},
	"oestrogen":   entry{"017", "estrogen"},
	"paediatric":  entry{"018", "pediatric"},
	"defence":     entry{"019", "defense"},
	"licence":     entry{"020", "license"},
	"offence":     entry{"021", "offense"},
	"pretence":    entry{"022", "pretense"},
	"analogue":    entry{"023", "analog"},
	"catalogue":   entry{"024", "catalog"},
	"dialogue":    entry{"025", "dialog"},
}
