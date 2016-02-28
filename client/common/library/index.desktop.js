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

import {createPlaylist, addSongToPlaylist, removeFromPlaylist, deletePlaylist, playSong, playPlaylist} from '../actions/library'

import {deriveAllSongs} from '../reducers/library'

import type {Playlist, Song} from '../constants/types/ft'
import ReactList from 'react-list'

class Library extends Component {
  props: {
    errorMessage: ?string,
    playlists: Array<Playlist>,
    createPlaylist: (name: string) => void,
    path: any,
    playPlaylist: () => void,
    playSong: () => void,
  };

  state: {
    openedPlaylistWindow: boolean,
    showPlaylistsDialog: boolean,
    songsToAdd: Array<Song>,
    playlistName: string,
    sortedBy: 'title' | 'artist' | 'genre' | 'album',
    sortedByInverse: boolean
  };

  constructor (props: any) {
    super(props)
    this.state = {openedPlaylistWindow: false, playlistName: '', sortedBy: 'title', sortedByInverse: false, showPlaylistsDialog: false, songsToAdd: []}
  }

  renderItem (song, key) {
    const colStyle = {flex: 1, overflow: 'scroll'}
    return (
      <div
        style={{...globalStyles.clickable, ...globalStyles.noSelect}} key={key}>
        <MenuItem innerDivStyle={{paddingRight: 0}} onTouchTap={this.props.playSong}>
          <div style={{display: 'flex'}}>
            <span style={colStyle}>{song.meta.title}</span>
            <span style={colStyle}>{song.meta.artist}</span>
            <span style={colStyle}>{song.meta.album}</span>
            <span style={colStyle}>{song.meta.genre}</span>
            <div style={{width: 50, position: 'relative'}}>
              <FlatButton style={{margin: 0, padding: 0, height: 50, width: 20, flex:1}} onClick={() => {}}>
                <i onClick={() => this.promptAddSongToPlaylist(song)} className="material-icons" style={{position: 'relative', right: 18, top: 5}}>library_music</i>
              </FlatButton>
            </div>
          </div>

        </MenuItem>
      </div>
    )
  }

  renderPlaylist (playlist, key) {
    return (
      <div
        style={{...globalStyles.clickable, ...globalStyles.noSelect}} key={key}>
        <MenuItem primaryText={playlist.name}
          onTouchTap={this.props.playPlaylist}
        />
      </div>
    )
  }

  cancelNewPlaylist () {
    this.setState({openedPlaylistWindow: false})
  }

  createPlaylist () {
    this.setState({openedPlaylistWindow: false})
    this.props.createPlaylist(this.state.playlistName)
  }

  changeSortedBy (newSortedBy) {
    if (this.state.sortedBy === newSortedBy) {
      this.setState({sortedByInverse: !this.state.sortedByInverse})
    }
    this.setState({sortedBy: newSortedBy})
  }

  promptAddSongToPlaylist (song: Song) {
    this.setState({showPlaylistsDialog: true, songsToAdd: [song]})
  }

  addSongToPlaylist (song: Song, playlist: Playlist) {
    this.props.addSongToPlaylist(song, playlist)
    this.setState({showPlaylistsDialog: false, songsToAdd: []})
  }

  cancelPromptAddSongToPlaylist (song: Song) {
    this.setState({showPlaylistsDialog: false, songsToAdd: []})
  }

  render () {
    if (this.props.playlists.length === 0) {
      return (
        <div
          style={{...globalStyles.noSelect}}>
          Looks like you got no music, señor :(
        </div>
      )
    }

    const playlists = this.props.playlists
    let songs = deriveAllSongs(this.props.playlists)
    songs = [...songs].sort((a: Song, b: Song) => {
      // $FlowIssue
      const result = a.meta[this.state.sortedBy] < b.meta[this.state.sortedBy]
      return this.state.sortedByInverse ? !result : result
    })
    const colStyle = {flex: 1, overflow: 'scroll', textAlign: 'left'}

    return (
      <div style={{...globalStyles.noSelect, ...globalStyles.flexBoxColumn, flex: 2}}>
        <span style={{fontSize: 42, textAlign: 'center'}}> Your Música </span>
        <div style={{display: 'flex', marginLeft: 16}}>
          <FlatButton labelStyle={{padding: 0}} style={colStyle} label={'Title'}
            onClick={() => this.changeSortedBy('title')}/>
          <FlatButton labelStyle={{padding: 0}} style={colStyle} label={'Artist'}
            onClick={() => this.changeSortedBy('artist')}/>
          <FlatButton labelStyle={{padding: 0}} style={colStyle} label={'Album'}
            onClick={() => this.changeSortedBy('album')}/>
          <FlatButton labelStyle={{padding: 0}} style={colStyle} label={'Genre'}
            onClick={() => this.changeSortedBy('genre')}/>
          <div style={{width: 50}}/>
        </div>
        <Divider/>

        <ReactList
          itemRenderer={(i, k) => this.renderItem(songs[i], k)}
          length={songs.length}
          type='uniform'/>

        <Divider/>
        <span style={{fontSize: 42, textAlign: 'center'}}> Playlists </span>
        <SweetTable
          icons={[]}
          onClick={() => {}}
          headers={[['playlist name', () => {}]]}
          colFormatter={p => [p.name]}
          items={playlists}
         />

        <div style={{marginTop: 'auto', alignSelf: 'flex-end', overflow: 'visible',
          marginRight: 20,
          marginBottom: 20,
        }}>
          <FloatingActionButton onClick={() => {this.setState({openedPlaylistWindow: true})}}>
            <ContentAdd />
          </FloatingActionButton>
          <div>
          <Dialog
            title="Name your new playlist"
            contentStyle={{marginBottom: 65}}
            onRequestClose={() => this.cancelNewPlaylist()}
            open={this.state.openedPlaylistWindow}>
            <div style={{display: 'flex', marginBottom: 20}}>
              <TextField
                style={{marginRight: 'auto', marginLeft: 'auto'}}
                hintText='Playlist Name'
                onChange={({target: {value: playlistName}}) => this.setState({playlistName})}
                floatingLabelText='Playlist Name'/>
            </div>
            <div style={{...globalStyles.flexBoxRow, justifyContent: 'flex-end'}}>
              <FlatButton label="Cancel" secondary={true} onTouchTap={() => this.cancelNewPlaylist()}/>
              <FlatButton label="Create" primary={true} onTouchTap={() => this.createPlaylist()}/>
            </div>
          </Dialog>
          <Dialog
            title="Add this song to a playlist"
            contentStyle={{marginBottom: 65}}
            onRequestClose={() => this.cancelPromptAddSongToPlaylist()}
            open={this.state.showPlaylistsDialog}>
            <div style={{display: 'flex', marginBottom: 20}}>
              <SweetTable
                icons={[]}
                onClick={(i) => {this.addSongToPlaylist(this.state.songsToAdd[0], playlists[i])}}
                headers={[['playlist name', () => {}]]}
                colFormatter={p => [p.name]}
                items={playlists}
               />
            </div>
            <div style={{...globalStyles.flexBoxRow, justifyContent: 'flex-end'}}>
              <FlatButton label="Cancel" secondary={true} onTouchTap={() => this.cancelPromptAddSongToPlaylist()}/>
              <FlatButton label="Add" primary={true} onTouchTap={() => this.cancelPromptAddSongToPlaylist()}/>
            </div>
          </Dialog>
          </div>
        </div>
      </div>
    )
  }
}

export default connect(
  s => {
    const {errorMessage, playlists} = s.library
    return {errorMessage, playlists}
  },
  d => (bindActionCreators({createPlaylist, deletePlaylist, removeFromPlaylist, addSongToPlaylist, playSong, playPlaylist}, d))
)(Library)

const styles = {
}
