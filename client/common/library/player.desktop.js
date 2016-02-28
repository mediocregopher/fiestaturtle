import {connect} from 'react-redux'
import Paper from 'material-ui/lib/paper'
import React, {Component} from 'react'
import FloatingActionButton from 'material-ui/lib/floating-action-button'
import {globalStyles} from '../styles/style-guide'

class Player extends Component {
  props: {
    sourceUrl: string,
    onNext: () => void,
    onPrev: () => void
  };

  state: {
    paused: booleaan
  };

  constructor (props) {
    super(props)
    this.state = {paused: true}
  }

  render () {
    return (
      <div style={{...globalStyles.flexBoxRow, alignItems: 'center', justifyContent: 'space-around', width: 230,
       marginRight: 'auto', marginLeft: 'auto', overflow: 'visible'}}>

        <FloatingActionButton mini>
          Previous
        </FloatingActionButton>

        <FloatingActionButton>
          {this.state.paused ? 'Play' : 'Pause'}
        </FloatingActionButton>

        <FloatingActionButton mini>
          Next
        </FloatingActionButton>

      </div>
    )
  }
}

export default connect(
  s => ({sourceUrl: s.library.nowPlayingSourceUrl})
)(Player)
