# Fiesta Turtle

## Services

### Brian

Server for connecting clients

### Richard

Search server

### Client

Gui for searching for and playing songs

## Install

- Install ipfs

```
$ go get -d github.com/ipfs/go-ipfs
$ cd $GOPATH/src/github.com/ipfs/go-ipfs
$ make toolkit_upgrade
$ make install
```

- Start ipfs

```
$ ipfs start
```

- Install, build, and run Brian

```
$ cd fiesta-turtle/brian
$ go get
$ go build
$ ./brian
```

- Run Richard

```
$ cd fiesta-turtle/richard
$ go run main.go
```

- Install, build, and run the client

Note: make sure you're on at Node v5.0 or above. Electron won't build correctly with older versions.

```
$ cd fiesta-turtle/client/desktop
$ npm install
$ npm run hot-server
$ npm run start-hot
```
