onpar
---

A dwm bar that is on par with other advanced bars such as i3blocks and polybar

## Features

- Lightweight
- Written in Go which is more maintainable than C
- Currently shows date and time, battery level, volume level, DarkSky weather forecast, and keyboard layout
- Can schedule running each function indepdently

## Installation

- You need a working go environment
- `go build onpar.go` to build the binary
- Add the binary to your path
- Add the path to the binary in your `.xinitrc`

## TODOs

- [X] Make it configurable
- [] Add colors
- [] Add icons
- [] Revisit the use of goroutines and the possibility to use alarm signals instead
- [] Refactor the code

## License

See [LICENSE](https://github.com/aonemd/onpar/blob/master/LICENSE).
