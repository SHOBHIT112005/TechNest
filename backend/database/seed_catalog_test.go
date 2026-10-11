package database

import "testing"

func TestCatalogSize(t *testing.T) {
	all := append(append([]seedProduct{}, seedData...), generatedProducts()...)
	if len(all) < 500 {
		t.Fatalf("want 500+ products, got %d", len(all))
	}
	seen := map[string]bool{}
	cats := map[string]bool{}
	for _, p := range all {
		if seen[p.Name] {
			t.Fatalf("duplicate name %q", p.Name)
		}
		seen[p.Name] = true
		if p.Price <= 0 {
			t.Fatalf("bad price for %q", p.Name)
		}
		for _, c := range p.Categories {
			cats[c] = true
		}
	}
	t.Logf("%d products, %d categories", len(all), len(cats))
}
