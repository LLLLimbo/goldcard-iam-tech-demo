let http = require('zitadel/http')

function inform(ctx, api) {
    http.fetch('http://message-receiver:17011/rcv', {
        method: 'POST',
        body: {
            "authError": ctx.v1.authError,
            "name": ctx.v1.authRequest.username,
            "id": ctx.v1.authRequest.userId,
            "browserInfo": ctx.v1.authRequest.browserInfo
        }
    });
}