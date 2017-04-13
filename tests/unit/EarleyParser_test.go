package tests

import (
	"testing"
	"nli-go/lib/importer"
	"fmt"
	"nli-go/lib/parse"
	"nli-go/lib/parse/earley"
    "sort"
)

func TestEarleyParser(test *testing.T) {

	internalGrammarParser := importer.NewInternalGrammarParser()
	grammar := internalGrammarParser.CreateGrammar(`[
		rule: s(P) -> np(E) vp(P),			sense: subject(P, E);
		rule: np(E) -> nbar(E);
		rule: np(E) -> det(D) nbar(E),      sense: determiner(E, D);
		rule: nbar(E) -> noun(E);
		rule: nbar(E) -> adj(E) nbar(E);
		rule: vp(P) -> verb(P);
	]`)

	lexicon := internalGrammarParser.CreateLexicon(`[
		form: 'the',			pos: det,            sense: isa(E, the);
		form: 'a',  			pos: det;
		form: 'shy',			pos: adj;
		form: 'small',			pos: adj;
		form: 'boy',			pos: noun,			sense: isa(E, boy);
		form: 'girl',			pos: noun,			sense: isa(E, girl);
		form: 'cries',  		pos: verb,  		sense: predication(E, cry);
		form: 'sings',			pos: verb,			sense: predication(E, sing);
	]`)

	rawInput := "the small shy girl sings"
	tokenizer := parse.NewTokenizer()

	parser := earley.NewParser(grammar, lexicon)
	relationizer := earley.NewRelationizer(lexicon)

	wordArray := tokenizer.Process(rawInput)

	tree, _ := parser.Parse(wordArray)
	relations := relationizer.Relationize(tree)

	if relations.String() != "[subject(S5, E5) determiner(E5, D5) isa(D5, the) isa(E5, girl) predication(S5, sing)]" {
		test.Error(fmt.Sprintf("Relations: %v", relations))
	}
	if tree.String() != "[s [np [det the] [nbar [adj small] [nbar [adj shy] [nbar [noun girl]]]]] [vp [verb sings]]]" {
		test.Error(fmt.Sprintf("tree: %v", tree.String()))
	}
}

func TestEarleyParserSuggest(t *testing.T) {

    internalGrammarParser := importer.NewInternalGrammarParser()
    grammar := internalGrammarParser.CreateGrammar(`[
		rule: s(P) -> np(E) vp(P),			sense: subject(P, E);
		rule: np(E) -> nbar(E);
		rule: np(E) -> det(D) nbar(E),      sense: determiner(E, D);
		rule: nbar(E) -> noun(E);
		rule: nbar(E) -> adj(E) nbar(E);
		rule: vp(P) -> verb(P);
	]`)

    lexicon := internalGrammarParser.CreateLexicon(`[
		form: 'the',			pos: det,            sense: isa(E, the);
		form: 'a',  			pos: det;
		form: 'shy',			pos: adj;
		form: 'small',			pos: adj;
		form: 'boy',			pos: noun,			sense: isa(E, boy);
		form: 'girl',			pos: noun,			sense: isa(E, girl);
		form: 'cries',  		pos: verb,  		sense: predication(E, cry);
		form: 'sings',			pos: verb,			sense: predication(E, sing);
	]`)

    parser := earley.NewParser(grammar, lexicon)

    var tests = []struct {
        input []string
        want []string
    } {
        {[]string{}, []string{"a", "boy", "girl", "shy", "small", "the"} },
        {[]string{"the"}, []string{"boy", "girl", "shy", "small"} },
        {[]string{"the", "shy"}, []string{"boy", "girl", "shy", "small"} },
        {[]string{"the", "shy", "boy"}, []string{"cries", "sings"} },
        {[]string{"the", "shy", "boy", "sings"}, []string{} },
    }

    for _, test := range tests {

        suggests := parser.Suggest(test.input)
        sort.Strings(suggests)
        suggestsString := fmt.Sprintf("%v", suggests)
        wantString := fmt.Sprintf("%v", test.want)

        if suggestsString != wantString {
            t.Errorf("EarleyParserSuggest: got %v, want %v", suggestsString, wantString)
        }
    }
}