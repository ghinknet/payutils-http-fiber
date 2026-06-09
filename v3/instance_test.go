package httpfiber

import (
	"net/http"
	"net/http/httptest"
	"testing"

	stderrors "errors"

	"github.com/gofiber/fiber/v3"

	"go.gh.ink/payutils/v3/errors"
)

func TestNewInstance_RejectsUnsupported(t *testing.T) {
	_, err := Driver{}.NewInstance(123)
	if !stderrors.Is(err, errors.ErrUnsupportedInstance) {
		t.Errorf("err = %v, want ErrUnsupportedInstance", err)
	}
}

func TestNewInstance_AcceptsRouterAndRoutesPost(t *testing.T) {
	app := fiber.New()

	inst, err := Driver{}.NewInstance(app)
	if err != nil {
		t.Fatalf("NewInstance(*fiber.App) error: %v", err)
	}

	called := false
	inst.Post("/pay/callback", func(w http.ResponseWriter, _ *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("success"))
	})

	req := httptest.NewRequest(http.MethodPost, "/pay/callback", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test error: %v", err)
	}
	if !called {
		t.Error("registered handler was not invoked")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}
