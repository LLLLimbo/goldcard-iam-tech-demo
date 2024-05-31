const authConfig = {
    authority: 'http://zitadel:8080/', //Replace this with your issuer URL
    client_id: '269224945731567623@demo', //Replace this with your client id
    redirect_uri: 'http://localhost:3000/',
    response_type: 'code',
    scope: 'openid profile email urn:zitadel:iam:org:project:id:Demo:aud', //Replace PROJECT_ID with the id of the project where the API resides.
    post_logout_redirect_uri: 'http://localhost:3000/',
    response_mode: 'query',
    code_challenge_method: 'S256',
  };

 export default authConfig;
