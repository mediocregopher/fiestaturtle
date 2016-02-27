import Window from './window'
import {ipcMain} from 'electron'
import resolveRoot from '../resolve-root'
import hotPath from '../hot-path'

export default function () {
  const mainWindow = new Window(
    resolveRoot(`renderer/index.html?src=${hotPath('index.bundle.js')}`), {
      title: 'Fiesta Turtle',
      show: true
    }
  )

  ipcMain.on('showMain', () => {
    mainWindow.show(true)
  })

  return mainWindow
}
