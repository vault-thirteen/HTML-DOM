package htmldom

import (
	"testing"

	"github.com/vault-thirteen/tester"
	"golang.org/x/net/html"
)

func Test_GetSiblingNodeByTag(t *testing.T) {

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
	result = GetSiblingNodeByTag(nil, TagDiv)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Sibling Match.
	result = GetSiblingNodeByTag(divA, TagDiv)
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-b",
		},
		{
			Key: AttributeId,
			Val: "id-object-b",
		},
	})

	// Test #3. No Matches.
	result = GetSiblingNodeByTag(divA, TagB)
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Sibling.
	var tmp *html.Node
	tmp = GetSiblingNodeByTag(divA1, TagDiv) // -> divA2.
	aTest.MustBeDifferent(tmp, (*html.Node)(nil))
	aTest.MustBeEqual(tmp.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-2",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-2",
		},
	})
	tmp = GetSiblingNodeByTag(tmp, TagDiv) // -> divA3.
	aTest.MustBeDifferent(tmp, (*html.Node)(nil))
	aTest.MustBeEqual(tmp.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-3",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-3",
		},
	})
	result = GetSiblingNodeByTag(tmp, TagDiv)
	aTest.MustBeEqual(result, (*html.Node)(nil))
	// This last Iteration is required for the full Test Coverage.
	result = GetSiblingNodeByTag(tmp.NextSibling, TagDiv)
	aTest.MustBeEqual(result, (*html.Node)(nil))
}

func Test_GetSiblingNodeByTagAndClass(t *testing.T) {

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
	result = GetSiblingNodeByTagAndClass(nil, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Sibling Match.
	result = GetSiblingNodeByTagAndClass(divA, TagDiv, "class-object-b")
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-b",
		},
		{
			Key: AttributeId,
			Val: "id-object-b",
		},
	})

	// Test #3. No Matches.
	result = GetSiblingNodeByTagAndClass(divA, TagDiv, "class-object-fake")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Sibling.
	var tmp *html.Node
	tmp = GetSiblingNodeByTagAndClass(divA1, TagDiv, "class-object-a-2") // -> divA2.
	aTest.MustBeDifferent(tmp, (*html.Node)(nil))
	aTest.MustBeEqual(tmp.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-2",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-2",
		},
	})
	tmp = GetSiblingNodeByTagAndClass(tmp, TagDiv, "class-object-a-3") // -> divA3.
	aTest.MustBeDifferent(tmp, (*html.Node)(nil))
	aTest.MustBeEqual(tmp.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-3",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-3",
		},
	})
	result = GetSiblingNodeByTagAndClass(tmp, TagDiv, "any")
	aTest.MustBeEqual(result, (*html.Node)(nil))
	// This last Iteration is required for the full Test Coverage.
	result = GetSiblingNodeByTagAndClass(tmp.NextSibling, TagDiv, "any")
	aTest.MustBeEqual(result, (*html.Node)(nil))
}

func Test_GetSiblingNodeByTagAndId(t *testing.T) {

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
	result = GetSiblingNodeByTagAndId(nil, TagDiv, "")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #2. First Sibling Match.
	result = GetSiblingNodeByTagAndId(divA, TagDiv, "id-object-b")
	aTest.MustBeDifferent(result, (*html.Node)(nil))
	aTest.MustBeEqual(result.Data, TagDiv)
	aTest.MustBeEqual(result.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-b",
		},
		{
			Key: AttributeId,
			Val: "id-object-b",
		},
	})

	// Test #3. No Matches.
	result = GetSiblingNodeByTagAndId(divA, TagDiv, "id-object-fake")
	aTest.MustBeEqual(result, (*html.Node)(nil))

	// Test #4. Null Sibling.
	var tmp *html.Node
	tmp = GetSiblingNodeByTagAndId(divA1, TagDiv, "id-object-a-2") // -> divA2.
	aTest.MustBeDifferent(tmp, (*html.Node)(nil))
	aTest.MustBeEqual(tmp.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-2",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-2",
		},
	})
	tmp = GetSiblingNodeByTagAndId(tmp, TagDiv, "id-object-a-3") // -> divA3.
	aTest.MustBeDifferent(tmp, (*html.Node)(nil))
	aTest.MustBeEqual(tmp.Attr, []html.Attribute{
		{
			Key: AttributeClass,
			Val: "class-object-a-3",
		},
		{
			Key: AttributeId,
			Val: "id-object-a-3",
		},
	})
	result = GetSiblingNodeByTagAndId(tmp, TagDiv, "any")
	aTest.MustBeEqual(result, (*html.Node)(nil))
	// This last Iteration is required for the full Test Coverage.
	result = GetSiblingNodeByTagAndId(tmp.NextSibling, TagDiv, "any")
	aTest.MustBeEqual(result, (*html.Node)(nil))
}
