package kblib

import (
	"context"

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
