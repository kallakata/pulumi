package test

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/stretchr/testify/assert"
)

type sdkTest struct {
	Directory   string
	Description string
	Skip        codegen.StringSet
}

var sdkTests = []sdkTest{
	{
		Directory:   "external-resource-schema",
		Description: "External resource schema",
	},
	{
		Directory:   "nested-module",
		Description: "Nested module",
	},
	{
		Directory:   "nested-module-thirdparty",
		Description: "Third-party nested module",
	},
	{
		Directory:   "plain-schema-gh6957",
		Description: "Repro for #6957",
	},
	{
		Directory:   "resource-args-python",
		Description: "Resource args with same named resource and type",
	},
	{
		Directory:   "simple-enum-schema",
		Description: "Simple schema with enum types",
	},
	{
		Directory:   "simple-plain-schema",
		Description: "Simple schema with plain properties",
	},
	{
		Directory:   "simple-plain-schema-with-root-package",
		Description: "Simple schema with root package set",
	},
	{
		Directory:   "simple-resource-schema",
		Description: "Simple schema with local resource properties",
	},
	{
		Directory:   "simple-resource-schema-custom-pypackage-name",
		Description: "Simple schema with local resource properties and custom Python package name",
	},
	{
		Directory:   "simple-methods-schema",
		Description: "Simple schema with methods",
		Skip:        codegen.NewStringSet("docs", "dotnet"),
	},
}

// TestSDKCodegen runs the complete set of SDK code generation tests against a particular language's code generator.
//
// An SDK code generation test consists of a schema and a set of expected outputs for each language. Each test is
// structured as a directory that contains that information:
//
//     test-directory/
//         schema.json
//         language-0
//         ...
//         language-n
//
// The schema is the only piece that must be manually authored. Once the schema has been written, the expected outputs
// can be generated by running `PULUMI_ACCEPT=true go test ./..." from the `pkg/codegen` directory.
func TestSDKCodegen(t *testing.T, language string, genPackage GenPkgSignature) {
	testDir := filepath.Join("..", "internal", "test", "testdata")

	for _, tt := range sdkTests {
		t.Run(tt.Description, func(t *testing.T) {
			if tt.Skip.Has(language) {
				t.Skip()
				return
			}

			files, err := GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.Directory, "schema.json"), genPackage)
			assert.NoError(t, err)

			dir := filepath.Join(testDir, tt.Directory)

			if RewriteFilesWhenPulumiAccept(t, dir, language, files) {
				return
			}

			expectedFiles, err := LoadBaseline(dir, language)
			assert.NoError(t, err)

			ValidateFileEquality(t, files, expectedFiles)
		})
	}
}
