package color

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ColorTestSuite struct {
	suite.Suite
}

func TestColorTestSuite(t *testing.T) {
	suite.Run(t, new(ColorTestSuite))
}

func (s *ColorTestSuite) TestColor() {
	for key, value := range Palette {
		assert.Equal(
			s.T(),
			Color(key, "hi"),
			ColorStubStart+value+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		)
	}
}

func (s *ColorTestSuite) TestIntegration() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["yellow"]+ColorStubEnd+
			"y"+
			ColorStubStart+Palette["blue"]+ColorStubEnd+
			"b1"+ColorStubStart+Palette["default"]+ColorStubEnd+
			ColorStubStart+Palette["yellow"]+ColorStubEnd+
			ColorStubStart+Palette["red"]+ColorStubEnd+
			"r"+ColorStubStart+Palette["default"]+ColorStubEnd+
			ColorStubStart+Palette["yellow"]+ColorStubEnd+
			ColorStubStart+Palette["blue"]+ColorStubEnd+
			"b2"+ColorStubStart+Palette["default"]+ColorStubEnd+
			ColorStubStart+Palette["yellow"]+ColorStubEnd+
			"y"+ColorStubStart+Palette["default"]+ColorStubEnd,
		Yellow(fmt.Sprintf("y%s%s%sy", Blue("b1"), Red("r"), Blue("b2"))),
	)
}

func (s *ColorTestSuite) TestDefault() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["default"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Default("hi"),
	)
}

func (s *ColorTestSuite) TestBold() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["bold"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Bold("hi"),
	)
}

func (s *ColorTestSuite) TestDim() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["dim"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Dim("hi"),
	)
}

func (s *ColorTestSuite) TestItalics() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["italics"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Italics("hi"),
	)
}

func (s *ColorTestSuite) TestUnderline() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["underline"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Underline("hi"),
	)
}

func (s *ColorTestSuite) TestBlack() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["black"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Black("hi"),
	)
}

func (s *ColorTestSuite) TestGray() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["gray"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Gray("hi"),
	)
}

func (s *ColorTestSuite) TestGrey() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["grey"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Grey("hi"),
	)
}

func (s *ColorTestSuite) TestRed() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["red"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Red("hi"),
	)
}

func (s *ColorTestSuite) TestLightRed() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lred"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightRed("hi"),
	)
}

func (s *ColorTestSuite) TestGreen() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["green"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Green("hi"),
	)
}

func (s *ColorTestSuite) TestLightGreen() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lgreen"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightGreen("hi"),
	)
}

func (s *ColorTestSuite) TestYellow() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["yellow"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Yellow("hi"),
	)
}

func (s *ColorTestSuite) TestLightYellow() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lyellow"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightYellow("hi"),
	)
}

func (s *ColorTestSuite) TestBlue() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["blue"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Blue("hi"),
	)
}

func (s *ColorTestSuite) TestLightBlue() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lblue"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightBlue("hi"),
	)
}

func (s *ColorTestSuite) TestViolet() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["violet"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Violet("hi"),
	)
}

func (s *ColorTestSuite) TestLightViolet() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lviolet"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightViolet("hi"),
	)
}

func (s *ColorTestSuite) TestCyan() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["cyan"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		Cyan("hi"),
	)
}

func (s *ColorTestSuite) TestLightCyan() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lcyan"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightCyan("hi"),
	)
}

func (s *ColorTestSuite) TestLightGray() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lgray"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightGray("hi"),
	)
}

func (s *ColorTestSuite) TestLightGrey() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["lgrey"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		LightGrey("hi"),
	)
}

func (s *ColorTestSuite) TestWhite() {
	assert.Equal(
		s.T(),
		ColorStubStart+Palette["white"]+ColorStubEnd+"hi"+ColorStubStart+"0"+ColorStubEnd,
		White("hi"),
	)
}
