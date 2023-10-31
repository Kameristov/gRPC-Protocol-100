package port

import (
	"bytes"
	"fmt"
	"time"

	"go.bug.st/serial"
)

type Port struct {
	portType string
	protocol string
	portName string
	port     serial.Port
}

func New(pName, pType, pProtocol string) Porter {
	p := &Port{
		portName: pName,
		portType: pType,
		protocol: pProtocol,
	}

	return p
}

func (p *Port) Conect() error {

	mode := &serial.Mode{}

	if p.portType == "RS-232" {
		switch p.protocol {
		case "1C":
			mode.BaudRate = 57600
			mode.Parity = serial.NoParity
			mode.StopBits = serial.OneStopBit
		case "2":
			mode.BaudRate = 4800
			mode.Parity = serial.EvenParity
			mode.StopBits = serial.OneStopBit
		case "Stndr":
			mode.BaudRate = 19200
			mode.Parity = serial.SpaceParity
			mode.StopBits = serial.OneStopBit
		}
	}

	var err error
	p.port, err = serial.Open(p.portName, mode)
	if err != nil {
		return fmt.Errorf("error opening port: %v", err)
	}
	p.port.SetReadTimeout(time.Microsecond * 500)
	return nil
}

func (p *Port) Disconect() error {
	if p.port == nil {
		return fmt.Errorf("port doesn't open")
	}

	err := p.port.Close()
	if err != nil {
		return fmt.Errorf("error closing port: %v", err)
	}
	return nil
}

func (p *Port) Communicate(data []byte) ([]byte, error) {

	err := p.Send(data)
	if err != nil {
		return nil, err
	}
	recv, err := p.Recv()
	if err != nil {
		return nil, err
	}
	return recv, nil
}

func (p *Port) Send(data []byte) error {
	_, err := p.port.Write(data)
	if err != nil {
		return fmt.Errorf("error writing port: %v", err)
	}
	fmt.Printf("Send: % 02X\n", data)
	return nil
}

func (p *Port) Recv() ([]byte, error) {
	var b bytes.Buffer
	buff := make([]byte, 100)
	for {
		n, err := p.port.Read(buff)
		if err != nil {

			break
		}
		if n == 0 {

			break
		}
		b.Write(buff[:n])
	}
	fmt.Printf("Recv: % 02X\n", b.Bytes())
	return b.Bytes(), nil
}
