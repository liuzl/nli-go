package parse

import (
	"nli-go/lib/mentalese"
	"regexp"
	"strings"
)

type Lexicon struct {
	allItems     []LexItem
	lexItems     map[string][]LexItem
	regexps      map[string][]LexItem
	senseBuilder SenseBuilder
}

func NewLexicon() *Lexicon {
	return &Lexicon{
		allItems:     []LexItem{},
		lexItems:     map[string][]LexItem{},
		regexps:      map[string][]LexItem{},
		senseBuilder: NewSenseBuilder(),
	}
}

func (lexicon *Lexicon) AddLexItem(lexItem LexItem) {

	form := lexItem.Form
	partOfSpeech := lexItem.PartOfSpeech

	lexicon.allItems = append(lexicon.allItems, lexItem)

	if lexItem.IsRegExp {

		_, found := lexicon.regexps[partOfSpeech]
		if !found {
			lexicon.regexps[partOfSpeech] = []LexItem{}
		}

		lexicon.regexps[partOfSpeech] = append(lexicon.regexps[partOfSpeech], lexItem)

	} else {

		_, found := lexicon.lexItems[form]
		if !found {
			lexicon.lexItems[form] = []LexItem{}
		}

		lexicon.lexItems[form] = append(lexicon.lexItems[form], lexItem)
	}
}

func (lexicon *Lexicon) GetLexItem(word string, partOfSpeech string) (LexItem, bool) {

	// try the word as is
	lexItems, found := lexicon.lexItems[word]

	// try the lowercase version
	if !found {
		lexItems, found = lexicon.lexItems[strings.ToLower(word)]
	}

	// try the regular expressions
	if !found {
		regExps, regExpFound := lexicon.regexps[partOfSpeech]
		if regExpFound {
			for _, regExpItem := range regExps {

				expression, _ := regexp.Compile(regExpItem.Form)
				if expression.FindString(word) != "" {

					termType := mentalese.Term_stringConstant
					if partOfSpeech == "number" {
						termType = mentalese.Term_number
					}

					from := mentalese.Term{TermType: mentalese.Term_variable, TermValue: "Form"}
					to := mentalese.Term{TermType: termType, TermValue: word}

					sense := lexicon.senseBuilder.ReplaceTerm(regExpItem.RelationTemplates, from, to)

					return LexItem{
						Form:              word,
						PartOfSpeech:      regExpItem.PartOfSpeech,
						RelationTemplates: sense}, true
				}
			}
		}
	}

	if found {
		for _, lexItem := range lexItems {
			if lexItem.PartOfSpeech == partOfSpeech {
				return lexItem, true
			}
		}
	}

	return LexItem{}, false
}

func (lexicon *Lexicon) GetWordForms(partOfSpeech string) []string {
	forms := []string{}

	for form, lexItems := range lexicon.lexItems {
		for _, lexItem := range lexItems {
			if lexItem.PartOfSpeech == partOfSpeech {
				forms = append(forms, form)
			}
		}
	}

	return forms
}

func (lexicon *Lexicon) ImportFrom(fromLexicon *Lexicon) {
	for _, lexItem := range fromLexicon.allItems {
		lexicon.AddLexItem(lexItem)
	}
}
