
# A go webhook client for rocket.chat

A golang [rocket.chat](https://rocket.chat) hook :rocket:

## Usage

```sh
# Setup webhook in administration area
$ export ROCKET_HOOK="http://<rocket.domain>/hooks/<token>"
$ gocket "Ahoi!"
```

## Development

```sh
$ go run cmd/main.go Ahoi
```
