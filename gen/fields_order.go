package gen

import (
	"errors"
	"os"

	"github.com/ogen-go/ogen/json"
	"github.com/wheissd/mkgo/internal/entity"
)

type fieldsOrder map[string]map[string]int

func updateFieldsOrder(sch *entity.Schema) error {
	fName := "mkgo/fields_order_gen.json"
	edgeFName := "mkgo/edge_order_gen.json"
	fOrderFile, err := os.ReadFile(fName)
	parseFOrder, parseEOrder := true, true
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		parseFOrder = false
	}
	edgesFile, err := os.ReadFile(edgeFName)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		parseEOrder = false
	}
	fOrder, edgeOrder := make(fieldsOrder), make(fieldsOrder)
	if parseFOrder {
		err = json.Unmarshal(fOrderFile, &fOrder)
		if err != nil {
			return err
		}
	}
	if parseEOrder {
		err = json.Unmarshal(edgesFile, &edgeOrder)
		if err != nil {
			return err
		}
	}
	for entityKey, ent := range sch.Entities {
		if _, ok := fOrder[ent.Name]; !ok {
			fOrder[ent.Name] = make(map[string]int)
		}
		for fieldKey, field := range ent.Fields {
			if _, ok := fOrder[ent.Name][field.Name]; !ok {
				fOrder[ent.Name][field.Name] = len(fOrder[ent.Name]) + 1
			}
			sch.Entities[entityKey].Fields[fieldKey].Order = fOrder[ent.Name][field.Name]
		}
		if _, ok := edgeOrder[ent.Name]; !ok {
			edgeOrder[ent.Name] = make(map[string]int)
		}
		for edgeI, edge := range sch.Entities[entityKey].Edges {
			if _, ok := edgeOrder[ent.Name][edge.Name]; !ok {
				edgeOrder[ent.Name][edge.Name] = len(edgeOrder[ent.Name]) + 1
			}
			sch.Entities[entityKey].Edges[edgeI].Order = edgeOrder[ent.Name][edge.Name]
		}
	}

	fOrderBytes, err := json.Marshal(fOrder)
	if err != nil {
		return err
	}
	fErr := os.WriteFile(fName, fOrderBytes, 0744)
	if fErr != nil {
		return fErr
	}

	edgeOrderBytes, err := json.Marshal(edgeOrder)
	if err != nil {
		return err
	}
	edgeErr := os.WriteFile(edgeFName, edgeOrderBytes, 0744)
	if edgeErr != nil {
		return edgeErr
	}
	return nil
}
