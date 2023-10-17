# wlk
Based on abandoned lxn/walk and fork tailscale/walk, simplified and modernized

Majority of this code can be sourced to:
- [tailscale/walk](https://github.com/tailscale/walk)
- [tailscale/win](https://github.com/tailscale/win)
- [lxn/walk](https://github.com/lxn/walk)

Check out the [LICENSE](LICENSE) file for more information, or [AUTHORS](AUTHORS) for a list of contributors.

## Why Wlk?

Windows users love the GUI. OSX and linux users tend to prefer cli, especially if devs. I don't need a cross-platform, electron/HTML powered GUI. I just need a windows app that renders what go is doing. Walk felt the right way to take it, wlk takes it a bit further.

## Examples

Check out # [examples](examples) to check out how the library looks.

There are *many, many* warnings for how old walk is. I'm slowly going through and cleaning them up.

## Dark Mode ?

This is work in progress.

- [x] Dark Mode is opt in via `walk.SetDarkModeAllowed(true)`. This is disabled by default, and should be called before any other walk functions to ensure proper painting.
- [x] If your program need to detect if Dark Mode is enabled, use `walk.IsDarkModeEnabled()`
- [x] Dark Mode is detected via registry
- [ ] Theme changing (Light Mode / Dark Mode) subscription/notification is not implemented
- [ ] Overhaul all components to default to dark mode if enabled on initialization
