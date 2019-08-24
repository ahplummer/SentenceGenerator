copy: build
	cp ./corpora/data/animals/donkeys.json .
	cp ./corpora/data/words/nouns.json .
	cp ./corpora/data/words/adjs.json .
	cp ./corpora/data/words/personal_nouns.json .
	cp ./corpora/data/words/verbs.json .
	cp ./corpora/data/words/prepositions.json .
	cp ./corpora/data/animals/common.json .

build:
	go build sentencegenerator.go