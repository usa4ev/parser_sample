package handlers

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type company struct {
	Name string
	CEO  string
	INN  string
	KPP  string
}

func ScrapCompany(data io.Reader) (company, error) {
	cmp := company{}
	page, err := html.Parse(data)
	if err != nil {
		return cmp, fmt.Errorf("failed to parse data: %v", err)
	}

	main, err := getNodeMain(page)
	if err != nil {
		return cmp, err
	}

	scrapCompanyFields(main, &cmp)

	return cmp, nil
}

func scrapCEO(node *html.Node) (string, bool) {
	if text := execText(node); text != "Руководитель" {
		return "", false
	}

	var target *html.Node
	for target = node.NextSibling; target != nil; target = target.NextSibling {
		for _, attr := range target.Attr {
			if attr.Key == "class" && attr.Val == "company-info__text" &&
				target.FirstChild != nil && target.FirstChild.FirstChild != nil {

				if text := execText(target.FirstChild.FirstChild); text != "" {
					return text, true
				}
			}
		}
	}

	return "", false
}

func scrapInnKpp(node *html.Node) (string, string, bool) {
	if text := execText(node); text != "ИНН/КПП" {
		return "", "", false
	}

	var inn, kpp string
	for curNode := node.NextSibling; curNode != nil; curNode = curNode.NextSibling {
		if target := getChildByID(curNode, "clip_inn"); target != nil {
			if text := execText(target); text != "" {
				inn = text
				continue
			}
		}

		if target := getChildByID(curNode, "clip_kpp"); target != nil {
			if text := execText(target); text != "" {
				kpp = text
				continue
			}
		}
	}

	return inn, kpp, inn != "" && kpp != ""
}

func scrapName(node *html.Node) (string, bool) {
	if text := execText(node); text != "" {
		return text, true
	}

	return "", false
}

func scrapCompanyFields(node *html.Node, company *company) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		// scrap nodes
		for _, attr := range child.Attr {
			switch {
			case attr.Key == "class" && attr.Val == "company-info__title":
				if inn, kpp, ok := scrapInnKpp(child); ok {
					company.INN = inn
					company.KPP = kpp
				}

				if ceo, ok := scrapCEO(child); ok {
					company.CEO = ceo
				}
			case attr.Key == "class" && attr.Val == "company-name":
				if name, ok := scrapName(child); ok {
					company.Name = name
				}
			}

		}

		scrapCompanyFields(child, company)
	}
}

// getNodeMain finds node with id == main in body/main node
func getNodeMain(node *html.Node) (*html.Node, error) {
	htmlNode := node.FirstChild
	for {
		if htmlNode == nil || htmlNode.DataAtom == atom.Html {
			break
		}
		htmlNode = htmlNode.NextSibling
	}

	if htmlNode == nil {
		return nil, fmt.Errorf("failed to find htmlNode node attemting to find main node")
	}

	body := htmlNode.FirstChild
	for {
		if body == nil || body.DataAtom == atom.Body {
			break
		}
		body = body.NextSibling
	}

	if body == nil {
		return nil, fmt.Errorf("failed to find body node attemting to find main node")
	}

	wrapper := getChildByID(body, "wrapper")
	if wrapper == nil {
		return nil, fmt.Errorf("failed to find wrapper node attemting to find main node")
	}

	main := getChildByID(wrapper, "main")
	if main == nil {
		return nil, fmt.Errorf("failed to find main node")
	}

	return main, nil
}

// getChildByID returns first child with id == val for given node.
func getChildByID(node *html.Node, val string) *html.Node {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		for _, attr := range child.Attr {
			if attr.Key == "id" && attr.Val == val {
				return child
			}
		}
	}

	return nil
}

func execText(node *html.Node) string {
	if textNode := getTextNode(node); textNode != nil {
		return textNode.Data
	}

	return ""
}

func getTextNode(node *html.Node) *html.Node {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			return child
		}
	}

	return nil
}
