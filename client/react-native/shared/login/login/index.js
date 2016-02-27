import React, {Component} from 'react'
import {connect} from 'react-redux'
import {cancelLogin, relogin} from '../../actions/login'
import {navigateTo} from '../../actions/router'
import Render from './index.render'

class Login extends Component {
  render () {
    return (
      <Render
        onLogin={this.props.onLogin}
        onBack={this.props.onBack}
        onSignup={this.props.onSignup}
        serverURI={this.props.serverURI}
        users={this.props.users}
        lastUser={this.props.lastUser}
      />
    )
  }

  static parseRoute (store, currentPath, nextPath) {
    return {componentAtTop: {}}
  }
}

Login.propTypes = Render.propTypes

export default connect(
  store => {
    const users = store.login && store.login.configuredAccounts && store.login.configuredAccounts.map(c => c.username) || []
    let lastUser = store.config && store.config.status && store.config.status.user && store.config.status.user.username

    if (users.indexOf(lastUser) === -1 && users.length) {
      lastUser = users[0]
    }

    return {
      serverURI: /* store.config && store.config.config && store.config.config.serverURI */'https://keybase.io',
      users, lastUser
    }
  },
  dispatch => {
    return {
      onLogin: (user, passphrase, store) => dispatch(relogin(user, passphrase, store)),
      onBack: () => dispatch(cancelLogin()),
      onSignup: () => dispatch(navigateTo(['signup']))
    }
  }
)(Login)
