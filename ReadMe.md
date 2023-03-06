# ECE 49595OSS Project
Our idea revolves around the use of new-age technologies and Identity proxy methods to grant
access to users and engineers to their applications. Unlike using VPNs which grant access to a
LAN once verified, using Zero Trust Network Access is a technology we plan to use in our idea
to offer limited access to only users who have been approved. ZTNA provides a “never trust,
always verify” least-privilege approach. In this approach, the application is not exposed to the
internet or does not exist on the global subnet but only when a user is authenticated, a pipeline
is opened to the application to offer user-specific, restricted, and access to the user. In
conclusion, our idea offers an Identity-based identification and access control instead of
IP-based access control.

## Status
- UI: Frontend v0.1
  - List Screen
  - Google Auth
  - Angular Material Design
- MongoDB for registered users and app permissions
- GoLang API Endpoints
  - /users/getApps (returns list of accessible apps for user, requires JWT token)
  - /users/login (returns JWT token in string format of authorized user)
- Integrated end-2-end
  

## Environment Variables
|Variable Name|Variable Description|Default Value|
|----------------|----------------------------|------------|
|PORT|Port used for API|8000|
|SECRET_KEY|Secret key used in JWT encoding||
|TOKEN_DURATION|Token duration (hours) for JWT token|24|
|REFRESH_TOKEN_DURATION|Token duration (hours) for JWT refresh token|196|
|MONGODB_URL|URL for hosted MongoDB||

