package markdown

import (
	"fmt"
	"io"
	"strings"

	"github.com/thegrumpyape/markdown/internal"
)

type Markdown struct {
	body []string
	dest io.Writer
	err  error
}

func NewMarkdown(w io.Writer) *Markdown {
	return &Markdown{
		body: []string{},
		dest: w,
	}
}

func (m *Markdown) Error() error {
	return m.err
}

func (m *Markdown) Build() error {
	if _, err := fmt.Fprint(m.dest, m.String()); err != nil {
		if m.err != nil {
			return fmt.Errorf("failed to write markdown text: %w: %s", err, m.err.Error())
		}
		return fmt.Errorf("failed to write markdown text: %v", err)
	}
	return m.err
}

func (m *Markdown) String() string {
	return strings.Join(m.body, internal.LineFeed())
}

func (m *Markdown) Paragraph(text string) *Markdown {
	m.body = append(m.body, text)
	return m
}

func (m *Markdown) header(text string, size int) *Markdown {
	sizeString := strings.Repeat("#", size)
	m.body = append(m.body, fmt.Sprintf("%s %s", sizeString, text))
	return m
}

func (m *Markdown) H1(text string) *Markdown {
	return m.header(text, 1)
}

func (m *Markdown) H2(text string) *Markdown {
	return m.header(text, 2)
}

func (m *Markdown) H3(text string) *Markdown {
	return m.header(text, 3)
}

func (m *Markdown) H4(text string) *Markdown {
	return m.header(text, 4)
}

func (m *Markdown) H5(text string) *Markdown {
	return m.header(text, 5)
}

func (m *Markdown) H6(text string) *Markdown {
	return m.header(text, 6)
}

func (m *Markdown) BulletList(text ...string) *Markdown {
	for _, v := range text {
		m.body = append(m.body, fmt.Sprintf("- %s", v))
	}
	return m
}

func (m *Markdown) OrderedList(text ...string) *Markdown {
	for i, v := range text {
		m.body = append(m.body, fmt.Sprintf("%d. %s", i+1, v))
	}
	return m
}

func (m *Markdown) Blockquote(text string) *Markdown {
	m.body = append(m.body, fmt.Sprintf("> %s", text))
	return m
}

func (m *Markdown) CodeBlock(text string, lang string) *Markdown {
	m.body = append(m.body, fmt.Sprintf("```%s%s%s%s```", lang, internal.LineFeed(), text, internal.LineFeed()))
	return m
}

func (m *Markdown) Rule() *Markdown {
	m.body = append(m.body, "---")
	return m
}

type TableSet struct {
	header []string
	rows   [][]string
}

func (t *TableSet) String() string {
	result := []string{}
	result = append(result, fmt.Sprintf("| %s |", strings.Join(t.header, " | ")))
	result = append(result, fmt.Sprintf("|%s", strings.Repeat("---|", len(t.header))))

	for _, row := range t.rows {
		result = append(result, fmt.Sprintf("| %s |", strings.Join(row, " | ")))
	}

	return strings.Join(result, internal.LineFeed())
}

func (m *Markdown) Table(t TableSet) *Markdown {
	m.body = append(m.body, t.String())
	return m
}

type TaskListSet struct {
	Checked bool
	Text    string
}

func (m *Markdown) TaskList(set []TaskListSet) *Markdown {
	for _, v := range set {
		if v.Checked {
			m.body = append(m.body, fmt.Sprintf("- [x] %s", v.Text))
		} else {
			m.body = append(m.body, fmt.Sprintf("- [ ] %s", v.Text))
		}
	}
	return m
}

func Link(text, url string) string {
	return fmt.Sprintf("[%s](%s)", text, url)
}

func URL(url string) string {
	return fmt.Sprintf("<%s>", url)
}

func Image(alt, url string) string {
	return fmt.Sprintf("![%s](%s)", alt, url)
}

func Bold(text string) string {
	return fmt.Sprintf("**%s**", text)
}

func Italic(text string) string {
	return fmt.Sprintf("*%s*", text)
}

func BoldItalic(text string) string {
	return fmt.Sprintf("***%s***", text)
}

func Code(text string) string {
	return fmt.Sprintf("`%s`", text)
}

func Highlight(text string) string {
	return fmt.Sprintf("==%s==", text)
}

func Strikethrough(text string) string {
	return fmt.Sprintf("~~%s~~", text)
}

func Subscript(text string) string {
	return fmt.Sprintf("~%s~", text)
}

func Superscript(text string) string {
	return fmt.Sprintf("^%s^", text)
}

func Emoji(shortcode string) string {
	return fmt.Sprintf(":%s:", shortcode)
}
