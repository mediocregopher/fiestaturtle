// @flow

import * as Constants from '../constants/upload'
import fs from 'fs'
import mm from 'musicmetadata'

import type {TypedAsyncAction, TypedAction} from '../constants/types/flux'
import type {CheckFileMeta, ShowFileMeta, UpdateFileMeta, UploadFinished, SongMetaData, Reset} from '../constants/upload'

function checkingFileMeta (files: Array<string>): CheckFileMeta {
  return {
    type: Constants.checkingFileMeta,
    payload: {files}
  }
}

export function showFileMeta (files: Array<string>): TypedAsyncAction<ShowFileMeta | CheckFileMeta> {
  return (dispatch, getState) => {
    dispatch(checkingFileMeta(files))
    files.map(f => {
      mm(fs.createReadStream(f), function (err, metadata) {
        if (err) {
          dispatch({type: Constants.showFileMeta, error: true, payload: {error: err}})
        } else {
          dispatch({type: Constants.showFileMeta, payload: {metadata}})
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

      setTimeout(() => dispatch({type: Constants.uploadFinished, payload: {}}), 3e3)
    }
  }
}

export function reset (metadata: SongMetaData): Reset {
  return {type: Constants.reset, payload: {}}
}
