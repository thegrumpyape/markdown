package markdown

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/thegrumpyape/markdown/internal"
)

var valueMismatchError = "value mismatch (-want +got):\n%s"

func TestMarkdownParagraph(t *testing.T) {
	t.Parallel()

	t.Run("success_Paragraph()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.Paragraph("Text")
		want := []string{"Text"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownHeader(t *testing.T) {
	t.Parallel()

	t.Run("success_H1()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H1("Header")
		want := "# Header"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_H2()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H2("Header")
		want := "## Header"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_H3()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H3("Header")
		want := "### Header"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_H4()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H4("Header")
		want := "#### Header"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_H5()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H5("Header")
		want := "##### Header"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_H6()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.H6("Header")
		want := "###### Header"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownBulletList(t *testing.T) {
	t.Parallel()

	t.Run("success_BulletList()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.BulletList("Item1", "Item2", "Item3")
		want := []string{"- Item1", "- Item2", "- Item3"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownOrderedList(t *testing.T) {
	t.Parallel()

	t.Run("success_OrderedList()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.OrderedList("Item1", "Item2", "Item3")
		want := []string{"1. Item1", "2. Item2", "3. Item3"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownBlockquote(t *testing.T) {
	t.Parallel()

	t.Run("success_Blockquote()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.Blockquote("Quote")
		want := "> Quote"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownCodeBlock(t *testing.T) {
	t.Parallel()

	t.Run("success_CodeBlock()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.CodeBlock("test code", "go")
		want := []string{fmt.Sprintf("```go%stest code%s```", internal.LineFeed(), internal.LineFeed())}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownRule(t *testing.T) {
	t.Parallel()

	t.Run("success_Rule()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.Rule()
		want := []string{"---"}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownTable(t *testing.T) {
	t.Parallel()

	t.Run("success_Table()", func(t *testing.T) {
		t.Parallel()

		table := TableSet{header: []string{"Test", "Table"}, rows: [][]string{{"testing", "tables"}, {"testing", "tables"}}}
		m := NewMarkdown(os.Stdout)
		m.Table(table)

		want := "| Test | Table |" + internal.LineFeed() + "|---|---|" + internal.LineFeed() + "| testing | tables |" + internal.LineFeed() + "| testing | tables |"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownTaskList(t *testing.T) {
	t.Parallel()

	t.Run("success_TaskList()", func(t *testing.T) {
		t.Parallel()

		taskList := []TaskListSet{{Checked: true, Text: "checked"}, {Checked: false, Text: "unchecked"}}
		m := NewMarkdown(os.Stdout)
		m.TaskList(taskList)

		want := "- [x] checked" + internal.LineFeed() + "- [ ] unchecked"
		got := m.String()

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestMarkdownLF(t *testing.T) {
	t.Parallel()

	t.Run("success_LF()", func(t *testing.T) {
		t.Parallel()

		m := NewMarkdown(os.Stdout)
		m.LF()

		want := []string{"  "}
		got := m.body

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}

func TestSyntax(t *testing.T) {
	t.Parallel()

	t.Run("success_Link()", func(t *testing.T) {
		t.Parallel()

		want := "[Example](https://example.com)"
		got := Link("Example", "https://example.com")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_URL()", func(t *testing.T) {
		t.Parallel()

		want := "<https://example.com>"
		got := URL("https://example.com")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Image()", func(t *testing.T) {
		t.Parallel()

		want := "![Example](https://example.com)"
		got := Image("Example", "https://example.com")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Bold()", func(t *testing.T) {
		t.Parallel()

		want := "**Text**"
		got := Bold("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Italic()", func(t *testing.T) {
		t.Parallel()

		want := "*Text*"
		got := Italic("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_BoldItalic()", func(t *testing.T) {
		t.Parallel()

		want := "***Text***"
		got := BoldItalic("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Code()", func(t *testing.T) {
		t.Parallel()

		want := "`Text`"
		got := Code("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Highlight()", func(t *testing.T) {
		t.Parallel()

		want := "==Text=="
		got := Highlight("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Strikethrough()", func(t *testing.T) {
		t.Parallel()

		want := "~~Text~~"
		got := Strikethrough("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Subscript()", func(t *testing.T) {
		t.Parallel()

		want := "~Text~"
		got := Subscript("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Superscript()", func(t *testing.T) {
		t.Parallel()

		want := "^Text^"
		got := Superscript("Text")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})

	t.Run("success_Emoji()", func(t *testing.T) {
		t.Parallel()

		want := ":joy:"
		got := Emoji("joy")

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(valueMismatchError, diff)
		}
	})
}
