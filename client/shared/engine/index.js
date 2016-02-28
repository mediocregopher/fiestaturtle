// Handles sending requests to native mobile (then go) and back

import engine from './index.native'
import Transport from './transport'
import rpc from 'framed-msgpack-rpc'
import {printRPC} from '../local-debug'
const {client: {Client: RpcClient}} = rpc

import setupLocalLogs, {logLocal} from '../util/local-log'

import {Buffer} from 'buffer'
import NativeEventEmitter from '../common-adapters/native-event-emitter'
import windowsHack from './windows-hack'
import {log} from '../native/log/logui'

import {constants} from '../constants/types/keybase-v1'
import {printOutstandingRPCs} from '../local-debug'

class Engine {
  onConnect () {
  }

  listenOnConnect (key, f) {
  }

  getSessionID () {
  }

  setupListener () {
  }

  // Bind a single callback to a method.
  // Newer calls to this will overwrite older listeners, and older listeners will not get called.
  listenGeneralIncomingRpc (params) {
  }

  removeGeneralIncomingRpc (method) {
  }

  listenServerInit (method, listener) {
  }

  _serverInitIncomingRPC (method, param, response) {
  }

  _rpcWrite (data) {
  }

  _wrapResponseOnceOnly (method, param, response) {
  }

  _generalIncomingRpc (method, param, response) {
  }

  _hasNoHandler (method, callMap, generalIncomingRpcMap) {
  }

  _rpcIncoming (payload) {
  }

  setFailOnError () {
  }

  // Make an RPC and call callbacks in the incomingCallMap
  // (name of call, {arguments object}, {methodName: function(params, response)}, function(err, data)
  // Use rpc() instead
  rpc_unchecked (method, param, incomingCallMap, callback) {
  }

  rpc (params) {
  }

  cancelRPC (sessionID) {
  }

  reset () {
  }
}

export default new Engine()
