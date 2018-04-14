const basePath = window.location.hostname === 'localhost' ? 'localhost:3000' : ''

export default {
  createRoomPath: () => (
    `${basePath}/create_new_room`
  )
}