package os_RU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y 'аз'", Long: "d MMMM, y 'аз'", Medium: "dd MMM y 'аз'", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y 'аз'", Long: "d MMMM, y 'аз'", Medium: "dd MMM y 'аз'", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "Май", 11: "Нояб.", 12: "Дек.", 6: "Июнь", 3: "Март.", 4: "Апр.", 7: "Июль", 8: "Авг.", 9: "Сент.", 10: "Окт.", 2: "Февр.", 1: "Янв."}, Narrow: ut.CalendarMonthFormatNameValue{6: "И", 9: "С", 7: "И", 11: "Н", 1: "Я", 4: "А", 2: "Ф", 12: "Д", 3: "М", 5: "М", 8: "А", 10: "О"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "Ноябрь", 1: "Январь", 5: "Май", 9: "Сентябрь", 4: "Апрель", 8: "Август", 12: "Декабрь", 2: "Февраль", 6: "Июнь", 10: "Октябрь", 3: "Мартъи", 7: "Июль"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Хцб", 1: "Крс", 2: "Дцг", 3: "Ӕрт", 4: "Цпр", 5: "Мрб", 6: "Сбт"}, Narrow: ut.CalendarDayFormatNameValue{1: "К", 2: "Д", 3: "Ӕ", 4: "Ц", 5: "М", 6: "С", 0: "Х"}, Short: ut.CalendarDayFormatNameValue{}, Wide: ut.CalendarDayFormatNameValue{3: "Ӕртыццӕг", 4: "Цыппӕрӕм", 5: "Майрӕмбон", 6: "Сабат", 0: "Хуыцаубон", 1: "Къуырисӕр", 2: "Дыццӕг"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "ӕмбисбоны размӕ", "pm": "ӕмбисбоны фӕстӕ"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
