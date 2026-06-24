# notifications-fcm

> **Firebase Cloud Messaging** (HTTP v1) push channel for [togo](https://to-go.dev) **notifications**.

Sends a `PushNotification` to a recipient's device tokens via FCM v1 (OAuth from a service account).

## Install

```bash
togo install togo-framework/notifications-fcm
```

```ini
FCM_PROJECT_ID=your-firebase-project
GOOGLE_APPLICATION_CREDENTIALS=/path/service-account.json
```

Registers the `fcm` notifications channel. MIT
