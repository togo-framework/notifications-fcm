<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/notifications-fcm</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/notifications-fcm"><img src="https://pkg.go.dev/badge/github.com/togo-framework/notifications-fcm.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/notifications-fcm
```

<!-- /togo-header -->

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

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
