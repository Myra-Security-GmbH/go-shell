package impl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
	"github.com/Myra-Security-GmbH/myra-shell/command"
	"github.com/Myra-Security-GmbH/myra-shell/context"
	"github.com/Myra-Security-GmbH/myra-shell/tag"
	"github.com/Myra-Security-GmbH/myra-shell/util"
)

//
// CreateCommand ...
//
type CreateCommand struct {
	*command.BaseCommand
}

//
// Completer ..,
//
func (cmd *CreateCommand) Completer(ctx context.Container, input string) []string {
	return []string{}
}

//
// Execute ...
//
func (cmd *CreateCommand) Execute(
	ctx context.Container,
	buffer *string,
) (uint, error) {
	var err error
	var line string
	dors := ctx.FindSelection(context.AreaDomain, context.AreaSubDomain)
	connector := ctx.GetAPIConnector()

	var t interface{}

	switch ctx.GetSelection() {
	case context.AreaDomain:
		t = &vo.DNSRecordVO{}

	case context.AreaRedirect:
		t = &vo.RedirectVO{ExpertMode: true}

	case context.AreaIPFilter:
		t = &vo.IPFilterVO{}

	case context.AreaCache:
		t = &vo.CacheSettingVO{}

	case context.AreaErrorPage:
		t = &vo.ErrorPageVO{}

	default:
		return 1, errors.New("Command is not available in this context")
	}

	typ := reflect.TypeOf(t).Elem()

	if typ.Kind() != reflect.Struct {
		return 1, errors.New("Invalid given data type")
	}

	fields := make([]string, typ.NumField())
	tags := make(map[string]*tag.Create)

	for i := 0; i < typ.NumField(); i++ {
		fieldName := typ.Field(i).Name
		tags[fieldName] = tag.GetCreateTag(typ.Field(i).Tag)

		if tags[fieldName] == nil || tags[fieldName].Ignore {
			continue
		}

		fields[tags[fieldName].Order] = fieldName
	}

	fmt.Printf("%+v\n", fields)
	fmt.Println(len(cmd.GetParams()))
	fmt.Println(len(fields))

	if len(cmd.GetParams()) != len(fields) {
		fmt.Print("Usage create")

		for _, fieldName := range fields {
			if strings.Trim(fieldName, " \n") == "" {
				continue
			}

			fmt.Print(" [", fieldName)

			if tags[fieldName].Type == tag.TagCreateFile {
				fmt.Print("(filepath)")
			}

			fmt.Print("]")
		}

		fmt.Println()
		return 1, nil
	}

	data := make(map[string]string)

	for argNum, fieldName := range fields {
		if fieldName == "" {
			continue
		}

		if cmd.IsFlagSet("interactive") {
			//c.readLine.SetPrompt(fieldName + "> ")
			//line, err = c.readLine.Readline()

			//if err != nil {
			//	output.Println(err)
			//	return 1
			//}
		} else {
			line = cmd.GetParameter(argNum, "")
		}

		data[fieldName] = line
	}

	buildEntityFromData(data, tags, &t)

	err = connector.SaveEntity(dors.Identifier(), t)

	if err != nil {
		return 1, err
	}

	return 0, nil
}

func buildEntityFromData(data map[string]string, metadata map[string]*tag.Create, entity *interface{}) {
	var err error

	val := reflect.ValueOf(*entity).Elem()

	for fieldName, value := range data {
		f := val.FieldByName(fieldName)

		if !f.CanSet() {
			continue
		}

		if metadata[fieldName].Type == tag.TagCreateFile {
			fp, err := os.Open(value)

			if err != nil {
				fmt.Println("Could not open file [", value, "]", err)
				return
			}

			fileValue, err := ioutil.ReadAll(fp)

			if err != nil {
				fmt.Println("Error reading file [", value, "]", err)
				return
			}

			value = string(fileValue)

			fmt.Println("filecontent:", value)
		}

		switch f.Kind() {
		case reflect.String:
			f.SetString(value)

		case reflect.Bool:
			var v bool
			v, err = strconv.ParseBool(value)
			if err != nil {
				fmt.Println("Could not parse given input assuming false")
				v = false
			}
			f.SetBool(v)

		case reflect.Int:
			var v int64
			v, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				fmt.Println("Could not parse given input assuming 0")
				v = 0
			}

			f.SetInt(v)

		case reflect.Uint:
			var v uint64
			v, err = strconv.ParseUint(value, 10, 64)
			if err != nil {
				fmt.Println("Could not parse given input assuming 0")
				v = 0
			}

			f.SetUint(v)
		}
	}
}

func init() {
	argumentDefinitions := []command.ArgumentDefintion{
		&command.GenericArgumentDefinition{
			Description: util.NewTextTemplate(""),
			Example:     "4.5*",
			Name:        "filter",
			Optional:    true,
		},
	}

	flagDefintions := []command.FlagDefintion{
		&command.GenericFlagDefinition{
			Description: util.NewTextTemplate("prompt before every removal"),
			Name:        "verbose",
			ShortName:   "v",
		},
	}

	cmd := &CreateCommand{
		BaseCommand: command.NewBaseCommand(
			command.CommandCreate,
			argumentDefinitions,
			flagDefintions,
		),
	}

	command.RegisterCommand(cmd)
}
