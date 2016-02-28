// @flow

import * as Constants from '../constants/library'

import type {Song, Playlist} from '../constants/types/ft'
import type {TypedAsyncAction} from '../constants/types/flux'
import type {SwapPlaylist, RemoveFromPlaylist, CreatePlaylist, DeletePlaylist} from '../constants/library'

import {rpc} from '../util/rpc'

export function addSongToPlaylist (song: Song, playlist: Playlist): TypedAsyncAction<SwapPlaylist> {
  return dispatch => {
    const nextPlaylist = {...playlist, songs: playlist.songs.concat(song)}
    // TODO delete playlist from brian and create it again
    return deletePlaylistCall(playlist)
      .then(() => createPlaylistCall(nextPlaylist))
      .then(p => dispatch({type: Constants.swapPlaylist, payload: {playlist: p, oldPlaylist: playlist}}))
  }
}

export function removeFromPlaylist (song: Song, playlist: Playlist): TypedAsyncAction<SwapPlaylist> {
  return dispatch => {
    const nextPlaylist = {...playlist, songs: playlist.songs.filter(s => s !== song)}
    return deletePlaylistCall(playlist)
      .then(() => createPlaylistCall(nextPlaylist))
      .then(p => dispatch({type: Constants.swapPlaylist, payload: {playlist: p, oldPlaylist: playlist}}))
  }
}

export function createPlaylist (name: string): TypedAsyncAction<CreatePlaylist> {
  console.log('creating playlist', name)
  return (dispatch, getState) => {
    if (getState().library.playlists.filter(p => name === p.name).length > 0) {
      dispatch({type: Constants.createPlaylist, error: true, payload: {error: 'name already taken'}})
    } else {
      const playlist: Playlist = {
        version: 1,
        name,
        songs: []
      }
      createPlaylistCall(playlist).then(p => {
        dispatch({type: Constants.createPlaylist, payload: {playlist: p}})
      })
    }
  }
}

export function deletePlaylist (playlist: Playlist): TypedAsyncAction<DeletePlaylist> {
  return dispatch => {
    deletePlaylistCall(playlist)
      .then(() => dispatch({type: Constants.deletePlaylist, payload: {playlist}}))
  }
}


// TODO
export function playSong (song: Song): TypedAsyncAction<DeletePlaylist> {
  return dispatch => {
    // TODO get song url
    dispatch({type: 'library:playSong', payload: {url: ''}})
  }
}

export function playPlaylist (playlist: Playlist): TypedAsyncAction<DeletePlaylist> {
  return dispatch => {
  }
}

export function deleteSongFromLibrary (song: Song): any {
  return (dispatch, getState) => {
    const lib = getState().library.playlists.filter(p => p.name === 'Library')[0]
    const newSongs = lib.songs.filter(s => s !== song)
    swapPlaylistCall(lib, {...lib, songs: newSongs}).then(() => {
      dispatch(getUser())
    })
  }
}

export function getUser (): any {
  return dispatch => {
    return rpc('GetUser', {}).then(({error, result}) => {
      if (error) {
        console.error("Error fetching user:", error)
      } else {
        console.log('user result:', result)
        const playlists = result.user.playlists || []
        if (playlists.length === 0) {
          createPlaylistCall({
            version: 1,
            name: 'Library',
            songs: []
          }).then(() => setTimeout(getUser, 100))
        } else {
          dispatch({type: 'library:updatePlaylists', payload: {playlists}})
          dispatch(getPlaylists(playlists))
        }
      }
    })
  }
}

// TODO delete playlist from brian and create it again
function createPlaylistCall (playlist: Playlist): Promise {
  return rpc('CreatePlaylist', {playlist})
}

export function swapPlaylistCall (oldPlaylist: Playlist, newPlaylist: Playlist): Promise {
  return rpc('CreatePlaylist', {playlist: newPlaylist, replaces: oldPlaylist.id})
}

function deletePlaylistCall (playlist: Playlist): Promise {
  return rpc('DeletePlaylist', {playlistID: playlist.id})
}

window.deletePlaylist = deletePlaylistCall

export function getPlaylists (shellPlaylists: any): Promise {
  return dispatch => {
    return Promise.all(shellPlaylists.map(p => {
      console.log('here2')
      return rpc('GetPlaylist', {playlistID: p.id}).then(({result, error}) => {
        console.log('here3', error, result)
        if (error) {
          console.error('error: ', error)
          return
        }
        return result.playlist
      })
    })).then((newPlaylists) => {
      dispatch({type: 'library:updatePlaylists', payload: {playlists: newPlaylists}})
    })
  }
}
