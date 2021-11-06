package templates

import (
	_ "embed"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/csrf"
)

var (
	//go:embed login.html
	loginHtml string
	//go:embed welcome.html
	welcomeHtml string
	//go:embed consent.html
	consentHtml string
)

var (
	loginHtmlTemplate   = template.Must(template.New("html").Parse(loginHtml))
	welcomeHtmlTemplate = template.Must(template.New("html").Parse(welcomeHtml))
	consentHtmlTemplate = template.Must(template.New("html").Parse(consentHtml))
)

type (
	consentModel struct {
		Scopes    []string
		CsrfField template.HTML
	}
)

func LoginHtml(w io.Writer, r *http.Request) error {
	if err := loginHtmlTemplate.Execute(w, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	}); err != nil {
		return err
	}
	return nil
}

func WelcomeHtml(w io.Writer, userName string) error {
	if err := welcomeHtmlTemplate.Execute(w, userName); err != nil {
		return err
	}
	return nil
}

func ConsentHtml(w io.Writer, r *http.Request, scopes []string) error {
	model := consentModel{
		CsrfField: csrf.TemplateField(r),
		Scopes:    scopes,
	}
	if err := consentHtmlTemplate.Execute(w, model); err != nil {
		return err
	}
	return nil
}
