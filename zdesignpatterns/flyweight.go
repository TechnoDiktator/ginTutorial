package zdesignpatterns


/*

Conceptual Example
In a game of Counter-Strike, Terrorist and Counter-Terrorist have a different type of dress. For simplicity, let’s assume that both Terrorist and Counter-Terrorists have one dress type each. The dress object is embedded in the player object as below.

Below is the struct for a player. We can see that the dress object is embedded in the player struct:

type player struct {
    dress      dressA
    playerType string // Can be T or CT
    lat        int
    long       int
}
Let’s say there are 5 Terrorists and 5 Counter-Terrorist, so a total of 10 players. Now there are two options concerning dress.

Each of the 10 player objects creates a different dress object and embeds them. A total of 10 dress objects will be created.

We create two dress objects:

Single Terrorist Dress Object: This will be shared across 5 Terrorists.
Single Counter-Terrorist Dress Object: This will be shared across 5 Counter-Terrorist.
As you can see in Approach 1, a total of 10 dress objects are created while in approach 2 only 2 dress objects are created. The second approach is what we follow in the Flyweight design pattern. The two dress objects which we created are called the flyweight objects.

The Flyweight pattern takes out the common parts and creates flyweight objects. These flyweight objects (dress) can then be shared among multiple objects (player). This drastically reduces the number of dress objects, and the good part is that even if you create more players, only two dress objects will be sufficient.

In the flyweight pattern, we store the flyweight objects in the map field. Whenever the other objects which share the flyweight objects are created, then flyweight objects are fetched from the map.

Let’s see what parts of this arrangement will be intrinsic and extrinsic states:

Intrinsic State: Dress in the intrinsic state as it can be shared across multiple Terrorist and Counter-Terrorist objects.

Extrinsic State: Player location and the player weapon are an extrinsic state as they are different for every object.

 dressFactory.go: Flyweight factory
package main

import "fmt"

const (
    //TerroristDressType terrorist dress type
    TerroristDressType = "tDress"
    //CounterTerrroristDressType terrorist dress type
    CounterTerrroristDressType = "ctDress"
)

var (
    dressFactorySingleInstance = &DressFactory{
        dressMap: make(map[string]Dress),
    }
)

type DressFactory struct {
    dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
    if d.dressMap[dressType] != nil {
        return d.dressMap[dressType], nil
    }

    if dressType == TerroristDressType {
        d.dressMap[dressType] = newTerroristDress()
        return d.dressMap[dressType], nil
    }
    if dressType == CounterTerrroristDressType {
        d.dressMap[dressType] = newCounterTerroristDress()
        return d.dressMap[dressType], nil
    }

    return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
    return dressFactorySingleInstance
}
 dress.go: Flyweight interface
package main

type Dress interface {
    getColor() string
}
 terroristDress.go: Concrete flyweight object
package main

type TerroristDress struct {
    color string
}

func (t *TerroristDress) getColor() string {
    return t.color
}

func newTerroristDress() *TerroristDress {
    return &TerroristDress{color: "red"}
}
 counterTerroristDress.go: Concrete flyweight object
package main

type CounterTerroristDress struct {
    color string
}

func (c *CounterTerroristDress) getColor() string {
    return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
    return &CounterTerroristDress{color: "green"}
}
 player.go: Context
package main

type Player struct {
    dress      Dress
    playerType string
    lat        int
    long       int
}

func newPlayer(playerType, dressType string) *Player {
    dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
    return &Player{
        playerType: playerType,
        dress:      dress,
    }
}

func (p *Player) newLocation(lat, long int) {
    p.lat = lat
    p.long = long
}
 game.go: Client code
package main

type game struct {
    terrorists        []*Player
    counterTerrorists []*Player
}

func newGame() *game {
    return &game{
        terrorists:        make([]*Player, 1),
        counterTerrorists: make([]*Player, 1),
    }
}

func (c *game) addTerrorist(dressType string) {
    player := newPlayer("T", dressType)
    c.terrorists = append(c.terrorists, player)
    return
}

func (c *game) addCounterTerrorist(dressType string) {
    player := newPlayer("CT", dressType)
    c.counterTerrorists = append(c.counterTerrorists, player)
    return
}
 main.go: Client code
package main

import "fmt"

func main() {
    game := newGame()

    //Add Terrorist
    game.addTerrorist(TerroristDressType)
    game.addTerrorist(TerroristDressType)
    game.addTerrorist(TerroristDressType)
    game.addTerrorist(TerroristDressType)

    //Add CounterTerrorist
    game.addCounterTerrorist(CounterTerrroristDressType)
    game.addCounterTerrorist(CounterTerrroristDressType)
    game.addCounterTerrorist(CounterTerrroristDressType)

    dressFactoryInstance := getDressFactorySingleInstance()

    for dressType, dress := range dressFactoryInstance.dressMap {
        fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
    }
}
 output.txt: Execution result
DressColorType: ctDress
DressColor: green
DressColorType: tDress
DressColor: red





*/