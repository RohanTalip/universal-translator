package el

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} - {0}", Long: "{1} - {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "Ιούλ", 8: "Αύγ", 9: "Σεπ", 10: "Οκτ", 11: "Νοέ", 1: "Ιαν", 2: "Φεβ", 3: "Μάρ", 12: "Δεκ", 4: "Απρ", 5: "Μάι", 6: "Ιούν"}, Narrow: ut.CalendarMonthFormatNameValue{3: "Μ", 5: "Μ", 8: "Α", 9: "Σ", 11: "Ν", 1: "Ι", 4: "Α", 6: "Ι", 7: "Ι", 10: "Ο", 12: "Δ", 2: "Φ"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{7: "Ιούλιος", 8: "Αύγουστος", 9: "Σεπτέμβριος", 10: "Οκτώβριος", 1: "Ιανουάριος", 2: "Φεβρουάριος", 3: "Μάρτιος", 4: "Απρίλιος", 5: "Μάιος", 6: "Ιούνιος", 11: "Νοέμβριος", 12: "Δεκέμβριος"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Τρί", 3: "Τετ", 4: "Πέμ", 5: "Παρ", 6: "Σάβ", 0: "Κυρ", 1: "Δευ"}, Narrow: ut.CalendarDayFormatNameValue{4: "Π", 5: "Π", 6: "Σ", 0: "Κ", 1: "Δ", 2: "Τ", 3: "Τ"}, Short: ut.CalendarDayFormatNameValue{0: "Κυ", 1: "Δε", 2: "Τρ", 3: "Τε", 4: "Πέ", 5: "Πα", 6: "Σά"}, Wide: ut.CalendarDayFormatNameValue{6: "Σάββατο", 0: "Κυριακή", 1: "Δευτέρα", 2: "Τρίτη", 3: "Τετάρτη", 4: "Πέμπτη", 5: "Παρασκευή"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "μεσημ.", "evening1": "απόγ.", "night1": "βράδυ", "am": "π.μ.", "pm": "μ.μ.", "morning1": "πρωί"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "μ.μ.", "morning1": "πρωί", "afternoon1": "μεσημέρι", "evening1": "απόγευμα", "night1": "βράδυ", "am": "π.μ."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "βράδυ", "am": "π.μ.", "pm": "μ.μ.", "morning1": "πρωί", "afternoon1": "μεσημέρι", "evening1": "απόγευμα"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "μετά Χριστόν", Abbrev: "μ.Χ.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "προ Χριστού", Abbrev: "π.Χ.", Narrow: ""}}}}
