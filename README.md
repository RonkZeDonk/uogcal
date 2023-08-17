# UoG Course Calendar

[UoG Course Calendar](https://uogcal.ronkzd.xyz) is an automated schedule to
calendar app. Inputting labs and lectures into google calendar is tedious.
Automate this process into a couple clicks.

Future plans include adding assignment/test dates to your calendar with the help
of volunteer contributors.

## Running

- Migrate your Postgres database to the latest version in the `./migrations`
  folder
- Simply run `make` to create a executable. You can use the GOOS and GOARCH
  environment variables to pick a different os/architecture

## Developing

Run two different terminal sessions:

- on session A run `make watchStyles`
  - this will run the Tailwind CLI to compile a CSS bundle
- on session B run the go project (ie. `go run cmd/main.go`)
  - you can set the `ENV` environment variable to `development` to enable
    hot-reloading for .hbs files (ie. `ENV=development go run cmd/main.go`)

## Production Information

This code targets super low end hardware. More specifically a Raspberry Pi 1B+
(basically a full sized Pi with the power of a RPi Zero) running Postgres 11.