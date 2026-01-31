package main

import "testing"

func TestIsValidEmail(t *testing.T) {
	t.Run("valid email with mx, spf, dmarc", func(t *testing.T) {
		var email = []struct {
			name 	   string
			email        string
			hasMX        bool
			hasSPF       bool
			hasDMARC     bool
			
		}{
			{
			name:     "Valid Gmail domain",
			email:    "gmail.com",
			hasMX:    true,
			hasSPF:   true,
			hasDMARC: true,
		},
		{
			name:     "Valid Yahoo domain",
			email:    "yahoo.com",
			hasMX:    true,
			hasSPF:   true,
			hasDMARC: true,
		},
		{
			name:     "Invalid domain",
			email:    "invalid-domain-xyz-123.com",
			hasMX:    false,
			hasSPF:   false,
			hasDMARC: false,
		},
		{
			name:     "Malformed email (no domain)",
			email:    "notanemail",
			hasMX:    false,
			hasSPF:   false,
			hasDMARC: false,
		},
		{
			name:     "Empty string",
			email:    "",
			hasMX:    false,
			hasSPF:   false,
			hasDMARC: false,
		},
		}
		for _, e := range email {
			gotHasMX, gotHasSPF, gotHasDMARC, _, _ := isValidEmail(e.email)
			if gotHasMX != e.hasMX {
				t.Errorf("got %v want %v for hasMX", gotHasMX, e.hasMX)
			}
			if gotHasSPF != e.hasSPF {
				t.Errorf("got %v want %v for hasSPF", gotHasSPF, e.hasSPF)
			}
			if gotHasDMARC != e.hasDMARC {
				t.Errorf("got %v want %v for hasDMARC", gotHasDMARC, e.hasDMARC)
			}
			
		}
	})
}

func Benchmark(b *testing.B){
	testEmails:=[]string {
		"gmail.com",
		"yahoo.com",
		"invalid-domain-xyz-123.com",
		"notanemail",
		"",
	}
	for _, email := range testEmails {
		b.Run(email, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isValidEmail(email)
			}
		})
	}

}