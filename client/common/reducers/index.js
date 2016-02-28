/* @flow */

import {combineReducers} from 'redux'
import {subReducer as router} from './router'
import upload from './upload'
import library from './library'

const combinedReducer = combineReducers({
  router,
  upload,
  library
})

let reducer = combinedReducer

export default reducer
