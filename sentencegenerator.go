package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sentencegenerator/structs"
)

func GetArticle(word string) string{
	article := "a"
	switch word[0]{
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		article = "an"
	}
	return article
}
//func homePage(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//	fmt.Println("Endpoint Hit: homePage")
//}
func GetSentence(w http.ResponseWriter, r *http.Request){
	dir,_ := os.Getwd()
	js := structs.JsonStruct{}
	js.LoadStruct(dir + "/corpora/data/animals/donkeys.json")
	js.LoadStruct(dir + "/corpora/data/words/nouns.json")
	js.LoadStruct(dir + "/corpora/data/words/adjs.json")
	js.LoadStruct(dir + "/corpora/data/words/personal_nouns.json")
	js.LoadStruct(dir + "/corpora/data/words/verbs.json")
	js.LoadStruct(dir + "/corpora/data/words/prepositions.json")
	js.LoadStruct(dir + "/corpora/data/animals/common.json")
	//fmt.Printf("Length is %d", len(donkey.Donkeys))
	noun := js.GetRandomNoun()
	adjective := js.GetRandomAdjective()
	pnoun := js.GetRandomPersonalNoun()
	pastverb := js.GetRandomPastTenseVerb()
	preposition := js.GetRandomPreposition()
	adjective2 := js.GetRandomAdjective()
	animal := js.GetRandomAnimal()
	sentence := fmt.Sprintf("The %s %s %s %s %s %s the %s %s.", pnoun, pastverb, GetArticle(adjective), adjective, noun,
		preposition, adjective2, animal)
	fmt.Fprintf(w, sentence)
	fmt.Printf("Sentence Generated and returned: %s", sentence)

}

func handleRequests() {
	portPtr := flag.String("port", "", "Port to use")
	flag.Parse()
	if *portPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	http.HandleFunc("/", GetSentence)
	fmt.Printf("Started Webserver on port %s.", *portPtr)
	log.Fatal(http.ListenAndServe(":" + *portPtr, nil))
}
func main() {
	handleRequests()
}