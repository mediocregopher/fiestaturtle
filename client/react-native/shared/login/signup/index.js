/* @flow */

import React from 'react'
import ErrorText from '../error.render'

// Signup Components
import InviteCode from './invite-code'
import UsernameEmailForm from './username-email-form'
import PassphraseSignup from './passphrase'
import DeviceName from './device-name'
import Success from './success'

import {Map} from 'immutable'
import type {URI} from '../../reducers/router'

export default function signupRouter (currentPath: Map<string, string>, uri: URI): any {
  // Fallback (for debugging)
  let form = <ErrorText currentPath={currentPath} />

  const path = currentPath.get('path')

  const {component: Component, props} = currentPath.get('parseRoute') || {}
  if (Component) {
    form = <Component {...props}/>
  } else {
    switch (path) {
      case 'signup':
      case 'inviteCode':
        form = <InviteCode/>
        break
      case 'usernameAndEmail':
        form = <UsernameEmailForm/>
        break
      case 'passphraseSignup':
        form = <PassphraseSignup/>
        break
      case 'deviceName':
        form = <DeviceName/>
        break
      case 'paperkey':
      case 'success':
        form = <Success/>
        break
    }
  }

  return {
    componentAtTop: {
      component: () => form,
      hideBack: true
    },
    parseNextRoute: signupRouter
  }
}
