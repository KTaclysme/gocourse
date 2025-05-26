package greeter

import (
	"testing"
)

func TestGreet(t *testing.T) {
	type testCase struct {
		lang     Language
		want     string
		wantErr  bool
		errValue string
	}
	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"German": {
			lang: "de",
			want: "Hallo welt",
		},
		"Empty": {
			lang:     "",
			wantErr:  true,
			errValue: `unsupported language`,
		},
		"Unsupported language": {
			lang:     "es",
			wantErr:  true,
			errValue: `unsupported language "es"`,
			want:     `Unsupported language "de"`,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Greet(test.lang)

			if test.wantErr {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				if err.Error() != test.errValue {
					t.Errorf("Expected error: %q, got %q", test.errValue, err.Error())
				}

				if got != "" {
					t.Errorf("Expected empty string for 'got' when an error occurs, but got: %q", got)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error, but got: %v", err)
				}
				if got != test.want {
					t.Errorf("Expected: %q, got: %q", test.want, got)
				}
			}
		})
	}
}

func TestGreet_French(t *testing.T) {
	lang := Language("fr")
	want := "Bonjour le monde"

	got, err := Greet(lang)
	if err != nil {
		t.Fatalf("TestGreet_French failed: did not expect an error, but got: %v", err) // Check for unexpected error
	}
	if got != want {
		t.Errorf("expected: %q, got %q", want, got)
	}
}
