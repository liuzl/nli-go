package tests

import (
	"testing"
	"nli-go/lib/importer"
	"nli-go/lib/mentalese"
)

func TestRelationTransformer(t *testing.T) {

	parser := importer.NewInternalGrammarParser()

	// "name all customers"
	relationSet, _, _ := parser.CreateRelationSet(`[
		instance_of(E2, name)
		predicate(S1, name)
		object(S1, E1)
		instance_of(E1, customer)
		determiner(E1, D1)
		instance_of(D1, all)
	]`)

	tests := []struct {
		transformations string
		wantExtracted string
		wantReplaced string
		wantAppended string
	} {
		{
			`[
				task(A, B), subject(Y) :- predicate(A, X), object(A, Y), determiner(Y, Z), instance_of(Z, B)
				done() :- predicate(A, X), object(A, Y), determiner(Y, Z), instance_of(Z, B)
				magic(A, X) :- predicate(A, X), predicate(X, A)
			]`,
			"[task(S1, all) subject(E1) done()]", "[instance_of(E2, name) instance_of(E1, customer) task(S1, all) subject(E1) done()]",
			"[instance_of(E2, name) predicate(S1, name) object(S1, E1) instance_of(E1, customer) determiner(E1, D1) instance_of(D1, all)  task(S1, all) subject(E1) done()]",
		},
	}
	for _, test := range tests {

		transformations, _, _ := parser.CreateTransformations(test.transformations)

		transformer := mentalese.NewRelationTransformer(transformations)

		wantExtracted, _, _ := parser.CreateRelationSet(test.wantExtracted)
		wantReplaced, _, _ := parser.CreateRelationSet(test.wantReplaced)
		wantAppended, _, _ := parser.CreateRelationSet(test.wantAppended)

		extractedResult := transformer.Extract(relationSet)
		replacedResult := transformer.Replace(relationSet)
		appendedResult := transformer.Append(relationSet)

		if extractedResult.String() != wantExtracted.String() || replacedResult.String() != wantReplaced.String() || appendedResult.String() != wantAppended.String() {
			t.Errorf("RelationTransformer: got\n%v\n%v\n%v,\nwant\n%v\n%v\n%v", extractedResult, replacedResult, appendedResult, wantExtracted, wantReplaced, wantAppended)
		}
	}
}
