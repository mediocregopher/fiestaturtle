// @flow

export type SongMeta = {
  title: string,
  artist: Array<string>,
  album: string,
  albumartist: Array<string>,
  year: string,
  track: NoOf,
  disk: NoOf,
  genre: Array<string>,
  picture: Array<SongImage>,
  duration: number,
}

export type Base64 = string
export type BlockID = string
export type NodeID = string
export type bytes = any

export type Uploaded = {
  id?: BlockID,
  uploaderID?: NodeID,
  uploaderSig?: bytes
}

export type SongImage = {
  format: string,
  song: Base64,
  id: ?BlockID
}

export type NoOf = {
  no: number,
  of: number
}

export type Playlist = {
  version: number,
  name: string,
  songs: Array<Song>
} & Uploaded

export type UploadedSongMeta = {
  dataID: BlockID
}

export type Song = {
  version: number,
  meta: SongMeta,
  UploadedSongMeta: UploadedSongMeta,
  url: string
} & Uploaded

export type User = {
  version: number,
  name: string,
  playlist: Array<Playlist>
} & Uploaded
