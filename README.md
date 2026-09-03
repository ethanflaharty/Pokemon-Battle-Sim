# Pokemon CLI Battle Simulator

This is a Pokemon Battle Simulator that takes place in your CLI, based off of the Generation 9 battle mechanics. 

## Motivations

I began building this while I was about 2/3 of the way through my studies on Boot.dev at a point where I was doubting how much I had actually learned so far. I figured this would be a good way to test my skills and have fun doing it because I am a huge Pokemon fan so I began building. After putting ~10 hours into this I decided to pause and continue it as my capstone project because I had really enjoyed the time I had put in at that time and knew I was going to need something for my capstone project and moved on having gotten many ideas of where to take this whilst doing it. I still plan on building this to improve the project as well as my skills and potentially get to a point where this could be an actual "fan game" type of project.

## Quick Start

### Requirements

- Go 1.26.3 or later

### Running the Simulator

1. Clone the repository:

```bash
git clone https://github.com/ethanflaharty/Pokemon-Battle-Sim.git
```

2. Navigate to the project directory

```bash
cd Pokemon-Battle-Sim
```

3. Run the simulator using "go run ."

```bash
go run .
```

## Usage
- After running the project using:

```bash
go run .
```

- Play the battle sim by inputing the number of the move you want to use each turn

## Contributing

### Clone the Repo

```bash
git clone https://github.com/ethanflaharty/Pokemon-Battle-Sim.git
```

### Build the compiled binary

```bash
go build
```

### Run the program

```bash
go run .
```

## Current Status
- The current version has the majority of parts of an individual pokemon implemented including:
    - Base Stats
    - Types
    - IVs
    - EVs
    - Natures
    - Battle Stats (aka the Stats shown in game)
- Also has Stat Changes, Status Conditions added tied to pokemon
- The damage calculation is the same as what Bulbapedia shows from the games, however it is within the calculateMove function that returns a MoveResult struct that allows it to have parts that can be booleans to make parts of the function cleaner in my opinion, for example the "Hit bool" allows it to check into more niche cases if it doesn hit and it can more easily be passed around
- Weather and terrains functions and changes are split around the packages
    - Processing happens in the battle package due to it tending to need both the pokemondata and battledata packages due to sandstorm and grassy terrain specifically damaging and healing the pokemon respectively
    - The damage changes/multipliers are calculated as a helper function in the damageResults.go alongside essentially every damage changing variables
    - They are initialized in the battleData package however, due to things like moves needing the BattleState in some cases for some status moves

## Future Improvements
- The next planned improvement is adding JSON of all of the persistent information as stated earlier so it'll be easier for future improvements by one having a list of what I might want to implement next, two having necesary information saved and unchanging so it pulls the info instead of relying on multiple different human inputs if others are using it. 
- Once that is done beginning to get abilities implemented to continue to build on the current pokemon structs that are built
- Adding teams and switching is a planned
- Getting to a point in the future where I can get it so everytime you run it, it chooses either a wild pokemon randomly or a trainer with randomly selected pokemon to make it more engaging
- After I began this I had a thought of eventually getting this to become a game within the CLI where you move through a map and traverse through a region like a normal pokemon game and battle through. This being my first major project I've realized how long this will take on my own and still plan on doing this at some point in some way, potentially making my own Pokemon "fan game" because that sounds more funny and would allow me to begin practicing long term development and game development in some way especially since I've been playing Heartgold and just looking at those sprites almost makes me want to make like an updated version of that with the current battle mechanics but I don't know when that might come to fruition.