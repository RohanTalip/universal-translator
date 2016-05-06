package it_CH

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "dic", 5: "mag", 7: "lug", 10: "ott", 8: "ago", 9: "set", 1: "gen", 3: "mar", 4: "apr", 11: "nov", 2: "feb", 6: "giu"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 1: "G", 9: "S", 7: "L", 8: "A", 11: "N", 2: "F", 3: "M", 4: "A", 5: "M", 6: "G", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "novembre", 12: "dicembre", 8: "agosto", 10: "ottobre", 3: "marzo", 4: "aprile", 5: "maggio", 6: "giugno", 7: "luglio", 9: "settembre", 1: "gennaio", 2: "febbraio"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "dom", 1: "lun", 2: "mar", 3: "mer", 4: "gio", 5: "ven", 6: "sab"}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "G"}, Short: ut.CalendarDayFormatNameValue{6: "sab", 0: "dom", 1: "lun", 2: "mar", 3: "mer", 4: "gio", 5: "ven"}, Wide: ut.CalendarDayFormatNameValue{0: "Domenica", 1: "Lunedì", 2: "Martedì", 3: "Mercoledì", 4: "Giovedì", 5: "Venerdì", 6: "Sabato"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "mezzogiorno", "pm": "PM", "morning1": "mattina", "afternoon1": "pomeriggio", "evening1": "sera", "night1": "notte", "midnight": "mezzanotte"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "mattina", "afternoon1": "pomeriggio", "evening1": "sera", "night1": "notte", "midnight": "mezzanotte", "am": "m.", "noon": "mezzogiorno", "pm": "p."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "mattina", "afternoon1": "pomeriggio", "evening1": "sera", "night1": "notte", "midnight": "mezzanotte", "am": "AM", "noon": "mezzogiorno"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
