# wlk


[![GoDoc](https://godoc.org/github.com/xackery/wlk?status.svg)](https://godoc.org/github.com/xackery/wlk) [![Go Report Card](https://goreportcard.com/badge/github.com/xackery/wlk)](https://goreportcard.com/report/github.com/xackery/wlk)


Based on lxn/walk which hasn't been updated for 3+ years, as well contributions to the fork tailscale/walk, simplified and modernized

Majority of this code can be sourced to:
- [tailscale/walk](https://github.com/tailscale/walk)
- [tailscale/win](https://github.com/tailscale/win)
- [lxn/walk](https://github.com/lxn/walk)
- [lxn/win](https://github.com/lxn/win)

Check out the [LICENSE](LICENSE) file for more information, or [AUTHORS](AUTHORS) for a list of contributors.

I don't take credit for the majority of this code, as the above refs show there has been an incredible amount of work prior to my copy.

## Why Wlk?

Wlk is a direct approach to rendering components via win32 API calls using WinForms. It starts fast and has minimum dependencies.

It **does not** have cross-platform support. It is exclusively for windows.

## Examples

Check out # [examples](examples/README.md) to see how the library looks.

There are *many, many* warnings for how old walk is. We are slowly going through and cleaning them up.

## Dark Mode Support

This is work in progress.

- [x] Dark Mode is opted into by using `walk.SetDarkModeAllowed(true)`. This is disabled by default, and should be called before any other walk functions to ensure proper painting
- [x] If your program needs to detect if Dark Mode is enabled, use `walk.IsDarkModeEnabled()`
- [x] Dark Mode is detected via registry
- [ ] Theme changing (Light Mode / Dark Mode) subscription/notification is not yet implemented
- [ ] Overhaul all components to default to dark mode if enabled on initialization
- [ ] Allow custom theming.