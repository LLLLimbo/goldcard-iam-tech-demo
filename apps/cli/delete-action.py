import http.client
import argparse
import json

parser = argparse.ArgumentParser(description='Post request')
parser.add_argument('--auth', type=str, help='Authorization header')
parser.add_argument('--orgid', type=str, help='x-zitadel-orgid header')
parser.add_argument('--actionId', type=str, help='action id')

args = parser.parse_args()

headers = {
  'Authorization': 'Bearer ' + args.auth,
  'x-zitadel-orgid': args.orgid,
  'Content-Type': 'application/json'
}

data = {}

conn = http.client.HTTPConnection("zitadel:8080")

conn.request("DELETE", "/management/v1/actions/"+ args.actionId, body=json.dumps(data), headers=headers)

resp = conn.getresponse()
print(resp.read())

conn.close()