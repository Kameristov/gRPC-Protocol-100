package protocol100r5

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gRPC-Protocol-100/pkg/port"
	"strconv"
)

type Command struct {
	tareValue int32
	port      port.Porter
}

func New(p port.Porter) Commander {
	c := &Command{port: p}
	return c
}

type ScaleParData struct {
	P_Max   string // Максимальная нагрузка
	P_Min   string // Минимальная нагрузка
	P_e     string // Поверочный интервал весов
	P_T     string // Максимальная масса тары
	Fix     string // Параметр фиксации веса
	Calcode string // Код юстировки
	PO_Ver  string // Версия ПО
	PO_Summ string // Контрольная сумма ПО
}

func BytesToScalePar(data []byte) ScaleParData {
	sliceBytes := bytes.Split(data, []byte{0x0D, 0x0A})
	sp := ScaleParData{
		P_Max:   string(sliceBytes[0]),
		P_Min:   string(sliceBytes[1]),
		P_e:     string(sliceBytes[2]),
		P_T:     string(sliceBytes[3]),
		Fix:     string(sliceBytes[4]),
		Calcode: string(sliceBytes[5]),
		PO_Ver:  string(sliceBytes[6]),
		PO_Summ: string(sliceBytes[7]),
	}
	return sp
}

// Получить параметры весов
func (c *Command) GetScalePar() (string, error) {

	res, err := PortCycle(c.port, []byte{cmd_get_scale_par})
	if err != nil {
		return "", fmt.Errorf("error GetScalePar: %v", err)
	}

	scalePar := BytesToScalePar(res)
	str := fmt.Sprintf("%+v", scalePar)
	return str, err
}

// Получить массу на весах
func (c *Command) GetMassa() (string, error) {
	res, err := PortCycle(c.port, []byte{cmd_get_massa})
	if err != nil {
		return "", fmt.Errorf("error GetMassa: %v", err)
	}
	str := fmt.Sprintf("%d", binary.LittleEndian.Uint32(res[1:5]))
	fmt.Println(res, str, err)
	return str, err
}

// Применить значение тары
func (c *Command) SetTare() error {
	var tareVal []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(tareVal, uint32(c.tareValue))

	_, err := PortCycle(c.port, bytes.Join([][]byte{{cmd_set_tare}, tareVal}, []byte{}))
	if err != nil {
		return fmt.Errorf("error SetTare: %v", err)
	}

	return nil
}

// Изменить вес тары
func (c *Command) SetTareValue(val string) error {
	tareVal, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return fmt.Errorf("error parsing string to int: %v", err)
	}
	c.tareValue = int32(tareVal)
	return nil
}

// Установить нулевое значение весов
func (c *Command) SetZero() error {
	_, err := PortCycle(c.port, []byte{cmd_set_zero})
	return err
}

func PortCycle(p port.Porter, data []byte) ([]byte, error) {
	request, err := cmdBuildRequest(data)
	if err != nil {
		return []byte{}, fmt.Errorf("error building request: %v", err)
	}
	
	response, err := p.Communicate(request)
	if err != nil {
		return []byte{}, fmt.Errorf("error port communication: %v", err)
	}
	
	res, err := cmdParseResponse(response)
	if err != nil {
		return []byte{}, fmt.Errorf("error parse response: %v", err)
	}
	return res, nil
}
