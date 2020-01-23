package graphql_test

import (
	"context"
	"log"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"go.saastack.io/jaal/graphql"
	"go.saastack.io/jaal/internal"
	"go.saastack.io/jaal/schemabuilder"
)

type GatewayType int

const (
	GatewayType_Vehicle GatewayType = iota
	GatewayType_Asset
)

//TODO: Add a test case for testing union object nil error
func TestUnionType(t *testing.T) {
	type Vehicle struct {
		Name  string
		Speed int64
	}
	type Asset struct {
		Name         string
		BatteryLevel int64
	}

	type Gateway struct {
		schemabuilder.Union

		*Vehicle
		*Asset
	}

	schema := schemabuilder.NewSchema()
	query := schema.Query()
	schema.Enum(GatewayType(0), map[string]GatewayType{
		"vehicle": 0,
		"asset":   1,
	})

	query.FieldFunc("gateway", func(args struct{ Type GatewayType }) (*Gateway, error) {
		if args.Type == GatewayType_Vehicle {
			return &Gateway{
				Vehicle: &Vehicle{Name: "a", Speed: 50},
			}, nil
		}

		return &Gateway{
			Asset: &Asset{Name: "b", BatteryLevel: 5},
		}, nil
	})

	obj := schema.Object("Asset", Asset{})
	obj.FieldFunc("name", func(in Asset) string {
		return in.Name
	})
	obj.FieldFunc("batteryLevel", func(in Asset) int64 {
		return in.BatteryLevel
	})

	obj = schema.Object("Vehicle", Vehicle{})
	obj.FieldFunc("name", func(in Vehicle) string {
		return in.Name
	})
	obj.FieldFunc("speed", func(in Vehicle) int64 {
		return in.Speed
	})

	builtSchema := schema.MustBuild()

	ctx := context.Background()

	q, err := graphql.Parse(`
		{
			asset: gateway(type: "asset") { __typename ... on Asset { name batteryLevel } ... on Vehicle { name speed } }
			vehicle: gateway(type: "vehicle") { __typename ... on Asset { name batteryLevel } ... on Vehicle { name speed } }
		}
	`, map[string]interface{}{"var": float64(3)})

	if err := graphql.ValidateQuery(ctx, builtSchema.Query, q.SelectionSet); err != nil {
		t.Error(err)
	}

	e := graphql.Executor{}

	result, err := e.Execute(ctx, builtSchema.Query, nil, q)
	if err != nil {
		t.Error(err)
	}

	if d := pretty.Compare(internal.AsJSON(result), internal.ParseJSON(`
		{"vehicle": { "name": "a", "speed": 50, "__typename": "Vehicle" }, "asset": { "name": "b", "batteryLevel": 5, "__typename": "Asset" }}`)); d != "" {
		t.Errorf("expected did not match result: %s", d)
	}
}

type UnionPart1 struct{ OtherThing string }
type UnionPart2 struct{ Thing string }

type UnionMarkerPtrType struct {
	*schemabuilder.Union

	*UnionPart1
	*UnionPart2
}

// TestBadUnionMarkerPtr tests the union with pointer to schemabuilder.Union
func TestBadUnionMarkerPtr(t *testing.T) {
	schema := schemabuilder.NewSchema()
	query := schema.Query()
	query.FieldFunc("union", func() (*UnionMarkerPtrType, error) {
		return nil, nil
	})

	obj := schema.Object("UnionPart1", UnionPart1{})
	obj.FieldFunc("otherThing", func(in UnionPart1) string {
		return in.OtherThing
	})

	obj = schema.Object("UnionPart2", UnionPart2{})
	obj.FieldFunc("thing", func(in UnionPart2) string {
		return in.Thing
	})

	_, err := schema.Build()
	if err == nil {
		t.Fatalf("expected error, received nil")
	}
	if !strings.Contains(err.Error(), "bad method union on type schemabuilder.query: UnionMarkerPtrType not registered as object") {
		t.Errorf("expected error, received %s", err.Error())
	}
}

type UnionWithNonAnonymousPtrType struct {
	Something *schemabuilder.Union

	*UnionPart1
	*UnionPart2
}

func TestBadUnionNonAnonymousPtr(t *testing.T) {
	schema := schemabuilder.NewSchema()
	query := schema.Query()
	query.FieldFunc("union", func() (*UnionWithNonAnonymousPtrType, error) {
		return nil, nil
	})

	obj := schema.Object("UnionPart1", UnionPart1{})
	obj.FieldFunc("otherThing", func(in UnionPart1) string {
		return in.OtherThing
	})

	obj = schema.Object("UnionPart2", UnionPart2{})
	obj.FieldFunc("thing", func(in UnionPart2) string {
		return in.Thing
	})

	_, err := schema.Build()
	if err == nil {
		t.Fatalf("expected error, received nil")
	}

	if !strings.Contains(err.Error(), "bad method union on type schemabuilder.query: UnionWithNonAnonymousPtrType not registered as object") {
		t.Errorf("expected error, received %s", err.Error())
	}
}

type UnionNonAnonymousMembersType struct {
	schemabuilder.Union

	A *UnionPart1
	B *UnionPart2
}

func TestBadUnionNonAnonymousMembers(t *testing.T) {
	schema := schemabuilder.NewSchema()
	query := schema.Query()
	query.FieldFunc("union", func() (*UnionNonAnonymousMembersType, error) {
		return nil, nil
	})

	obj := schema.Object("UnionPart1", UnionPart1{})
	obj.FieldFunc("otherThing", func(in UnionPart1) string {
		return in.OtherThing
	})

	obj = schema.Object("UnionPart2", UnionPart2{})
	obj.FieldFunc("thing", func(in UnionPart2) string {
		return in.Thing
	})

	_, err := schema.Build()
	if err == nil {
		t.Fatalf("expected error, received nil")
	}

	if !strings.Contains(err.Error(), "union type member types must be anonymous") {
		t.Errorf("expected error, received %s", err.Error())
	}
}

func TestNonPointerOneHot(t *testing.T) {
	type UnionType struct {
		schemabuilder.Union

		UnionPart1
		UnionPart2
	}

	schema := schemabuilder.NewSchema()
	query := schema.Query()
	query.FieldFunc("union", func() (*UnionType, error) {
		return nil, nil
	})
	obj := schema.Object("UnionPart1", UnionPart1{})
	obj.FieldFunc("otherThing", func(in UnionPart1) string {
		return in.OtherThing
	})

	obj = schema.Object("UnionPart2", UnionPart2{})
	obj.FieldFunc("thing", func(in UnionPart2) string {
		return in.Thing
	})

	_, err := schema.Build()
	if err == nil {
		t.Fatalf("expected error, received nil")
	}

	if !strings.Contains(err.Error(), "union type member must be a pointer to a struct") {
		t.Errorf("expected error, received %s", err.Error())
	}
}

func TestBadUnionNonOneHot(t *testing.T) {
	type UnionType struct {
		schemabuilder.Union

		*UnionPart1
		*UnionPart2
	}

	schema := schemabuilder.NewSchema()
	query := schema.Query()
	query.FieldFunc("union", func() (*UnionType, error) {
		return &UnionType{UnionPart1: &UnionPart1{}, UnionPart2: &UnionPart2{}}, nil
	})

	obj := schema.Object("UnionPart1", UnionPart1{})
	obj.FieldFunc("otherThing", func(in UnionPart1) string {
		return in.OtherThing
	})

	obj = schema.Object("UnionPart2", UnionPart2{})
	obj.FieldFunc("thing", func(in UnionPart2) string {
		return in.Thing
	})

	builtSchema := schema.MustBuild()
	ctx := context.Background()

	q, err := graphql.Parse(`{ union { __typename } }`, map[string]interface{}{"var": float64(3)})
	if err != nil {
		panic(err)
	}

	if err := graphql.ValidateQuery(ctx, builtSchema.Query, q.SelectionSet); err != nil {
		t.Error(err)
	}

	t.Run("", func(t *testing.T) {
		e := graphql.Executor{}
		_, err := e.Execute(ctx, builtSchema.Query, nil, q)
		if err == nil {
			t.Error("expected err, received nil")
		}

		if !strings.Contains(err.Error(), "union type field should only return one value") {
			t.Errorf("expected err, received %s", err.Error())
		}
	})

}

func TestUnionList(t *testing.T) {
	type UnionType struct {
		schemabuilder.Union

		*UnionPart1
		*UnionPart2
	}

	schema := schemabuilder.NewSchema()
	query := schema.Query()
	query.FieldFunc("list", func() ([]*UnionType, error) {
		return []*UnionType{
			&UnionType{UnionPart2: &UnionPart2{"b"}},
			&UnionType{UnionPart1: &UnionPart1{"a"}},
		}, nil
	})

	obj := schema.Object("UnionPart1", UnionPart1{})
	obj.FieldFunc("otherThing", func(in UnionPart1) string {
		return in.OtherThing
	})

	obj = schema.Object("UnionPart2", UnionPart2{})
	obj.FieldFunc("thing", func(in UnionPart2) string {
		return in.Thing
	})

	builtSchema := schema.MustBuild()
	ctx := context.Background()

	q, err := graphql.Parse(`{ list { ... on UnionPart1 { otherThing } ... on UnionPart2 { thing } } }`, map[string]interface{}{"var": float64(3)})

	if err := graphql.ValidateQuery(ctx, builtSchema.Query, q.SelectionSet); err != nil {
		t.Error(err)
	}

	e := graphql.Executor{}
	result, err := e.Execute(ctx, builtSchema.Query, nil, q)
	if err != nil {
		t.Errorf("expected no error, received %s", err.Error())
	}

	log.Println(internal.AsJSON(result))

	if d := pretty.Compare(internal.AsJSON(result), internal.ParseJSON(`
		{ "list": [{"thing": "b"}, { "otherThing": "a" } ] }`)); d != "" {
		t.Errorf("expected did not match result: %s", d)
	}
}

func TestUnionStruct(t *testing.T) {
	type UnionType struct {
		schemabuilder.Union

		*UnionPart1
		*UnionPart2
	}

	type WrapperType struct {
		X *UnionType
	}

	schema := schemabuilder.NewSchema()
	query := schema.Query()
	query.FieldFunc("wrapper", func() (*WrapperType, error) {
		return &WrapperType{
			X: &UnionType{UnionPart2: &UnionPart2{"b"}},
		}, nil
	})

	obj := schema.Object("WrapperType", WrapperType{})
	obj.FieldFunc("x", func(in WrapperType) *UnionType {
		return in.X
	})

	obj = schema.Object("UnionPart1", UnionPart1{})
	obj.FieldFunc("otherThing", func(in UnionPart1) string {
		return in.OtherThing
	})

	obj = schema.Object("UnionPart2", UnionPart2{})
	obj.FieldFunc("thing", func(in UnionPart2) string {
		return in.Thing
	})

	builtSchema := schema.MustBuild()
	ctx := context.Background()

	q, err := graphql.Parse(`{ wrapper { x {... on UnionPart1 { otherThing } ... on UnionPart2 { thing } } } }`, map[string]interface{}{"var": float64(3)})

	if err := graphql.ValidateQuery(ctx, builtSchema.Query, q.SelectionSet); err != nil {
		t.Error(err)
	}

	e := graphql.Executor{}
	result, err := e.Execute(ctx, builtSchema.Query, nil, q)
	if err != nil {
		t.Errorf("expected no error, received %s", err.Error())
	}

	log.Println(internal.AsJSON(result))

	if d := pretty.Compare(internal.AsJSON(result), internal.ParseJSON(`
		{ "wrapper": { "x": { "thing": "b"} } }`)); d != "" {
		t.Errorf("expected did not match result: %s", d)
	}
}
