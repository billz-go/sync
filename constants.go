package sync

const (
	DefaultTimeZone = "Asia/Tashkent"
)
const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	OfflineFileName = "offline.txt"
	GmtOffset       = 5
	OfflineMaxHours = 24

	HTTPTimeoutInSeconds = 30

	ParameterCannotBeEmpty = "parameter <<%s>> cannot be empty error"

	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

var (
	ValidationErrors = map[string]map[string]string{
		"NO_ITEMS":                {"en": "", "uz": "", "ru": "В чеке нет товаров"},
		"VAT_SUM_INCORRECT":       {"en": "", "uz": "", "ru": "Не корректная общая сумма НДС"},
		"VAT_ITEM_SUM_INCORRECT":  {"en": "", "uz": "", "ru": "Не корректная сумма НДС для товара "},
		"TOTAL_SUM_INCORRECT":     {"en": "", "uz": "", "ru": "Цена товаров и сумма оплат не совпадают"},
		"INCOMLETE_ITEM_DATA":     {"en": "", "uz": "", "ru": "Переданы не все поля в товаре "},
		"NO_INTERNET":             {"en": "", "uz": "", "ru": "Проверьте интернет соединение"},
		"INCORRECT_TIMEZONE":      {"en": "", "uz": "", "ru": "Не корректное время на компьютере"},
		"TOO_LONG_OFFLINE_PERIOD": {"en": "", "uz": "", "ru": "Восстановите подключение к интернету, вы оффлайн более 24 часов"},
	}
)
