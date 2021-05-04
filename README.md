## Running instruction
- install go: 1.14 (might work with other versions as well)
- install following library for sending notifications
`go get -u github.com/gen2brain/beeep`
`go get -u github.com/faiface/beep`
`go get .`

- To find open slots for 18+ citizens for today,  run following command
- Using PIN: `go run main.go pin 485001 1`

- To find open slots for 18+ citizens for next 5 days,  run following command
- Using PIN: `go run main.go pin 485001 1 5`

- To find open slots for 18+ citizens for next 5 days and alert by playing manual song(plays for 2 sec),  run following command
- Using PIN: `go run main.go pin 485001 1 5 example.mp3`

pin: always `pin` for finding slot using pin
pin value: value of pin for exampl: 485001
frequency: it defines frequency of querying in minutes
- Using : `go run main.go did 513 1`
did: always `did` for finding slot using pin
did value: value of pin for example: 513

frequency: it defines frequency of querying in minutes

## Configuration 
- To blacklist a center with `centre_id`, add it to the list names `BLACKLIST` in `main.go`