// @flow

import * as Constants from '../constants/library'

import type {Song, Playlist} from '../constants/types/ft'
import type {TypedAsyncAction} from '../constants/types/flux'
import type {SwapPlaylist, RemoveFromPlaylist, CreatePlaylist, DeletePlaylist} from '../constants/library'

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
export function playSong (playlist: Playlist): TypedAsyncAction<DeletePlaylist> {
  return dispatch => {
  }
}

export function playPlaylist (playlist: Playlist): TypedAsyncAction<DeletePlaylist> {
  return dispatch => {
  }
}


// TODO delete playlist from brian and create it again
function createPlaylistCall (playlist: Playlist): Promise {
  return new Promise((resolve, reject) => {
    resolve(playlist)
  })
}

function deletePlaylistCall (playlist: Playlist): Promise {
  return new Promise((resolve, reject) => {
    resolve()
  })
}
