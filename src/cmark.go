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

type CMark struct {
	logger     logging.Logger
	extensions []Extension
}

func New(extensions ...Extension) *CMark {
	exists := slices.ContainsFunc(extensions, func(ext Extension) bool {
		return ext.Name() == "markdown"
	})

	if !exists {
		extensions = append(extensions, markdown.New())
	}

	return &CMark{nil, extensions}
}

func (self *CMark) WithLogging(logger logging.Logger) *CMark {
	self.logger = logger
	return self
}

func (self *CMark) Parse(src []byte) (html.Node, error) {
	document := html.Fragment()
	ptr := tokens.Ptr(src)
	ptr.Next()

	for {
		if ptr.Eof() {
			break
		}

		node, err := self.ParseBlock(ptr)

		if err != nil {
			return document, err
		}

		if node == nil {
			continue
		}

		htmlNode, ok := node.(html.Node)

		if !ok {
			continue
		}

		document.Push(htmlNode)
	}

	return document, nil
}

func (self *CMark) ParseDir(path string) ([]html.Node, error) {
	self.println(path)
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
	self.println(path)
	src, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return self.Parse(src)
}

func (self *CMark) ParseBlock(ptr *tokens.Pointer) (html.Node, error) {
	if ptr.Eof() {
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
	if ptr.Eof() {
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

func (self *CMark) ParseSyntax(ptr *tokens.Pointer, extension string, name string) (html.Node, error) {
	self.printf("parsing '%s->%s'...", extension, name)
	self.print("done\n")
	return nil, nil
}

func (self CMark) print(message string) {
	if self.logger == nil {
		return
	}

	self.logger.Debug(message)
}

func (self CMark) printf(format string, args ...any) {
	if self.logger == nil {
		return
	}

	self.logger.Debugf(format, args...)
}

func (self CMark) println(message string) {
	if self.logger == nil {
		return
	}

	self.logger.Debugln(message)
}
