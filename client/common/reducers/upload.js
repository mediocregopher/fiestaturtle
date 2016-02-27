// @flow

import * as Constants from '../constants/upload'

import type {UploadActions, SongMetaData} from '../constants/upload'

export type UploadState = {
  phase: 'TODO' | 'userUpload' | 'checking' | 'metaEdit' | 'uploadToService' | 'done',
  checkingMeta: boolean,
  files: Array<string>, // file paths
  filesMetaData: Array<SongMetaData>,
  metaReadError: Array<string>,
  editMessage: string,
  editingMetaDataIndex: number
}

const initialState: UploadState = {
  phase: 'TODO',
  checkingMeta: false,
  files: [],
  filesMetaData: [],
  metaReadError: [],
  editMessage: ['¡Be a Good Persona!', '¡Be everyone\'s best amigo!', '¡Meta-Money!'][Math.floor(Math.random()*3)],
  editingMetaDataIndex: -1
}

export default function (state: UploadState = initialState, action: UploadActions): UploadState {
  switch (action.type) {
    case Constants.checkingFileMeta:
      if (action.error) {
        return state
      } else {
        return {
          ...state,
          checkingMeta: true,
          files: action.payload.files,
          filesMetaData: [],
          metaReadError: []
        }
      }
    case Constants.showFileMeta:
      if (action.error) {
        return {
          ...state,
          metaReadError: [...state.metaReadError, action.payload.error]
        }
      } else {
        return {
          ...state,
          checkingMeta: (state.filesMetaData.length + 1) !== state.files.length,
          filesMetaData: [...state.filesMetaData, action.payload.metadata],
          editingMetaDataIndex: (state.filesMetaData.length + 1) === state.files.length ? 0 : state.editingMetaDataIndex,
        }
      }

    case Constants.updateFileMeta:
      if (action.error) {
        return state
      } else {
        const copy = [...state.filesMetaData]
        copy[state.editingMetaDataIndex] = action.payload.metadata
        return {
          ...state,
          filesMetaData: copy,
          editingMetaDataIndex: (state.editingMetaDataIndex + 1 === state.files.length ? -1 : state.editingMetaDataIndex + 1),
          phase: state.editingMetaDataIndex + 1 === state.files.length ? 'uploadToService' : state.phase
        }
      }

    case Constants.uploadFinished:
      if (action.error) {
        console.error('err', action.error)
        return state
      } else {
        return {
          ...initialState,
          phase: 'done'
        }
      }

    case Constants.reset:
      return initialState

    default:
      return state
  }
}
