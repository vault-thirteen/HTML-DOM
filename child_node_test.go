package htmldom

import (
	"testing"

	"github.com/vault-thirteen/tester"
	"golang.org/x/net/html"
)

func Test_GetChildNodeByTag(t *testing.T) {

	// Preparation.
	var aTest = tester.New(t)
	var err error
	var containerNode *html.Node
	containerNode, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	var divA *html.Node
	divA = GetChildNodeByTag(containerNode, TagDiv)
	aTest.MustBeDifferent(divA, (*html.Node)(nil))
	var divA1 *html.Node
	divA1 = GetChildNodeByTag(divA, TagDiv)
	aTest.MustBeDifferent(divA1, (*html.Node)(nil))
	var result *html.Node

	// Test #1. Fool Check.
	result = GetChildNodeByTag(nil, TagDiv)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Child Match.
	result = GetChildNodeByTag(containerNode, TagDiv)
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a",
		},
		{
			Key: AttributeId,
			Val: "id-object-a",
		},
	})

	// Test #3. No Matches.
	result = GetChildNodeByTag(containerNode, TagB)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Child.
	result = GetChildNodeByTag(divA1, TagDiv)
	aTest.MustBeEqual(result, (*html.Node)(nil))
}

func Test_GetChildNodeByTagAndClass(t *testing.T) {

	// Preparation.
	var aTest = tester.New(t)
	var containerNode *html.Node
	var err error
	containerNode, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	var divA *html.Node
	divA = GetChildNodeByTagAndClass(containerNode, TagDiv, "class-object-a")
	aTest.MustBeDifferent(divA, (*html.Node)(nil))
	aTest.MustBeEqual(divA.Data, TagDiv)
	aTest.MustBeEqual(divA.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a",
		},
		{
			Key: AttributeId,
			Val: "id-object-a",
		},
	})
	var divA1 *html.Node
	divA1 = GetChildNodeByTagAndClass(divA, TagDiv, "class-object-a-1")
	aTest.MustBeDifferent(divA1, (*html.Node)(nil))
	var result *html.Node

	// Test #1. Fool Check.
	result = GetChildNodeByTagAndClass(nil, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Child Match.
	result = GetChildNodeByTagAndClass(divA, TagDiv, "class-object-a-3")
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-3",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-3",
		},
	})

	// Test #3. No Matches.
	result = GetChildNodeByTagAndClass(containerNode, TagB, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Child.
	result = GetChildNodeByTagAndClass(divA1, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))
}

func Test_GetChildNodeByTagAndId(t *testing.T) {

	// Preparation.
	var aTest = tester.New(t)
	var containerNode *html.Node
	var err error
	containerNode, err = _test_prepareTestNode()
	aTest.MustBeNoError(err)
	var divA *html.Node
	divA = GetChildNodeByTagAndId(containerNode, TagDiv, "id-object-a")
	aTest.MustBeDifferent(divA, (*html.Node)(nil))
	aTest.MustBeEqual(divA.Data, TagDiv)
	aTest.MustBeEqual(divA.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a",
		},
		{
			Key: AttributeId,
			Val: "id-object-a",
		},
	})
	var divA1 *html.Node
	divA1 = GetChildNodeByTagAndId(divA, TagDiv, "id-object-a-1")
	aTest.MustBeDifferent(divA1, (*html.Node)(nil))
	var result *html.Node

	// Test #1. Fool Check.
	result = GetChildNodeByTagAndId(nil, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Child Match.
	result = GetChildNodeByTagAndId(divA, TagDiv, "id-object-a-3")
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-3",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-3",
		},
	})

	// Test #3. No Matches.
	result = GetChildNodeByTagAndId(containerNode, TagB, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Child.
	result = GetChildNodeByTagAndId(divA1, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))
}
