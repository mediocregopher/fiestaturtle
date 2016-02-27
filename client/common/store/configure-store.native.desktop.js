import {createStore, compose} from 'redux'
import DevTools from '../../desktop/renderer/redux-dev-tools'

export default f => {
  return compose(f, DevTools.instrument())(createStore)
}
