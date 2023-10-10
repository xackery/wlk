# Progressindicator Example


TODO: fix progress indicator.
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x20 pc=0x2c7be5]

goroutine 1 [running, locked to thread]:
github.com/xackery/wlk/walk.(*ContainerBase).CreateLayoutItem(...)
        C:/src/wlk/walk/container.go:119
github.com/xackery/wlk/walk.CreateLayoutItemsForContainerWithContext({0x3823b8?, 0xc0000a6380?}, 0xc00009e430)
        C:/src/wlk/walk/layout.go:65 +0x17e
github.com/xackery/wlk/walk.CreateLayoutItemsForContainer({0x3823b8, 0xc0000a6380})
        C:/src/wlk/walk/layout.go:51 +0x91
github.com/xackery/wlk/walk.(*FormBase).startLayout(0xc00000d400)
        C:/src/wlk/walk/form.go:685 +0x79
github.com/xackery/wlk/walk.(*FormBase).WndProc(0xc00000d400, 0x0?, 0x47, 0x0, 0x266c5ff760)
        C:/src/wlk/walk/form.go:792 +0x468
github.com/xackery/wlk/walk.(*Dialog).WndProc(0x2fa5a0?, 0xc0001886c0?, 0x290438?, 0xc000107920?, 0xc000107620?)
        C:/src/wlk/walk/dialog.go:260 +0xbe
github.com/xackery/wlk/walk.defaultWndProc(0x1e308f?, 0x107a18?, 0xc000107a20?, 0x1e308f?)
        C:/src/wlk/walk/window.go:2207 +0xa5
syscall.SyscallN(0x7ffebcb1ea90?, {0xc000107cd0?, 0x3?, 0x0?})
        C:/Program Files/Go/src/runtime/syscall_windows.go:544 +0x107
syscall.Syscall(0xc000188240?, 0x303620?, 0x2dc529?, 0x2ce740?, 0x2dc527?)
        C:/Program Files/Go/src/runtime/syscall_windows.go:482 +0x35
github.com/xackery/wlk/win.ShowWindow(0x290438?, 0x8)
        C:/src/wlk/win/user32.go:3431 +0x56
github.com/xackery/wlk/walk.setWindowVisible(...)
        C:/src/wlk/walk/window.go:1421
github.com/xackery/wlk/walk.(*WindowBase).SetVisible(0xc00000d400, 0x1)
        C:/src/wlk/walk/window.go:1385 +0x5a
github.com/xackery/wlk/walk.(*FormBase).Show(0xc00000d400)
        C:/src/wlk/walk/form.go:566 +0x122
github.com/xackery/wlk/walk.(*Dialog).Show(0xc00000d400)
        C:/src/wlk/walk/dialog.go:192 +0x195
github.com/xackery/wlk/walk.(*Dialog).Run(0xc00000d400)
        C:/src/wlk/walk/dialog.go:235 +0x18
main.RunMyDialog({0x0, 0x0})
        C:/src/wlk/examples/progressindicator/pi.go:76 +0x7de
main.main()
        C:/src/wlk/examples/progressindicator/pi.go:16 +0x17
        
![Alt text](image.png)