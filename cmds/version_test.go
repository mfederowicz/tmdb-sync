package cmds

import "testing"

func TestResolveAPIVersion(t *testing.T) {
	tests := []struct {
		name    string
		v3, v4  bool
		want    apiVersion
		wantErr bool
	}{
		{"default is v3", false, false, apiV3, false},
		{"explicit v3", true, false, apiV3, false},
		{"explicit v4", false, true, apiV4, false},
		{"both is an error", true, true, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolveAPIVersion(tt.v3, tt.v4)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
