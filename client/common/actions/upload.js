// @flow

import * as Constants from '../constants/upload'
import fs from 'fs'
import mm from 'musicmetadata'

import {rpc} from '../util/rpc'

import {swapPlaylistCall} from './library'

import type {TypedAsyncAction, TypedAction} from '../constants/types/flux'
import type {CheckFileMeta, ShowFileMeta, UpdateFileMeta, UploadFinished, SongMetaData, Reset} from '../constants/upload'

function checkingFileMeta (files: Array<string>): CheckFileMeta {
  return {
    type: Constants.checkingFileMeta,
    payload: {files}
  }
}

function bufferToBase64(buf) {
  var binstr = Array.prototype.map.call(buf, function (ch) {
    return String.fromCharCode(ch);
  }).join('');
  return btoa(binstr);
}

export function showFileMeta (files: Array<string>): TypedAsyncAction<ShowFileMeta | CheckFileMeta> {
  return (dispatch, getState) => {
    dispatch(checkingFileMeta(files))
    files.map(f => {
      mm(fs.createReadStream(f), function (err, metadata) {
        if (err) {
          dispatch({type: Constants.showFileMeta, error: true, payload: {error: err}})
        } else {
          dispatch({type: Constants.showFileMeta, payload: {path: f, metadata}})
        }
      })
    })
  }
}

export function updateFileMeta (metadata: SongMetaData): TypedAsyncAction<UpdateFileMeta | UploadFinished> {
  return (dispatch, getState) => {
    dispatch({type: Constants.updateFileMeta, payload: {metadata}})
    if (getState().upload.phase === 'uploadToService') {
      // TODO upload to brian

      const songsMeta = getState().upload.filesMetaData
      const songsUploaded = songsMeta.map(({path, metadata}) => {
        const picture = metadata.picture.map(p => ({...p, data: bufferToBase64(p.data)}))
        const song = {
          meta: {...metadata, picture},
          version: 1,
        }

        return rpc('UploadSongData', {song, songPath: path})
          .then(({result, error}) => {
            console.log('after upload', result, error)
            if (error) {
              console.error('error uploading', error)
            } else {
              return result.song
            }
          })
      })

      Promise.all(songsUploaded).then((songs) => {
        const library = getState().library.playlists.filter(p => p.name === 'Library')[0]
        const newLibrary = {...library, songs: library.songs.concat(songs)}
        console.log('done uploading songs')
        swapPlaylistCall(library, newLibrary).then(({result, error}) => {
          if (error) {
            console.error('error uploading', error)
          } else {
            const newLibrary = result.playlist
            const newPlaylists = getState().library.playlists.map(p => p.name === 'Library' ? newLibrary : p)
            dispatch({type: 'library:updatePlaylists', payload: {playlists: newPlaylists}})
            dispatch({type: Constants.uploadFinished, payload: {}})
          }
        })
      })
    }
  }
}

export function reset (): Reset {
  return {type: Constants.reset, payload: {}}
}
