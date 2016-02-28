// @flow
import React, {Component} from 'react'
import {connect} from 'react-redux'
import FileProcessor from 'react-file-processor'
import RaisedButton from 'material-ui/lib/raised-button'
import {bindActionCreators} from 'redux'

import {navigateUp} from '../actions/router'
import {showFileMeta, updateFileMeta, reset} from '../actions/upload'


import EditMetaData from './editMetaData.desktop'

import {globalStyles} from '../styles/style-guide'

class Upload extends Component {
  handleClick() {
    this.refs.myFileInput.chooseFile();
  }

  handleFileSelect(e, files) {
    console.log(e, files)
    console.log('paths:', files.map(f => f.path))
    this.props.showFileMeta(files.map(f => f.path))
  }

  render () {
    if (this.props.upload.metaReadError.length > 0) {
      return (
        <div style={{...globalStyles.flexBoxColumn, marginTop: 40, marginLeft: 'auto', marginRight: 'auto', alignItems: 'center'}}>
          <div> Uploading error {this.props.upload.metaReadError}</div>
        </div>
      )
    }

    if (this.props.upload.checkingMeta) {
      return (
        <div style={{...globalStyles.flexBoxColumn, marginTop: 40, marginLeft: 'auto', marginRight: 'auto', alignItems: 'center'}}>
          <div> Checking meta data... {this.props.upload.filesMetaData.length + 1} of {this.props.upload.files.length}</div>
        </div>
      )
    }

    if (this.props.upload.filesMetaData.length === this.props.upload.files.length && this.props.upload.files.length > 0 && this.props.upload.phase !== 'uploadToService') {
      const {editingMetaDataIndex, filesMetaData, editMessage} = this.props.upload
      return (
        <EditMetaData editMessage={editMessage} index={editingMetaDataIndex} songMetas={filesMetaData.map(m => m.metadata)} nextOrDone={(songMeta) => {this.props.nextOrDone(songMeta)}}/>
      )
    }

    if (this.props.upload.phase === 'uploadToService') {
      return (
        <div> Uploading to service now </div>
      )
    }

    if (this.props.upload.phase === 'done') {
      return (
        <div style={{...globalStyles.flexBoxColumn, marginTop: 40, marginLeft: 'auto', marginRight: 'auto', alignItems: 'center'}}>
          <div> Done! </div>
          <RaisedButton primary label='To my library!' style={{marginTop: 40}}
            onClick={() => {this.props.reset(); this.props.navUp()}}/>
        </div>
      )
    }

    return (
      <div style={{...globalStyles.flexBoxColumn, marginTop: 40, marginLeft: 'auto', marginRight: 'auto', alignItems: 'center'}}>

        <span style={{fontSize: 40}}>¡⌛ de upload!</span>
        <span style={{fontSize: 20, lineHeight: '30px', marginTop: 20}}>or drag n' drop anywhere on this page</span>

        <FileProcessor
          ref='myFileInput'
          onFileSelect={(e, files) => this.handleFileSelect(e, files)}>
          <RaisedButton primary label='Open Local File' style={{marginTop: 40}}
            onClick={() => this.handleClick()}/>
        </FileProcessor>
      </div>
    )
  }
}

export default connect(
  ({upload}) => ({upload}),
  (dispatch) => ({
    ...(bindActionCreators({showFileMeta}, dispatch)),
    nextOrDone: (m) => dispatch(updateFileMeta(m)),
    reset: () => dispatch(reset()),
    navUp: () => dispatch(navigateUp())
  }),
  (s, d, o) => ({...s, ...d, ...o}),
)(Upload)
