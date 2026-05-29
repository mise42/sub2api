package service

import (
	"context"
	"testing"

	"github.com/resend/resend-go/v3"
)

type fakeResendEmailsSender struct {
	request *resend.SendEmailRequest
	options *resend.SendEmailOptions
}

func (f *fakeResendEmailsSender) SendWithOptions(_ context.Context, params *resend.SendEmailRequest, options *resend.SendEmailOptions) (*resend.SendEmailResponse, error) {
	f.request = params
	f.options = options
	return &resend.SendEmailResponse{Id: "email_123"}, nil
}

func TestResendProviderSendUsesSDKRequestAndIdempotencyKey(t *testing.T) {
	fake := &fakeResendEmailsSender{}
	provider := &ResendProvider{
		config: &ResendConfig{
			APIKey:   "re_test",
			From:     "noreply@example.com",
			FromName: "Sub2API",
		},
		emails: fake,
	}

	err := provider.Send(context.Background(), EmailMessage{
		To:             "user@example.com",
		Subject:        "Hello",
		HTML:           "<p>Hello</p>",
		IdempotencyKey: "verify-code/user@example.com/123",
	})
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	if fake.request == nil {
		t.Fatal("expected SDK send request")
	}
	if got, want := fake.request.From, "Sub2API <noreply@example.com>"; got != want {
		t.Fatalf("From = %q, want %q", got, want)
	}
	if got, want := fake.request.To, []string{"user@example.com"}; len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("To = %v, want %v", got, want)
	}
	if got, want := fake.request.Subject, "Hello"; got != want {
		t.Fatalf("Subject = %q, want %q", got, want)
	}
	if got, want := fake.request.Html, "<p>Hello</p>"; got != want {
		t.Fatalf("Html = %q, want %q", got, want)
	}
	if fake.options == nil {
		t.Fatal("expected SDK send options")
	}
	if got, want := fake.options.IdempotencyKey, "verify-code/user@example.com/123"; got != want {
		t.Fatalf("IdempotencyKey = %q, want %q", got, want)
	}
}

func TestNormalizeEmailProviderDefaultsToSMTP(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "", want: EmailProviderSMTP},
		{input: "smtp", want: EmailProviderSMTP},
		{input: "resend", want: EmailProviderResend},
		{input: "RESEND", want: EmailProviderResend},
		{input: "unknown", want: EmailProviderSMTP},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := NormalizeEmailProvider(tt.input); got != tt.want {
				t.Fatalf("NormalizeEmailProvider(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
