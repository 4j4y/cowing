## Running instruction
- install go: 1.14 (might work with other versions as well)
- install following library for sending notifications
`go get -u github.com/gen2brain/beeep`
- To find open slots for 18+ citizens for today,  run following command

### Using PIN: 
`go run main.go pin 485001 1`

- To find open slots for 18+ citizens for next 5 days,  run following command
- Using PIN: `go run main.go pin 485001 1 5`

pin: always `pin` for finding slot using pin
pin value: value of pin for exampl: 485001
frequency: it defines frequency of querying in minutes

### Using District ID: 
`go run main.go did 513 1`

- To find open slots for 18+ citizens for next 5 days,  run following command
- Using PIN: `go run main.go pin 485001 1 5`

did: always `did` for finding slot using pin
did value: value of did for example: 513
frequency: it defines frequency of querying in minutes

#### How to find District ID
- go to cowin app, open console>network, select district, click search, see it in the query

## Configuration 
- To blacklist a center search following comment:
`// blacklist a center` and add your center like follownig with `&&` operator with existing clause
`covidData.Centers[i].CenterID != 582783`
