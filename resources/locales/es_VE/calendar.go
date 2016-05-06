package es_VE

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "jun.", 7: "jul.", 8: "ago.", 9: "sept.", 12: "dic.", 1: "ene.", 2: "feb.", 3: "mar.", 4: "abr.", 5: "may.", 10: "oct.", 11: "nov."}, Narrow: ut.CalendarMonthFormatNameValue{8: "A", 9: "S", 4: "A", 2: "F", 3: "M", 5: "M", 6: "J", 7: "J", 10: "O", 11: "N", 1: "E", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{9: "septiembre", 11: "noviembre", 12: "diciembre", 4: "abril", 8: "agosto", 5: "mayo", 6: "junio", 7: "julio", 1: "enero", 2: "febrero", 3: "marzo", 10: "octubre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar."}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "X", 4: "J"}, Short: ut.CalendarDayFormatNameValue{1: "Lu", 2: "Ma", 3: "Mi", 4: "Ju", 5: "Vi", 6: "Sa", 0: "Do"}, Wide: ut.CalendarDayFormatNameValue{5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles", 4: "jueves"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m."}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "noon": "m.", "evening1": "tarde", "night1": "noche", "am": "a. m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "a. m.", "pm": "p. m.", "evening1": "tarde", "night1": "noche", "noon": "mediodía", "morning1": "madrugada", "morning2": "mañana"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
