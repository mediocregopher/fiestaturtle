// @flow

import * as Constants from '../constants/library'
import type {Song, Playlist} from '../constants/types/ft'
import type {LibraryActions} from '../constants/library'

const dummyPlaylists = []

export type LibraryState = {
  playlists: Array<Playlist>,
  errorMessage: ?string
}

const initialState: LibraryState = {
  playlists: [],
  errorMessage: null
}

export function deriveAllSongs (playlists: Array<Playlist>): Array<Song> {
  return playlists.reduce((acc, p) => acc.concat(p.songs), [])
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
          playlists: state.playlists.map(p => p === oldPlaylist ? playlist : p)
        }
      }

    default:
      return state
  }
}
