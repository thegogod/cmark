package cmark

import (
	"slices"

	"github.com/thegogod/cmark/extensions/markdown"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/logging"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

var log = logging.Console("cmark")

type CMark struct {
	last       string
	extensions []Extension
}

func New(extensions ...Extension) *CMark {
	exists := slices.ContainsFunc(extensions, func(ext Extension) bool {
		return ext.Name() == "markdown"
	})

	if !exists {
		extensions = append(extensions, markdown.New())
	}

	return &CMark{"", extensions}
}

func (self *CMark) Parse(src []byte) (html.Node, error) {
	document := html.Fragment()
	ptr := tokens.Ptr(src)

	for {
		if ptr.Iter.Curr.Kind() == 0 {
			break
		}

		node, err := self.ParseBlock(ptr)

		if err != nil {
			return document, err
		}

		if node == nil {
			continue
		}

		document.Push(node)
	}

	return document, nil
}

func (self *CMark) ParseBlock(ptr *tokens.Pointer) (html.Node, error) {
	if ptr.Iter.Curr.Kind() == 0 {
		return nil, nil
	}

	tx := tx.New(ptr)

	for _, ext := range self.extensions {
		ptr.Ext = ext.Name()
		node, err := ext.ParseBlock(self, ptr)

		if err == nil {
			return node, err
		}

		tx.Rollback()
	}

	return nil, nil
}

func (self *CMark) ParseInline(ptr *tokens.Pointer) (html.Node, error) {
	if ptr.Iter.Curr.Kind() == 0 {
		return nil, nil
	}

	tx := tx.New(ptr)

	for _, ext := range self.extensions {
		ptr.Ext = ext.Name()
		node, err := ext.ParseInline(self, ptr)

		if node == nil && err == nil {
			return nil, nil
		}

		if node != nil && err == nil {
			return node, err
		}

		tx.Rollback()
	}

	return nil, nil
}
