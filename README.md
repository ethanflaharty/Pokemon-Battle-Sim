# Pokemon CLI Battle Simulator

This is a Pokemon Battle Simulator that takes place in your CLI.

## Features 

- Pokemon's Battle Loop
- Uses the math from the Pokemon games based on the Bulbapedia information
- Can use whatever Pokemon you want from the start in the current state
- Features functional battle states such as weathers and terrains
- Features Status Conditions that work just like the ones in the original game

## How to Run

### Requirements

- Go 1.26.3 or later

### Running the Simulator

1. Clone the repository:

    ```bash
    git clone https://github.com/ethanflaharty/Pokemon-Battle-Sim.git

2. Navigate to the project directory

    ```bash
    cd Pokemon-Battle-Sim

3. Run the simulator using "go run ."

    ```bash
    go run .

## Project Structure

The structure:
- The main file/package uses all of the current package implementations
- The battle package within the battleSystem folder is the next in line so to say as it takes information from pokemondata and battledata
- Then we trickle down to pokemondata that takes info from battledata
- And finally for what is currently implemented battledata is the baseline that doesn't use any of the other packages apart of this project specifically
- The persistentData folder holds the JSON data that isn't currently implemented but is the part I'm working on implementing. This will hold the data of the Pokemon "world" that won't change, hench persistent

## How It Works
- After running "go run ." it will initiate the battle with the wild pokemon which in the current state is set to a Charmander and it will use your only pokemon Bulbasaur allowing you to choose your move. 
- When damage is taken it will tell you how much damage was done and how much HP is left
- Repeating that, once one of the pokemon faints it will print what pokemon fainted and end the program
- If you want to change the pokemon, moves or any of the details of your pokemon in the current state in the main.go file update either the wild or ally pokemon that is in the battle

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