export function rpc (method: string, param: Object) {
  return jsonRPC( {"jsonrpc": "2.0", "method": "Brian." + method, "params": param, "id": 1})
    .then(r => r.json())
}

function jsonRPC (body: any) {
  const msg = JSON.stringify({
    ...body
  })
  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json")
  myHeaders.append('Access-Control-Allow-Origin', '*')
  return fetch(
    'http://localhost:4567/rpc',
    {
      headers: myHeaders,
      method: 'POST',
      body: msg
    }
  )
}
