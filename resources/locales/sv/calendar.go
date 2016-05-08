package sv

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y-MM-dd"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y-MM-dd"}}, Time: ut.CalendarDateFormat{Full: "'kl'. HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "feb.", 5: "maj", 6: "juni", 9: "sep.", 8: "aug.", 10: "okt.", 11: "nov.", 12: "dec.", 1: "jan.", 3: "mars", 4: "apr.", 7: "juli"}, Narrow: ut.CalendarMonthFormatNameValue{1: "J", 3: "M", 6: "J", 8: "A", 11: "N", 2: "F", 4: "A", 5: "M", 7: "J", 9: "S", 10: "O", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "november", 1: "januari", 2: "februari", 4: "april", 5: "maj", 7: "juli", 8: "augusti", 10: "oktober", 12: "december", 3: "mars", 6: "juni", 9: "september"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "ons", 4: "tors", 5: "fre", 6: "lör", 0: "sön", 1: "mån", 2: "tis"}, Narrow: ut.CalendarDayFormatNameValue{6: "L", 0: "S", 1: "M", 2: "T", 3: "O", 4: "T", 5: "F"}, Short: ut.CalendarDayFormatNameValue{3: "on", 4: "to", 5: "fr", 6: "lö", 0: "sö", 1: "må", 2: "ti"}, Wide: ut.CalendarDayFormatNameValue{1: "måndag", 2: "tisdag", 3: "onsdag", 4: "torsdag", 5: "fredag", 6: "lördag", 0: "söndag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "efterm.", "evening1": "kväll", "night1": "natt", "midnight": "midnatt", "am": "f.m.", "pm": "e.m.", "morning1": "morgon", "morning2": "förm."}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "em", "morning1": "morg.", "morning2": "förm.", "afternoon1": "efterm.", "evening1": "kväll", "night1": "natt", "midnight": "midn.", "am": "fm"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"morning1": "morgon", "morning2": "förmiddag", "afternoon1": "eftermiddag", "evening1": "kväll", "night1": "natt", "midnight": "midnatt", "am": "förmiddag", "pm": "eftermiddag"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "efter Kristus", Abbrev: "e.Kr.", Narrow: "e.Kr."}, BC: ut.CalendarEraFormatNames{Full: "före Kristus", Abbrev: "f.Kr.", Narrow: "f.Kr."}}}}
