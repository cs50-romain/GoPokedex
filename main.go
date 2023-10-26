package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"	
	"net/http"
	"os"
	"strings"
	//"text/tabwriter"
)

type Pokemons struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID           int    `json:"id"`
	IsMainSeries bool   `json:"is_main_series"`
	Name         string `json:"name"`
	Names        []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEntries []struct {
		EntryNumber    int `json:"entry_number"`
		PokemonSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"pokemon_entries"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	VersionGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_groups"`
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
	} `json:"forms"`
	GameIndices []struct {
		Version   struct {
			Name string `json:"name"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
			} `json:"move_learn_method"`
			VersionGroup struct {
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []any  `json:"past_abilities"`
	PastTypes     []any  `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       string `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  string `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      string `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func getPokemonsData(location string) Pokemons{
	var pokemons Pokemons
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/" + location)

	if err != nil {
		fmt.Println(err)
		return pokemons 
	}

	respData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return pokemons
	}

	if err := json.Unmarshal(respData, &pokemons); err != nil {
		fmt.Println(err)
	}

	return pokemons
}

func parseCMD(str string) {
	str = strings.Trim(str, "\n")
	args := strings.Split(str, " ")
	cmd := args[0]

	if cmd == "exit" {
		fmt.Println("Gotta catch'em all! Goodbye Trainer...")
		os.Exit(0)
	} else {
		if cmd == "check" {
			arg := args[1]
			pokemon := getPokemonData(arg)
			displayPokemonData(pokemon)
		} else if cmd == "location" { 
			arg := args[1]
			pokemonsData := getPokemonsData(arg)
			displayPokemonsNameByLocation(pokemonsData)
		} else {
			fmt.Println("Invalid input")
		}
	}
}
	
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err:", err)
		}
		parseCMD(input)
	}
}

func displayPokemonsNameByLocation(pokemonsData Pokemons) {
	locationPokemonsArray := []string{}

	for _, pokemon := range pokemonsData.PokemonEntries {
		locationPokemonsArray = append(locationPokemonsArray, pokemon.PokemonSpecies.Name)
	}

	quicksort(locationPokemonsArray, 0, len(locationPokemonsArray)-1)

	newlineCounter := 0
	for _, pokemon := range locationPokemonsArray {
		fmt.Printf("%20s", pokemon)
		newlineCounter++
		if newlineCounter == 3 {
			fmt.Println()
			newlineCounter = 0
		}
	}
	fmt.Println()
}

func getPokemonData(pokemon string) Pokemon {
	var pokemonData Pokemon
	url := "http://pokeapi.co/api/v2/pokemon/" + pokemon
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("First err:", err)
		return pokemonData
	}
	
	respData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Second err:", err)
		return pokemonData
	}

	if err := json.Unmarshal(respData, &pokemonData); err != nil {
		fmt.Println("Third err:", err)
		return pokemonData
	}

	return pokemonData
}

func displayPokemonData(pokemon Pokemon){
	fmt.Println(pokemon.Name)
	fmt.Println(pokemon.Height * 10, "cm")
	fmt.Println(pokemon.Weight / 10, "kg")

	if len(pokemon.Types) > 0 {
		fmt.Print("Types: ")
		for i := 0; i < len(pokemon.Types); i++ {
			fmt.Print(pokemon.Types[i].Type.Name + "; ")
		}
	}

	fmt.Println("\nStats:")
	for i := 0; i < len(pokemon.Stats); i++ {
		fmt.Printf("\t- %s:%d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}
}

func displayPokemonMoves(pokemonData Pokemon) {
	pokemonDataArray := []string{}

	for _, move := range pokemonData.Moves {
		pokemonDataArray = append(pokemonDataArray, move.Move.Name)
	}

	quicksort(pokemonDataArray, 0, len(pokemonDataArray)-1)
	
	for _, move := range pokemonDataArray {
		fmt.Println(move)
	}
}

func displayPokemons(pokemonArray []string) {
	for _, pokemon := range pokemonArray {
		fmt.Printf("Pokemon: %s\n", pokemon)
	}
}

func quicksort(array []string, low int, high int) {
	if low < high {
		pivot := partition(array, low, high)
		quicksort(array, low, pivot-1)
		quicksort(array, pivot+1, high)
	}
}

func partition(array []string, low int, high int) int {
	x := array[high]
	i := low

	for j := low; j < high; j++ {
		if array[j] <= x {
			array[j], array[i] = array[i], array[j]
			i++
		}
	}

	array[i], array[high] = array[high], array[i]
	return i
}
