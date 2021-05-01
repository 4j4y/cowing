## running instruction
- install go: 1.14 (might work with other versions as well)
- install following library for sending notifications
`go get -u github.com/gen2brain/beeep`

## Configuration 
- To blacklist a center search following comment:
`// blacklist a center` and add your center like follownig with `&&` operator with existing clause
`covidData.Centers[i].CenterID != 582783`
