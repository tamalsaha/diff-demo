package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	apps "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/util/jsonmergepatch"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
)

var patchTypes = []string{"json", "merge", "strategic"}

func main() {
	var (
		src string
		dst string
		t   string
	)
	cmd := &cobra.Command{
		Use:               "diff",
		Short:             "Generate patch",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			srcYAML, err := ioutil.ReadFile(src)
			if err != nil {
				log.Fatalln(err)
			}
			srcJson, err := yaml.YAMLToJSON(srcYAML)
			if err != nil {
				log.Fatalln(err)
			}

			dstYAML, err := ioutil.ReadFile(dst)
			if err != nil {
				log.Fatalln(err)
			}
			dstJson, err := yaml.YAMLToJSON(dstYAML)
			if err != nil {
				log.Fatalln(err)
			}

			var patch []byte
			switch t {
			case "strategic":
				patch, err = strategicpatch.CreateTwoWayMergePatch(srcJson, dstJson, apps.Deployment{})
			case "merge":
				patch, err = jsonmergepatch.CreateThreeWayJSONMergePatch(srcJson, dstJson, srcJson)
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(patch))
		},
	}
	cmd.Flags().StringVarP(&src, "src", "s", src, "source file")
	cmd.Flags().StringVarP(&dst, "dst", "d", dst, "desired file")
	cmd.Flags().StringVarP(&t, "type", "t", "strategic", fmt.Sprintf("The type of patch being provided; one of %v", strings.Join(patchTypes, ",")))

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
