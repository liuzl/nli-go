package mentalese

import (
	"nli-go/lib/common"
	"sort"
)

// Turns a set of unscoped (or partially scoped) relations into scoped relations
// Like this
//
//     quant(S1, [ isa(S1, parent) ], D1, [ isa(D1, every) ], [
//         quant(O1, [ isa(O1, child) ], D2, [ isa(D2, 2) ], [
//             have_child(S1, O1)
//         ])
//     ])
//
// This part was mainly influenced by quantifier scoping of the Core Language Engine
// Specifically the article "Quantifier Scoping in the SRI Core Language Engine" by Douglas B. Moran

type QuantifierScoper struct {
	log *common.SystemLog
}

func NewQuantifierScoper(log *common.SystemLog) QuantifierScoper {
	return QuantifierScoper{log: log}
}

func (scoper QuantifierScoper) Scope(relations RelationSet) RelationSet {

	// collect all quantifications
	quantifications, nonQuantifications := scoper.collectQuantifications(relations)

	// sort the quantifications by hard constraints and preferences
	sort.Sort(QuantificationArray(quantifications))

	// nest the quantifications to create scopes
	scopedRelations := scoper.scopeQuantifications(quantifications)

	// fill in the other relations at the outermost position where they still are scoped.
	scoper.addNonQuantifications(&scopedRelations, len(quantifications), nonQuantifications)

	return scopedRelations
}

func (scoper QuantifierScoper) collectQuantifications(relations RelationSet) (QuantificationArray, RelationSet) {
	quantifications := QuantificationArray{}
	nonQuantifications := RelationSet{}
	for _, relation := range relations {
		if relation.Predicate == Predicate_Quantification {
			quantifications = append(quantifications, relation)
		} else {
			nonQuantifications = append(nonQuantifications, relation)
		}
	}
	return quantifications, nonQuantifications
}

func (scoper QuantifierScoper) scopeQuantifications(quantifications QuantificationArray) RelationSet {

	scope := RelationSet{}

	for i := len(quantifications) - 1; i >= 0; i-- {

		quantification := quantifications[i]

		quant := Relation{
			Predicate: Predicate_Quant,
			Arguments: []Term{
				quantification.Arguments[0],
				quantification.Arguments[1],
				quantification.Arguments[2],
				quantification.Arguments[3],
				{TermType: Term_relationSet, TermValueRelationSet: scope},
			},
		}

		scope = RelationSet{quant}
	}

	return scope
}

func (scoper QuantifierScoper) addNonQuantifications(scopedRelations *RelationSet, depth int, nonQuantifications RelationSet) {

	for _, nonQuantification := range nonQuantifications {

		scope := scopedRelations
		nonQuantificationScope := scope

		for d := 0; d < depth; d++ {

			scopedRelation := (*scope)[0]
			rangeVariable := scopedRelation.Arguments[Quantification_RangeVariableIndex]

			scope = &scopedRelation.Arguments[Quantification_ScopeIndex].TermValueRelationSet

			if scoper.variableMatches(nonQuantification, rangeVariable) {
				nonQuantificationScope = scope
			}
		}

		*nonQuantificationScope = append(*nonQuantificationScope, nonQuantification)
	}
}

func (scoper QuantifierScoper) variableMatches(relation Relation, variable Term) bool {
	match := false

	for _, argument := range relation.Arguments {
		if argument.Equals(variable) {
			match = true
			break
		}
	}

	return match
}
