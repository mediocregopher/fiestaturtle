// @flow

import type {TypedAction} from '../constants/types/flux'

export type SongMetaData = {
  artist : Array<string>,
  album : ?string,
  albumartist : Array<string>,
  title : ?string,
  year : ?string,
  track : ?{ no : number, of : number },
  disk : ?{ no : number, of : number },
  genre : Array<string>,
  picture : Array<{format : string, data : any}>,
  duration : number //in seconds
}

export const checkingFileMeta = 'upload:checkingFileMeta'
export type CheckFileMeta = TypedAction<'upload:checkingFileMeta', {files: Array<string>}, {}>

export const showFileMeta = 'upload:showFileMeta'
export type ShowFileMeta = TypedAction<'upload:showFileMeta', {metadata: SongMetaData}, {error: any}>

export const updateFileMeta = 'upload:updateFileMeta'
export type UpdateFileMeta = TypedAction<'upload:updateFileMeta', {metadata: SongMetaData}, {error: any}>

export const uploadFinished = 'upload:finished'
export type UploadFinished = TypedAction<'upload:finished', {}, {error: any}>

export const reset = 'upload:reset'
export type Reset = TypedAction<'upload:reset', {}, {}>

export type UploadActions = CheckFileMeta | ShowFileMeta | UpdateFileMeta | UploadFinished | Reset
