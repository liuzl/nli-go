package tests

import (
	"testing"
	//"nli-go/lib/importer"
	//"nli-go/lib/central"
	//"nli-go/lib/knowledge"
	//"nli-go/lib/parse"
	//"nli-go/lib/mentalese"
	//"nli-go/lib/generate"
	//"nli-go/lib/parse/earley"
)

// Mimics some of SHRDLU's functions, but in the cli-go way

// NOTE: tests not working!

func TestBlocksWorld(t *testing.T) {

	//	internalGrammarParser := importer.NewInternalGrammarParser()
	//	matcher := mentalese.NewRelationMatcher()
	//	transformer := mentalese.NewRelationTransformer(matcher)
	//
	//	grammar := internalGrammarParser.CreateGrammar(`[
	//		rule: s(P) -> np(E) vp(P),                         sense: subject(P, E);
	//		rule: s(P1) -> auxDo(X) np(E1) verb(P1) np(E2),    sense: subject(P1, E1) object(P1, E2);
	//		rule: np(E) -> nbar(E);
	//		rule: np(E) -> det(E) nbar(E);
	//		rule: nbar(E) -> noun(E);
	//		rule: nbar(E) -> adj(E) nbar(E);
	//		rule: vp(P) -> verb(P);
	//	]`)
	//
	//	lexicon := internalGrammarParser.CreateLexicon(`[
	//		form: 'does',          pos: auxDo;
	//		form: 'the',           pos: det;
	//		form: 'red',           pos: adj,        sense: instance_of(E, red);
	//		form: 'blue',          pos: adj,        sense: instance_of(E, blue);
	//		form: 'block',         pos: noun,       sense: instance_of(E, block);
	//		form: 'support',       pos: verb,       sense: predication(E, support);
	//	]`)
	//
	//	parser := earley.NewParser(grammar, lexicon)
	//	relationizer := earley.NewRelationizer(lexicon)
	//
	//	genericSense2domainSpecificSense := internalGrammarParser.CreateTransformations(`[
	//		predication(P, support) subject(P, A) object(P, B) => support(A, B);
	//		instance_of(A, block) => is(A, block);
	//		instance_of(A, red) => color(A, red);
	//		instance_of(A, blue) => color(A, blue);
	//	]`)
	//
	//	question2answer := internalGrammarParser.CreateTransformations(`[
	//		support(A, B) => answer(yes);
	//	]`)
	//
	//
	//	// domain specific sense to database relations
	//	domainSpecific2database := internalGrammarParser.CreateDbMappings(`[
	//		isa(A, B) ->> is(A, B);
	//		support(A, B) ->> support(A, B);
	//		color(A, B) ->> color(A, B);
	//	]`)
	//
	//	// domain knowledge
	//	databaseRelations := internalGrammarParser.CreateRelationSet(`[
	//		is(block1, block)
	//		is(block2, block)
	//		support(block1, block2)
	//		color(block1, red)
	//		color(block2, blue)
	//	]`)
	//
	//	factBase1 := knowledge.NewInMemoryFactBase(databaseRelations, domainSpecific2database)
	//
	//	problemSolver := central.NewProblemSolver(matcher)
	//	problemSolver.AddFactBase(factBase1)
	//
	//	domainSpecific2generic := internalGrammarParser.CreateTransformations(`[
	//		answer(yes) => adverb(P1, yes);
	//	]`)
	//
	//	generationGrammar := internalGrammarParser.CreateGenerationGrammar(`[
	//		rule: s(P) -> adverb(P),    condition: adverb(P);
	//	]`)
	//
	//	generationLexicon := internalGrammarParser.CreateGenerationLexicon(`[
	//		form: 'yes',                pos: adverb,         condition: adverb(E, yes);
	//	]`)
	//
	//	generator := generate.NewGenerator(generationGrammar, generationLexicon)
	//
	//	surfaceRepresentation := generate.NewSurfaceRepresentation()
	//
	//	// -----------------------
	//
	////common.LoggerActive = true
	//
	//	tests := []struct {
	//		input string
	//		want string
	//	} {
	//		{"Does the red block support the blue block?", "Yes"},
	//	}
	//
	//	for _, test := range tests {
	//
	//		rawInput := test.input
	//		tokenizer := parse.NewTokenizer()
	//		wordArray := tokenizer.Process(rawInput)
	//		parseTree, _ := parser.Parse(wordArray)
	//		genericSense := relationizer.Relationize(parseTree)
	//		dsSense := transformer.Extract(genericSense2domainSpecificSense, genericSense)
	//
	//		questionSense := transformer.Extract(question2answer, dsSense)
	//
	//		domainSpecificAnswerSense := problemSolver.Solve(questionSense)
	//
	//		if (len(domainSpecificAnswerSense) == 0) {
	////			t.Errorf("Blocks World: expected %s, got %v", test.want, domainSpecificAnswerSense)
	//			continue
	//		}
	//
	//		genericAnswerSense := transformer.Extract(domainSpecific2generic, domainSpecificAnswerSense[0]);
	//
	//		answerWords := generator.Generate(genericAnswerSense)
	//
	//		result := surfaceRepresentation.Create(answerWords)
	//
	//		if result != test.want {
	//			t.Errorf("%s: got %s, want %s", test.input, result, test.want)
	//		}
	//	}
}
