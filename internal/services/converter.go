package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHtml(md string) (ConvertedPost, error) {
	convertedPost := ConvertedPost{}
	fm := Frontmatter{}

	mdBody, err := frontmatter.MustParse(
		strings.NewReader(md),
		&fm,
	)
	if err != nil {
		return ConvertedPost{}, err
	}

	if err := fm.Validate(); err != nil {
		return ConvertedPost{}, err
	}

	convertedPost.Frontmatter, err = json.Marshal(&fm)
	if err != nil {
		return ConvertedPost{}, err
	}

	convertedPost.Title = fm.Title

	// Create markdown parser with extensions.
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	bodyAST := p.Parse(mdBody)

	test := bodyAST.GetChildren()
	for t := range test {
		fmt.Printf("%v", t)
	}

	// Create HTML renderer with extensions.
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	convertedPost.Body = string(markdown.Render(bodyAST, renderer))

	return convertedPost, nil
}

type ConvertedPost struct {
	Title       string
	Body        string
	Frontmatter []byte
}

type Frontmatter struct {
	Title       string   `yaml:"title" json:"title"`
	Description string   `yaml:"description" json:"description"`
	Tags        []string `yaml:"tags" json:"tags"`
}

func (fm Frontmatter) Validate() error {
	if strings.TrimSpace(fm.Title) == "" {
		return errors.New("frontmatter: title is required")
	}

	if strings.TrimSpace(fm.Description) == "" {
		return errors.New("frontmatter: description is required")
	}

	if len(fm.Tags) == 0 {
		return errors.New("frontmatter: at least one tag is required")
	}

	for _, tag := range fm.Tags {
		if strings.TrimSpace(tag) == "" {
			return errors.New(
				"frontmatter: tags must not contain empty values",
			)
		}
	}

	return nil
}
