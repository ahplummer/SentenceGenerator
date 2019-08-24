package structs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)
type Verb struct {
	Present string `json:"present"`
	Past string `json:"past"`
}
type JsonStruct struct {
	Donkeys []string `json:"donkeys"`
	Nouns []string `json:"nouns"`
	Adjectives []string `json:"adjs"`
	PersonalNouns []string `json:"personalNouns"`
	Verbs []Verb `json:"verbs"`
	Prepositions []string `json:"prepositions"`
	Animals []string `json:"animals"`
}
func (js *JsonStruct) LoadStruct(jsonFile string) error {
	file, _ := ioutil.ReadFile(jsonFile)
	err := json.Unmarshal([]byte(file), &js)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
func (js *JsonStruct) GetRandomDonkey() (string) {
	rand.Seed(time.Now().UnixNano())
	return js.Donkeys[rand.Intn(len(js.Donkeys))]
}
func (js *JsonStruct) GetRandomNoun() (string) {
	rand.Seed(time.Now().UnixNano())
	return js.Nouns[rand.Intn(len(js.Nouns))]
}
func (js *JsonStruct) GetRandomAdjective() (string) {
	rand.Seed(time.Now().UnixNano())
	return js.Adjectives[rand.Intn(len(js.Adjectives))]
}
func (js *JsonStruct) GetRandomPersonalNoun() (string) {
	rand.Seed(time.Now().UnixNano())
	return js.PersonalNouns[rand.Intn(len(js.PersonalNouns))]
}
func (js *JsonStruct) GetRandomPastTenseVerb() (string) {
	rand.Seed(time.Now().UnixNano())
	return js.Verbs[rand.Intn(len(js.Verbs))].Past
}
func (js *JsonStruct) GetRandomPreposition() (string) {
	rand.Seed(time.Now().UnixNano())
	return js.Prepositions[rand.Intn(len(js.Prepositions))]
}
func (js *JsonStruct) GetRandomAnimal() (string) {
	rand.Seed(time.Now().UnixNano())
	return js.Animals[rand.Intn(len(js.Animals))]
}