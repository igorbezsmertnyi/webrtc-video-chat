# WebRTC video chat
___

The application builds on [Simple Peer](https://github.com/feross/simple-peer) node module, for opening peer connection between browsers and Go signal server with connect via WebSocket protocol to a user and sending command necessary for connection is done.

### Start Application

1. Clone the repostitory
2. Install nodejs
3. Install golang
4. Install postgresql
5. Run following command:

> `go install`<br/>
> `go run main.go`<br/>
> `npm install`<br/>
> `npm run dev`<br/>

Make sure that `psql` service working. Also, you need to have `postgres` user and `postgres` dbname.

For start `postgres` you able to use Docker, by run `docker-compose up` command.