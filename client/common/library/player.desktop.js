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
    console.log('asdf')
  }

  componentWillReceiveProps (nextProps) {
    if (this.props.sourceUrl !== nextProps.sourceUrl && nextProps.sourceUrl != null) {
      console.log('here!!')
      this.setState({paused: false})
    }
  }

  componentDidUpdate (oldProps, oldState) {
    if (oldState !== this.state) {
      console.log('next state', this.state.paused)
      if (this.state.paused) {
        this.refs.playerRef.pause()
      } else {
        this.refs.playerRef.play()
      }
    }
  }

  render () {
    console.log('source', this.props.sourceUrl)
    return (
      <div style={{...globalStyles.flexBoxRow, alignItems: 'center', justifyContent: 'space-around', width: 230,
       marginRight: 'auto', marginLeft: 'auto', overflow: 'visible'}}>

        <FloatingActionButton mini onClick={this.props.onPrev}>
          Previous
        </FloatingActionButton>

        <FloatingActionButton onClick={() => {this.setState({paused: !this.state.paused})}}>
          {this.state.paused ? 'Play' : 'Pause'}
        </FloatingActionButton>

        <FloatingActionButton mini onClick={this.props.onNext}>
          Next
        </FloatingActionButton>

        {this.props.sourceUrl && <audio ref='playerRef' src={this.props.sourceUrl}/>}

      </div>
    )
  }
}

export default connect(
  s => ({sourceUrl: s.library.nowPlayingSourceUrl})
)(Player)
