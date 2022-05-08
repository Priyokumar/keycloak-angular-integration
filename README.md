# keycloak-angular-integration

## To add aud field in the access token, you need to do manually in the keyclock admin console.

- Create a client scope eg. my-client-scope
- Add a mapper in the newly created my-client-scope
- Mapper Type: Audience
- Included Client Audience: Select the client that you want to add "aud"
- Add to access code: ON

