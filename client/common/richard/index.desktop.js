// @flow
import React, {Component} from 'react'
import {connect} from 'react-redux'
import {globalStyles} from '../styles/style-guide'
import MenuItem from 'material-ui/lib/menus/menu-item'
import FloatingActionButton from 'material-ui/lib/floating-action-button'
import ContentAdd from 'material-ui/lib/svg-icons/content/add'
import Dialog from 'material-ui/lib/dialog'
import FlatButton from 'material-ui/lib/flat-button'
import TextField from 'material-ui/lib/text-field';
import {bindActionCreators} from 'redux'
import Divider from 'material-ui/lib/divider'

import SweetTable from '../table.desktop'

import {deleteSongFromLibrary, createPlaylist, addSongToPlaylist, removeFromPlaylist, deletePlaylist, playSong, playPlaylist} from '../actions/library'

import {deriveAllSongs} from '../reducers/library'

import type {Playlist, Song} from '../constants/types/ft'
import ReactList from 'react-list'

export default class Library extends Component {
  render () {
    return (
      <div style={{
        marginLeft: 'auto',
        marginRight: 'auto',
        marginTop: 40
        }}>

        <div style={{marginTop: 20}}>
        <TextField
          hintText="Richard's addr"
          ref='input'
          onEnterKeyDown={() => {
            const v = this.refs.input && this.refs.input.getValue()
            const newRichardServerList = this.props.richardServers.concat({addr: v})
            this.props.onRichardServerChange(newRichardServerList)
          }}
          floatingLabelText="Richard's addr" />
        </div>

        <div style={{marginTop: 20}}>
          <span style={{fontSize: 42, textAlign: 'center'}}> Manage Richards </span>

          <SweetTable
            icons={[
              ['delete', i => {
                const rs = this.props.richardServers[i]
                const newRichardServerList = this.props.richardServers.filter(r => r !== rs)
                this.props.onRichardServerChange(newRichardServerList)
              }]
            ]}
            onClick={i => {}}
            headers={[
              ['Address', () => {}]
            ]}
            colFormatter={({addr}) => {
              return [addr]
            }}
            items={this.props.richardServers || []} />
        </div>
      </div>
    )
  }
}
