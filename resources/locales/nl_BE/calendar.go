package nl_BE

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/MM/yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "okt.", 6: "jun.", 12: "dec.", 1: "jan.", 2: "feb.", 3: "mrt.", 5: "mei", 7: "jul.", 8: "aug.", 11: "nov.", 4: "apr.", 9: "sep."}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 6: "J", 3: "M", 8: "A", 9: "S", 10: "O", 5: "M", 11: "N", 12: "D", 1: "J", 2: "F", 7: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "november", 4: "april", 7: "juli", 12: "december", 9: "september", 2: "februari", 3: "maart", 5: "mei", 6: "juni", 8: "augustus", 1: "januari", 10: "oktober"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "za", 0: "zo", 1: "ma", 2: "di", 3: "wo", 4: "do", 5: "vr"}, Narrow: ut.CalendarDayFormatNameValue{4: "D", 5: "V", 6: "Z", 0: "Z", 1: "M", 2: "D", 3: "W"}, Short: ut.CalendarDayFormatNameValue{1: "ma", 2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za", 0: "zo"}, Wide: ut.CalendarDayFormatNameValue{5: "vrijdag", 6: "zaterdag", 0: "zondag", 1: "maandag", 2: "dinsdag", 3: "woensdag", 4: "donderdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
