package kblib

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient"
	"github.com/killbill/kbcli/kbclient/tag_definition"
	"github.com/killbill/kbcli/kbmodel"
)

// GetTagDefinitions returns list of tag definitions
func GetTagDefinitions(ctx context.Context, c *kbclient.KillBill) (map[string]*kbmodel.TagDefinition, error) {
	resp, err := c.TagDefinition.GetTagDefinitions(ctx, &tag_definition.GetTagDefinitionsParams{})
	if err != nil {
		return nil, err
	}

	result := map[string]*kbmodel.TagDefinition{}
	for _, td := range resp.Payload {
		result[*td.Name] = td
	}

	return result, nil
}

// GetTagDefinition gets tag definition by name or id
func GetTagDefinition(ctx context.Context, c *kbclient.KillBill, idOrName string) (*kbmodel.TagDefinition, error) {
	var uuid strfmt.UUID
	keyOrID, isID := ParseKeyOrID(idOrName)
	if !isID {

		resp, err := c.TagDefinition.GetTagDefinitions(ctx, &tag_definition.GetTagDefinitionsParams{})
		if err != nil {
			return nil, err
		}
		for _, td := range resp.Payload {
			if *td.Name == idOrName {
				return td, nil
			}
		}
		return nil, nil
	} else {
		uuid = strfmt.UUID(keyOrID)
	}

	resp, err := c.TagDefinition.GetTagDefinition(ctx, &tag_definition.GetTagDefinitionParams{
		TagDefinitionID: uuid,
	})
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
