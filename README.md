# GO PING

Notificador de ping con systray y notificaciones

# Librerias necesarias

- go get -u github.com/getlantern/systray
- go get -u github.com/gen2brain/beeep
- go get -u github.com/go-toast/toast
- go get -u github.com/tadvi/systray

#Comandos de ejecucion

Para hacer debug en windows

```sh
$ GOOS=windows GOARCH=386 \
  CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
  go build
```

Para crear exe para windows

```sh
$ GOOS=windows GOARCH=386 \
  CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
  go build -ldflags -H=windowsgui
```

#Ejemplo

![alt text](http://drive.google.com/uc?export=view&id=17rU4wYyrQ_iJRdWyQTvPrcnC-d1OfmkA "Ejemplo")


License
----

MIT
