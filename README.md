# Marshall Speaker Bluetooth Keep

A tiny tool to keep the Marshall Bluetooth speaker connected, overcoming the 20-minute connection limit without playing any sounds.

## Reason

Marshall speakers disconnect after 20 minutes of inactivity, which can be inconvenient. This tool helps maintain the Bluetooth connection.

Reference: https://my.marshall.com/forum/question/9421/stanmore-bluetooth-connection-drops-every-20-minutes-while-connected-and-playing-music

## Compile

```shell
make build
```

## Usage

```shell
./build/bin/marshall-bluetooth-keep
```

### Linux Systemd Installation

```shell
cp ./build/bin/marshall-bluetooth-keep /usr/local/bin
cp marshall-bluetooth-keep.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable --now marshall-bluetooth-keep
```
