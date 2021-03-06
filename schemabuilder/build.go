package schemabuilder

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.appointy.com/jaal/graphql"
)

// schemaBuilder is a struct for holding all the graph information for types as
// we build out graphql types for our graphql schema.  Resolved graphQL "types"
// are stored in the type map which we can use to see sections of the graph.
type schemaBuilder struct {
	types        map[reflect.Type]graphql.Type
	objects      map[reflect.Type]*Object
	enumMappings map[reflect.Type]*EnumMapping
	typeCache    map[reflect.Type]cachedType // typeCache maps Go types to GraphQL datatypes
	inputObjects map[reflect.Type]*InputObject
}

// cachedType is a container for GraphQL datatype and the list of its fields
type cachedType struct {
	argType *graphql.InputObject
	fields  map[string]argField
}

// getType is the "core" function of the GraphQL schema builder.  It takes in a reflect type and builds the appropriate graphQL "type".
// This includes going through struct fields and attached object methods to generate the entire graphql graph of possible queries.
// This function will be called recursively for types as we go through the graph.
func (sb *schemaBuilder) getType(nodeType reflect.Type) (graphql.Type, error) {
	// Support scalars and optional scalars. Scalars have precedence over structs to have eg. time.Time function as a scalar.
	if typeName, values, ok := sb.getEnum(nodeType); ok {
		return &graphql.NonNull{Type: &graphql.Enum{Type: typeName, Values: values, ReverseMap: sb.enumMappings[nodeType].ReverseMap}}, nil
	}

	if typeName, ok := getScalar(nodeType); ok {
		return &graphql.NonNull{Type: &graphql.Scalar{Type: typeName}}, nil
	}
	if nodeType.Kind() == reflect.Ptr {
		if typeName, ok := getScalar(nodeType.Elem()); ok {
			return &graphql.Scalar{Type: typeName}, nil // XXX: prefix typ with "*"
		}
	}

	// Structs
	if nodeType.Kind() == reflect.Struct {
		if err := sb.buildStruct(nodeType); err != nil {
			return nil, err
		}
		return &graphql.NonNull{Type: sb.types[nodeType]}, nil
	}
	if nodeType.Kind() == reflect.Ptr && nodeType.Elem().Kind() == reflect.Struct {
		if err := sb.buildStruct(nodeType.Elem()); err != nil {
			return nil, err
		}
		return sb.types[nodeType.Elem()], nil
	}

	switch nodeType.Kind() {
	case reflect.Slice:
		elementType, err := sb.getType(nodeType.Elem())
		if err != nil {
			return nil, err
		}

		// Wrap all slice elements in NonNull.
		if _, ok := elementType.(*graphql.NonNull); !ok {
			elementType = &graphql.NonNull{Type: elementType}
		}

		return &graphql.NonNull{Type: &graphql.List{Type: elementType}}, nil

	default:
		return nil, fmt.Errorf("bad type %s: should be a scalar, slice, or struct type", nodeType)
	}
}

// getEnum gets the Enum type information for the passed in reflect.Type by looking it up in our enum mappings.
func (sb *schemaBuilder) getEnum(typ reflect.Type) (string, []string, bool) {
	if sb.enumMappings[typ] != nil {
		var values []string
		for mapping := range sb.enumMappings[typ].Map {
			values = append(values, mapping)
		}
		return typ.Name(), values, true
	}
	return "", nil, false
}

// getScalar grabs the appropriate scalar graphql field type name for the passed
// in variable reflect type.
func getScalar(typ reflect.Type) (string, bool) {
	for match, name := range scalars {
		if typesIdenticalOrScalarAliases(match, typ) {
			return name, true
		}
	}
	return "", false
}

var scalars = map[reflect.Type]string{
	reflect.TypeOf(bool(false)):                      "Boolean",
	reflect.TypeOf(int(0)):                           "Int",
	reflect.TypeOf(int8(0)):                          "Int",
	reflect.TypeOf(int16(0)):                         "Int",
	reflect.TypeOf(int32(0)):                         "Int",
	reflect.TypeOf(int64(0)):                         "Int",
	reflect.TypeOf(uint(0)):                          "Int",
	reflect.TypeOf(uint8(0)):                         "Int",
	reflect.TypeOf(uint16(0)):                        "Int",
	reflect.TypeOf(uint32(0)):                        "Int",
	reflect.TypeOf(uint64(0)):                        "Int",
	reflect.TypeOf(float32(0)):                       "Float",
	reflect.TypeOf(float64(0)):                       "Float",
	reflect.TypeOf(string("")):                       "String",
	reflect.TypeOf(ID{Value: ""}):                    "ID",
	reflect.TypeOf(Map{Value: ""}):                   "Map",
	reflect.TypeOf(Timestamp(timestamp.Timestamp{})): "Timestamp",
	reflect.TypeOf(Duration(duration.Duration{})):    "Duration",
	reflect.TypeOf(Bytes{Value: []byte{}}):           "Bytes",
}
