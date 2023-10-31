package protocol100r5

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/snksoft/crc"
)

// Формирование запроса
func cmdBuildRequest(data []byte) ([]byte, error) {
	if len(data) < 1 {
		return nil, errors.New("empty message")
	}

	var b bytes.Buffer

	// Header
	b.WriteByte(header0)
	b.WriteByte(header1)
	b.WriteByte(header2)

	// Len
	var dataLen []byte = make([]byte, 2)
	binary.LittleEndian.PutUint16(dataLen, uint16(len(data)))
	b.Write(dataLen)

	// Data
	b.Write(data)

	// CRC
	var hash []byte = make([]byte, 2)
	binary.LittleEndian.PutUint16(hash, uint16(crc.CalculateCRC(crc.CCITT, data)))
	b.Write(hash)

	return b.Bytes(), nil
}

// Парсинг ответа
func cmdParseResponse(data []byte) ([]byte, error) {
	// Проверка на длину сообщения
	if len(data) < 3 {
		return []byte{}, errors.New("bad response len ")
	}

	// Проверка header
	if data[0] != header0 || data[1] != header1 || data[2] != header2 {
		return []byte{}, errors.New("bad response")
	}

	// Провекра CRC
	ourHash := uint16(crc.CalculateCRC(crc.CCITT, data[5:len(data)-2]))
	inputHash := binary.LittleEndian.Uint16(data[len(data)-2:])
	if inputHash != ourHash {
		fmt.Printf("in % 02X | our % 02X", inputHash, ourHash)
		return []byte{}, errors.New("crc check failed")
	}

	// Длина сообщения
	//dataLen := binary.LittleEndian.Uint16(data[3:5])

	// Определение ответа
	switch data[5] {
	case cmd_nack:
		return []byte{}, errors.New("принята неизвестная команда")
	case cmd_nack_tare:
		return []byte{}, errors.New("невозможно установить тару")
	case cmd_error:
		return []byte{}, errors.New("error: " + getError(data[6]))
	case cmd_ack_scale_par:
		return data[5 : len(data)-2], nil
	case cmd_ack_massa:
		return data[5 : len(data)-2], nil
	case cmd_ack_set_tare:
		return data[5 : len(data)-2], nil
	case cmd_ack_set:
		return data[5 : len(data)-2], nil
	case cmd_ack_name:
		return data[5 : len(data)-2], nil
	case cmd_ack_set_ethernet:
		return data[5 : len(data)-2], nil
	case cmd_ack_wifi_ip:
		return data[5 : len(data)-2], nil
	case cmd_ack_wifi_ssid:
		return data[5 : len(data)-2], nil
	}

	return []byte{}, errors.New("unexpected command")
}

func getError(cmdErr byte) string {
	switch cmdErr {
	case cmd_error_type_0x07:
		return "Команда не поддерживается"
	case cmd_error_type_0x08:
		return "Нагрузка на весовом устройстве превышает НПВ"
	case cmd_error_type_0x09:
		return "Весовое устройство не в режиме взвешивания"
	case cmd_error_type_0x0A:
		return "Ошибка входных данных"
	case cmd_error_type_0x0B:
		return "Ошибка сохранения данных"
	case cmd_error_type_0x10:
		return "Интерфейс WiFi не поддерживается"
	case cmd_error_type_0x11:
		return "Интерфейс Ethernet не поддерживается"
	case cmd_error_type_0x15:
		return "Установка >0< невозможна"
	case cmd_error_type_0x17:
		return "Нет связи с модулем взвешивающим"
	case cmd_error_type_0x18:
		return "Установлена нагрузка на платформу при включении весового устройства"
	case cmd_error_type_0x19:
		return "Весовое устройство неисправно"
	case cmd_error_type_0xF0:
		return "Неизвестная ошибка"
	}
	return ""
}
