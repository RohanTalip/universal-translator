package de_LU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'um' {0}", Long: "{1} 'um' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 5: "Mai", 11: "Nov", 4: "Apr", 10: "Okt", 9: "Sep", 2: "Feb", 6: "Jun", 8: "Aug", 12: "Dez", 3: "Mär", 7: "Jul"}, Narrow: ut.CalendarMonthFormatNameValue{3: "M", 12: "D", 7: "J", 8: "A", 9: "S", 2: "F", 4: "A", 6: "J", 10: "O", 1: "J", 5: "M", 11: "N"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{8: "August", 11: "November", 1: "Januar", 3: "März", 10: "Oktober", 2: "Februar", 5: "Mai", 6: "Juni", 9: "September", 4: "April", 7: "Juli", 12: "Dezember"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mo", 2: "Di", 3: "Mi", 4: "Do", 5: "Fr", 6: "Sa", 0: "So"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "M", 2: "D", 3: "M", 4: "D", 5: "F", 6: "S"}, Short: ut.CalendarDayFormatNameValue{1: "Mo.", 2: "Di.", 3: "Mi.", 4: "Do.", 5: "Fr.", 6: "Sa.", 0: "So."}, Wide: ut.CalendarDayFormatNameValue{5: "Freitag", 6: "Samstag", 0: "Sonntag", 1: "Montag", 2: "Dienstag", 3: "Mittwoch", 4: "Donnerstag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "Abend", "midnight": "Mitternacht", "pm": "nachm.", "morning1": "Morgen", "afternoon1": "Mittag", "afternoon2": "Nachmittag", "night1": "Nacht", "am": "vorm.", "morning2": "Vormittag"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning2": "Vormittag", "am": "vorm.", "evening1": "Abend", "afternoon1": "Mittag", "midnight": "Mitternacht", "afternoon2": "Nachmittag", "night1": "Nacht", "pm": "nachm.", "morning1": "Morgen"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "Mittag", "evening1": "Abend", "pm": "nachm.", "am": "vorm.", "afternoon2": "Nachmittag", "night1": "Nacht", "midnight": "Mitternacht", "morning1": "Morgen", "morning2": "Vormittag"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
