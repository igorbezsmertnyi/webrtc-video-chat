export default {
  socket: null,
  connect: () => {
    this.socket = new WebSocket('ws://localhost:3000/channel')
  },
  send: data => {
    this.socket.send(data)
  }
}