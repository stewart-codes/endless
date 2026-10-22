package train

import "testing"

func TestChainForAuthorUsesModulo(t *testing.T) {
	chains := []MarkovChain{{}, {}, {}}
	cases := []struct {
		author string
		want   int
	}{
		{"Arlo Mills", 0},
		{"Joe Goetz", 1},
		{"Billy Goetz", 2},
		{"Marybeth Trott", 0},
		{"Charlie Davis", 1},
	}
	for _, tc := range cases {
		got := ChainForAuthor(tc.author, chains)
		if got != chains[tc.want] {
			t.Fatalf("ChainForAuthor(%q): got chain index mismatch", tc.author)
		}
	}
}
