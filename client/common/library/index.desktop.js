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

import Player from './player.desktop'

import SweetTable from '../table.desktop'

import {deleteSongFromLibrary, createPlaylist, addSongToPlaylist, removeFromPlaylist, deletePlaylist, playSong, playPlaylist} from '../actions/library'

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
    sortedByInverse: boolean,
    queue: Array<any>,
    playingIndex: number
  };

  constructor (props: any) {
    super(props)
    this.state = {openedPlaylistWindow: false, playlistName: '', sortedBy: 'title', sortedByInverse: false, showPlaylistsDialog: false, songsToAdd: [],
      queue: [],
      playingIndex: -1
    }
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

  addSongToQueue (song: Song) {
    this.setState({queue: this.state.queue.concat(song)})
    if (this.state.playingIndex === -1) {
      this.props.playSong(song)
      this.setState({playingIndex: 0})
    }
  }

  playNextSongInQueue () {
    if (this.state.queue.length > this.state.playingIndex + 1) {
      this.props.playSong(this.state.queue[this.state.playingIndex + 1])
      this.setState({playingIndex: this.state.playingIndex + 1})
    }
  }

  playPrevSongInQueue () {
    if (this.state.playingIndex - 1 >= 0) {
      this.props.playSong(this.state.queue[this.state.playingIndex - 1])
      this.setState({playingIndex: this.state.playingIndex - 1})
    }
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

        <SweetTable
          icons={[
            ['library_music', i => this.promptAddSongToPlaylist(songs[i])],
            ['delete', i => this.props.deleteSongFromLibrary(songs[i])]
          ]}
          onClick={i => this.addSongToQueue(songs[i])}
          headers={[
            ['Title', () => this.changeSortedBy('title')],
            ['Artist', () => this.changeSortedBy('artist')],
            ['Album', () => this.changeSortedBy('album')],
            ['Genre', () => this.changeSortedBy('genre')]
          ]}
          colFormatter={({meta: {title, artist, album, genre}}) => [title, artist, album, genre]}
          items={songs} />

        {this.state.queue.length > 0 && (
          <span style={globalStyles.flexBoxColumn}>
          <span style={{fontSize: 42, textAlign: 'center'}}> ¿Por Queue? </span>
          <SweetTable
            icons={[
              ['clear', i => console.log('remove from q')]
            ]}
            onClick={i => {}}
            headers={[
              ['Title', () => {}],
              ['Artist', () => {}],
            ]}
            colFormatter={song => {
              const {meta: {title, artist, album, genre}} = song
              if (song === this.state.queue[this.state.playingIndex]) {
                return ['-> ' + title, artist]
              }
              return [title, artist]
            }}
            items={this.state.queue} />
            </span>
        )}
        <span style={{fontSize: 42, textAlign: 'center'}}> Playlists </span>
        <SweetTable
          icons={[]}
          onClick={() => {}}
          headers={[['playlist name', () => {}]]}
          colFormatter={p => [p.name]}
          items={playlists}
         />

        <div style={{...globalStyles.flexBoxRow, marginTop: 'auto'}}>
          <div style={{flex: 2, overflow: 'visible'}}>
            <Player
              onTrackEnd={() => this.playNextSongInQueue()}
              onNext={() => this.playNextSongInQueue()}
              onPrev={() => this.playPrevSongInQueue()}/>
          </div>
          <div style={{marginTop: 'auto', marginLeft: 'auto', overflow: 'visible',
            marginRight: 20,
            marginBottom: 20,
          }}>
            <FloatingActionButton onClick={() => {this.setState({openedPlaylistWindow: true})}}>
              <ContentAdd />
            </FloatingActionButton>
          </div>
        </div>
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
    )
  }
}

export default connect(
  s => {
    const {errorMessage, playlists} = s.library
    return {errorMessage, playlists}
  },
  d => (bindActionCreators({deleteSongFromLibrary, createPlaylist, deletePlaylist, removeFromPlaylist, addSongToPlaylist, playSong, playPlaylist}, d))
)(Library)

const styles = {
}
