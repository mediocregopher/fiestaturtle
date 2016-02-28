// @flow

import Immutable, {Set} from 'immutable'

import * as Constants from '../constants/library'
import type {Song, Playlist} from '../constants/types/ft'
import type {LibraryActions} from '../constants/library'

import {katyPerry, jBieb} from '../util/demoArt'

const dummySongs: Array<Song> = [
  {
    version: 1,
    url: '',
    meta: {
      title: 'bieber fieber',
      artist: ['J. Biebs'],
      album: '¡Bielieve!',
      albumartist: [],
      year: '20xx',
      track: null,
      disk: null,
      genre: ['Human Music'],
      picture: [{
        format: 'png',
        data: jBieb,
        id: null
      }],
      duration: 400
    }
  },
  {
    version: 1,
    url: '',
    meta: {
      title: 'Works of Fire',
      artist: ['K. Perry'],
      album: 'Let it Burn',
      albumartist: [],
      year: '20xx',
      track: null,
      disk: null,
      genre: ['Human Music'],
      picture: [{
        format: 'png',
        data: katyPerry,
        id: null
      }],
      duration: 400
    }
  }
]

const dummyPlaylists: Array<Playlist> = [
  {version: 1, name: 'Sweet tunes', songs: dummySongs},
  {version: 1, name: 'Other Sweet tunes', songs: dummySongs}
]

export type LibraryState = {
  playlists: Array<Playlist>,
  nowPlayingSourceUrl: ?string,
  errorMessage: ?string
}

const initialState: LibraryState = {
  nowPlayingSourceUrl: null,
  playlists: dummyPlaylists,
  errorMessage: null
}

export function deriveAllSongs (playlists: Array<Playlist>): Array<Song> {
  return playlists.map(p => Immutable.fromJS(p)).reduce((acc, p) => acc.union(p.get('songs')), Set()).toJS()
}

export default function (state: LibraryState = initialState, action: LibraryActions): LibraryState {
  switch (action.type) {
    case Constants.createPlaylist:
      if (action.error) {
        return {
          ...state,
          errorMessage: action.payload.error
        }
      } else {
        return {
          ...state,
          playlists: state.playlists.concat(action.payload.playlist)
        }
      }

    case Constants.deletePlaylist:
      if (action.error) {
        return {
          ...state,
          errorMessage: action.payload.error
        }
      } else {
        const playlistToDelete = action.payload.playlist
        return {
          ...state,
          playlists: state.playlists.filter(p => p !== playlistToDelete)
        }
      }

    case Constants.swapPlaylist:
      if (action.error) {
        return {
          ...state,
          errorMessage: action.payload.error
        }
      } else {
        const {oldPlaylist, playlist} = action.payload
        return {
          ...state,
          playlists: state.playlists.map(p => p.id === oldPlaylist.id ? playlist : p)
        }
      }

    case 'library:updatePlaylists':
      if (action.error) {
        return {
          ...state,
          errorMessage: action.payload.error
        }
      } else {
        const {playlists} = action.payload
        return {
          ...state,
          playlists: playlists.map(p => p.songs ? p : {...p, songs: []})
        }
      }

    case 'library:playSong':
      if (action.error) {
        return {
          ...state,
          errorMessage: action.payload.error
        }
      } else {
        const {playlists} = action.payload
        return {
          ...state,
          nowPlayingSourceUrl: action.payload.url
        }
      }

    default:
      return state
  }
}
