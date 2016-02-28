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
  props: {
    onSearchEnter: (q: string) => void,
    searchResults: Array<any>,
    onItemClick: (i: number) => void
  };

  render () {
    console.log('songs', this.props.songResults)
    return (
      <div style={{
        marginLeft: 'auto',
        marginRight: 'auto',
        marginTop: 40
        }}>

        <div style={{marginTop: 20}}>
        <TextField
          floatingLabelText="Search Query"
          ref='input'
          onEnterKeyDown={() => {
            const v = this.refs.input && this.refs.input.getValue()
            this.props.onSearchEnter(v)
          }}
          />
        </div>

        {this.props.songResults.length > 0 &&
          <div style={{marginTop: 20}}>
            <span style={{fontSize: 42, textAlign: 'center'}}> Search Results </span>

            <SweetTable
              icons={[
                ['library_add', i => {
                  const s = this.props.songResults[i]
                  this.props.onSongAdd(s)
                }]
              ]}
              onClick={i => {}}
              headers={[
                ['title', () => {}],
                ['artist', () => {}]
              ]}
              colFormatter={({meta: {title, artist}}) => {
                return [title, artist]
              }}
              items={this.props.songResults || []} />
          </div>
        }

        <FlatButton secondary label={'Back'} onClick={this.props.onBack} />
      </div>
    )
  }
}
