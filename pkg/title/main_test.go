package title

import (
	"fmt"
	"strings"
	"testing"
)

type testSpec struct {
	input  string // Test input.
	output string // Expected output.
}

// Validate title.Title() results.
func TestTitle(t *testing.T) {
	tests := []testSpec{
		{
			// Test if endings with "'s", punctuated abbreviations and chars
			// "åäö" are handled correctly.
			input:  "göran's secret quote, a.b.r. and åå ää öö test...",
			output: "Göran's Secret Quote, A.B.R. And Åå Ää Öö Test...",
			//        ^^  ^^                 ^ ^      ^^ ^^ ^^
		},
	}

	for _, test := range tests {
		result := Title(test.input)
		if result != test.output {
			t.Errorf("Title(%s) FAILED\nExpected: %s\nReceived: %s\n",
				test.input, test.output, result)
		} else {
			t.Logf("Title(%s) SUCCEDED\nResult: %s\n", test.input, result)
		}
	}
}

// Validate title.Titleize() results.
func TestTitleize(t *testing.T) {
	test := testSpec{}
	tests := []testSpec{}

	// Test if first and last word gets titlecased.
	test = testSpec{
		input:  "the the the",
		output: "The the The",
	}
	tests = append(tests, test)

	// Test if all words in lowercase list gets lowered.
	lower := strings.Join(lowercase, " ")
	test = testSpec{
		input:  fmt.Sprintf("FIRST %s LAST", strings.ToUpper(lower)),
		output: fmt.Sprintf("First %s Last", lower),
	}
	tests = append(tests, test)

	// Test if punctuated abbreviation does not trigger new sentence.
	test = testSpec{
		input:  "first a.b.r. the last",
		output: "First A.B.R. the Last",
	}
	tests = append(tests, test)

	// Test if punctuated digits does trigger new sentence.
	test = testSpec{
		input:  "first 1.2.3. the last",
		output: "First 1.2.3. The Last",
	}
	tests = append(tests, test)

	// Test if all chars in end_sentence list triggers new sentence.
	for _, char := range end_sentence {
		test := testSpec{
			input:  fmt.Sprintf("first %s the last", char),
			output: fmt.Sprintf("First %s The Last", char),
		}
		tests = append(tests, test)
	}

	// Test if lowercase list is ignored following bracket or quote.
	sidestep_chars := []string{"(", "{", "[", "<", "\"", "'"}
	for _, char := range sidestep_chars {
		test := testSpec{
			input:  fmt.Sprintf("first %sthe last", char),
			output: fmt.Sprintf("First %sThe Last", char),
		}
		tests = append(tests, test)
	}

	// Test if song and album related titleization is correct.
	music_tests := []testSpec{
		{
			input:  "anna von hausswolff - the truth, the glow, the fall",
			output: "Anna von Hausswolff - The Truth, The Glow, The Fall",
			//            ^^^            ^^^^^      ^^^^^     ^^^^^
		},
		{
			input:  "bob marley & the wailers - no woman, no cry (live)",
			output: "Bob Marley & The Wailers - No Woman, No Cry (Live)",
			//                  ^^^^^                            ^^^^^
		},
		{
			input:  "danzig - satan (from satan's sadists)",
			output: "Danzig - Satan (From Satan's Sadists)",
			//                      ^^^^^      ^^
		},
		{
			input:  "tears for fears @ rule the world: the greatest hits",
			output: "Tears for Fears @ Rule the World: The Greatest Hits",
			//             ^^^              ^^^      ^^^^^
		},
		{
			input:  "sepultura - r.i.p. (rest in pain)",
			output: "Sepultura - R.I.P. (Rest in Pain)",
			//                     ^^^        ^^
		},
	}
	tests = append(tests, music_tests...)

	for _, test := range tests {
		result := Titleize(test.input)
		if result != test.output {
			t.Errorf("Titleize(%s) FAILED\nExpected: %s\nReceived: %s\n",
				test.input, test.output, result)
		} else {
			t.Logf("Titleize(%s) SUCCEDED\nResult: %s\n", test.input, result)
		}
	}
}
