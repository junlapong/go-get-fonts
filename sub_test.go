package main

import (
	"testing"
)

const cssWoff = `
@font-face {
	font-family: 'Sarabun';
	font-style: italic;
	font-weight: 400;
	src: local('Sarabun Italic'), local('Sarabun-Italic'), url(https://fonts.gstatic.com/s/sarabun/v7/DtVhJx26TKEr37c9aBB5nXwP.woff) format('woff');
 }
  @font-face {
	font-family: 'Sarabun';
	font-style: normal;
	font-weight: 200;
	src: local('Sarabun ExtraLight'), local('Sarabun-ExtraLight'), url(https://fonts.gstatic.com/s/sarabun/v7/DtVmJx26TKEr37c9YNpoilss7Q.woff) format('woff');
 }
  @font-face {
	font-family: 'Sarabun';
	font-style: normal;
	font-weight: 400;
	src: local('Sarabun Regular'), local('Sarabun-Regular'), url(https://fonts.gstatic.com/s/sarabun/v7/DtVjJx26TKEr37c9aBVJmQ.woff) format('woff');
 }
`

func TestParesCSSWoff(t *testing.T) {

	t.Run("TestParesCSSWoff", func(t *testing.T) {

		err := paresFontFace(cssWoff)

		if err != nil {
			t.Errorf("got %v", err)
		}
	})
}

const cssFontFace = `@font-face {
	font-family: 'Sarabun';
	font-style: italic;
	font-weight: 400;
	src: local('Sarabun Italic'), local('Sarabun-Italic'), url(https://fonts.gstatic.com/s/sarabun/v7/DtVhJx26TKEr37c9aBB5nXwP.woff) format('woff');
}`

func TestParseFontFamily(t *testing.T) {

	t.Run("TestParseFontFamily", func(t *testing.T) {

		err := parseFontFamily(cssFontFace)

		if err != nil {
			t.Errorf("got %v", err)
		}
	})
}

func TestParseFontStyle(t *testing.T) {

	t.Run("TestParseFontStyle", func(t *testing.T) {

		err := parseFontStyle(cssFontFace)

		if err != nil {
			t.Errorf("got %v", err)
		}
	})
}

func TestParseFontWeight(t *testing.T) {

	t.Run("TestParseFontWeight", func(t *testing.T) {

		err := parseFontWeight(cssFontFace)

		if err != nil {
			t.Errorf("got %v", err)
		}
	})
}

func TestParseFontURL(t *testing.T) {

	t.Run("TestParseFontURL", func(t *testing.T) {

		err := parseFontURL(cssFontFace)

		if err != nil {
			t.Errorf("got %v", err)
		}
	})
}
