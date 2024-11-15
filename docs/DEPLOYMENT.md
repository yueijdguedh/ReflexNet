## ReflexNet Deployment Guide

This guide covers deploying a ReflexNet node in various environments.

## Prerequisites

- Go 1.21 or higher
- Git
- Make
- 16GB RAM minimum
- 100GB SSD storage

## Building from Source

```bash
git clone https://github.com/yueijdguedh/ReflexNet.git
cd ReflexNet
make build
```

The binary will be located at `bin/reflexnetd`.

## Initialize Node

```bash
reflexnetd init <moniker> --chain-id reflexnet-1
```

## Configure Genesis

Download the genesis file:

```bash
curl https://raw.githubusercontent.com/yueijdguedh/ReflexNet/main/genesis.json > ~/.reflexnet/config/genesis.json
```

## Configure Peers

Edit `~/.reflexnet/config/config.toml`:

```toml
persistent_peers = "peer1@ip1:26656,peer2@ip2:26656"
```

## Start Node

```bash
reflexnetd start
```

## Run as Service

Create systemd service file at `/etc/systemd/system/reflexnetd.service`:

```ini
[Unit]
Description=ReflexNet Node
After=network.target

[Service]
Type=simple
User=reflexnet
ExecStart=/usr/local/bin/reflexnetd start
Restart=on-failure
RestartSec=10
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl enable reflexnetd
sudo systemctl start reflexnetd
```

## Validator Setup

Create validator:

```bash
reflexnetd tx staking create-validator \
  --amount=10000000mcell \
  --pubkey=$(reflexnetd tendermint show-validator) \
  --moniker="<moniker>" \
  --chain-id=reflexnet-1 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --from=<key-name>
```

## Docker Deployment

Build Docker image:

```bash
docker build -t reflexnet:latest .
```

Run container:

```bash
docker run -d \
  --name reflexnet-node \
  -v ~/.reflexnet:/root/.reflexnet \
  -p 26656:26656 \
  -p 26657:26657 \
  reflexnet:latest
```

## Monitoring

Configure Prometheus endpoint in `config.toml`:

```toml
prometheus = true
prometheus_listen_addr = ":26660"
```

## Backup

Backup private keys:

```bash
cp ~/.reflexnet/config/priv_validator_key.json ~/backup/
cp ~/.reflexnet/config/node_key.json ~/backup/
```

## Security

- Use firewall to restrict ports
- Keep software updated
- Use key management service for validators
- Enable sentry node architecture for validators

## Troubleshooting

Check logs:

```bash
journalctl -u reflexnetd -f
```

Check sync status:

```bash
reflexnetd status | jq .SyncInfo
```

## Upgrade

```bash
# Stop node
sudo systemctl stop reflexnetd

# Build new version
git pull
make build

# Copy binary
sudo cp bin/reflexnetd /usr/local/bin/

# Start node
sudo systemctl start reflexnetd
```

