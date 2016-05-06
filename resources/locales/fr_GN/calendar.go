package fr_GN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "mars", 4: "avr.", 5: "mai", 9: "sept.", 11: "nov.", 2: "févr.", 7: "juil.", 10: "oct.", 12: "déc.", 1: "janv.", 6: "juin", 8: "août"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 10: "O", 12: "D", 6: "J", 1: "J", 3: "M", 4: "A", 9: "S", 2: "F", 5: "M", 8: "A", 11: "N"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "février", 8: "août", 7: "juillet", 10: "octobre", 1: "janvier", 4: "avril", 6: "juin", 9: "septembre", 3: "mars", 5: "mai", 11: "novembre", 12: "décembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim."}, Narrow: ut.CalendarDayFormatNameValue{3: "M", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L", 2: "M"}, Short: ut.CalendarDayFormatNameValue{6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me", 4: "je", 5: "ve"}, Wide: ut.CalendarDayFormatNameValue{4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
