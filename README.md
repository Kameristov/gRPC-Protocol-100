# gRPC-Protocol-100

## Задание

 Написать драйвер устройства [(Протокол 100)](./docs/Protokol_100%20(r%205)%202018%20V3.pdf) ,который передает данные, используя gRPC ([proto файл](./internal/grpcserver/stream/stream.proto)).

## Настройка

Файл [Config.json](./config.json) содержит настройки сервиса. Файл должен распологаться в одном катологе с сервисом.

- `"Level"` - уровень логирования:
  - `Error`
  - `Warn`
  - `Info`
  - `Debug`
- `"SerialPort"` - название последовательного порта:
  - `COMx` - windows;
- `Type`- тип подключения:
  - `USB`
  - `RS-232`
- `Protocol` тип протокола при передаче по RS-232:
  - `1C`- BAUD RATE = 57600, Parity = none, Stop = 1;
  - `2` - BAUD RATE = 4800, Parity = even, Stop = 1;
  - `Stndr` - BAUD RATE = 19200, Parity = space, Stop = 1;
- `Address` - порт gRPC.

## Сборка

Сборка бинарника под Windows

```makefile
make build
```

После сборки в корне проекта появится бинарный файл сервиса `gRPC-Protocol-100.exe`

## Запуск

```makefile
make run
```

Собирает и запускает сервис.
