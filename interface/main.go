package main

import "fmt"

type Formatter interface {
	Format() string
}

type plainText struct {
	text string
}

func (text plainText) Format() string {
	return text.text
}

type bold struct {
	text string
}

func (text bold) Format() string {
	return "**" + text.text + "**"
}

type code struct {
	text string
}

func (text code) Format() string {
	return "`" + text.text + "`"
}

type italic struct {
	text string
}

func (text italic) Format() string {
	return "_" + text.text + "_"
}

type chainFormatter struct {
	text *string
}

func (c *chainFormatter) AddFormatter(formatter Formatter) {
	*c.text = formatter.Format()
}

func main() {
	text := "hello world"

	plain := plainText{text: text}
	fmt.Println("plain text: ", plain.Format())

	boldText := bold{text: text}
	fmt.Println("bold text: ", boldText.Format())

	italicText := italic{text: text}
	fmt.Println("italic text: ", italicText.Format())

	codeText := code{text: text}
	fmt.Println("code text: ", codeText.Format())

	chain := chainFormatter{text: &text}
	chain.AddFormatter(bold{text: *chain.text})
	chain.AddFormatter(italic{text: *chain.text})
	chain.AddFormatter(code{text: *chain.text})
	fmt.Println("chain formatted text: ", text)
}
