# lets-zeditor
Try the zed editor that supports vim mode.

# Install @ Artix Linux
```bash
sudo pacman -S zed
```

## Use CLI
```bash
zeditor file
```
However, it seems that the window that starts up does not close properly with `:wq`... I would be happy if vim mode was officially supported.

## joke
```bash
$ go build -gcflags '-N -l -m' main.go
# command-line-arguments
./main.go:8:2: moved to heap: a
./main.go:19:11: ... argument does not escape
./main.go:19:12: *g escapes to heap
./main.go:20:11: ... argument does not escape
./main.go:20:12: "Hello, Zed!\n" escapes to heap
```
