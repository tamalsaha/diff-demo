package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	apps "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
)

func main() {
	var (
		src string
		dst string
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

			patch, err := strategicpatch.CreateTwoWayMergePatch(srcJson, dstJson, apps.Deployment{})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(patch))
		},
	}
	cmd.Flags().StringVarP(&src, "src", "s", src, "source file")
	cmd.Flags().StringVarP(&dst, "dst", "d", dst, "desired file")

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
