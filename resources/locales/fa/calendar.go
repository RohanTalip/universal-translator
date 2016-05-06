package fa

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "y/M/d"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}، ساعت {0}", Long: "{1}، ساعت {0}", Medium: "{1}،\u200f {0}", Short: "{1}،\u200f {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "نوامبر", 12: "دسامبر", 5: "مه", 7: "ژوئیه", 8: "اوت", 4: "آوریل", 6: "ژوئن", 9: "سپتامبر", 10: "اکتبر", 1: "ژانویه", 2: "فوریه", 3: "مارس"}, Narrow: ut.CalendarMonthFormatNameValue{2: "ف", 4: "آ", 8: "ا", 11: "ن", 1: "ژ", 3: "م", 5: "م", 6: "ژ", 7: "ژ", 9: "س", 10: "ا", 12: "د"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "فوریه", 3: "مارس", 8: "اوت", 11: "نوامبر", 1: "ژانویه", 5: "مه", 6: "ژوئن", 7: "ژوئیه", 9: "سپتامبر", 10: "اکتبر", 12: "دسامبر", 4: "آوریل"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "دوشنبه", 2: "سه\u200cشنبه", 3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه", 6: "شنبه", 0: "یکشنبه"}, Narrow: ut.CalendarDayFormatNameValue{2: "س", 3: "چ", 4: "پ", 5: "ج", 6: "ش", 0: "ی", 1: "د"}, Short: ut.CalendarDayFormatNameValue{4: "۵ش", 5: "ج", 6: "ش", 0: "۱ش", 1: "۲ش", 2: "۳ش", 3: "۴ش"}, Wide: ut.CalendarDayFormatNameValue{0: "یکشنبه", 1: "دوشنبه", 2: "سه\u200cشنبه", 3: "چهارشنبه", 4: "پنجشنبه", 5: "جمعه", 6: "شنبه"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "ظ", "pm": "ب.ظ.", "morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ."}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب", "midnight": "نیمه\u200cشب", "am": "ق.ظ.", "noon": "ظهر", "pm": "ب.ظ."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "نیمه\u200cشب", "am": "قبل\u200cازظهر", "noon": "ظهر", "pm": "بعدازظهر", "morning1": "صبح", "afternoon1": "عصر", "evening1": "سر شب", "night1": "شب"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "میلادی", Abbrev: "م.", Narrow: "م"}, BC: ut.CalendarEraFormatNames{Full: "قبل از میلاد", Abbrev: "ق.م.", Narrow: "ق"}}}}
