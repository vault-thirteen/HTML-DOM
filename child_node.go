package htmldom

import (
	"golang.org/x/net/html"
)

// GetChildNodeByTag searches for the nearest child node having the specified
// tag name.
func GetChildNodeByTag(
	startingPoint *html.Node,
	tagName string,
) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}

	for {
		if childNode.Data == tagName {
			return childNode
		} else {
			childNode = childNode.NextSibling

			if childNode == nil {
				return nil
			}
		}
	}
}

// GetChildNodeByTagAndClass searches for the nearest child node having the
// specified tag name and class.
func GetChildNodeByTagAndClass(
	startingPoint *html.Node,
	tagName string,
	className string,
) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}

	for {
		var classValue string
		var classExists bool
		classValue, classExists = GetNodeAttributeValue(childNode, AttributeClass)

		if (childNode.Data == tagName) &&
			(classExists) &&
			(classValue == className) {
			return childNode
		} else {
			childNode = childNode.NextSibling

			if childNode == nil {
				return nil
			}
		}
	}
}

// GetChildNodeByTagAndId searches for the nearest child node having the
// specified tag name and id.
func GetChildNodeByTagAndId(
	startingPoint *html.Node,
	tagName string,
	idName string,
) (childNode *html.Node) {
	if startingPoint == nil {
		return nil
	}

	childNode = startingPoint.FirstChild
	if childNode == nil {
		return nil
	}

	for {
		var idValue string
		var idExists bool
		idValue, idExists = GetNodeAttributeValue(childNode, AttributeId)

		if (childNode.Data == tagName) &&
			(idExists) &&
			(idValue == idName) {
			return childNode
		} else {
			childNode = childNode.NextSibling

			if childNode == nil {
				return nil
			}
		}
	}
}
