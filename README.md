# Steam Redirector Go

This is an experimental reimplementation of the [Mod Organizer 2 Linux
Installer's Steam Redirector](https://github.com/Furglitch/modorganizer2-linux-installer/tree/6.0.6/steam-redirector).
This is written in Go, as opposed to the original in C.

This implementation has no dependencies beyond the Go compiler and does not
need different code for Linux and Windows string and file handling.

It also creates the file steam-redirector.log in the game folder, which includes
errors from this program as well as the stdout/stderr of the game.

To compile:

| OS / Arch      | Command                              |
|----------------|--------------------------------------|
| Native         | `go build`                           |
| Windows 64-bit | `GOOS=windows GOARCH=amd64 go build` |
| Windows 32-bit | `GOOS=windows GOARCH=386 go build`   |

Notes:

Build as DLL:
```bash
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -buildmode=c-shared -o steam-redirector.dll
```

Dump version.dll:
```bash
winedump spec test.dll -I . -f ~/'.steam/steam/steamapps/common/Proton 10.0/files/share/default_pfx/drive_c/windows/syswow64/version.dll'
```