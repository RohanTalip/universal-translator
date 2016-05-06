package fr_CH

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH.mm:ss 'h' zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "avr.", 7: "juil.", 8: "août", 3: "mars", 2: "févr.", 5: "mai", 9: "sept.", 11: "nov.", 6: "juin", 10: "oct.", 12: "déc.", 1: "janv."}, Narrow: ut.CalendarMonthFormatNameValue{6: "J", 11: "N", 1: "J", 3: "M", 9: "S", 12: "D", 2: "F", 7: "J", 5: "M", 8: "A", 10: "O", 4: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mai", 10: "octobre", 11: "novembre", 4: "avril", 6: "juin", 8: "août", 7: "juillet", 12: "décembre", 3: "mars", 2: "février", 9: "septembre", 1: "janvier"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim.", 1: "lun."}, Narrow: ut.CalendarDayFormatNameValue{2: "M", 3: "M", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L"}, Short: ut.CalendarDayFormatNameValue{5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me", 4: "je"}, Wide: ut.CalendarDayFormatNameValue{5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "matin", "afternoon1": "après-midi"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
