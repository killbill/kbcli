package commands

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/killbill/kbcli/kbcmd/kblib"

	"github.com/killbill/kbcli/kbclient/tag_definition"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/cmdlib/args"
	"github.com/killbill/kbcli/kbmodel"
	"github.com/urfave/cli"
)

var (
	createTagDefinitionProperties args.Properties
)

var tagDefinitionFormatter = cmdlib.Formatter{
	Columns: []cmdlib.Column{
		{
			Name: "NAME",
			Path: "$.name",
		},
		{
			Name: "ID",
			Path: "$.id",
		},
		{
			Name: "IS_CONTROL",
			Path: "$.isControlTag",
		},
		{
			Name: "DESCRIPTION",
			Path: "$.description",
		},
	},
}

func listTagDefinitions(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 0 {
		return cmdlib.ErrorInvalidArgs
	}

	resp, err := o.Client().TagDefinition.GetTagDefinitions(ctx, &tag_definition.GetTagDefinitionsParams{})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)
	return nil
}

func getTagDefinition(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}
	idOrName := o.Args[0]
	td, err := kblib.GetTagDefinition(ctx, o.Client(), idOrName)
	if err != nil {
		return err
	}
	if td == nil {
		return fmt.Errorf("tag definition %s not found", idOrName)
	}
	o.Print(td)
	return nil
}

func createTagDefinition(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) < 3 {
		return cmdlib.ErrorInvalidArgs
	}
	name := o.Args[0]
	description := o.Args[1]
	var applicableTypes []string

	if o.Args[2] == "all" {
		applicableTypes = kblib.GetAllSupportedTagTargets()
	} else {
		applicableTypes = strings.Split(o.Args[2], ",")
		for _, at := range applicableTypes {
			if !kblib.IsValidTagTarget(at) {
				return fmt.Errorf("Invalid type: %s. Must be one of %s", at, strings.Join(kblib.GetAllSupportedTagTargets(), ","))
			}
		}
	}

	tda := &kbmodel.TagDefinition{}
	err := args.LoadProperties(tda, createTagDefinitionProperties, o.Args[3:])
	if err != nil {
		return err
	}

	t := &kbmodel.TagDefinition{
		Name:                  &name,
		Description:           &description,
		ApplicableObjectTypes: applicableTypes,
		IsControlTag:          tda.IsControlTag,
	}
	resp, err := o.Client().TagDefinition.CreateTagDefinition(ctx, &tag_definition.CreateTagDefinitionParams{
		Body: t,
		ProcessLocationHeader: true,
	})
	if err != nil {
		return err
	}
	o.Print(resp.Payload)
	return nil
}

func deleteTagDefinitions(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}
	idOrName := o.Args[0]
	td, err := kblib.GetTagDefinition(ctx, o.Client(), idOrName)
	if err != nil {
		return err
	}
	if td == nil {
		return fmt.Errorf("tag definition %s not found", idOrName)
	}

	_, err = o.Client().TagDefinition.DeleteTagDefinition(ctx, &tag_definition.DeleteTagDefinitionParams{
		TagDefinitionID: td.ID,
	})

	return err
}

func registerTagDefinitionCommands(r *cmdlib.App) {
	cmdlib.AddFormatter(reflect.TypeOf(&kbmodel.TagDefinition{}), tagDefinitionFormatter)
	createTagDefinitionProperties = args.GetProperties(&kbmodel.TagDefinition{})
	createTagDefinitionProperties.Get("IsControlTag").Default = "False"

	// Register top level command
	r.Register("", cli.Command{
		Name:  "tags",
		Usage: "tag related commands",
	}, nil)

	r.Register("tags", cli.Command{
		Name:  "list",
		Usage: "list tag definitions",
	}, listTagDefinitions)

	r.Register("tags", cli.Command{
		Name:      "get",
		Usage:     "get tag definition by name or id",
		ArgsUsage: "[NAME|ID]",
	}, getTagDefinition)

	r.Register("tags", cli.Command{
		Name:  "create",
		Usage: "create new tag definition",
		ArgsUsage: fmt.Sprintf(`NAME DESCRIPTION APPLICABLE_TYPES [IsControlTag=true/false]

   APPLICABLE_TYPES: comma separated values of tag targets. (use 'all' for all targets)
   Valid targets are:
   
   %s

   For e.g.,
     kbcmd tags create premium_account "Account has high transaction volume" ACCOUNT,INVOICE`,
			strings.Join(kblib.GetAllSupportedTagTargets(), "\n   ")),
	}, createTagDefinition)

	r.Register("tags", cli.Command{
		Name:  "delete",
		Usage: "delete tag definition",
		ArgsUsage: `[ID|NAME]
   For e.g.,
      kbcmd tags delete mytag
      (or)
      kbcmd tags delete d3cb5fcc-2004-47ee-ace3-ce7002c909af`,
	}, deleteTagDefinitions)
}
