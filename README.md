# Chat Service
to communicate with chat service, subscribe to websocket 
and give it access_token to create client with your userId
``
ws://localhost:8000/ws/subscribe?access_token={token}
``

### message body
body of websocket communication would be like this

Parameter | Type
:-:|:-:
`action` | String
`type` | String 
`id` | String
`body` | Interface

{"action":"Message","type":"user","id":"16","body":6}

body has two action
* Message
* Join

use `Message` to send message to user, group , etc...

use `Join` to join groups or channel with specific id

body has three type too
* user
* group
* channel

use `user` to send message to specific user

use `group` to send message or join to the group

use `channel` to send message or join to the channel

* id

use `id` to specify which user or group or channel 
you want to send message or join to.