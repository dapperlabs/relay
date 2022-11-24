[![Go Report](https://goreportcard.com/badge/github.com/dapperlabs/relay)](https://goreportcard.com/report/github.com/dapperlabs/relay)
[![License](https://img.shields.io/github/license/dapperlabs/relay.svg)](https://github.com/dapperlabs/relay/blob/master/LICENSE)

[![GoDoc](https://godoc.org/github.com/dapperlabs/relay?status.svg)](https://godoc.org/github.com/dapperlabs/relay)

# mail-relay

SMTP server to use for forwarding messages via HTTP-based services in environments not supporting SMTP outbound directly

# How this works
Relay accepts SMTP connection to handle message with invoking Mailgun API with the message it got. Simply :)
In depth in current implementation relay sends message synchronously via Mailgun Go's client.

# How to use
```
docker run -it \
  -e RELAY_ADDR=:2025 \
  -e RELAY_DOMAIN="<domain>" \
  -e RELAY_MAILGUN_PRIVATE_KEY="<mailgun private key>" \
  -e RELAY_MAX_IDLE_SECONDS=300 \
  -e RELAY_MAX_MESSAGE_BYTES=1048576 \
  -e RELAY_MAX_RECIPIENTS=50 \
  teran/relay
```


## License
MIT
