onpar
---

A dwm bar that is on par with other advanced bars such as i3blocks and polybar

## Features

- The ability to schedule running each function independently
- The extensibility to write any script and use its STDOUT and show it on the bar
- Currently shows date and time, battery level, volume level, DarkSky weather forecast, and keyboard layout

## Installation

- You need a working Go environment and [FontAwesome](https://fontawesome.com) & [icons-in-terminal](https://github.com/sebastiencs/icons-in-terminal) to show the icons
- Build the Go components in `components/` using `go build component_name.go`
- Add your own components
- You need to use the full path to the components folder or add the components folder to your $PATH
- `go build onpar.go` to build the main binary
- Add the binary to your $PATH
- Add the path to the binary in your `.xinitrc`

## TODOs

- [X] Make it configurable
- [] Add colors
- [X] Add icons
- [] Revisit the use of goroutines and the possibility to use alarm signals instead
- [X] Refactor the code

## License

See [LICENSE](https://github.com/aonemd/onpar/blob/master/LICENSE).
