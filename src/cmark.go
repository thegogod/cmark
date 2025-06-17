package cmark

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/thegogod/cmark/extensions/markdown"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/logging"
	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

var log = logging.Console("cmark")

type CMark struct {
	extensions []Extension
}

func New(extensions ...Extension) *CMark {
	exists := slices.ContainsFunc(extensions, func(ext Extension) bool {
		return ext.Name() == "markdown"
	})

	if !exists {
		extensions = append(extensions, markdown.New())
	}

	return &CMark{extensions}
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

func (self *CMark) ParseDir(path string) ([]html.Node, error) {
	log.Debugln(path)
	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	nodes := []html.Node{}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			v, err := self.ParseDir(entryPath)

			if err != nil {
				return nil, err
			}

			nodes = append(nodes, v...)
		} else if strings.HasSuffix(entry.Name(), ".md") {
			v, err := self.ParseFile(entryPath)

			if err != nil {
				return nil, err
			}

			nodes = append(nodes, v)
		}
	}

	return nodes, nil
}

func (self *CMark) ParseFile(path string) (html.Node, error) {
	log.Debugln(path)
	src, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return self.Parse(src)
}

func (self *CMark) ParseBlock(ptr *tokens.Pointer) (html.Node, error) {
	if ptr.Iter.Curr.Kind() == 0 {
		return nil, nil
	}

	var node html.Node = nil
	var err error = nil

	tx := tx.New(ptr)

	for _, ext := range self.extensions {
		node, err = ext.ParseBlock(self, ptr)

		if node != nil && err == nil {
			break
		}

		tx.Rollback()
	}

	return node, err
}

func (self *CMark) ParseInline(ptr *tokens.Pointer) (html.Node, error) {
	if ptr.Iter.Curr.Kind() == 0 {
		return nil, nil
	}

	var node html.Node = nil
	var err error = nil

	tx := tx.New(ptr)

	for _, ext := range self.extensions {
		node, err = ext.ParseInline(self, ptr)

		if node != nil && err == nil {
			break
		}

		tx.Rollback()
	}

	return node, err
}
