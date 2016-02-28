// @flow
import React, {Component} from 'react'
import {connect} from 'react-redux'

import {globalStyles} from './styles/style-guide'
import {navigateUp, routeAppend} from './actions/router'

import RaisedButton from 'material-ui/lib/raised-button'
import Divider from 'material-ui/lib/divider'


import Upload from './upload/index.desktop'
import Library from './library/index.desktop'

type Props = {
  username: ?string,
}

const water = `
    <svg version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="300px" height="300px" viewBox="0 0 300 300" enable-background="new 0 0 300 300" xml:space="preserve">
      <path fill="#04ACFF" id="waveShape" d="M300,300V2.5c0,0-0.6-0.1-1.1-0.1c0,0-25.5-2.3-40.5-2.4c-15,0-40.6,2.4-40.6,2.4
	c-12.3,1.1-30.3,1.8-31.9,1.9c-2-0.1-19.7-0.8-32-1.9c0,0-25.8-2.3-40.8-2.4c-15,0-40.8,2.4-40.8,2.4c-12.3,1.1-30.4,1.8-32,1.9
	c-2-0.1-20-0.8-32.2-1.9c0,0-3.1-0.3-8.1-0.7V300H300z"/>
    </svg>
`

class Nav extends Component {
  componentDidMount () {
  }

  render () {
    console.log('route', this.props.route)

    const {pathSoFar, navUp} = this.props
    const firstPath = pathSoFar.first() && pathSoFar.first().get('path')
    console.log('path', pathSoFar)

    if (firstPath === 'search') {
      return <div path={pathSoFar.rest()} onClick={navUp}> to search </div>
    } else if (firstPath === 'upload') {
      return <Upload path={pathSoFar.rest()} onClick={navUp}/>
    }

    return (
      <div style={{...globalStyles.flexBoxColumn, flex: 1}}>

      <div style={{...globalStyles.flexBoxColumn, marginTop: 40, marginLeft: 'auto', marginRight: 'auto', alignItems: 'center'}}>

        <div id={'banner'} style={{marginLeft: 80, marginRight: 80}}>
          <div style={{fontStyle: 'italic', fontSize: 32, textAlign: 'center', marginTop: 20, marginBottom: -88, zIndex: 10, position: 'relative'}}>¡Fiesta Turtle!</div>
          <div className={'fill'} dangerouslySetInnerHTML={{__html: water}}/>
        </div>

        <div style={{...globalStyles.flexBoxRow, justifyContent: 'space-around', marginTop: 20, minWidth: 280}}>
          <div style={globalStyles.flexBoxColumn}>
            <RaisedButton label='Search' primary onClick={this.props.onSearch}/>
          </div>
          <div style={globalStyles.flexBoxColumn}>
            <RaisedButton label='Upload' secondary onClick={this.props.onUpload}/>
          </div>
        </div>

      </div>
      <Divider style={{marginTop: 20}}/>
      <div style={{marginTop: 20, flex: 2, display: 'flex'}}>
        <Library path={pathSoFar} />
      </div>
      </div>
    )
  }
}

export default connect(
  (state) => {
    const pathSoFar = state.router.get('uri').slice(2)
    return {route: state.router.get('uri'), pathSoFar}
  },
  (dispatch) => ({
    onUpload: () => dispatch(routeAppend('upload')),
    onSearch: () => dispatch(routeAppend('search')),
    navUp: () => dispatch(navigateUp())
  }),
  (s, d, o) => ({
    ...s, ...d,
    username: o.username || 'mediocregopher'
  }),
)(Nav)
