package tests

import (
	"testing"
	"nli-go/lib/example3"
)

func TestSimpleGoalSpecification(test *testing.T) {

	// relations
	internalGrammarParser := example3.NewSimpleInternalGrammarParser()

	// who did Kurt Cobain marry?
	// Note: this representation is rubbish :) I will get to that later
	// non-domain specific
	genericSense, _, _ := internalGrammarParser.CreateRelationSet(`[
		predication(S1, marry)
		tense(S1, past)
		subject(S1, E1)
		object(S1, E2)
		info_request(E3)
		name(E2, 'Kurt Cobain')
	]`)

	// transform the generic sense into a domain specific sense. Leaving out material, but making it more compact
	domainSpecificAnalysis, _, _ := internalGrammarParser.CreateTransformations(`[
		married_to(A, B) :- predication(P1, marry), subject(P1, A), object(P1, B)
		name(A, N) :- name(A, N)
	]`)

	// create domain specific representation
	transformer := example3.NewSimpleRelationTransformer(domainSpecificAnalysis)
	domainSpecificSense := transformer.Extract(genericSense)

	// goal specification
	// if X was the request, Y is what the user really wants to know
	// turn X into Y
	// name ALL X for whom holds that X was married to Y
	// he / she was married to a, b, and c
	// name(B, N) will have multiple bindings for B and N
	// de operator :- is hier niet handig. Het gaat niet om een implicatie, maar om een interpretatie

	// je loopt er nu al tegenaan dat je makkelijk predicaten uit het verkeerde domein gebruikt. dat moet niet kunnen

	// date er direct een antwoord-template beschikbaar is, is omdat die eerder bedacht is; als er nog geen beschikbaar is,
	// moet die misschien bedacht worden; dat is een meta-probleem

	domainSpecificGoalAnalysis, _, _ := internalGrammarParser.CreateQAPairs(`[
		{
			Q: married_to(A, B), info_request(B)
			A: married_to(A, B), gender(A, G), name(B, N)
		}
	]`)

	// A: married_to(A, B), person(A, _, G, _), person(B, N, _, _)
	// 			RQ: B

	// optimalisatie-fase: doe eerst de meest gerestricteerde doelen (bv name(A, 'Kurt Cobain')

	transformer = example3.NewSimpleRelationTransformer(domainSpecificGoalAnalysis)
	goalSense := transformer.Extract(domainSpecificSense)

	rules, _, _ := internalGrammarParser.CreateRules(`
		married_to(X, Y) :- married_to(Y, X)
	`)
	ruleBase1 := example3.NewSimpleRuleBase(rules)

	ds2db, _, _ := internalGrammarParser.CreateRules(`[
		married_to(A, B) :- marriages(A, B)
		name(A, N) :- person(A, N, _, _)
		gender(A, male) :- person(A, _, 'm', _)
		gender(A, female) :- person(A, _, 'f', _)
	]`)

	// note! db specific
	facts, _, _ := internalGrammarParser.CreateRelationSet(`[
		marriages(14, 11, '1992')
		person(11, 'Kurt Cobain', 'M', '1967')
		person(14, 'Courtney Love', 'F', '1964')
	]`)
	factBase1 := example3.NewSimpleFactBase(facts, ds2db)

	// produce response
	problemSolver := example3.NewSimpleProblemSolver()
	problemSolver.AddKnowledgeBase(factBase1)
	problemSolver.AddKnowledgeBase(ruleBase1)
	domainSpecificResponseSense := problemSolver.solve(goalSense)

	//// turn domain specific response into generic response
	//specificResponseSpec, _, _ := internalGrammarParser.CreateTransformations(`[
	//	predication(S1, marry), object(S1, E2), subject(S1, A), object(S1, B) :- marry(A, B)
	//	name(A, N) :- name(A, N)
	//	gender(A, N) :- gender(A, N)
	//]`)
	//transformer = example3.NewSimpleRelationTransformer(specificResponseSpec)
	//genericResponseSense := transformer.Extract(domainSpecificResponseSense)
	//
	if len(domainSpecificResponseSense.GetRelations()) != 1 {
		test.Error("Wrong number of results")
	}

}