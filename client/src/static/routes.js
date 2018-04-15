const basePath = window.location.hostname === 'localhost' ? 'http://localhost:3000' : ''

export default {
  createRoomPath: () => (
    `${basePath}/api/create_room`
  ),
  getRoomPath: slug => (
    `${basePath}/api/get_room/${slug}`
  ),
  connectToRoomPath: slug => (
    `${basePath}/api/room/${slug}`
  )
}