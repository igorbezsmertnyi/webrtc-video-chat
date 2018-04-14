import Peer from 'simple-peer'

const WebRTC = {}
const connection = {
  initiator: true,
  trickle: false
}
const peer = new Peer(connection)

const start = () => {
  peer.on('signal', e => {
    console.log(e)
  })

  const video = document.getElementById('videoStream')

  window.URL = window.URL || window.webkitURL
  window.MediaSource = window.MediaSource || window.WebKitMediaSource
  navigator.getMedia = (navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia)

  const handleSuccess = stream => {
    video.srcObject = stream
    // pc.addStream(stream)

    setTimeout(() => 
      video.play()
        .catch(err => logError(err))
      , 1000)
  }

  const localDescCreated = desc => {
    // pc.setLocalDescription(desc, () => {
    //   // signalingChannel.send(JSON.stringify({
    //   //   'sdp': pc.localDescription
    //   // }))
    // }, logError)
  }

  const logError = err => {
    console.error(`${err.name}: ${err.message}`)
  }

  navigator.mediaDevices.getUserMedia({ video: true, audio: false })
    .then(handleSuccess)
    .catch(logError)
}

WebRTC.start = start

export default WebRTC