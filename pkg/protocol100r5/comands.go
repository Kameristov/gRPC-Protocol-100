package protocol100r5

// Заголовочная последовательность
const (
	header0 byte = 0xF8
	header1 byte = 0x55
	header2 byte = 0xCE
)

// Список команд
// Get
const (
	cmd_get_scale_par byte = 0x75 // Запрос параметров весового устройства
	cmd_get_massa     byte = 0x23 // Запрос значения массы нетто, массы тары, флагов стабильности, установки нуля и тары
	cmd_get_name      byte = 0x20 // Запрос имени и ID весового устройства
	cmd_get_ethernet  byte = 0x2d // Запрос параметров Ethernet
	cmd_get_wifi_ip   byte = 0x33 // Запрос IP-параметров подключения к сети Wi-Fi
	cmd_get_wifi_ssid byte = 0x3A // Запрос параметров доступа к сети Wi-Fi
)

// Set
const (
	cmd_set_tare byte = 0xA3 // Установить массу тары на весовом устройстве

	cmd_set_zero      byte = 0x72 // Установить >0<
	cmd_set_name      byte = 0x22 // Установить имя весового устройства
	cmd_set_ethernet  byte = 0x2D // Установить параметры Ethernet
	cmd_set_wifi_ip   byte = 0x31 // Установить IP-параметры подключения к сети Wi-Fi
	cmd_set_wifi_ssid byte = 0x3C // Установить параметры доступа к сети Wi-Fi
)

// Ack
const (
	cmd_ack_scale_par    byte = 0x76 // Параметры весового устройства.
	cmd_ack_massa        byte = 0x24 // Масса нетто, масса тары, флаги стабильности, установки нуля и тары с весового устройства
	cmd_ack_set_tare     byte = 0x12 // Команда установки тары выполнена успешно
	cmd_ack_set          byte = 0x27 // Команда выполнена успешно
	cmd_ack_name         byte = 0x21 // Передача имени и ID весового устройств
	cmd_ack_set_ethernet byte = 0x2E // Передача параметров Ethernet
	cmd_ack_wifi_ip      byte = 0x34 // Передача IP-параметров подключения к сети Wi-Fi
	cmd_ack_wifi_ssid    byte = 0x3B // Передача параметров 	доступа к сети Wi-Fi
)

// Err
const (
	cmd_error byte = 0x28 // Ошибка выполнения команды
)

// Nack
const (
	cmd_nack      byte = 0xF0 // Принята неизвестная команда
	cmd_nack_tare byte = 0x15 // Невозможно установить тару
)

// Error types
const (
	cmd_error_type_0x07 byte = 0x07 // Команда не поддерживается
	cmd_error_type_0x08 byte = 0x08 // Нагрузка на весовом устройстве превышает НПВ
	cmd_error_type_0x09 byte = 0x09 // Весовое устройство не в режиме взвешивания
	cmd_error_type_0x0A byte = 0x0A // Ошибка входных данных
	cmd_error_type_0x0B byte = 0x0B // Ошибка сохранения данных
	cmd_error_type_0x10 byte = 0x10 // Интерфейс WiFi не поддерживается
	cmd_error_type_0x11 byte = 0x11 // Интерфейс Ethernet не поддерживается
	cmd_error_type_0x15 byte = 0x15 // Установка >0< невозможна
	cmd_error_type_0x17 byte = 0x17 // Нет связи с модулем взвешивающим
	cmd_error_type_0x18 byte = 0x18 // Установлена нагрузка на платформу при включении весового устройства
	cmd_error_type_0x19 byte = 0x19 // Весовое устройство неисправно
	cmd_error_type_0xF0 byte = 0xF0 // Неизвестная ошибка
)
