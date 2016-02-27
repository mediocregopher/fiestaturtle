/* @flow */

import {combineReducers} from 'redux'
import {subReducer as router} from './router'
import upload from './upload'

const combinedReducer = combineReducers({
  router,
  upload
})

let reducer = combinedReducer

export default reducer
