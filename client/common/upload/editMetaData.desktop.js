// @flow
import React, {Component} from 'react'
import {globalStyles} from '../styles/style-guide'

import TextField from 'material-ui/lib/text-field';
import RaisedButton from 'material-ui/lib/raised-button'

import type {SongMetaData} from '../constants/upload'

export default class Upload extends Component {
  props: {
    editMessage: string,
    songMetas: Array<SongMetaData>,
    index: number,
    nextOrDone: (meta: SongMetaData) => void
  };

  state: SongMetaData;

  constructor (props: any) {
    super(props)

    this.state = props.songMetas[this.props.index]
  }

  render () {
    const {songMetas} = this.props
    const songMeta = songMetas[this.props.index]
    return (
      <div style={{...globalStyles.flexBoxColumn, marginTop: 40, marginBottom: 40, marginLeft: 'auto', marginRight: 'auto', alignItems: 'center'}}>
        <div style={{fontStyle: 'italic', marginLeft: 20, marginRight: 20, marginBottom: 20, fontSize: 40}}>
          {this.props.editMessage}
        </div>
        <span style={{marginLeft: 20, marginRight: 20}}>
          Update Song MetaData for: {songMeta.title} ({this.props.index + 1} of {songMetas.length})
        </span>
        <div style={globalStyles.flexBoxColumn}>
          <TextField
            floatingLabelText={"Title"}
            onChange={({target: {value: title}}) => this.setState({title})}
            defaultValue={songMeta.title}/>
          <TextField
            floatingLabelText={"Artist"}
            onChange={({target: {value: artist}}) => this.setState({artist})}
            defaultValue={songMeta.artist}/>
          <TextField
            floatingLabelText={"Album"}
            onChange={({target: {value: album}}) => this.setState({album})}
            defaultValue={songMeta.album}/>
          <TextField
            floatingLabelText={"Genre (Optional)"}
            onChange={({target: {value: genre}}) => this.setState({genre})}
            defaultValue={songMeta.genre}/>
          <TextField
            floatingLabelText={"Album Artists (Optional)"}
            onChange={({target: {value: albumartist}}) => this.setState({albumartist})}
            defaultValue={songMeta.albumartist}/>
          <TextField
            floatingLabelText={"Year (Optional)"}
            onChange={({target: {value: year}}) => this.setState({year})}
            defaultValue={songMeta.year}/>
        </div>
        <div style={globalStyles.flexBoxColumn}>
          <RaisedButton primary label={songMetas.length - 1 === this.props.index ? 'Done' : 'Next song'} style={{marginTop: 40}}
            onClick={() => this.props.nextOrDone(this.state)}/>
        </div>
      </div>
    )
  }
}
