OnePlay
===================================
High performance and low latency cloud gaming with OnePlay



Features
--------------
- WebRTC streaming protocol (base on [pion/webrtc](https://github.com/pion/webrtc))
- Codec:
    - video codec: H264, HEVC (in progress)
    - audio codec: OPUS 
- Low-level API:  
    - Screen capture: D3D11 Output Duplication (D3D11 API)
    - Encoder: NVENC (through libavcodec)
    - IPC-protocol: RTP over UDP (through libavformat) 
- Client app:
    - Browser (for demo purpose)
    - Window native app (in progress)     

Join us
-----------
[Architecture](https://miro.com/app/board/uXjVOOfX5j0=/?share_link_id=961338824424) |
[Document](https://oneplay.notion.site/Project-document-c02bb0e93db04b2ca1a69053590f4613)


Subproject
-------------------------
  - [screencoder]() - Battery included screen capturer-encoder-packetizer. 
  - [hid-proxy]() - simulate user's interaction on another devices
  - [webrtc-proxy]() - WebRTC protocol data and multimedia stream transport agent.
  - [webrtc-signaling]() - necessary internet infrastructure for [webrtc-proxy]().
  - [webrtc-browser]() - access [webrtc-proxy]() from your browser device





Licensing
-----------
is distributed uner GNU General Public License v3

Contributing
-----------
- If you want to contribute to this repository email us at dohuy@oneplay.in