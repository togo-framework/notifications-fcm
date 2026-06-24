// Package fcm is a Firebase Cloud Messaging (HTTP v1) push channel for togo
// notifications. Sends a PushNotification to the recipient's device tokens.
// Install: `togo install togo-framework/notifications-fcm`.
// Env: FCM_PROJECT_ID, GOOGLE_APPLICATION_CREDENTIALS (service-account JSON).
package fcm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2/google"

	"github.com/togo-framework/notifications"
	"github.com/togo-framework/togo"
)

const scope = "https://www.googleapis.com/auth/firebase.messaging"

func init() {
	notifications.RegisterChannel("fcm", func(k *togo.Kernel) notifications.Channel {
		return &channel{projectID: os.Getenv("FCM_PROJECT_ID"), client: &http.Client{Timeout: 15 * time.Second}}
	})
}

type channel struct {
	projectID string
	client    *http.Client
}

func (c *channel) token(ctx context.Context) (string, error) {
	creds, err := google.FindDefaultCredentials(ctx, scope)
	if err != nil {
		return "", err
	}
	t, err := creds.TokenSource.Token()
	if err != nil {
		return "", err
	}
	return t.AccessToken, nil
}

func (c *channel) Send(ctx context.Context, to notifications.Notifiable, n notifications.Notification) error {
	pn, ok := n.(notifications.PushNotification)
	if !ok {
		return nil
	}
	tokens := to.RoutePushTokens()
	if len(tokens) == 0 || c.projectID == "" {
		return nil
	}
	msg := pn.ToPush(to)
	access, err := c.token(ctx)
	if err != nil {
		return err
	}
	endpoint := "https://fcm.googleapis.com/v1/projects/" + c.projectID + "/messages:send"
	for _, device := range tokens {
		payload := map[string]any{"message": map[string]any{
			"token":        device,
			"notification": map[string]string{"title": msg.Title, "body": msg.Body},
			"data":         msg.Data,
		}}
		buf, _ := json.Marshal(payload)
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(buf))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+access)
		resp, err := c.client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode >= 300 {
			b, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			return fmt.Errorf("notifications-fcm: status %d: %s", resp.StatusCode, string(b))
		}
		resp.Body.Close()
	}
	return nil
}
