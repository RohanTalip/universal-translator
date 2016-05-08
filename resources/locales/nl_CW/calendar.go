package nl_CW

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "apr.", 6: "jun.", 3: "mrt.", 5: "mei", 7: "jul.", 10: "okt.", 12: "dec.", 9: "sep.", 1: "jan.", 2: "feb.", 8: "aug.", 11: "nov."}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 1: "J", 2: "F", 3: "M", 11: "N", 5: "M", 6: "J", 8: "A", 12: "D", 4: "A", 7: "J", 9: "S"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "februari", 4: "april", 5: "mei", 7: "juli", 8: "augustus", 11: "november", 1: "januari", 6: "juni", 10: "oktober", 9: "september", 12: "december", 3: "maart"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "zo", 1: "ma", 2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za"}, Narrow: ut.CalendarDayFormatNameValue{1: "M", 2: "D", 3: "W", 4: "D", 5: "V", 6: "Z", 0: "Z"}, Short: ut.CalendarDayFormatNameValue{2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za", 0: "zo", 1: "ma"}, Wide: ut.CalendarDayFormatNameValue{4: "donderdag", 5: "vrijdag", 6: "zaterdag", 0: "zondag", 1: "maandag", 2: "dinsdag", 3: "woensdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
