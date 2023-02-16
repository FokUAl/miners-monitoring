package service

import (
	"fmt"
	"strings"

	"github.com/FokUAl/miners-monitoring/internal/repository"
	"github.com/FokUAl/miners-monitoring/pkg"
	"golang.org/x/net/html"
)

type InfoService struct {
	repo repository.Info
}

func NewInfoService(repo repository.Info) *InfoService {
	return &InfoService{
		repo: repo,
	}
}

func (s *InfoService) ParsingFile(path string) error {
	doc, err := html.Parse(strings.NewReader(path))
	if err != nil {
		return nil
	}

	tempNode := pkg.GetElementById(doc, "cbi")
	tempText := pkg.RenderNode(tempNode)

	fmt.Println(tempText)
	return nil
}
