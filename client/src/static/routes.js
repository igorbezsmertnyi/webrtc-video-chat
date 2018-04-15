const basePath = window.location.hostname === 'localhost' ? 'http://localhost:3000' : ''

export default {
  createRoomPath: () => (
    `${basePath}/api/create_room`
  ),
  connectToRoomPath: slug => (
    `${basePath}/api/room/${slug}`
  )
}