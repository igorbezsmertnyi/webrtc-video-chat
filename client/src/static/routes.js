const basePath = window.location.hostname === 'localhost' ? 'http://localhost:3000' : ''

const wsSchema = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
const baseSocket = window.location.hostname === 'localhost' ? `${wsSchema}//localhost:3000` : `${wsSchema}//${host}`

export default {
  createRoomPath: () => (
    `${basePath}/api/create_room`
  ),
  getRoomPath: slug => (
    `${basePath}/api/get_room/${slug}`
  ),
  connectToRoomPath: slug => (
    `${basePath}/api/room/${slug}`
  ),
  socketPath: slug => (
    `${baseSocket}/channel/${slug}`
  )
}