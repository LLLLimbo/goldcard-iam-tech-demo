import http.client
import argparse
import json

parser = argparse.ArgumentParser(description='Post request')
parser.add_argument('--auth', type=str, help='Authorization header')
parser.add_argument('--orgid', type=str, help='x-zitadel-orgid header')
parser.add_argument('--actionIds', nargs='+', type=str, help='action ids')

args = parser.parse_args()

headers = {
  'Authorization': 'Bearer ' + args.auth,
  'x-zitadel-orgid': args.orgid,
  'Content-Type': 'application/json'
}

data = {
  "actionIds": args.actionIds
}

conn = http.client.HTTPConnection("zitadel:8080")

conn.request("POST", "/management/v1/flows/3/trigger/1", body=json.dumps(data), headers=headers)

resp = conn.getresponse()
print(resp.read())

conn.close()