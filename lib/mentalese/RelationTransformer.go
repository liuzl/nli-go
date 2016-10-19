package mentalese

import (
	"nli-go/lib/common"
)

type RelationTransformer struct {
	matcher         RelationMatcher
}

// using transformations
func NewRelationTransformer() *RelationTransformer {
	return &RelationTransformer{matcher: RelationMatcher{}}
}

// return the original relations, but replace the ones that have matched with their replacements
func (transformer *RelationTransformer) Replace(transformations []RelationTransformation, relationSet RelationSet) RelationSet {

	matchedIndexes, replacements := transformer.matchAllTransformations(transformations, relationSet)
	newRelations := RelationSet{}

	for i, oldRelation := range relationSet  {
		if !common.IntArrayContains(matchedIndexes, i) {
			newRelations = append(newRelations, oldRelation)
		}
	}

	newRelations = append(newRelations, replacements...)

	return newRelations
}

// Try to match all transformations to relationSet, and return the replacements that resulted from the transformations
func (transformer *RelationTransformer) Extract(transformations []RelationTransformation, relationSet RelationSet) (RelationSet) {

	_, replacements := transformer.matchAllTransformations(transformations, relationSet)

	return replacements
}

// only add the replacements to the original relations
func (transformer *RelationTransformer) Append(transformations []RelationTransformation, relationSet RelationSet) RelationSet {

	_, replacements := transformer.matchAllTransformations(transformations, relationSet)

	newRelations := RelationSet{}
	newRelations = append(newRelations, relationSet...)
	newRelations = append(newRelations, replacements...)

	return newRelations
}

// Attempts all transformations on all relations
// Returns the indexes of the matched relations, and the replacements that were created, each in a single set
func (transformer *RelationTransformer) matchAllTransformations(transformations []RelationTransformation, haystackSet RelationSet) ([]int, RelationSet){

	common.LogTree("matchAllTransformations", haystackSet)

	matchedIndexes := []int{}
	replacements := RelationSet{}

	for _, transformation := range transformations {

		// each transformation application is completely independent from the others
		bindings, newIndexes, match := transformer.matcher.MatchSequenceToSet(transformation.Pattern, haystackSet, Binding{})
		if match {
			matchedIndexes = append(matchedIndexes, newIndexes...)
			for _, binding := range bindings {
				replacements = append(replacements, transformer.createReplacements(transformation.Replacement, binding)...)
			}
		}
	}

	matchedIndexes = common.IntArrayDeduplicate(matchedIndexes)

	common.LogTree("matchAllTransformations", matchedIndexes, replacements)

	return matchedIndexes, replacements
}

func (transformer *RelationTransformer) createReplacements(relations RelationSet, bindings Binding) RelationSet {

	replacements := RelationSet{}

	common.LogTree("createReplacements", relations, bindings)

	for _, relation := range relations {

		newRelation := Relation{}
		newRelation.Predicate = relation.Predicate

		for _, argument := range relation.Arguments {

			arg := argument

			if argument.IsVariable() {
				value, found := bindings[argument.String()]
				if found {
					arg = value
				} else {
					// replacement could not be bound, leave variable unchanged
				}
			}

			newRelation.Arguments = append(newRelation.Arguments, arg)
		}

		replacements = append(replacements, newRelation)
	}

	common.LogTree("createReplacements", replacements)

	return replacements
}