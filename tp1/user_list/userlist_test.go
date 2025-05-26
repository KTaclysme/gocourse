package userlist

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

type testCase struct {
	action Action
	nom    Nom
	tel    Tel
	want   string
}

func TestUserlist(t *testing.T) {
	users = make(map[Nom]Tel)

	var tests = []testCase{
		{
			action: "ajouter",
			nom:    "Keke",
			tel:    "012345689",
			want:   "User Keke added with phone number 012345689\n",
		},
		{
			action: "ajouter",
			nom:    "Keke",
			tel:    "012345689",
			want:   "User Keke already exists with phone number 012345689\n",
		},
		{
			action: "modifier",
			nom:    "Keke",
			tel:    "012345690",
			want:   "User Keke updated with phone number 012345690\n",
		},
		{
			action: "rechercher",
			nom:    "Keke",
			want:   "User Keke found with phone number 012345690\n",
		},
		{
			action: "supprimer",
			nom:    "Keke",
			want:   "User Keke deleted\n",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w

			fn, ok := Actions[test.action]
			if !ok {
				t.Fatalf("Unknown action: %s", test.action)
			}
			fn(test.nom, test.tel)
			w.Close()

			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)
			got := buf.String()

			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
