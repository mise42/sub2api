# Sub2API

Sub2API provides an API gateway and user/admin operations surface. This glossary captures project-specific language used across operational settings and user-facing email flows.

## Language

**Email Service**:
The system capability that sends transactional emails such as verification codes, password resets, alerts, and reports. An Email Service has exactly one active Email Provider.
_Avoid_: SMTP settings

**Email Provider**:
The configured delivery channel used by the Email Service. Current provider choices are SMTP and Resend; SMTP is the default when no provider has been selected.
_Avoid_: Mail method, email type

**SMTP Provider**:
An Email Provider that delivers mail through an SMTP server. SMTP host, port, username, password, and TLS mode belong only to this provider.
_Avoid_: Email service

**Resend Provider**:
An Email Provider that delivers mail through the Resend HTTP API. Resend API credentials and sender identity belong only to this provider.
_Avoid_: Resend SMTP

**Provider Credential**:
A secret used by one Email Provider to authenticate with its delivery channel. Provider Credentials are not shared across providers.
_Avoid_: Reused password, shared mail secret

**Sender Identity**:
The email address and display name used as the sender for transactional emails. Each Email Provider owns its own Sender Identity because provider verification rules can differ.
_Avoid_: From config

**Test Email**:
A real transactional email sent to an administrator-selected recipient to verify the active or drafted Email Provider configuration.
_Avoid_: Connection test

**SMTP Connection Test**:
An SMTP Provider-only diagnostic that checks whether the configured SMTP server can be reached and authenticated without sending an email.
_Avoid_: Email provider test

**Draft Email Provider Configuration**:
An unsaved Email Provider configuration that an administrator is editing and may use for a Test Email before saving it.
_Avoid_: Temporary SMTP settings

## Example Dialogue

Developer: Which Email Provider should this deployment use?
Domain expert: Use the Resend Provider because the hosting environment blocks outbound SMTP ports.

Developer: Should I change the SMTP Provider settings for that?
Domain expert: No. SMTP Provider settings are only for SMTP delivery. Configure the Resend Provider and keep the Email Service pointed at Resend.

Developer: Can the Resend Provider reuse the SMTP Provider password field?
Domain expert: No. Each Email Provider owns its Provider Credential separately, even if the admin switches between providers.

Developer: Can Resend reuse the SMTP sender address?
Domain expert: No. Resend has its own Sender Identity because it may require a different verified sending domain.

Developer: Should every Email Provider have a connection test?
Domain expert: No. Test Email is the generic verification. SMTP Connection Test is only an SMTP Provider diagnostic.

Developer: Can a Test Email use settings that have not been saved yet?
Domain expert: Yes. A Draft Email Provider Configuration may be tested before it becomes the active Email Provider configuration.

Developer: What Email Provider should an existing deployment use before an administrator chooses one?
Domain expert: Use the SMTP Provider by default so existing mail settings continue to mean the same thing.

Developer: Does adding Resend change what existing SMTP settings mean?
Domain expert: No. Existing SMTP settings continue to describe the SMTP Provider even after the Email Service gains more provider choices.
