let http = require('zitadel/http')

function inform(ctx, api) {
    http.fetch('http://message:17011/rcv', {
        method: 'POST',
        body: {
            "authError": ctx.v1.authError,
            "name": ctx.v1.user.username,
            "id": ctx.v1.user.id,
            "remoteIp": ctx.v1.authRequest.browserInfo.remoteIp,
            "userAgent": ctx.v1.authRequest.browserInfo.userAgent,
        }
    });
}