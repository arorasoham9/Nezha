# ECE 49595OSS Project
A Zero Trust Network Access, “never trust, always verify”, proxy  that offers limited application and server access to only users who have been explicitly approved. The client application is not exposed to the internet and does not exist on the global subnet. Only when a user is authenticated with Google and authorized with the proxy, a pipeline is opened to the application to offer user-specific, restricted access to the application. 

## Status
The repository is a work in progress. The following have been implemented: 
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

