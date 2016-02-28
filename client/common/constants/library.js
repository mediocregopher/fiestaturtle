// @flow

import type {TypedAction} from '../constants/types/flux'
import type {Song, Playlist} from '../constants/types/ft'

export const addSongToPlaylist = 'library:addSongToPlaylist'
export type AddSongToPlaylist = TypedAction<'library:addSongToPlaylist', {song: Song, playlist: Playlist}, {error: string}>

export const removeFromPlaylist = 'library:removeFromPlaylist'
export type RemoveFromPlaylist = TypedAction< 'library:removeFromPlaylist', {song: Song, playlist: Playlist}, {error: string}>

export const createPlaylist = 'library:createPlaylist'
export type CreatePlaylist = TypedAction<'library:createPlaylist', {playlist: Playlist}, {error: string}>

export const deletePlaylist = 'library:deletePlaylist'
export type DeletePlaylist = TypedAction<'library:deletePlaylist', {playlist: Playlist}, {error: string}>

export const swapPlaylist = 'library:swapPlaylist'
export type SwapPlaylist = TypedAction<'library:swapPlaylist', {oldPlaylist: Playlist, playlist: Playlist}, {error: string}>

export type LibraryActions = AddSongToPlaylist | RemoveFromPlaylist | CreatePlaylist | DeletePlaylist | SwapPlaylist
