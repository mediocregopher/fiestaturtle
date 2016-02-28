// @flow
import React, {Component} from 'react'
import {connect} from 'react-redux'
import {globalStyles} from '../styles/style-guide'

import type {Playlist, Song} from '../constants/types/ft'

class Library extends Component {
  props: {
    errorMessage: ?string,
    playlists: Playlist,
    path: any
  };

  render () {
    if (this.props.playlists.length === 0) {
      return <div> Looks like you got no music, señor :(</div>
    }
    return (
      <div> music goes here </div>
    )
  }
}

export default connect(
  s => {
    const {errorMessage, playlists} = s.library
    return {errorMessage, playlists}
  }
)(Library)
