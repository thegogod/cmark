package html_test

import (
	"testing"

	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/maps"
)

func TestElement(t *testing.T) {
	scope := ast.NewScope()

	t.Run("should render", func(t *testing.T) {
		el := html.Elem("div")
		el.SetAttr("id", "123")
		el.SetStyles(maps.Pair("color", "red"), maps.Pair("display", "block"))
		html := string(el.Render(scope))

		if html != `<div id="123" style="color: red;display: block;"></div>` {
			t.Error(html)
		}
	})

	t.Run("should render pretty", func(t *testing.T) {
		el := html.Elem("div")
		el.SetAttr("id", "123")
		el.SetAttr("style", "color: red;display: block;")
		html := string(el.RenderPretty(scope, " "))

		if html != "<div id=\"123\" style=\"color: red;display: block;\">\n</div>" {
			t.Error(html)
		}
	})

	t.Run("should render with children", func(t *testing.T) {
		el := html.Elem("div")
		el.SetAttr("id", "123")
		el.AddClass("main")
		el.Push(
			html.P().WithAttr("id", "1").Push("hello world!"),
			html.Elem("input").WithAttr("value", "test").Void(),
		)

		html := string(el.Render(scope))

		if html != `<div id="123" class="main"><p id="1">hello world!</p><input value="test" /></div>` {
			t.Error(html)
		}
	})

	t.Run("should render with children pretty", func(t *testing.T) {
		el := html.Elem("div")
		el.SetAttr("id", "123")
		el.SetAttr("class", "main")
		el.Push(
			html.P().WithAttr("id", "1").Push("hello world!"),
			html.Elem("input").WithAttr("value", "test").Void(),
		)

		html := string(el.RenderPretty(scope, "\t"))

		if html != "<div id=\"123\" class=\"main\">\n\t<p id=\"1\">hello world!</p>\n\t<input value=\"test\" />\n</div>" {
			t.Error(html)
		}
	})

	t.Run("should render with children nested pretty", func(t *testing.T) {
		el := html.Elem("div")
		el.SetAttr("id", "123")
		el.SetAttr("class", "main")
		el.Push(
			html.P().WithAttr("id", "1").
				Push("hello world!").
				Push(html.Elem("span").Push("hi!")),
			html.Elem("input").WithAttr("value", "test").Void(),
		)

		html := string(el.RenderPretty(scope, "\t"))

		if html != "<div id=\"123\" class=\"main\">\n\t<p id=\"1\">\n\t\thello world!\n\t\t<span>hi!</span>\n\t</p>\n\t<input value=\"test\" />\n</div>" {
			t.Error(html)
		}
	})

	t.Run("select", func(t *testing.T) {
		t.Run("should select by id", func(t *testing.T) {
			el := html.Div(
				html.P(
					html.Span("test").WithId("123"),
				),
			)

			nodes := el.Select(html.HasAttr("id", "123"))

			if len(nodes) != 1 {
				t.Fatal(nodes)
			}
		})

		t.Run("should select by class", func(t *testing.T) {
			el := html.Div(
				html.P(
					html.Span("test").WithClass("main"),
				),
				html.Div().WithClass("main"),
			)

			nodes := el.Select(html.HasClass("main"))

			if len(nodes) != 2 {
				t.Fatal(nodes)
			}
		})

		t.Run("should select by tag", func(t *testing.T) {
			el := html.Div(
				html.P(
					html.Span("test").WithClass("main"),
				),
				html.Span().WithClass("main"),
			)

			nodes := el.Select(html.HasTag("span"))

			if len(nodes) != 2 {
				t.Fatal(nodes)
			}
		})

		t.Run("should select by id and class", func(t *testing.T) {
			el := html.Div(
				html.P(
					html.Img().WithSrc("https://www.google.com").WithId("123").WithClass("main"),
					html.Span("test").WithClass("main"),
				),
				html.Div().WithClass("main"),
			)

			nodes := el.Select(
				html.HasClass("main"),
				html.HasAttr("id", "123"),
			)

			if len(nodes) != 1 {
				t.Fatal(nodes)
			}
		})

		t.Run("should select by id or class", func(t *testing.T) {
			el := html.Div(
				html.P(
					html.Img().WithSrc("https://www.google.com").WithId("123"),
					html.Span("test").WithClass("main"),
				),
				html.Div().WithClass("main"),
			)

			nodes := el.Select(html.Or(
				html.HasClass("main"),
				html.HasAttr("id", "123"),
			))

			if len(nodes) != 3 {
				t.Fatal(nodes)
			}
		})

		t.Run("should select pseudo element", func(t *testing.T) {
			el := html.Div(
				html.P(
					html.Img().WithSrc("https://www.google.com").WithId("123"),
					html.MetaData().Set("a", 102.3).Set("b", "hi"),
					html.Span("test").WithClass("main"),
				),
				html.Div().WithClass("main"),
			)

			nodes := el.Select(html.HasTag(":metadata"))

			if len(nodes) != 1 {
				t.Fatal(nodes)
			}

			metadata := nodes[0].(html.MetaDataElement)

			if metadata.Get("a") != 102.3 || metadata.Get("b") != "hi" {
				t.Fatal(metadata)
			}
		})
	})
}
