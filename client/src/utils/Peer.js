import Peer from 'simple-peer'
import Socket from './Socket'

const connection = {
  initiator: true,
  trickle: false
}

const peer = new Peer(connection)

Socket.connect()

peer.on('signal', e => {
  console.log(e)
  Socket.send(e)
})