import http.client
import argparse
import json

parser = argparse.ArgumentParser(description='Post request')
parser.add_argument('--auth', type=str, help='Authorization header')
parser.add_argument('--orgid', type=str, help='x-zitadel-orgid header')

args = parser.parse_args()

headers = {
  'Authorization': 'Bearer ' + args.auth,
  'x-zitadel-orgid': args.orgid,
  'Content-Type': 'application/json'
}

data = {
  "name": "inform",
  "script":"let http=require('zitadel/http');function inform(ctx,api){http.fetch('http://message:17011/rcv',{method:'POST',body:{\"authError\":ctx.v1.authError,\"name\":ctx.v1.user.username,\"id\":ctx.v1.user.id,\"remoteIp\":ctx.v1.authRequest.browserInfo.remoteIp,\"userAgent\":ctx.v1.authRequest.browserInfo.userAgent}});}\n",
  "timeout": "5s",
  "allowedToFail": True
}

conn = http.client.HTTPConnection("zitadel:8080")

conn.request("POST", "/management/v1/actions", body=json.dumps(data), headers=headers)

resp = conn.getresponse()
print(resp.read())

conn.close()