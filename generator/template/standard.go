package template

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type HelperFuncs struct {
	globalParameters map[string]string
}

type ErrUndefinedGlobalParameter struct {
	targetParameterName string
}

func (e ErrUndefinedGlobalParameter) Error() string {
	return fmt.Sprintf("global parameter undefined: %s", e.targetParameterName)
}

var (
	errUndefinedGlobalParameter = ErrUndefinedGlobalParameter{}
)

func (h *HelperFuncs) RequestParameter() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimRight(text, "\n")
	return text, nil
}

func (h *HelperFuncs) RequestGlobalParameter(name string) (string, error) {
	if param, ok := h.globalParameters[name]; ok {
		return param, nil
	}
	errUndefinedGlobalParameter.targetParameterName = name
	return "", errUndefinedGlobalParameter
}

func (HelperFuncs) Println(text string) (string) {
	fmt.Println(text)
	return ""
}

func Processor(in io.Reader, out io.Writer, globalParams map[string]string) error {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	tmplFile := string(buf)

	helper := HelperFuncs{globalParameters: globalParams}

	funcMap := template.FuncMap{
		"RequestParameter": func(description string) (string, error) {
			fmt.Println(description)
			return helper.RequestParameter()
		},
		"SetGlobalParameter": func(name string) (string, error) {
			return helper.RequestGlobalParameter(name)
		},
		"Println": func(text string) string {
			return helper.Println(text)
		},
	}

	tmpl, err := template.New("tmpl").Funcs(funcMap).Parse(tmplFile)
	if err != nil {
		return err
	}
	var buff bytes.Buffer

	err = tmpl.Execute(&buff, &helper)
	if err != nil {
		return err
	}
	out.Write(buff.Bytes())

	return nil
}
