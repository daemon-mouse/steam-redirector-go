# Steam Redirector Go

This is an experimental reimplementation of the [Mod Organizer 2 Linux
Installer's Steam Redirector](https://github.com/Furglitch/modorganizer2-linux-installer/tree/6.0.6/steam-redirector).
This is written in Go, as opposed to the original in C.

This implementation has no dependencies beyond the Go compiler and does not
need different code for Linux and Windows string and file handling.

It also has two feature improvements over the original:

- It creates the file steam-redirector.log in the game folder, which
  includes errors from this program as well as the stdout/stderr of the game.
- It allows multiple arguments to be passed to the wrapped program instead of
  only one.

To compile:

| OS / Arch      | Command                              |
|----------------|--------------------------------------|
| Native         | `go build`                           |
| Windows 64-bit | `GOOS=windows GOARCH=amd64 go build` |
| Windows 32-bit | `GOOS=windows GOARCH=386 go build`   |
