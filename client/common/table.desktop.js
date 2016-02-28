// @flow

import React, {Component} from 'react'
import {globalStyles} from './styles/style-guide'
import MenuItem from 'material-ui/lib/menus/menu-item'
import FloatingActionButton from 'material-ui/lib/floating-action-button'
import ContentAdd from 'material-ui/lib/svg-icons/content/add'
import Dialog from 'material-ui/lib/dialog'
import FlatButton from 'material-ui/lib/flat-button'
import TextField from 'material-ui/lib/text-field';
import {bindActionCreators} from 'redux'
import Divider from 'material-ui/lib/divider'
import ReactList from 'react-list'

export default class SweetTable extends Component {
  props: {
    onClick: (index: number) => void,
    icons: [string, (rowIndex: number) => void],
    headers: [string, () => void],
    colFormatter: (item: any) => Array<string>,
    items: Array<any>
  };

  renderItem (item: any, key: any, index: number) {
    const colStyle = {flex: 1, overflow: 'scroll'}
    return (
      <div
        style={{...globalStyles.clickable, ...globalStyles.noSelect}} key={key}>
        <MenuItem innerDivStyle={{paddingRight: 0}} onTouchTap={() => this.props.onClick(index)}>
          <div style={{display: 'flex'}}>
            {this.props.colFormatter(item).map(i => <span key={i} style={colStyle}>{i}</span>)}
            {this.props.icons.map(([name, onClick]) => (
              <div style={{width: 30, position: 'relative'}}>
                <FlatButton key={name} style={{margin: 0, padding: 0, height: 50, width: 20, flex: 1}} onClick={() => onClick(index)}>
                  <i className='material-icons' style={{position: 'relative', right: 30, top: 5}}>{name}</i>
                </FlatButton>
              </div>
            ))}
          </div>
        </MenuItem>
      </div>
    )
  }

  render () {
    const colStyle = {flex: 1, overflow: 'scroll', textAlign: 'left'}

    return (
      <div style={{...globalStyles.noSelect, ...globalStyles.flexBoxColumn, flex: 2}}>
        <div style={{display: 'flex', marginLeft: 16}}>
          {this.props.headers.map(([name, onClick]) => (
            <FlatButton key={name} labelStyle={{padding: 0}} style={colStyle} label={name} onClick={onClick}/>
          ))}
          <div style={{width: 30 * this.props.icons.length}}/>
        </div>
        <Divider/>

        <ReactList
          itemRenderer={(i, k) => this.renderItem(this.props.items[i], k, i)}
          length={this.props.items.length}
          type='uniform'/>

      </div>
    )
  }
}

