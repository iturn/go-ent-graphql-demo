// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (t *Todo) Children(ctx context.Context) ([]*Todo, error) {
	result, err := t.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = t.QueryChildren().All(ctx)
	}
	return result, err
}

func (t *Todo) Parent(ctx context.Context) (*Todo, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}